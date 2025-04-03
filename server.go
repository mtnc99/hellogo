package main

import (
    "fmt"
    "net/http"
)

func Apresentacao(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Trabalho Final de Administração de Containers com RedHat OpenShift")
}

func main() {
    http.HandleFunc("/", Apresentacao)
    http.ListenAndServe(":8080", nil)
}
