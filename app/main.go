package main

import (
	"flag"
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/trust-net/go-sdb-node/api/bootnodes"
	"github.com/trust-net/go-sdb-node/api"
)

func main() {
	port := flag.String("port", "8888", "port to listen on")
	addr := flag.String("addr", "127.0.0.1", "address to listen on")
	flag.Parse()
	log.Printf("Starting app listening on %s:%s ...", *addr, *port)
	h := bootnodes.NewBootNodesHandler(*addr)
	r := mux.NewRouter()
	r.Path("/ping").Methods("GET").Handler(api.NewHandler(HealthCheck))
	r.Path("/bootnodes/start").Methods("POST").Handler(api.NewHandler(h.PostBootnodesStart))
	r.Path("/bootnodes/stop").Methods("POST").Handler(api.NewHandler(h.PostBootnodesStop))
	r.Path("/bootnodes").Methods("GET").Handler(api.NewHandler(h.GetBootnodes))
    r.Use(requestLogger)
    srv := &http.Server{
        Handler:      r,
        Addr:         *addr + ":" + *port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}

func requestLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.Method + " " + r.RequestURI + " call from " + r.RemoteAddr)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}

func HealthCheck(r *http.Request) (api.ApiResponse, api.Error) {
    return "Hello World!", nil
}
