package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"PieceIndex",
		"GET",
		"/piece",
		PieceIndex,
	},
	Route{
		"PieceShow",
		"GET",
		"/piece/{pieceId}",
		PieceShow,
	},
	Route{
		"PieceUpdate",
		"PUT",
		"/piece/{pieceId}",
		PieceUpdate,
	},
	Route{
		"PieceGridIndex",
		"GET",
		"/gridpiece",
		PieceGridIndex,
	},
	Route{
		"GridPieceShow",
		"GET",
		"/gridpiece/{pieceId}",
		GridPieceShow,
	},
	Route{
		"GridPieceUpdate",
		"PUT",
		"/gridpiece/{pieceId}",
		GridPieceUpdate,
	},
}
