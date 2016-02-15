package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func initializeDatabase(db gorm.DB) {
	db.CreateTable(&Piece{})
	db.Set("gorm:taple_operations", "ENGINE=InnoDB").CreateTable(&Piece{})

	piece := Piece{Type: WhiteQueen, Top: 100, Left: 100}
	db.NewRecord(piece)
	db.Create(&piece)
}

var DB gorm.DB

func main() {
	var err error
	DB, err = gorm.Open("postgres", "user=webapp password=arequipe dbname=chess_game sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	DB.DB()

	// Disable table name's pluralization
	DB.SingularTable(true)

	initializeDatabase(DB)

	var pieces Pieces
	DB.Find(&pieces)
	log.Println(pieces)
	go h.run()

	router := NewRouter()
	router.HandleFunc("/ws", WebsocketHandler)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("build/"))))

	log.Println("Listening at port 3000")
	http.ListenAndServe(":3000", router)
}
