package controllers

import (
	"net/http"
	"tienda-celulares/config"
	"tienda-celulares/models"
	"github.com/gin-gonic/gin"
)

// GET /productos
func ObtenerProductos(c *gin.Context) {
	var productos []models.Producto
	result := config.DB.Find(&productos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, productos)
}

// GET /productos/:id
func ObtenerProductoPorID(c *gin.Context) {
	id := c.Param("id")
	var producto models.Producto
	if err := config.DB.First(&producto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}

// POST /productos
func CrearProducto(c *gin.Context) {
	var producto models.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&producto)
	c.JSON(http.StatusCreated, producto)
}

// PUT /productos/:id
func ActualizarProducto(c *gin.Context) {
	id := c.Param("id")
	var producto models.Producto

	// Buscar el producto por ID
	if err := config.DB.First(&producto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	var input models.Producto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Actualizar campos
	producto.Nombre = input.Nombre
	producto.Descripcion = input.Descripcion
	producto.Precio = input.Precio
	producto.Stock = input.Stock
	producto.ImagenURL = input.ImagenURL

	if err := config.DB.Save(&producto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el producto"})
		return
	}

	c.JSON(http.StatusOK, producto)
}


// DELETE /productos/:id
func EliminarProducto(c *gin.Context) {
	id := c.Param("id")
	var producto models.Producto
	if err := config.DB.Delete(&producto, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Producto eliminado"})
}
