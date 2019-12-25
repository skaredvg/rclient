// rclient project main.go
package main

import (
	_ "bufio"
	"fmt"
	_ "io"
	"io/ioutil"
	"net/http"
)

func main() {
	c := http.Client{}

	for {
		a := ""
		fmt.Print("Введите запрос: ")
		n, err := fmt.Scan(&a)
		fmt.Println(a)
		if err != nil {
			fmt.Println("err=", err.Error())
		}
		if n == 0 || a == "exit" {
			break
		}
		resp, _ := c.Get("http://127.0.0.1:8080/" + a)
		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("err: ", err.Error())
		}
		fmt.Println("resp=", string(b))
	}

}
