package columns

import "github.com/Omelman/task-management/internal/repo"

func DeleteProject(id int) error {
	err := repo.Get().Columns().Delete(id)
	if err != nil {
		return err
	}
	return nil
}
