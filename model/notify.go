package model

import "time"

type NotifyType string

const (
	NotifyTest NotifyType = "test"
	NotifyEvent NotifyType = "event"
	NotifyError NotifyType = "error"
)

func (nt NotifyType) Valid() bool {
	if nt == NotifyTest || nt == NotifyEvent || nt == NotifyError {
		return true
	}
	return false
}

type Notify struct {
	Title string `json:"title"`
	Content string `type:varchar(1024)" json:"content"`
	ExtraData string `gorm:"type:varchar(1024)" json:"extra_data"`
	System string `json:"system"`
	StackInfo string `gorm:"type:varchar(1024)" json:"stack_info"`
	NType NotifyType `json:"n_type"`

	ID string `gorm:"unique_index" json:"id"`
	FromUrl string `json:"from_url"`
	FromIP string `json:"from_ip"`
	Region string `json:"region"`
	CreatedAt time.Time `json:"created_at"`

	ProjectID string `gorm:"index" json:"project_id"`
}