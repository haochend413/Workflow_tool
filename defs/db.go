package defs

import (
	"gorm.io/gorm"
)

// structure for db data storage

type DB_Data struct {
	NoteData      []*Note
	TopicData     []*Topic
	DailyTaskData []*DailyTask
}

// struct for single message
type Note struct {
	// use the unique ID as indicator
	gorm.Model
	Content string
	Topics  []*Topic `gorm:"many2many:note_topics;"`
}

type Topic struct {
	gorm.Model
	Topic string
	Notes []*Note `gorm:"many2many:note_topics;"`
}

type DailyTask struct {
	gorm.Model
	Task    string
	Success bool
}
