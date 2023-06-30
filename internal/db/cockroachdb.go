package db

import (
	"CRUDWithCockroach/internal/config"
	"time"

	"github.com/jpillora/backoff"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// CockroachDB represent gorm DB
	CockroachDB *gorm.DB
	// StopTickerCh signal for closing channel
	StopTickerCh chan bool
)

// InitializeCockroachConn :nodoc:
func InitializeCockroachConn() {
	conn, err := openCockroachConn(config.DatabaseDSN())
	if err != nil {
		log.WithField("databaseDSN", config.DatabaseDSN()).Fatal("failed to connect cockroach database: ", err)
	}

	CockroachDB = conn
	StopTickerCh = make(chan bool)

	go checkConnection(time.NewTicker(config.CockroachPingInterval()))
}

func openCockroachConn(dsn string) (*gorm.DB, error) {
	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	conn.SetMaxIdleConns(config.CockroachMaxIdleConns())
	conn.SetConnMaxLifetime(config.CockroachConnMaxLifetime())
	conn.SetMaxOpenConns(config.CockroachMaxOpenConns())

	return db, nil
}

func checkConnection(ticker *time.Ticker) {
	for {
		select {
		case <-StopTickerCh:
			ticker.Stop()
			return
		case <-ticker.C:
			if _, err := CockroachDB.DB(); err != nil {
				reconnectCockroachConn()
			}
		}
	}
}

func reconnectCockroachConn() {
	b := backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    100 * time.Millisecond,
		Max:    1 * time.Second,
	}

	for b.Attempt() < config.RetryAttempts {
		conn, err := openCockroachConn(config.DatabaseDSN())
		if err != nil {
			log.WithField("databaseDSN", config.DatabaseDSN()).Error("failed to connect cockroach database: ", err)
		}

		if conn != nil {
			CockroachDB = conn
			break
		}
		time.Sleep(b.Duration())
	}

	if b.Attempt() >= config.RetryAttempts {
		log.Fatal("maximum retry to connect database")
	}
	b.Reset()
}
