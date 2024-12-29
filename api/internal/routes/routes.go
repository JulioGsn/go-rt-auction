package routes

import (
	"net/http"

	"github.com/juliogsn/job-not-found/internal/handlers/jobs"
)

func Router() {
	http.HandleFunc("/jobs", jobs.List)
}
