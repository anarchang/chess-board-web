package main

type PieceType int

const (
	WhiteQueen PieceType = iota
	BlackQueen
	WhiteKing
	BlackKing
	WhiteRook
	BlackRook
	WhiteBishop
	BlackBishop
	WhiteKnight
	BlackKnight
	WhitePawn
	BlackPawn
)

type Piece struct {
	Id   int       `json:"id"`
	Type PieceType `json:"pieceType"`
	Top  int       `json:"top"`
	Left int       `json:"left"`
}

type Pieces []Piece
