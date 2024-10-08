package main

import (
	"fmt"
	"log"
	"net/http"

	"api.em.GO/src/config"
	"api.em.GO/src/middlewares"
	"api.em.GO/src/router/rotas"
)

func main() {
    config.Carregar()
    r := rotas.Gerar()

    handler := middlewares.SetupCORS(r)

    fmt.Printf("API RODANDO NA PORTA %d\n", config.Porta)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), handler))
}
