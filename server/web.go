package server

import (
	"fmt"
	"net/http"
)

/*
func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}
*/

func StartServer(webIP string, webPort string) {
	//http.HandleFunc("/", handleHome)
	http.Handle("/", http.FileServer(http.Dir("web/")))
	fmt.Printf("Starting server at %s:%s\n", webIP, webPort)
	http.ListenAndServe(fmt.Sprintf("%s:%s", webIP, webPort), nil)
}
