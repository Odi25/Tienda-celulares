package models
type Producto struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
	ImagenURL   string  `json:"imagen_url"`
}
