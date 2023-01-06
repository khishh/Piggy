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
	plaidimpl "github.com/khishh/personal-finance-app/pkg/plaid"
	"github.com/khishh/personal-finance-app/pkg/repository"
	plaidcore "github.com/plaid/plaid-go/v3/plaid"
	"gorm.io/gorm"
)

const defaultPort = "8080"

var (
	// psql vars
	POSTGRES_HOST             = ""
	POSTGRES_PORT             = ""
	POSTGRES_PASS             = ""
	POSTGRES_USER             = ""
	POSTGRES_SSLMODE          = ""
	POSTGRES_DB_NAME          = ""
	db               *gorm.DB = nil

	// plaid vars
	PLAID_CLIENT_ID                          = ""
	PLAID_SECRET                             = ""
	PLAID_ENV                                = ""
	PLAID_PRODUCTS                           = ""
	PLAID_COUNTRY_CODES                      = ""
	PLAID_REDIRECT_URI                       = ""
	APP_PORT                                 = ""
	client              *plaidcore.APIClient = nil
)

var environments = map[string]plaidcore.Environment{
	"sandbox":     plaidcore.Sandbox,
	"development": plaidcore.Development,
	"production":  plaidcore.Production,
}

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	POSTGRES_PASS = os.Getenv("POSTGRES_PASS")
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_SSLMODE = os.Getenv("POSTGRES_SSLMODE")
	POSTGRES_DB_NAME = os.Getenv("POSTGRES_DB_NAME")

	PLAID_CLIENT_ID = os.Getenv("PLAID_CLIENT_ID")
	PLAID_SECRET = os.Getenv("PLAID_SECRET")

	if PLAID_CLIENT_ID == "" || PLAID_SECRET == "" {
		log.Fatal("Error: PLAID_SECRET or PLAID_CLIENT_ID is not set. Did you copy .env.example to .env and fill it out?")
	}

	PLAID_ENV = os.Getenv("PLAID_ENV")
	PLAID_PRODUCTS = os.Getenv("PLAID_PRODUCTS")
	PLAID_COUNTRY_CODES = os.Getenv("PLAID_COUNTRY_CODES")
	PLAID_REDIRECT_URI = os.Getenv("PLAID_REDIRECT_URI")
	APP_PORT = os.Getenv("APP_PORT")

	// set defaults
	if PLAID_PRODUCTS == "" {
		PLAID_PRODUCTS = "transactions"
	}
	if PLAID_COUNTRY_CODES == "" {
		PLAID_COUNTRY_CODES = "US"
	}
	if PLAID_ENV == "" {
		PLAID_ENV = "sandbox"
	}
	if APP_PORT == "" {
		APP_PORT = "8000"
	}
	if PLAID_CLIENT_ID == "" {
		log.Fatal("PLAID_CLIENT_ID is not set. Make sure to fill out the .env file")
	}
	if PLAID_SECRET == "" {
		log.Fatal("PLAID_SECRET is not set. Make sure to fill out the .env file")
	}

	fmt.Printf("%s %s\n", PLAID_CLIENT_ID, PLAID_SECRET)

	// create Plaid client
	configuration := plaidcore.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", PLAID_CLIENT_ID)
	configuration.AddDefaultHeader("PLAID-SECRET", PLAID_SECRET)
	configuration.UseEnvironment(environments[PLAID_ENV])
	client = plaidcore.NewAPIClient(configuration)
	fmt.Printf("%+v\n", client)

}

func main() {

	config := &database.Config{
		Host:     POSTGRES_HOST,
		Port:     POSTGRES_PORT,
		Password: POSTGRES_PASS,
		User:     POSTGRES_USER,
		SSLMode:  POSTGRES_SSLMODE,
		DBName:   POSTGRES_DB_NAME,
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

	router := gin.Default()

	router.Use(gin.Logger())

	router.Use(corsMiddleWare())
	router.GET("/", playgroundHandler(db))
	router.POST("/api", graphqlHandler(db))
	router.POST("/api/create_link_token", plaidBaseHandler(client, plaidimpl.CreateLinkToken))
	router.POST("/api/create_access_token", plaidBaseHandler(client, plaidimpl.CreateAccessToken))
	router.POST("/api/transactions/sync", plaidBaseHandler(client, plaidimpl.GetTransactionsSync))
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
		ItemRepository: repository.NewItemService(db),
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

func plaidBaseHandler(client *plaidcore.APIClient, routerHandler func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("PlaidClient", client)
		c.Set("CountryCodes", PLAID_COUNTRY_CODES)
		c.Set("RedirectUri", PLAID_REDIRECT_URI)
		c.Set("PlaidProducts", PLAID_PRODUCTS)
		routerHandler(c)
	}
}
