package lib

import (
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(env *Env, logger *Logger) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		env.Db.Host, env.Db.User, env.Db.Password, env.Db.DbName, env.Db.Port, env.Db.SslMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})
	if err != nil {
		logger.Fatal("failed to connect database:", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Info("Url: ", dsn)
		logger.Panic(err)
	}
	sqlDB.SetMaxOpenConns(env.Db.MaxOpenConns)
	sqlDB.SetMaxIdleConns(env.Db.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(env.Db.MaxIdleTime) * time.Minute)
	logger.Info("Connected to the database")
	return &Database{
		DB: db,
	}
}

func ParsePgError(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		switch pgErr.ConstraintName {
		case "users_username_key":
			return "username has been registered"
		case "users_email_key":
			return "email has been registered"
		}
	}
	return ""
}
