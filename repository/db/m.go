package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// 定义 Note 结构体作为数据库模型
type Note struct {
	Rank    uint   `gorm:"type:uint"`
	Content string `gorm:"primaryKey;type:varchar(52)"` // primaryKey 放第一个才能用主键进行删除
	Done    bool   `gorm:"default:false"`
}

type NoteApp struct {
	db *gorm.DB
}

func NewNoteApp() *NoteApp {
	// 初始化数据库
	db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// 自动迁移创建表
	db.AutoMigrate(&Note{})

	return &NoteApp{db: db}
}

// AddNote 添加新的便签
func (a *NoteApp) AddNote(content string) {
	var noteCount int64
	a.db.Model(&Note{}).Unscoped().Count(&noteCount)
	// SELECT count(*) FROM "Note"
	note := Note{Rank: uint(noteCount + 1), Content: content}
	a.db.Create(&note)
	fmt.Println("Note added:", note)
}

// 获取所有便签
func (a *NoteApp) GetNotes() []Note {
	var notes []Note
	a.db.Find(&notes)
	fmt.Println("GetNotes: ", notes)
	return notes
}

// 删除便签
func (a *NoteApp) DeleteNote(content string) {
	// a.db.Where("content = ?", content).Delete(&Note{})
	a.db.Delete(&Note{}, content)
	fmt.Println("Note deleted with ID:", content)
}

// 标记便签为完成
func (a *NoteApp) Done(content string) {
	var note Note
	a.db.First(&note, content)
	note.Done = !note.Done
	a.db.Save(&note)
	fmt.Println("Note marked as done:", note)
}
