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
	// Leer variables de entorno
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Validación opcional para detectar si alguna variable viene vacía
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Panic("❌ Faltan variables de entorno. Verifica DB_HOST, DB_PORT, DB_USER, DB_PASSWORD y DB_NAME")
	}

	// Crear cadena de conexión
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		host, user, password, dbname, port,
	)

	// Conectar con GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("❌ Error al conectar a la base de datos:", err)
	}

	// Guardar la instancia
	DB = db
	fmt.Println("✅ Conexión exitosa a PostgreSQL")
}

