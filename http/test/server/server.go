// @program:     enet
// @file:        server.go
// @author:      edte
// @create:      2022-04-21 21:22
// @description:
package main

import (
	"fmt"
	"github.com/edte/enet/http"
)

func main() {
	http.HandleFunc("/", func(r *http.Request, w http.ResponseWriter) {
		fmt.Println(r.URL)
	})
	if err := http.Listen(":1234", nil); err != nil {
		panic(err)
	}
}
