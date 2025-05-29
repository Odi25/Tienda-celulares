package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"tienda-celulares/config"
	"tienda-celulares/models"
	"tienda-celulares/routes"
)

func main() {
	r := gin.Default()

	// Conexión y migración a la base de datos
	config.ConectarDB()
	config.DB.AutoMigrate(&models.Producto{}, &models.Usuario{}, &models.Compra{})

	// Registrar rutas de API
	routes.ConfigurarRutas(r)

	// Servir archivos estáticos (HTML, JS, CSS)
	r.Static("/web", "./static")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/web/index.html")
	})

	// ✅ Puerto dinámico para Render o 8080 por defecto local
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
