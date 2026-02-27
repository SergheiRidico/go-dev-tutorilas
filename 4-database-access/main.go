package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Connection properties
	cfg := mysql.NewConfig()
	cfg.User = "root"       // better os.Getenv("DBUSER")
	cfg.Passwd = "password" // better os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}

	fmt.Println("Connected to mysql database!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	album, err := albumByID(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Single album retrieved: %v\n", album)

	
	ZdobSiZdubAlbum := Album{
		Title:  "Ethnomecanica",
		Artist: "Zdob Şi Zdub",
		Price:  9.99,
	}

	id, err := addAlbum(ZdobSiZdubAlbum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Zdob Şi Zdub album was successsfully added under id: %d\n", id)

	err = deleteAlbumByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album with id: %d was successfully deleted from database\n", id)
}
