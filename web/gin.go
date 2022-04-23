// @program:     enet
// @file:        web.go
// @author:      edte
// @create:      2022-04-10 17:22
// @description:
package main

import (
	"encoding/json"
	"fmt"
	"github.com/edte/enet/http"
)

func main() {
	data := `{"Method":"get","URL":"/","Version":"","Header":null,"Body":null,"ContentLength":0,"TransferEncoding":null,"Close":false,"Host":"/","Form":null,"PostForm":null,"RemoteAddr":"127.0.0.1:1234"}`

	r := &http.Request{}
	if err := json.Unmarshal([]byte(data), r); err != nil {
		panic(err)
	}

	fmt.Println(r)

}
