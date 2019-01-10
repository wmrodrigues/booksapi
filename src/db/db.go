package db

import "database/sql"

//Database interface defines methods to be a database handler.
type Database interface {
	connect() error
	Ping() error
	Execute(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	MapScan(sqlString string) (map[string]interface{}, error)
	SliceMap(sqlString string) ([]map[string]interface{}, error)
	Close()
}

//Config is a simplified json struct to load database settings
type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
