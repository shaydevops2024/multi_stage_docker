package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type CalcInput struct {
    Num1 float64 `json:"num1"`
    Num2 float64 `json:"num2"`
}

type CalcOutput struct {
    Result float64 `json:"result"`
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
    // Allow CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        return
    }

    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var input CalcInput
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    output := CalcOutput{Result: input.Num1 + input.Num2}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(output)
}

func main() {
    http.HandleFunc("/calc", calcHandler)

    log.Println("Backend running on :8000")
    err := http.ListenAndServe("0.0.0.0:8000", nil)
    if err != nil {
        log.Fatal(err)
    }
}

