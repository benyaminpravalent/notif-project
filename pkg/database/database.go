package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/richardsahvic/jamtangan/pkg/config"
	"github.com/richardsahvic/jamtangan/pkg/logger"
)

var DB *sqlx.DB

// InitMySql initiates mysql connection and store it to DB
func InitMySql(ctx context.Context) {
	l := logger.GetLoggerContext(ctx, "database", "Connect")

	dsn := config.GetString("mysql_dsn")
	l.Info(dsn)

	dbConnection, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	err = dbConnection.Ping()
	if err != nil {
		panic(err.Error())
	}

	l.Info("Connected to MySQL")

	DB = dbConnection
}
