package main

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/get_products", GetProducts)
	http.ListenAndServe(":8080", server)
}

type Product struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func GetProducts(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "not alowed", http.StatusMethodNotAllowed)
	}
	products := make([]Product, 0)
	var prod Product
	for i := 0; i < 10; i++ {
		prod = Product{
			Name:  faker.Word(),
			Price: float32(rand.Int()%100000) / 100,
		}
		products = append(products, prod)
	}
	data, _ := json.Marshal(products)
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(resp, string(data))
	//resp.Header().Add("Test", "Hello from keys man")

}
