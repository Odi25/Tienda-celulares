package main

import (
	"github.com/gin-gonic/gin"
	"tienda-celulares/config"
	"tienda-celulares/models"
	"tienda-celulares/routes"
)

func main() {
	r := gin.Default()

	config.ConectarDB()
	config.DB.AutoMigrate(&models.Producto{}, &models.Usuario{}, &models.Compra{})

	// Registrar rutas de API primero
	routes.ConfigurarRutas(r)

	// Luego servir archivos estáticos desde /static
	r.Static("/web", "./static")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/web/index.html")
	})

	r.Run(":8080")
}
