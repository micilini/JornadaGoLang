package main

import(
	"net/http"
	"time"
	"log"
	"fmt"
)

func mensagemLogin(w http.ResponseWriter, r *http.Request){
	s := time.Now().Format("02/01/2006 03:04:05") // Aqui não usamos o m/d/y h:i:s como em outras linguagens de programação
	fmt.Fprintf(w, "<h1>Você se Logou: %s</h1>", s)
}

func main(){
	http.HandleFunc("/login", mensagemLogin)
	log.Println("Executando Código...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}