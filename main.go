package main

import (
	"fmt"
	"log"

	"github.com/snowflakedb/gosnowflake"
)

func main() {
	// Create a simple example that uses the gosnowflake package
	// This is just to demonstrate the dependency, not to actually connect to Snowflake

	// Create a DSN (Data Source Name) for Snowflake
	// This is just an example and won't actually connect
	cfg := gosnowflake.Config{
		Account:   "example",
		User:      "user",
		Password:  "password",
		Database:  "db",
		Schema:    "schema",
		Warehouse: "warehouse",
	}

	dsn, err := gosnowflake.DSN(&cfg)
	if err != nil {
		log.Fatalf("Failed to create DSN: %v", err)
	}

	fmt.Printf("Created Snowflake DSN: %s\n", dsn)
	fmt.Println("This is a sample application that references the gosnowflake package v1.14.0")
	fmt.Println("This version is not vulnerable to CVE-2023-34231")
}
