package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/project/notif-project/pkg/logger"
)

var DB *sqlx.DB

// InitMySql initiates mysql connection and store it to DB
func InitMySql(ctx context.Context) {
	l := logger.GetLoggerContext(ctx, "database", "Connect")

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	db, err := sqlx.Open(dialect, dbURI)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	l.Info("Connected to PostgreSQL")

	DB = db
}
