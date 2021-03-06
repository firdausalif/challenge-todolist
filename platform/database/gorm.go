package database

import (
	"fmt"
	"github.com/firdausalif/challenge-todolist/pkg/config"
	"github.com/firdausalif/challenge-todolist/platform/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// database instance
type DB struct {
	*gorm.DB
}

var defaultDB = &DB{}
var err error

// connect sets the db client of database using configuration
func (db *DB) connect(cfg *config.DB) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name,
	)

	gormConfig := &gorm.Config{
		// enhance performance config
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	}

	if config.DBCfg().Debug {
		//gormConfig.Logger = logger.Default.LogMode(logger.Info)
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
	pgdb.SetMaxOpenConns(10000)
	pgdb.SetMaxIdleConns(10000)
	pgdb.SetConnMaxLifetime(10 * time.Minute)

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
