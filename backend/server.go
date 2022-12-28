package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/khishh/personal-finance-app/graph"
	"github.com/khishh/personal-finance-app/pkg/database"
	"github.com/khishh/personal-finance-app/pkg/repository"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()
	config := &database.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASS"),
		User:     os.Getenv("POSTGRES_USER"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		DBName:   os.Getenv("POSTGRES_DB_NAME"),
	}
	log.Printf("%s %s\n", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	db, err := database.NewConnection(config)
	if err != nil {
		panic(err)
	}
	database.Migrate(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		BookRepository: repository.NewBookService(db),
		UserRepository: repository.NewUserService(db),
	}}))

	router.Handle("/graphiql", playground.Handler("GraphQL playground", "/api"))
	router.Handle("/api", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
