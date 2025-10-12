package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // --> init() --> register the driver (postgres in this case)
	// sql.Register("postgres", &PostgresDriver{})
)

// ORM
// 1 - model
// 2 - query builder to construct SQL queries
// 3 - db client / driver

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	connectionStr := "postgres://user:password@localhost:5430/mydatabase?sslmode=disable"

	db, err := sql.Open("postgres", connectionStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)

	id := insertProduct(db, Product{"apple watch", 399.99, true})
	log.Print(id)

	products := selectAllProducts(db)
	log.Print(products)
}

// Let's create product table

/* Product Table
- ID
- Name
- Price
- Available
- CreatedAt
*/

func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS products (
    	id SERIAL PRIMARY KEY,
    	name VARCHAR(100) NOT NULL,
    	price DECIMAL(6, 2) NOT NULL,
    	available BOOLEAN DEFAULT TRUE,
    	created_at TIMESTAMP DEFAULT NOW()
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO products (name, price, available) VALUES ($1, $2, $3) RETURNING id`

	var id int

	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func selectAllProducts(db *sql.DB) []Product {
	query := `SELECT name, price, available FROM products`
	data := []Product{}

	rows, err := db.Query(query)

	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No rows found: %v", err)
			return data
		}

		log.Fatal(err)
	}

	var name string
	var price float64
	var available bool

	for rows.Next() {
		err := rows.Scan(&name, &price, &available)

		if err != nil {
			log.Fatal(err)
		}

		data = append(data, Product{name, price, available})
	}

	return data
}
