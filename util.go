package main

import (
	"os"
	"strconv"
)

func getPort() int64 {
	str := os.Getenv("PORT")

	if str != "" {
		port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 0)
		if err != nil {
			panic(err)
		}

		return port
	} else {
		return 5000
	}
}
