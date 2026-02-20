package models

import "time"

type TaskLog struct {
	ID        int       `json:"id"`
	TaskID    int       `json:"task_id"`
	Action    string    `json:"action"`
	OldValue  string    `json:"old_value"`
	NewValue  string    `json:"new_value"`
	ChangedBy string    `json:"changed_by"`
	CreatedAt time.Time `json:"created_at"`
}
