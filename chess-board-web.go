package main

import (
	"flag"
	"fmt"
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
	// Flags
	var initDB = flag.Bool("initDB", false, "Initialize the Database")
	var port = flag.Int("port", 3000, "Port on which to run server")
	flag.Parse()

	var err error
	DB, err = gorm.Open("postgres", "user=webapp password=arequipe dbname=chess_game sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	DB.DB()

	// Disable table name's pluralization
	DB.SingularTable(true)

	if *initDB {
		initializeDatabase(DB)
	}

	var pieces Pieces
	DB.Find(&pieces)
	log.Println(pieces)
	go h.run()

	router := NewRouter()
	router.HandleFunc("/ws", WebsocketHandler)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("build/"))))

	log.Println("Listening at port", *port)
	portStr := fmt.Sprintf(":%d", *port)
	log.Println("Port: ", portStr)
	http.ListenAndServe(portStr, router)
}
