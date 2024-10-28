package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"core/internal/config"
	"core/internal/utils/mysql"

	sdkstr "github.com/flarehotspot/go-utils/strings"
	//
	// UNCOMMENT BELOW LINES WHEN DEBUGGING SQL QUERIES:
	//
	// "github.com/rs/zerolog"
	// "os"
	// sqldblogger "github.com/simukti/sqldb-logger"
	// "github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

type Database struct {
	mu sync.RWMutex
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	dbpass := sdkstr.Rand(8)
	dbname := fmt.Sprintf("flarehotspot_%s", sdkstr.Rand(8))

	err := mysql.SetupDb(dbpass, dbname)
	if err != nil {
		log.Println("Error installing mariadb: ", err)
		return nil, err
	}

	var db Database

	cfg, err := CreateDb()
	if err != nil {
		return nil, err
	}

	url := cfg.DbUrlString()
	log.Println("DB URL: ", url)

	// Ensure mysql starts up during boot before returning err
	openErrorCountThreshold := 5
	conn, err := sql.Open("mysql", url)
	for openErrorCount := 0; err != nil && openErrorCount < openErrorCountThreshold; openErrorCount++ {
		conn, err = sql.Open("mysql", url)
		time.Sleep(time.Second * 2)
	}
	if err != nil {
		return nil, err
	}

	//https://stackoverflow.com/questions/39980902/golang-mysql-error-packets-go33-unexpected-eof
	conn.SetConnMaxLifetime(time.Minute * 4)

	// UNCOMMENT BELOW LINES WHEN DEBUGGING SQL QUERIES:
	//
	// loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	// conn = sqldblogger.OpenDriver(
	// url,
	// conn.Driver(),
	// loggerAdapter,
	// sqldblogger.WithMinimumLevel(sqldblogger.LevelInfo),
	// sqldblogger.WithLogDriverErrorSkip(false),
	// sqldblogger.WithSQLQueryAsMessage(false),
	// sqldblogger.WithWrapResult(false),
	// sqldblogger.WithIncludeStartTime(false),
	// sqldblogger.WithPreparerLevel(sqldblogger.LevelInfo),
	// sqldblogger.WithQueryerLevel(sqldblogger.LevelInfo),
	// sqldblogger.WithExecerLevel(sqldblogger.LevelInfo),
	// )

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	db.db = conn
	return &db, nil
}

func (d *Database) SqlDB() (db *sql.DB) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.db
}

func (d *Database) SetSql(db *sql.DB) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.db = db
}

func CreateDb() (*config.DbConfig, error) {
	cfg, err := config.ReadDatabaseConfig()
	if err != nil {
		return cfg, err
	}

	log.Println("DB conn string: ", cfg.BaseConnStr())
	db, err := sql.Open("mysql", cfg.BaseConnStr())
	if err != nil {
		log.Println("Error opening database: ", err)
		return cfg, err
	}
	defer db.Close()

	log.Println("Creating database " + cfg.Database + "...")
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.Database)
	if err != nil {
		log.Println("Unable to create database:", err)
	}

	return cfg, nil
}
