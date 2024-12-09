package pg

import (
	"context"
	"core/internal/config"
	"log"
	"math"
	"math/big"
	"net"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func NumericToFloat64(numeric pgtype.Numeric) float64 {
	// Check for NaN
	if numeric.NaN {
		log.Println("numeric is NaN, returning 0")
		return math.NaN()
	}

	// Convert Int to *big.Float
	bigFloat := new(big.Float).SetInt(numeric.Int)

	// Apply the base-10 exponent
	scaleFactor := new(big.Float).SetFloat64(math.Pow10(int(numeric.Exp)))

	// Scale the value
	bigFloat.Mul(bigFloat, scaleFactor)

	// Convert to float64 (may lose precision for very large numbers)
	floatValue, _ := bigFloat.Float64()

	return floatValue
}

func Float64ToNumeric(value float64) pgtype.Numeric {
	var numeric pgtype.Numeric
	numeric.Scan(value)
	return numeric
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

func CreateDb(ctx context.Context, conn *pgx.Conn) (err error) {
	cfg, err := config.ReadDatabaseConfig()
	if err != nil {
		return
	}

	log.Println("Creating database " + cfg.Database + "...")
	_, err = conn.Exec(context.Background(), "CREATE DATABASE "+cfg.Database)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Println("Database already exists, skipping creation.")
			return nil
		}

		return err
	}

	return nil
}
