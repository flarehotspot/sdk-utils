package db

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"core/internal/config"
	"core/internal/db/sqlc"
	"core/internal/utils/pg"

	sdkstr "github.com/flarehotspot/go-utils/strings"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	mu      sync.RWMutex
	db      *pgxpool.Pool
	Queries sqlc.Queries
}

func NewDatabase() (*Database, error) {
	dbpass := sdkstr.Rand(8)
	dbname := strings.ToLower(fmt.Sprintf("flarehotspot_%s", sdkstr.Rand(8)))

	// Setup PostgreSQL server
	err := pg.SetupServer(dbpass, dbname)
	if err != nil {
		log.Println("Error installing postgres db: ", err)
		return nil, err
	}

	// Wait for the postgres server to be ready
	maxPortCheckTries := 30
	portCheckIndex := 0
	portOK := false
	for portCheckIndex < maxPortCheckTries {
		ok := CheckPostgresPort("localhost")
		if ok {
			portOK = true
			break
		} else {
			portCheckIndex++
			time.Sleep(1 * time.Second)
		}
	}

	if !portOK {
		panic("Unable to connect to the local postgres server!")
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
	for openErrorCount := 0; err != nil && openErrorCount < openErrorCountThreshold; openErrorCount++ {
		log.Println("Checking database connection...")
		pgPool, err = pgxpool.New(context.Background(), url)
		time.Sleep(time.Second * 2)
		log.Println("Error opening database: ", err)
	}
	if err != nil {
		log.Println("Error connecting to database.")
		return nil, err
	}

	// TODO: find an equivalent postgresql sql query debugging

	err = CheckDatabaseConnection(pgPool)
	if err != nil {
		return nil, err
	}

	db.Queries = *sqlc.New(pgPool)
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

	ctx := context.Background()
	log.Println("DB base conn string: ", cfg.BaseConnStr())
	connPool, err := pgx.Connect(ctx, cfg.BaseConnStr())
	if err != nil {
		log.Println("Error opening database: ", err)
		return cfg, err
	}
	defer connPool.Close(ctx)

	log.Println("Creating database " + cfg.Database + "...")
	_, err = connPool.Exec(context.Background(), "CREATE DATABASE "+cfg.Database)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Println("Unable to create database:", err)
			return nil, err
		}
		log.Println("Error creating database: ", err.Error())
		log.Println("Database already exists, skipping creation.")
	}

	return cfg, nil
}

func CheckPostgresPort(host string) bool {
	port := "5432"
	timeout := 2 * time.Second // Adjust timeout as needed

	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false // Port is not open
	}
	defer conn.Close()

	return true // Port is open
}
