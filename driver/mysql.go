package driver

import (
	"database/sql"
	"fmt"
)

// MySQLConfig ...
type MySQLConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Db       string
}

// ConnectToMySQL takes mysql config, forms the connection string and connects to mysql
func ConnectToMySQL(conf MySQLConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
