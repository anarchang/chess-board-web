package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func PieceIndex(w http.ResponseWriter, r *http.Request) {

	var pieces Pieces
	DB.Find(&pieces)

	json.NewEncoder(w).Encode(pieces)
}

func PieceGridIndex(w http.ResponseWriter, r *http.Request) {
	var gridPieces GridPieces
	DB.Find(&gridPieces)

	json.NewEncoder(w).Encode(gridPieces)
}

func PieceShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pieceId := vars["pieceId"]

	var piece Piece
	DB.Find(&piece, pieceId)

	json.NewEncoder(w).Encode(piece)
}

func GridPieceShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pieceId := vars["pieceId"]

	var piece GridPiece
	DB.Find(&piece, pieceId)

	json.NewEncoder(w).Encode(piece)
}

func PieceUpdate(w http.ResponseWriter, r *http.Request) {
	var updated_piece Piece
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		fmt.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(body, &updated_piece); err != nil {
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity

		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(updated_piece)

	vars := mux.Vars(r)
	pieceId := vars["pieceId"]

	var old_piece Piece
	t := DB.First(&old_piece, pieceId).Updates(updated_piece)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		fmt.Println(err)
	}

	// let the other clients know which piece was updated
	h.broadcast <- []byte(pieceId)
}

func GridPieceUpdate(w http.ResponseWriter, r *http.Request) {
	var updated_piece GridPiece
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		fmt.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(body, &updated_piece); err != nil {
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity

		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(updated_piece)

	vars := mux.Vars(r)
	pieceId := vars["pieceId"]

	var old_piece GridPiece
	t := DB.First(&old_piece, pieceId).Updates(updated_piece)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		fmt.Println(err)
	}

	// let the other clients know which piece was updated
	h.broadcast <- []byte(pieceId)
}
