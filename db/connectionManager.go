package db

import (
	"database/sql"
)

type ConnectionManager struct {
	db *Database
}

//NewMySqlConnectionManager inicia uma conexão com o banco de dados mysql instanciando um *ConnectionManager
func NewMySqlConnectionManager(config Config) *ConnectionManager {
	mysql := NewMySQL(config)
	return &ConnectionManager{db: mysql}
}

//NewPgSqlConnectionManager inicia uma conexão com o banco de dados postgres instanciando um *ConnectionManager
func NewPgSqlConnectionManager(config Config) *ConnectionManager {
	pgsql := NewPgSQL(config)
	return &ConnectionManager{db: pgsql}
}

// HandleTransaction gerencia o estado da conexão efetivando
// o commit ou o rollback de uma transação de acordo com o erro recebido
// 		tx *sql.Tx  é a transação que está sendo trabalhada pela conexão
//		err *error passar ponteiro do erro para ser analisado
func (c *ConnectionManager) HandleTransaction(tx *sql.Tx, err *error) error {
	defer c.Close()

	if *err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

// StartTransaction inicia uma transação com o banco de dados
func (c *ConnectionManager) StartTransaction() (*sql.Tx, error) {
	return c.db.StartTransaction()
}

// GetConnection retorna uma conexão pré-estabelecida com o banco de dados
func (c *ConnectionManager) GetConnection() *sql.DB {
	return c.db.Conn
}

// Close encerra a conexão com o banco de dados
func (c *ConnectionManager) Close() {
	c.db.Close()
}
