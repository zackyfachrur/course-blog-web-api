package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	dbuser := "baxaric:021Baxaric"
	dbname := "baxaric_course"
	dsn := dbuser + "@tcp(127.0.0.1:3306)/" + dbname

	// Buka koneksi database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Gagal koneksi ke database:", err)
		return nil, err
	}

	// Tes koneksi database
	err = db.Ping()
	if err != nil {
		log.Println("Gagal ping ke database:", err)
		return nil, err
	}

	fmt.Println("Koneksi ke database berhasil!")
	return db, nil
}
