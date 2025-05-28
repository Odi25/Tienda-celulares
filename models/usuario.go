package models

type Usuario struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `gorm:"unique" json:"correo"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
}
