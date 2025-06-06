package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func SoloAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        rol, err := c.Cookie("rol")
        if err != nil || rol != "admin" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Acceso solo para administradores"})
            return
        }
        c.Next()
    }
}