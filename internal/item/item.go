package item

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Name string `json:"name"`
}

func CreateNewItem(itemData []byte, writer http.ResponseWriter) {
	var item Item

	json.Unmarshal(itemData, &item)
	fmt.Println(item)
}
