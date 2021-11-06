package main

import (
	"app/graph"
	"app/databasepq"
	"app/graph/generated"
	"log"
	"net/http"
	"os"
	"database/sql"
	
	_"github.com/lib/pq"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := sql.Open("postgres", databasepq.DB_CONFIG)

	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: databasepq.DB{ Conn: db },
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
