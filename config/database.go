package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarDB() {
	host := os.Getenv("postgresql://tienda_user:8zUkdYLMhoTTLQXOf5VmIkfOaDGRRs6w@dpg-d0r8rkbuibrs73d0o40g-a/tienda_db_ussv")         // Ej: dpg-xxxxx.render.com
	user := os.Getenv("tienda_user")         // Ej: tienda_user
	password := os.Getenv("8zUkdYLMhoTTLQXOf5VmIkfOaDGRRs6w") // Ej: tu contraseña
	dbname := os.Getenv("tienda_db_ussv")       // Ej: tienda_db
	port := os.Getenv("5432")         // Ej: 5432

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("❌ Error al conectar a la base de datos:", err)
	}

	DB = db
	fmt.Println("✅ Conexión exitosa a PostgreSQL")
}
