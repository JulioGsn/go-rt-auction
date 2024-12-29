package jobs

import (
	"fmt"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Olá")
}

