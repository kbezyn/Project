package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type car struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Model  string `json:"model"`
	Run    int    `json:"run"`
	Owners byte   `json:"owners"`
}

func main() {
	url := "http://localhost:8080/cars"
	// метод запроса
	method := "POST"
	var newCar = []byte(`{"id": "5", "name": "Tank", "model": "300", "run": 1000, "owners": 0}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(newCar))
	// установить заголовок HTTP-запроса Content-Type
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	fmt.Println(string(url))

}
