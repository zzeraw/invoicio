package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"
)

type healthResponse struct {
    Status string `json:"status"`
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        _ = json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
    })

    srv := &http.Server{
        Addr:              ":" + port,
        Handler:           mux,
        ReadHeaderTimeout: 5 * time.Second,
    }

    log.Printf("listening on :%s", port)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("server error: %v", err)
    }
}
