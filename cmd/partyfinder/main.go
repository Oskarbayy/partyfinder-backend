package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Oskarbayy/partyfinder-backend/internal/db"
	"github.com/Oskarbayy/partyfinder-backend/internal/router"
	"github.com/Oskarbayy/partyfinder-backend/internal/users"
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

	userStore := db.NewUserStore(sqlDB)
	userSvc := users.NewService(userStore)
	userHandler := users.NewHandler(userSvc)

	r := router.New()
	userHandler.RegisterRoutes(r)

	// add listener here:
	// Minimal HTTP listener
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

	// Test user store \\
	// Test create USER
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	u := users.User{Name: "Navn1", Email: "Email1", Password: "Password1"}
	userStore.Create(ctx, &u)

	//Test find EMAIL
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel1()
	user, err := userStore.FindByEmail(ctx1, "Email1")
	fmt.Printf("\nUser: %+v\n", user)
}
