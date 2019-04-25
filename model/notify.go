package model

import "time"

type NotifyType string

const (
	NotifyTest NotifyType = "test"
	NotifyEvent NotifyType = "event"
	NotifyError NotifyType = "error"
)

type Notify struct {
	Title string `json:"title"`
	Content string `json:"content"`
	ExtraData string `json:"extra_data"`
	System string `json:"system"`
	StackInfo string `json:"stack_info"`
	NType NotifyType `json:"n_type"`

	ID string `gorm:"unique_index" json:"id"`
	FromUrl string `json:"from_url"`
	FromIP string `json:"from_ip"`
	Region string `json:"region"`
	CreatedAt time.Time `json:"created_at"`

	ProjectID string `json:"project_id"`
}