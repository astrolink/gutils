package db

import (
	"fmt"
	"log"

	_ "github.com/astrolink/gorm/dialects/mysql" //This is to get the mysql driver
	"github.com/astrolink/gutils/cache"
)

// NewMySQL makes a new instance of Database and connect to a MySQL database.
func NewMySQL(config Config) (*Database, error) {

	connectionLine := "%s:%s@tcp(%s:%d)/%s"
	connectionLine = fmt.Sprintf(connectionLine,
		config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabase())
	mysql := Database{
		ConnectionLine: connectionLine,
		Driver: "mysql",
	}
	var err error
	err = mysql.Connect()
	if err != nil {
		err = fmt.Errorf("error connecting to mysql - host: %s, port: %d, user: %s, database: %s | error: %v",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetDatabase(), err)
		log.Println(err)
	}
	return &mysql, err
}

// NewCachedMySQL makes a new instance of Database and connect to a MySQL database and Redis.
func NewCachedMySQL(config Config, cacheConfig cache.Config) (*Database, error) {
	connectionLine := "%s:%s@tcp(%s:%d)/%s"
	connectionLine = fmt.Sprintf(connectionLine,
		config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabase())
	mysql := Database{
		ConnectionLine: connectionLine,
		Driver: "mysql",
		CacheConfig: cacheConfig,
	}
	var err error
	err = mysql.Connect()
	if err != nil {
		err = fmt.Errorf("error connecting to mysql - host: %s, port: %d, user: %s, database: %s | error: %v",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetDatabase(), err)
		log.Println(err)
	}
	return &mysql, err
}

// NewMySQLReturningError Cria uma nova instância de Database e conecta a um banco de dados MySQL, retornando um erro se ocorrer.
func NewMySQLReturningError(config Config) (*Database, error) {
	connectionLine := "%s:%s@tcp(%s:%d)/%s"
	connectionLine = fmt.Sprintf(connectionLine,
		config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabase())
	mysql := Database{
		ConnectionLine: connectionLine,
		Driver:         "mysql",
	}
	var err error
	err = mysql.Connect()
	if err != nil {
		err = fmt.Errorf("error connecting to mysql - host: %s, port: %d, user: %s, database: %s | error: %v",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetDatabase(), err)
	}
	return &mysql, err
}

// NewCachedMySQLReturningError Cria uma nova instância de Database e conecta a um banco de dados MySQL com cache, retornando um erro se ocorrer.
func NewCachedMySQLReturningError(config Config, cacheConfig cache.Config) (*Database, error) {
	connectionLine := "%s:%s@tcp(%s:%d)/%s"
	connectionLine = fmt.Sprintf(connectionLine,
		config.GetUser(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabase())
	mysql := Database{
		ConnectionLine: connectionLine,
		Driver:         "mysql",
		CacheConfig:    cacheConfig,
	}
	var err error
	err = mysql.Connect()
	if err != nil {
		err = fmt.Errorf("error connecting to mysql - host: %s, port: %d, user: %s, database: %s | error: %v",
			config.GetHost(), config.GetPort(), config.GetUser(), config.GetDatabase(), err)
		log.Println(err)
	}
	return &mysql, err
}
