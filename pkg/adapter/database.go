package adapter

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DatabaseAdapter interface {
	CreateTableIfNotExits()
	Close()
}

type MySqlAdapter struct {
	db *sql.DB
}

func (m *MySqlAdapter) CreateTableIfNotExits() {
	result, err := m.db.Query("CREATE TABLE if not exists foo (time VARCHAR(5), power INT);")
	if err != nil {
		panic(err)
	}

	defer result.Close()
}

func (m *MySqlAdapter) Close() {
	defer m.db.Close()
}

func NewMySqlAdapter(user string, password string, host string) *MySqlAdapter {
	connection := fmt.Sprintf("%s:%s@tcp(%s:3306)/powerlevel", user, password, host)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySqlAdapter{
		db: db,
	}
}
