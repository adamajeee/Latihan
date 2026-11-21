package main

import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
fmt. Fprintln(w, "Hallo nama saya Adam Al Azhar")
}
http.HandleFunc("/", handlerIndex)

fmt.Println("server started at 43.173.13.18:9005")
http.ListenAndServe(":9005", nil)

}
