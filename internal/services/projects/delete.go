package projects

import (
	"github.com/Omelman/task-management/internal/repo"
)

func DeleteProject(id int) error {
	err := repo.Get().Projects().Delete(id)
	if err != nil {
		return err
	}
	return nil
}
