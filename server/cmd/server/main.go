package main

import (
	"flag"
	"github.com/BusyPeopleAPZ/APZ-3/server/db"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var portNumber = flag.Int("p", 8080, "HTTP port number")

func DatabaseConnection() (*pgxpool.Pool, error) {
	conn := &db.Connection{
		DbName: "balancersdb",
		Schema: "public",
		User: "postgres",
		Password: "root",
		Host: "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	flag.Parse()

	if server, fault := ComposeApiServer(HttpPortNumber(*portNumber)); fault == nil {
		go func() {
			log.Println("Starting VMS server...")
			fault := server.Start()
			if fault == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", fault)
			}
		}()

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if fault := server.Stop(); fault != nil && fault != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", fault)
		}
	} else {
		log.Fatalf("Cannot initialize VMS server: %s", fault)
	}
}
