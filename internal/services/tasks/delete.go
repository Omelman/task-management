package tasks

import "github.com/Omelman/task-management/internal/repo"

func DeleteTask(id int) error {
	err := repo.Get().Tasks().Delete(id)
	if err != nil {
		return err
	}
	return nil
}
