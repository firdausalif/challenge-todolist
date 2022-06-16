package database

import (
	"fmt"
	"github.com/firdausalif/challenge-todolist/pkg/config"
	"github.com/firdausalif/challenge-todolist/platform/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// database instance
type DB struct {
	*gorm.DB
}

var defaultDB = &DB{}
var err error

// connect sets the db client of database using configuration
func (db *DB) connect(cfg *config.DB) error {
	dsn := fmt.Sprintf("user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Host,
		cfg.Port,
		cfg.SslMode,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	gormConfig := &gorm.Config{}
	if config.DBCfg().Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db.DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return err
	}

	// gorm migration
	migrations.Migrate(db.DB)

	pgdb, err := db.DB.DB()
	if err != nil {
		return err
	}

	// connection pool settings
	pgdb.SetMaxOpenConns(cfg.MaxOpenConn)
	pgdb.SetMaxIdleConns(cfg.MaxIdleConn)
	pgdb.SetConnMaxLifetime(cfg.MaxConnLifetime)

	// Try to ping database.
	if err := pgdb.Ping(); err != nil {
		defer pgdb.Close() // close database connection
		return fmt.Errorf("can't sent ping to database, %w", err)
	}

	return nil
}

// GetDB returns db instance
func GetDB() *DB {
	return defaultDB
}

// ConnectDB sets the db client of database using default configuration
func ConnectDB() error {
	return defaultDB.connect(config.DBCfg())
}
