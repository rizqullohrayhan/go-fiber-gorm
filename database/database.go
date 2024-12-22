package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rizqullorayhan/go-fiber-gorm/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic(err)
	}

	// dsn := "root:@tcp(127.0.0.1:3306)/"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		port,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Koneksi gagal: %v", err)
	}

	// Membuat database jika belum ada
    sqlDB, _ := db.DB()
    _, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Config("DB_NAME")))
    if err != nil {
        log.Fatalf("Gagal membuat database: %v", err)
    }

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"),
	)
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Koneksi ke database gagal: %v", err)
    }
	log.Println("Koneksi ke database berhasil.")
}