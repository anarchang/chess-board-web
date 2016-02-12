package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB gorm.DB

func main() {
	var err error
	DB, err = gorm.Open("postgres", "user=postgres dbname=chess_game sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	DB.DB()

	// Disable table name's pluralization
	DB.SingularTable(true)
	DB.CreateTable(&Piece{})
	DB.Set("gorm:taple_operations", "ENGINE=InnoDB").CreateTable(&Piece{})

	piece := Piece{Type: WhiteQueen, Top: 100, Left: 100}
	DB.NewRecord(piece)
	DB.Create(&piece)

	var pieces Pieces
	DB.Find(&pieces)
	log.Println(pieces)
	go h.run()

	router := NewRouter()
	router.HandleFunc("/ws", WebsocketHandler)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("build/"))))

	log.Println("Listening at port 8080")
	http.ListenAndServe(":8080", router)
}
