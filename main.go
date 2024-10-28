package main

import (
	"fmt"
	"net/http"
	v1 "newsletter-app/pkg/api/v1"
	"newsletter-app/pkg/infrastructure/adapters/mongodb"
	"os"

	_ "newsletter-app/docs"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Newsletter API
// @version 1.0
// @description API to manage newsletters.
// @contact.name Gustavo Hernandez
// @contact.email gustavohdzmz@gmail.com
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	mongoUrl := os.Getenv("mongoUrl")
	fmt.Println("MongoDB:", mongoUrl)
	err = mongodb.Connect(mongoUrl)
	if err != nil {
		fmt.Println("Error al conectar a MongoDB:", err)
		return
	} else {
		fmt.Println("Conexión a la base de datos exitosa")
	}
	defer mongodb.Disconnect()

	router := v1.SetupRouter()

	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	router.Use(handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders))

	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	port := 8080
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Swagger UI disponible en http://localhost%s/docs/index.html\n", address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		fmt.Println(err)
	}
}
