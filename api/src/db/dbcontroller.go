package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type BreedCache struct {
	Query        string
	Data         string
	TimeCreation string
}

var (
	dbDriver = "mysql"
	dbUser   = "root"
	dbName   = "db_hostgator_challenge"
)

func dbConn(db_exists bool) (db *sql.DB) {
	var d *sql.DB
	var err error
	if db_exists {
		d, err = sql.Open(dbDriver, dbUser+":"+os.Getenv("MYSQL_ROOT_PASSWORD")+"@tcp(mysql:3306)/"+dbName+"?charset=utf8&parseTime=True&loc=Local&multiStatements=true")
	} else {
		d, err = sql.Open(dbDriver, dbUser+":"+os.Getenv("MYSQL_ROOT_PASSWORD")+"@tcp(mysql:3306)/?charset=utf8&parseTime=True&loc=Local&multiStatements=true")
	}
	if err != nil {
		panic(err.Error())
	}
	return d
}

func Select(query string) (res []BreedCache) {
	db := dbConn(true)
	rs, err := db.Query("SELECT query, data, time_creation FROM tb_breeds_cache WHERE query=?", query)
	if err != nil {
		panic(err.Error())
	}
	data_cached := BreedCache{}
	for rs.Next() {
		var query, data, time_creation string
		err = rs.Scan(&query, &data, &time_creation)
		if err != nil {
			panic(err.Error())
		}
		data_cached.Query = query
		data_cached.Data = data
		data_cached.TimeCreation = time_creation
		res = append(res, data_cached)
	}
	defer db.Close()
	return
}

func Insert(query string, data string) {
	db := dbConn(true)

	insForm, err := db.Prepare("INSERT INTO tb_breeds_cache(query, data) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(query, data)
	log.Println("INSERT: cache: " + query)

	defer db.Close()
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn(true)
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
}

func Migrate() {
	db := dbConn(false)

	sql, err := ioutil.ReadFile("./migrations/up.sql")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		db.Close()
		if r := recover(); r != nil {
			fmt.Println("Error occurred: ", r)
		}
	}()

	_, err = db.Exec(string(sql))
	if err != nil {
		panic(err.Error())
	}

	log.Println("Migration Completed.")
}

func CheckCacheResult(query string) (bool, []BreedCache) {
	data := Select(query)
	return len(data) > 0, data
}
