// @program:     enet
// @file:        client.go
// @author:      edte
// @create:      2022-04-21 21:22
// @description:
package main

import "github.com/edte/enet/http"

func main() {
	_, err := http.Get("/")
	if err != nil {
		panic(err)
	}
}
