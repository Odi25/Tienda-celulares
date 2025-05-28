package routes

import (
	"github.com/gin-gonic/gin"
	"tienda-celulares/controllers"
	"tienda-celulares/middleware"
)

func ConfigurarRutas(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/productos", controllers.ObtenerProductos)
		api.GET("/productos/:id", controllers.ObtenerProductoPorID)
		api.GET("/compras/:usuario_id", controllers.ObtenerComprasPorUsuario)
		api.POST("/login", controllers.Login)
		api.POST("/registro", controllers.Registro)
		api.GET("/debug/usuarios", controllers.DebugUsuarios)
		

		// Rutas protegidas
		api.POST("/productos", middleware.SoloAdmin(), controllers.CrearProducto)
		api.PUT("/productos/:id", middleware.SoloAdmin(), controllers.ActualizarProducto)
		api.DELETE("/productos/:id", middleware.SoloAdmin(), controllers.EliminarProducto)

		// Rutas para registrar compras
		api.POST("/compras", controllers.RegistrarCompra)
	}
}

