package config

import (
	"os"
	"strconv"
	"time"
)

// DB holds the DB configuration
type DB struct {
	Host            string
	Port            int
	SslMode         string
	Name            string
	User            string
	Password        string
	Debug           bool
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifetime time.Duration
}

var db = &DB{}

// DBCfg returns the default DB configuration
func DBCfg() *DB {
	return db
}

// LoadDBCfg loads DB configuration
func LoadDBCfg() {
	db.Host = os.Getenv("MYSQL_HOST")
	db.Port, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	db.User = os.Getenv("MYSQL_USER")
	db.Password = os.Getenv("MYSQL_PASSWORD")
	db.Name = os.Getenv("MYSQL_DBNAME")
	db.SslMode = os.Getenv("DB_SSL_MODE")
	db.Debug, _ = strconv.ParseBool(os.Getenv("DB_DEBUG"))
	db.MaxOpenConn, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	db.MaxIdleConn, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	lifeTime, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	db.MaxConnLifetime = time.Duration(lifeTime) * time.Second
}
