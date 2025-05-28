package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarDB() {
	dsn := "host=localhost user=tienda_user password=Sistemas dbname=tienda_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Error al conectar a la base de datos: " + err.Error())
	}
	fmt.Println("✅ Conexión exitosa a PostgreSQL")
	var base string
	DB.Raw("SELECT current_database()").Scan(&base)
	fmt.Println("📌 Conectado a la base de datos:", base)
	var user string
	DB.Raw("SELECT current_user").Scan(&user)
	fmt.Println("👤 Conectado como:", user)

}
