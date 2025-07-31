package main

import (
	"fmt"
	"log"
	"time"

	"github.com/snowflakedb/gosnowflake"
	clientv3 "go.etcd.io/etcd/client/v3"
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
	fmt.Println("This is a sample application that references the gosnowflake package v1.6.18")
	fmt.Println("This version is vulnerable to CVE-2023-34231")

	// Create a simple example that uses the etcd client package
	// This is just to demonstrate the dependency, not to actually connect to etcd
	etcdConfig := clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}

	// We're not actually creating a client, just demonstrating the dependency
	fmt.Println("\nThis application also references the etcd client package v3.5.9")
	fmt.Println("This version is vulnerable to GO-2024-2528 (GHSA-j86v-2vjr-fg8f)")
	fmt.Printf("Etcd config: %+v\n", etcdConfig)
}
