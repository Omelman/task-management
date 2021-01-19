package projects

import (
	"encoding/json"
	"github.com/Omelman/task-management/api/models"
	"github.com/Omelman/task-management/api/services/projects"
	"log"
	"net/http"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var req models.ProjectRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatal(err)
	}

	_, err := projects.CreateProject(r, req)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}
