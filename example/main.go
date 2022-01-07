package main

import (
	"bytes"
	"fmt"

	"github.com/ArunMurugan78/bencode"
)

func main() {
	data := map[string]interface{}{"name": "rose", "list": []interface{}{"pickachu", "anakin", "NYC", 78}}

	encodedData, _ := bencode.Encode(data)

	fmt.Println(string(encodedData)) // d4:listl8:pickachu6:anakin3:NYCi78ee4:name4:rosee

	decodedData, _ := bencode.Parse(bytes.NewReader(encodedData))

	fmt.Println(decodedData) // map[list:[pickachu anakin NYC 78] name:rose]
}
