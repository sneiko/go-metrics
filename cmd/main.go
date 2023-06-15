package main

import "metrics"

func main() {
	metrics.RunServer(metrics.Options{
		Host: "localhost",
		Port: 8080,
	})
}
