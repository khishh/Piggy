package main

import (
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khishh/personal-finance-app/graph"
	"github.com/khishh/personal-finance-app/pkg/database"
	"github.com/khishh/personal-finance-app/pkg/repository"
	"gorm.io/gorm"
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

	// router := chi.NewRouter()

	// // Add CORS middleware around every request
	// // See https://github.com/rs/cors for full option listing
	// router.Use(cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000"},
	// 	AllowCredentials: true,
	// 	Debug:            true,
	// }).Handler)

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
	// 	BookRepository: repository.NewBookService(db),
	// 	UserRepository: repository.NewUserService(db),
	// }}))

	// router.Handle("/graphiql", playground.Handler("GraphQL playground", "/api"))
	// router.Handle("/api", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, router))

	router := gin.Default()

	router.Use(gin.Logger())

	router.Use(corsMiddleWare())
	router.GET("/", playgroundHandler(db))
	router.POST("/api", graphqlHandler(db))
	router.Run(fmt.Sprintf(":%s", defaultPort))
}

func corsMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		log.Println(c.Request)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	handler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		BookRepository: repository.NewBookService(db),
		UserRepository: repository.NewUserService(db),
	}}))

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler(db *gorm.DB) gin.HandlerFunc {
	handler := playground.Handler("GraphQL playground", "/api")

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
