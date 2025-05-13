package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Oskarbayy/partyfinder-backend/internal/products"
	"github.com/Oskarbayy/partyfinder-backend/internal/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	DSN := os.Getenv("POSTGRES_DSN") // Data Source Name
	fmt.Println(DSN)

	// Setup DB
	sqlDB, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatal(err)
	}

	iProductRepository := products.NewProductRepository(sqlDB)
	productService := products.NewProductService(iProductRepository)
	productHandler := products.NewProductHandler(*productService)

	print(productHandler)

	r := router.New()

	// POST   /users          → createUser
	r.HandleFunc("/addProduct", productHandler.AddProduct).
		Methods(http.MethodPost)

	// add listener here:
	// Minimal HTTP listener
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
