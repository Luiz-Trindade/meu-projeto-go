package main

import (
	"fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router 		:= mux.NewRouter()
    serverPort 	:= 8080

    router.HandleFunc("/", func(w http.ResponseWriter, router *http.Request) {
        w.Write([]byte("Servidor Go rodando! ✅ \n"))
    })
    
    fmt.Printf("Servidor rodando na porta :%d\n", serverPort)
    
	addr := fmt.Sprintf(":%d", serverPort)
    http.ListenAndServe(addr, router)
}
