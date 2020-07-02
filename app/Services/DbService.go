package Services

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	driver string
	user string
	password string
	database string
}

var database *sql.DB

func GetDB() DB {
	return NewDB("mysql", "root", "", "blog")
}

func NewDB(driver, user, password, database string) DB {
	db := DB{driver, user, password, database}
	return db
}

func (c *DB) insert(table string, values string) {
	database, err := sql.Open(c.driver, c.user + ":" + c.password + "@/" + c.database)
	if err != nil {
		log.Println(err)
	}
	database.Query("insert into " + " " + table + " values (" + values + ")")
	defer database.Close()
}

func (c *DB) update(table string, values string) {
	database, err := sql.Open(c.driver, c.user + ":" + c.password + "@/" + c.database)
	if err != nil {
		log.Println(err)
	}
	database.Query("update " + " " + table + " set " + values)
	defer database.Close()
}

func (c *DB) delete(table string, values string) {
	database, err := sql.Open(c.driver, c.user + ":" + c.password + "@/" + c.database)
	if err != nil {
		log.Println(err)
	}
	database.Query("delete from " + " " + table + " where " + values)
	defer database.Close()
}

func (c *DB) query(queryString string) *sql.Rows {
	database, err := sql.Open(c.driver, c.user + ":" + c.password + "@/" + c.database)
	if err != nil {
		log.Println(err)
	}

	rows, err := database.Query(queryString)
	if err != nil {
		log.Println(err)
	}
	defer database.Close()

	return rows
}
