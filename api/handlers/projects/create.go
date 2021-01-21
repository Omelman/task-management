package projects

import (
	"github.com/Omelman/task-management/api/models"
	"github.com/Omelman/task-management/api/services/projects"
	"log"
	"net/http"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	req := models.ProjectRequest{
		ID:          1,
		ProjectName: "test_name",
		Description: "description",
	}

	_, err := projects.CreateProject(req)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}
