package comments

import "github.com/Omelman/task-management/internal/repo"

func DeleteComment(id int) error {
	err := repo.Get().Comments().Delete(id)
	if err != nil {
		return err
	}
	return nil
}
