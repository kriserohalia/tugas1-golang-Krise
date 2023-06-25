package konversiCelcius

import (
	"fmt"
	"log"
	"os"
	"tugas1-golang-Krise/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (a Connect) PostFahreinhet(hasil float64) {
	var angka models.Angka
	angka.Hasil = hasil
	a.ConnectDatabase().Create(angka)
}

func (a Connect) PostReamur(hasil float64) {
	var angka models.Angka
	angka.Hasil = hasil
	a.ConnectDatabase().Create(angka)
}

func (a Connect) PostKelvin(hasil float64) {
	var angka models.Angka
	angka.Hasil = hasil
	a.ConnectDatabase().Create(angka)
}

func (a Connect) PostRankin(hasil float64) {
	var angka models.Angka
	angka.Hasil = hasil
	a.ConnectDatabase().Create(angka)
}

type Connect struct {
	dbName  string
	dbPort  string
	dbPass  string
	dbUsser string
}

func (kr Connect) ConnectDatabase() *gorm.DB {
	erEnv := godotenv.Load()

	if erEnv != nil {
		log.Fatalln("Error to load env")
	}
	kr.dbName = os.Getenv("DB")
	kr.dbPort = os.Getenv("PORT")
	kr.dbPass = os.Getenv("PASS")
	kr.dbUsser = os.Getenv("USER")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", kr.dbUsser, kr.dbPass, kr.dbName, kr.dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error to Connection to Database")
	}
	db.AutoMigrate(&models.Angka{})
	return db

}
