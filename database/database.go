package database

import (
	"fmt"
	"log"

	"github.com/ehanz12/ai_habit_tracker/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//variable DB untuk menyimpan koneksi database
var DB *gorm.DB

func ConnectDB() {
	//ambil config database dari config package
	cfg := config.LoadEnv()
	//lakukan koneksi ke database menggunakan cfg
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	//inisialisasi variabel DB dengan koneksi database yang berhasil
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal Terhubung ke Database :", err)
	}
	DB = db
	fmt.Println("✅ Berhasil Terhubung ke Database")
}