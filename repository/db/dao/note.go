package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"schedule/repository/db/model"
)

type NoteDao struct {
	*gorm.DB
}

func NewNoteDao(cxt context.Context) *NoteDao {
	return &NoteDao{NewDBClient(cxt)}
}

func NewNoteDaoByDB(db *gorm.DB) *NoteDao {
	return &NoteDao{db}
}

// AddNote 添加新的便签
func (dao *NoteDao) AddNote(content string) {
	var noteCount int64
	dao.DB.Model(&model.Note{}).Unscoped().Count(&noteCount)
	// SELECT count(*) FROM "Note"
	note := model.Note{Rank: uint(noteCount + 1), Content: content}
	dao.DB.Create(&note)
	fmt.Println("Note added:", note)
}

// 获取所有便签
func (dao *NoteDao) GetNotes() []model.Note {
	var notes []model.Note
	// dao.DB.Order("rank").Find(&notes) // 根据rank排序返回数据 逆序desc
	dao.DB.Order("done, rank").Find(&notes) // 根据完成度 rank排序返回数据
	fmt.Println("GetNotes: ", notes)
	return notes
}

// 删除便签
func (dao *NoteDao) DeleteNote(content string) {
	dao.DB.Where("content = ?", content).Delete(&model.Note{})
	// dao.DB.Delete(&model.Note{}, content)
	fmt.Println("Note deleted with ID:", content)
}

// 标记便签为完成
func (dao *NoteDao) Done(content string) {
	var note model.Note
	// dao.DB.First(&note, content)
	// dao.DB.Where("content = ?", content).First(&note)
	dao.DB.First(&note, "content = ?", content)
	note.Done = !note.Done
	// fmt.Println("Done get note: ", note, content)
	dao.DB.Save(&note)
	// fmt.Println("Note marked as done:", note)
	// dao.DB.Model(&model.Note{}).Where("content = ?", content). Updates(&product).Error
}

func (dao *NoteDao) EditRank(newNote *model.Note) {
	var oddNote model.Note
	dao.DB.Where("content = ?", newNote.Content).First(&oddNote)

	fmt.Println("EditRank from", oddNote, "to", newNote)
	// 变大
	var notes []model.Note
	if newNote.Rank > oddNote.Rank {
		var noteCount int64
		dao.DB.Model(&model.Note{}).Unscoped().Count(&noteCount)
		newNote.Rank = min(newNote.Rank, uint(noteCount))
		dao.DB.Where("rank > ? AND rank <= ?", oddNote.Rank, newNote.Rank).Find(&notes)
		for _, note := range notes {
			note.Rank--
			dao.DB.Save(&note)
		}
	} else {
		newNote.Rank = max(newNote.Rank, 1)
		dao.DB.Where("rank >= ? AND rank < ?", newNote.Rank, oddNote.Rank).Find(&notes)
		for _, note := range notes {
			note.Rank++
			dao.DB.Save(&note)
		}
	}
	dao.DB.Save(&newNote)
}
