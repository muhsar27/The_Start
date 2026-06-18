package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	port := ":5000"

	fmt.Printf("Server is running and listening at http://localhost%s\n", port)
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/help", handleHelp)
	mux.HandleFunc("/ascii", handleAscii)
	mux.HandleFunc("/hello/", handleHelloParameterized)

	log.Fatal(http.ListenAndServe(port, mux))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func handleHelp(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("ASK QUESTIONS\n"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
}

func handleAscii(w http.ResponseWriter, r *http.Request) {

}
func handleHelloParameterized(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userList := params["user"]

	userName := "User"
	if len(userList) > 0 {
		userName = userList[0]
	}
	var output bytes.Buffer
	output.WriteString("Hello, ")
	output.WriteString(userName)
	output.WriteString("!\n")

	_, err := w.Write(output.Bytes())
	if err != nil {
		slog.Error("error writing response body", "err", err)
		return
	}

}
