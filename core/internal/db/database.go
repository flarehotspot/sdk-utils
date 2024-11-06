package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"core/internal/config"
	"core/internal/utils/mysql"

	sdkstr "github.com/flarehotspot/go-utils/strings"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	mu sync.RWMutex
	db *pgxpool.Pool
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

	dbConf, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	//https://stackoverflow.com/questions/39980902/golang-mysql-error-packets-go33-unexpected-eof
	dbConf.MaxConnLifetime = time.Minute * 4

	// Ensure postgresql starts up during boot before returning err
	openErrorCountThreshold := 5
	pgPool, err := pgxpool.NewWithConfig(context.Background(), dbConf)
	// conn, err := sql.Open("mysql", url)
	for openErrorCount := 0; err != nil && openErrorCount < openErrorCountThreshold; openErrorCount++ {
		pgPool, err = pgxpool.New(context.Background(), url)
		time.Sleep(time.Second * 2)
	}
	if err != nil {
		return nil, err
	}

	// TODO: find an equivalent postgresql sql query debugging

	err = CheckDatabaseConnection(pgPool)
	if err != nil {
		return nil, err
	}

	db.db = pgPool
	return &db, nil
}

func CheckDatabaseConnection(pool *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return pool.Ping(ctx)
}

func (d *Database) SqlDB() (db *pgxpool.Pool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.db
}

func (d *Database) SetSql(db *pgxpool.Pool) {
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
