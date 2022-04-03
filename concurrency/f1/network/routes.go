package network

import (
	"net/http"
)

func routes() /*http.Handler */ {
	//mux := pat.New()

	fs := http.FileServer(http.Dir("./web/html"))
	http.Handle("/", fs)
	//http.Handle("/", http.StripPrefix("/html/", fs))
	http.Handle("/ws", http.HandlerFunc(WsEndpoint))

	//return mux
}
