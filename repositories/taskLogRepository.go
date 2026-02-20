package repositories

import (
	"mini-jira-backend/configs"
	"mini-jira-backend/models"
)

// new log entry for task
func CreateTaskLog(log *models.TaskLog) (models.TaskLog, error) {
	_, err := configs.DB.Exec(`
		INSERT INTO task_logs (task_id, action, old_value, new_value, changed_by)
		VALUES ($1, $2, $3, $4, $5)
	`, log.TaskID, log.Action, log.OldValue, log.NewValue, log.ChangedBy)
	return *log, err
}

func ShowTaskLogs() ([]models.TaskLog, error) {
	rows, err := configs.DB.Query(`
		SELECT id, task_id, action, old_value, new_value, changed_by, created_at
		FROM task_logs
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var logs []models.TaskLog
	for rows.Next() {
		var log models.TaskLog
		err := rows.Scan(
			&log.ID,
			&log.TaskID,
			&log.Action,
			&log.OldValue,
			&log.NewValue,
			&log.ChangedBy,
			&log.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
