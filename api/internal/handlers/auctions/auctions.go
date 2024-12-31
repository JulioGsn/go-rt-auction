package auctions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Name      string  `json:"name"`
	BasePrice float32 `json:"basePrice"`
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	item := Item{
		Name:      "item 1",
		BasePrice: 25000.0,
	}

	itemJson, err := json.Marshal(item)

	if err != nil {
		fmt.Fprintf(w, "Error %d", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, string(itemJson))
}
