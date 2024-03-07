package main

import "encoding/json"

func main() {
	var object struct{}
	json.Marshal(&object)
}
