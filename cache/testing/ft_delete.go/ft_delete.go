package main

import (
	"fmt"

	Hermes "github.com/realTristan/Hermes/cache"
)

func main() {
	var cache *Hermes.Cache = Hermes.InitCache()

	// Initialize the FT cache
	cache.InitFT(-1, -1, map[string]bool{
		"name": true,
	})

	// Test Delete()
	var data = map[string]interface{}{
		"name": "tristan",
		"age":  17,
	}

	// print cache info
	cache.Info()

	// Set data
	cache.Set("user_id1", data)
	cache.Set("user_id1", data)
	cache.Set("user_id2", data)

	// print cache info
	cache.Info()

	// Delete data
	cache.Delete("user_id1")
	cache.Delete("user_id1")

	// Get data
	fmt.Println(cache.Get("user_id1"))
	fmt.Println(cache.Get("user_id2"))

	// Exists
	fmt.Println(cache.Exists("user_id1"))
	fmt.Println(cache.Exists("user_id2"))

	// Length
	fmt.Println(cache.Length())

	// Values
	fmt.Println(cache.Values())

	// Keys
	fmt.Println(cache.Keys())

	// Print the cache info
	cache.Info()
}
