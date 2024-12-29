package routes

import (
	"net/http"

	"github.com/juliogsn/rt-auction/internal/handlers/jobs"
)

func Router() {
	http.HandleFunc("GET /jobs", jobs.List)
}
