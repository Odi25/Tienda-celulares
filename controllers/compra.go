package controllers

import (
	"net/http"
	"time"
	"tienda-celulares/config"
	"tienda-celulares/models"
	"github.com/gin-gonic/gin"
)

// POST /api/compras
func RegistrarCompra(c *gin.Context) {
	var input models.Compra

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	input.Fecha = time.Now()

	if input.Cantidad <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cantidad debe ser mayor a 0"})
		return
	}

	// Buscar el producto
	var producto models.Producto
	if err := config.DB.First(&producto, input.ProductoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	// Verificar stock suficiente
	if producto.Stock < input.Cantidad {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stock insuficiente"})
		return
	}

	// Registrar compra
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar la compra"})
		return
	}

	// Descontar el stock
	producto.Stock -= input.Cantidad
	if err := config.DB.Save(&producto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el stock"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Compra registrada y stock actualizado",
		"compra":  input,
	})
}
func ObtenerComprasPorUsuario(c *gin.Context) {
	usuarioID := c.Param("usuario_id")
	var compras []models.Compra
	if err := config.DB.Where("usuario_id = ?", usuarioID).Find(&compras).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las compras"})
		return
	}
	c.JSON(http.StatusOK, compras)
}
