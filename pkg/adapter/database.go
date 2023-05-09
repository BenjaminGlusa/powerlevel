package adapter

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DatabaseAdapter interface {
	AddMeasurement(watt uint16)
	KwhToday() float32
	KwhThisMonth() float32
	KwhThisYear() float32
	KwhTotal() float32

	CreateTableIfNotExits()
	Close()
}

type MySqlAdapter struct {
	db *sql.DB
}

func (m *MySqlAdapter) CreateTableIfNotExits() {
	result, err := m.db.Query("CREATE TABLE if not exists `power` (`time` timestamp NOT NULL UNIQUE DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, `watt` int NOT NULL);")
	if err != nil {
		panic(err)
	}

	defer result.Close()
}

func (m *MySqlAdapter) AddMeasurement(watt uint16) {
	query := fmt.Sprintf("INSERT INTO `power` (`time`, `watt`) VALUES (now(), %d);", watt)
	result, err := m.db.Query(query)
	if err != nil {
		panic(err)
	}

	defer result.Close()
}

func (m *MySqlAdapter) KwhToday() float32 {
	var kWh float32
	err := m.db.QueryRow("SELECT ROUND(AVG(watt) / 1000,3) as kWh FROM power WHERE DATE(time) = CURDATE();").Scan(&kWh)
	if err != nil {
		panic(err)
	}

	return kWh
}


func (m *MySqlAdapter) KwhThisMonth() float32 {
	var kWh float32
	err := m.db.QueryRow("SELECT ROUND(AVG(watt) / 1000,3) as kWh FROM power WHERE UNIX_TIMESTAMP(time) BETWEEN UNIX_TIMESTAMP(LAST_DAY(CURDATE()) + INTERVAL 1 DAY - INTERVAL 1 MONTH) AND UNIX_TIMESTAMP(LAST_DAY(CURDATE()) + INTERVAL 1 DAY);").Scan(&kWh)
	if err != nil {
		panic(err)
	}

	return kWh
}

func (m *MySqlAdapter) KwhThisYear() float32 {
	var kWh float32
	err := m.db.QueryRow("SELECT ROUND(AVG(watt) / 1000,3) as kWh FROM power WHERE DATE(time) BETWEEN MAKEDATE(year(now()),1) AND MAKEDATE(year(now()),1) + interval 1 year - interval 1 day;").Scan(&kWh)
	if err != nil {
		panic(err)
	}

	return kWh
}

func (m *MySqlAdapter) KwhTotal() float32 {
	var kWh float32
	err := m.db.QueryRow("SELECT ROUND(AVG(watt) / 1000,3) as kWh FROM power;").Scan(&kWh)
	if err != nil {
		panic(err)
	}

	return kWh
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
