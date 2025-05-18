package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", func(w http.ResponseWriter, router *http.Request) {
        w.Write([]byte("Servidor Go rodando! ✅"))
    })
    http.ListenAndServe(":8080", router)
}
