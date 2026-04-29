package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// MongoDB connection and operations handler
type MongoDB struct {
	db *mgo.Database
	config Config
}

// NewMongoDB creates a new instance of MongoDB and connect to specified database
func NewMongoDB(config *Config) *MongoDB {
	m := MongoDB{config: *config}
	var err error
	tries := 0
	for tries < 10 {
		time.Sleep(time.Duration(tries) * time.Second) // increasing some time between tries
		err = m.connect()
		if err == nil {
			break
		}
		log.Println(err)
		tries++
		log.Println( "retrying database connection... try " + strconv.Itoa(tries) + "/10")
	}

	if err != nil {
		log.Println("error connecting to database err: " + err.Error())
		os.Exit(2)
	}

	return &m
}

// TestMongoDBConnection tries to connect to specified mongodb database
func TestMongoDBConnection(config *Config) error {
	m := MongoDB{config: *config}
	var err error

	err = m.connect()
	defer m.Close()

	if err != nil {
		log.Println(err)
	}

	return err
}

// connect open a connection to mongodb server
func (m *MongoDB) connect() error {
	var url string
	if m.config.GetUser() != "" {
		url = m.config.GetUser()
	}
	if m.config.GetPassword() != "" {
		url += ":" + m.config.GetPassword()
	}
	if m.config.GetUser() != "" || m.config.GetPassword() != "" {
		url += "@"
	}
	url += m.config.GetHost() + ":" + strconv.Itoa(m.config.GetPort())

	uri := "mongodb://" + url + "/" + m.config.GetDatabase()

	if rs, ok := m.config.(ReplicaSetConfig); ok {
		var params []string
		if v := rs.GetReplicaSet(); v != "" {
			params = append(params, "replicaSet="+v)
		}
		if v := rs.GetReadPreference(); v != "" {
			params = append(params, "readPreference="+v)
		}
		if v := rs.GetAuthSource(); v != "" {
			params = append(params, "authSource="+v)
		}
		if len(params) > 0 {
			uri += "?" + strings.Join(params, "&")
		}
	}

	session, err := mgo.Dial(uri)
	if err != nil {
		return err
	}

	m.db = session.DB(m.config.GetDatabase())
	m.db.Session = session
	return nil
}

// C exposes the collection property with its specific methods
func (m *MongoDB) C(name string) *mgo.Collection {
	return m.db.C(name)
}

// GetCollection exposes the collection property with its specific methods
func (m *MongoDB) GetCollection(name string) *mgo.Collection {
	return m.db.C(name)
}

// Copy copies the current database object
func (m *MongoDB) Copy() *MongoDB {
	db := mgo.Database{Session: m.db.Session.Copy(), Name: m.db.Name }
	return &MongoDB{config: m.config, db: &db}
}

// Close closes the database connection
func (m *MongoDB) Close() {
	m.db.Session.Close()
}
