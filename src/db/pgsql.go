package db

import (
	"common"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// PgSQL Connection and operations Handler.
type PgSQL struct {
	Connection *sql.DB
	db         Config
}

//Ping Tests the connection.
func (p *PgSQL) Ping() error {
	return p.Connection.Ping()
}

func (p *PgSQL) connect() error {
	var pgInfo string

	if p.db.Password == "" {
		pgInfo = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
			p.db.Host, p.db.Port, p.db.User, p.db.Database)
	} else {
		pgInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			p.db.Host, p.db.Port, p.db.User, p.db.Password, p.db.Database)
	}

	conn, err := sql.Open("postgres", pgInfo)
	if err != nil {
		return err
	}
	p.Connection = conn
	return p.Ping()
}

// NewPgSQL makes a new instance of PgSQL and connect to postgresql database.
func NewPgSQL(config Config) *PgSQL {
	pg := PgSQL{db: config}
	var err error
	try := 0
	for try < 10 {
		time.Sleep(time.Duration(try) * time.Second) //Increasing time between tries. Starting with 0.
		err = pg.connect()
		if err == nil {
			break
		}
		try++
		fmt.Printf("Error connecting to the database. Err: %s \n", err.Error())
		fmt.Println("Retrying database connection... Try " + strconv.Itoa(try) + "/10")
	}
	if err != nil {
		fmt.Printf("Error connecting to the database. Err: %s \n", err.Error())
		os.Exit(2)
	}
	return &pg
}

// Execute executes the query received with the given parameters.
func (p *PgSQL) Execute(query string, args ...interface{}) (sql.Result, error) {
	var result sql.Result

	stmtIns, err := p.Connection.Prepare(query)
	if err != nil {
		return result, err
	}
	defer stmtIns.Close()

	result, err = stmtIns.Exec(args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

//QueryRow gets the next Query Row
func (m *PgSQL) QueryRow(query string, args ...interface{}) *sql.Row {
	return m.Connection.QueryRow(query, args...)
}

//MapScan get first row in JSON format
func (m *PgSQL) MapScan(sqlString string) (map[string]interface{}, error) {
	rows, err := m.Connection.Query(sqlString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	entry := make(map[string]interface{})
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry = make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
	}
	return entry, nil
}

//SliceMap fetch all lines of given select
func (m *PgSQL) SliceMap(sqlString string) ([]map[string]interface{}, error) {
	rows, err := m.Connection.Query(sqlString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}

// Close is responsible for closing database connection
func (p *PgSQL) Close() {
	err := p.Connection.Close()
	common.PanicIfNotNil(err)
}
