//This package should import db and design functions, generate output for gui to render

//Lazy sync & writeback structure: controller should be able to write data from db to the local copy in main; (all sorts after managed by DB operations);
//And it should also be able to send gui data back to db to re-fresh the states;

package dbcontroller

import (
	"github.com/haochend413/mantis/defs"
)

// refresh database data; Run at quit or before specific functions
func (m *DBManager) RefreshDaily(data []*defs.DailyTask) error {
	if err := m.DataBases.DailyDB.SyncDailyTaskData(data); err != nil {
		return err
	}
	return nil
}

func (m *DBManager) RefreshNoteTopic(data *defs.DB_Data) error {
	m.DataBases.NoteDB.SyncTopicData(data.TopicData)
	return m.DataBases.NoteDB.SyncNoteData(data.NoteData)
}

// refresh database data; Run at quit or before specific functions
func (m *DBManager) RefreshAll(data *defs.DB_Data) error {
	if err := m.DataBases.DailyDB.SyncDailyTaskData(data.DailyTaskData); err != nil {
		return err
	}
	if err := m.DataBases.NoteDB.SyncTopicData(data.TopicData); err != nil {
		return err
	}
	return m.DataBases.NoteDB.SyncNoteData(data.NoteData)
}

// fetch database data, run at the Appinit
func (m *DBManager) FetchAll() *defs.DB_Data {
	var (
		history   []defs.Note
		topics    []defs.Topic
		dailytask []defs.DailyTask
	)

	// Preload relationships
	if err := m.DataBases.NoteDB.Db.Preload("Topics").Find(&history).Error; err != nil {
		return &defs.DB_Data{NoteData: []*defs.Note{}}
	}
	if err := m.DataBases.NoteDB.Db.Preload("Notes").Find(&topics).Error; err != nil {
		return &defs.DB_Data{TopicData: []*defs.Topic{}}
	}
	if err := m.DataBases.DailyDB.Db.Find(&dailytask).Error; err != nil {
		return &defs.DB_Data{DailyTaskData: []*defs.DailyTask{}}
	}

	notePtrs := make([]*defs.Note, 0, len(history))
	dailytaskPtrs := make([]*defs.DailyTask, 0, len(dailytask))
	topicPtrs := make([]*defs.Topic, 0, len(topics))
	for i := range history {
		notePtrs = append(notePtrs, &history[i])
	}
	for i := range dailytask {
		dailytaskPtrs = append(dailytaskPtrs, &dailytask[i])
	}
	for i := range topics {
		topicPtrs = append(topicPtrs, &topics[i])
	}

	return &defs.DB_Data{NoteData: notePtrs, DailyTaskData: dailytaskPtrs, TopicData: topicPtrs}
}
