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

	pieces := []Piece{
		Piece{Type: BlackRook, Top: 82, Left: 79},
		Piece{Type: BlackKnight, Top: 82, Left: 154},
		Piece{Type: BlackBishop, Top: 82, Left: 234},
		Piece{Type: BlackQueen, Top: 82, Left: 308},
		Piece{Type: BlackKing, Top: 82, Left: 383},
		Piece{Type: BlackBishop, Top: 82, Left: 458},
		Piece{Type: BlackKnight, Top: 82, Left: 533},
		Piece{Type: BlackRook, Top: 82, Left: 608},
		Piece{Type: BlackPawn, Top: 160, Left: 79},
		Piece{Type: BlackPawn, Top: 160, Left: 154},
		Piece{Type: BlackPawn, Top: 160, Left: 234},
		Piece{Type: BlackPawn, Top: 160, Left: 308},
		Piece{Type: BlackPawn, Top: 160, Left: 383},
		Piece{Type: BlackPawn, Top: 160, Left: 458},
		Piece{Type: BlackPawn, Top: 160, Left: 533},
		Piece{Type: BlackPawn, Top: 160, Left: 608},
		Piece{Type: WhiteRook, Top: 611, Left: 79},
		Piece{Type: WhiteKnight, Top: 611, Left: 154},
		Piece{Type: WhiteBishop, Top: 611, Left: 234},
		Piece{Type: WhiteQueen, Top: 611, Left: 308},
		Piece{Type: WhiteKing, Top: 611, Left: 383},
		Piece{Type: WhiteBishop, Top: 611, Left: 458},
		Piece{Type: WhiteKnight, Top: 611, Left: 533},
		Piece{Type: WhiteRook, Top: 611, Left: 608},
		Piece{Type: WhitePawn, Top: 533, Left: 79},
		Piece{Type: WhitePawn, Top: 533, Left: 154},
		Piece{Type: WhitePawn, Top: 533, Left: 234},
		Piece{Type: WhitePawn, Top: 533, Left: 308},
		Piece{Type: WhitePawn, Top: 533, Left: 383},
		Piece{Type: WhitePawn, Top: 533, Left: 458},
		Piece{Type: WhitePawn, Top: 533, Left: 533},
		Piece{Type: WhitePawn, Top: 533, Left: 608},
	}
	for _, piece := range pieces {
		db.NewRecord(piece)
		db.Create(&piece)
	}
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
