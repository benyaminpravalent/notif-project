package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/project/notif-project/pkg/config"
	"github.com/project/notif-project/pkg/logger"
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
