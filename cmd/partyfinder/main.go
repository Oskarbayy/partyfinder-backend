package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Oskarbayy/partyfinder-backend/config"
	"github.com/Oskarbayy/partyfinder-backend/internal/products"
	"github.com/Oskarbayy/partyfinder-backend/internal/router"
	_ "github.com/lib/pq"
)

func main() {

	sqlDB := config.GetDatabaseFromEnv()
	iProductRepository := products.NewProductRepository(sqlDB)
	productService := products.NewProductService(iProductRepository)
	productHandler := products.NewProductHandler(*productService)

	r := router.New()

	// POST
	r.HandleFunc("/addProduct", productHandler.AddProduct).
		Methods(http.MethodPost)

	// Minimal HTTP listener
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
