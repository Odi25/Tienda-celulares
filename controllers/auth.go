package controllers

import (
	"net/http"
	"tienda-celulares/config"
	"tienda-celulares/models"
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
)

func Login(c *gin.Context) {
	var input struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Elimina espacios
	correo := strings.TrimSpace(input.Correo)
	password := strings.TrimSpace(input.Password)

	fmt.Println("🧪 Intento de login:", correo, "/", password)

	var usuario models.Usuario
	if err := config.DB.Where("correo = ?", correo).First(&usuario).Error; err != nil {
		fmt.Println("❌ Correo no encontrado:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Correo incorrecto"})
		return
	}

	// Comparamos contraseña manualmente
	if usuario.Password != password {
		fmt.Println("❌ Contraseña incorrecta")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	c.SetCookie("rol", usuario.Rol, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Login exitoso",
		"rol":     usuario.Rol,
		"id":      usuario.ID,
	})
}


func Registro(c *gin.Context) {
	var nuevoUsuario models.Usuario

	if err := c.ShouldBindJSON(&nuevoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if nuevoUsuario.Rol == "" {
		nuevoUsuario.Rol = "cliente" // rol por defecto
	}

	// Verificar si ya existe el correo
	var existente models.Usuario
	if err := config.DB.Where("correo = ?", nuevoUsuario.Correo).First(&existente).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Correo ya registrado"})
		return
	}

	config.DB.Create(&nuevoUsuario)
	c.JSON(http.StatusCreated, gin.H{"mensaje": "Usuario registrado con éxito"})
}

func DebugUsuarios(c *gin.Context) {
	var usuarios []models.Usuario
	result := config.DB.Find(&usuarios)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}
