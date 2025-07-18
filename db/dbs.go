package db

import (
	"github.com/haochend413/mantis/db/dailydb"
	"github.com/haochend413/mantis/db/notedb"
)

// var DBs *DataBases

type DataBases struct {
	NoteDB  *notedb.NoteDB //Note, Topic, Note-Topic Link
	DailyDB *dailydb.DailyDB
}

func (DBs *DataBases) InitAll() {
	DBs.NoteDB = &notedb.NoteDB{}
	DBs.NoteDB.Db = InitNodeDB()

	DBs.DailyDB = &dailydb.DailyDB{}
	DBs.DailyDB.Db = InitDailyDB()
}

func (DBs *DataBases) CloseAll() {
	_ = DBs.NoteDB.Close()
	_ = DBs.DailyDB.Close()

}

// // gorm.Model definition
// type Model struct {
//   ID        uint           `gorm:"primaryKey"`
//   CreatedAt time.Time
//   UpdatedAt time.Time
//   DeletedAt gorm.DeletedAt `gorm:"index"`
// }
