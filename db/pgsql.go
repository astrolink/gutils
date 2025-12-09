package db

import (
	"fmt"
	"github.com/astrolink/gutils/cache"
	"log"

	_ "github.com/lib/pq" //Only Drivers
)

// NewPgSQL makes a new instance of PgSQL and connect to PostgresSQL database.
func NewPgSQL(config Config) (*Database, error) {
	var connectionLine string
	if config.GetPassword() == "" {
		connectionLine = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetDatabase())
	} else {
		connectionLine = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetPassword(), config.GetDatabase())
	}
	pg := Database{
		ConnectionLine: connectionLine,
		Driver:         "postgres",
	}
	var err error
	err = pg.Connect()
	if err != nil {
		err = fmt.Errorf("error connecting to postgres - host: %s, port: %d, user: %s, database: %s | error: %v",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetDatabase(), err)
		log.Println(err)
	}
	return &pg, err
}

// NewPgSQL makes a new instance of PgSQL and connect to PostgresSQL database and Redis.
func NewCachedPgSQL(config Config, cacheConfig cache.Config) (*Database, error) {
	var connectionLine string
	if config.GetPassword() == "" {
		connectionLine = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetDatabase())
	} else {
		connectionLine = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetPassword(), config.GetDatabase())
	}
	pg := Database{
		ConnectionLine: connectionLine,
		Driver:         "postgres",
		CacheConfig:    cacheConfig,
	}
	var err error
	err = pg.Connect()
	if err != nil {
		err = fmt.Errorf("error connecting to postgres - host: %s, port: %d, user: %s, database: %s | error: %v",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetDatabase(), err)
		log.Println(err)
	}
	return &pg, err
}
