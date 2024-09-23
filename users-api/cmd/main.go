package main

import "users-api/internal/router"

func main() {
	r := router.Init()

	if err := r.Run(":8080"); err != nil {
		return
	}
}
