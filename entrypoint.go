package main

import (
	_ "./ga-beacon"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("GABEACON_PORT")
	fmt.Println(port)
	http.ListenAndServe(strings.Join([]string{":", port},""), nil)
}
