package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type PortfolioData struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Skills      []string  `json:"skills"`
	Projects    []Project `json:"projects"`
}

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/")

	switch path {
	case "portfolio":
		handlePortfolio(w, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func handlePortfolio(w http.ResponseWriter, _ *http.Request) {
	data := PortfolioData{
		Title:       "Full Stack Developer",
		Description: "Passionate developer creating innovative solutions with modern technologies.",
		Skills:      []string{"Go", "JavaScript", "React", "Node.js", "Python", "Docker"},
		Projects: []Project{
			{
				Name:        "Portfolio Website",
				Description: "Modern portfolio built with Go backend and TypeScript",
			},
			{
				Name:        "E-commerce Platform",
				Description: "Full-stack e-commerce solution with real-time features",
			},
		},
	}

	json.NewEncoder(w).Encode(data)
}
