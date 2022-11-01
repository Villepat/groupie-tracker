package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server"
)

func main() {
	startServer()
	fmt.Println("Now listening to port: 8080")
}
