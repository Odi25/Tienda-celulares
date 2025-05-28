package models

import "time"

type Compra struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UsuarioID  uint      `json:"usuario_id"`             // ID del cliente
	ProductoID uint      `json:"producto_id"`            // ID del producto
	Cantidad   int       `json:"cantidad"`               // Cantidad de productos
	Fecha      time.Time `json:"fecha"`                  // Fecha de compra
}
