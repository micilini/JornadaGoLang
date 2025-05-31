package main

import (
    "html/template"
    "log"
    "net/http"
    "time"

    "github.com/go-chi/chi/v5"
)

type PaginaData struct {
    Titulo   string
    Mensagem string
    Hora     string
}

func main() {
    r := chi.NewRouter()

    // Carrega templates no início (de preferência, de forma global)
    tmpl := template.Must(template.ParseGlob("templates/*.html"))

    // Rota dinâmica com rota amigável: /ola/{nome}
    r.Get("/ola/{nome}", func(w http.ResponseWriter, r *http.Request) {
        nome := chi.URLParam(r, "nome")
        dados := PaginaData{
            Titulo:   "Serviço Dinâmico em Go",
            Mensagem: "Bem-vindo, " + nome + "!",
            Hora:     time.Now().Format("02/01/2006 15:04:05"),
        }
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        if err := tmpl.ExecuteTemplate(w, "pagina.html", dados); err != nil {
            http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        }
    })

    // Rota raiz responde com hora atual em JSON
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        hora := time.Now().Format(time.RFC3339)
        w.Write([]byte(`{"server_time":"` + hora + `"}`))
    })

    log.Println("Servidor dinâmico rodando em http://localhost:8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}