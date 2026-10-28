package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Run(addr string) {
  // No concurrency support yet, just getting the server up-and-running
  mux := http.NewServeMux()

  mux.HandleFunc("GET /ping", handlePing)

  srv := http.Server{
    Addr: addr,
    Handler: mux,
    ReadHeaderTimeout: 500 * time.Millisecond,
    ReadTimeout: 2 * time.Second,
    WriteTimeout: 30 * time.Second,
  }

  log.Printf("listening on %s", addr)

  if err := srv.ListenAndServe(); err != nil {
    log.Fatalf("server error: %v", err)
  }

  log.Print("shuting down...")
}

// TEMPORARY
func handlePing(w http.ResponseWriter, r *http.Request) {
  log.Print("GET /ping")
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  w.WriteHeader(http.StatusOK)

  fmt.Fprintln(w, "pong!")
}
