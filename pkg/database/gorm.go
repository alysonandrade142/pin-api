package database

import (
	"os"
	"pin-api/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Buscando variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Montando a string de conexão (DSN)
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword +
		" dbname=" + dbName + " port=" + dbPort + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.SetupLogger().Fatal("Falha ao conectar ao banco de dados: ", err)
	}

	return db
}
