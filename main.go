package main

import "gingo/routes"

func main() {
	router := routes.Routers()

	router.Run(":8060")
}
