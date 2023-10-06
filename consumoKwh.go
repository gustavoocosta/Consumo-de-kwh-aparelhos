package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Struct para representar um aparelho e seu consumo de energia por hora
type Appliance struct {
    Nome    string  `json:"nome"`
    Consumo float64 `json:"consumo_por_hora"`
}

// Mapa de exemplo para armazenar os aparelhos e seus consumos
var appliances = map[string]Appliance{
    "geladeira": {
        Nome:    "Geladeira",
        Consumo: 150.0, // Consumo em watts por hora
    },
    "tv": {
        Nome:    "TV",
        Consumo: 50.0,
    },
}

// Manipulador da rota principal
func handleConsumoEnergia(w http.ResponseWriter, r *http.Request) {
    // Calcular o consumo total de energia para todos os aparelhos
    totalConsumo := 0.0
    for _, appliance := range appliances {
        totalConsumo += appliance.Consumo
    }

    // Retornar o resultado como JSON
    result := map[string]float64{
        "total_consumo": totalConsumo,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

func main() {
    http.HandleFunc("/consumo-energia", handleConsumoEnergia)

    port := 8080
    fmt.Printf("Servidor em execução na porta %d...\n", port)
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
