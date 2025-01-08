package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/rose_shop")
	if err != nil {
		return fmt.Errorf("Error opening database: %v", err)
	}
	log.Println("Database connection opened")

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Error connecting to database: %v", err)
	}
	log.Println("Connected to database")
	return nil
}

func getRosePrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, ngrok-skip-browser-warning, Accept")
	date := r.URL.Query().Get("date")
	if date == "" {
		http.Error(w, "Please provide date", http.StatusBadRequest)
		return
	}
	log.Printf("Received request for date: %s", date)

	var price float64
	err := db.QueryRow("SELECT price FROM rose_prices WHERE date = ?", date).Scan(&price)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Price not found", http.StatusNotFound)
			log.Printf("Price not found for date: %s", date)
		} else {
			http.Error(w, "Error getting price", http.StatusInternalServerError)
			log.Printf("Error getting price for date: %s, error: %v", date, err)
		}
		return
	}
	log.Printf("Price for date %s is %f", date, price)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"success": true,
		"price":   price,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/rose-price", getRosePrice).Methods("GET")

	// 設置 CORS 中介軟體
	headers := handlers.AllowedHeaders([]string{"X-Requested-With",
		"Content-Type",
		"Authorization",
		"ngrok-skip-browser-warning",
		"Accept"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// HTTP server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handlers.CORS(headers, methods, origins)(r)))

}
