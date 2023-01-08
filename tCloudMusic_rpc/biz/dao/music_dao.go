package dao

import (
	"tCloudMusic_rpc/biz/model"

	"gorm.io/gorm"
)

type MusicDao struct {
	db *gorm.DB
}

// NewMusicDao: create music dao
func NewMusicDao(db *gorm.DB) *MusicDao {
	return &MusicDao{db: db}
}

// Insert: insert a music info
func (m *MusicDao) Insert(music *model.Music) (*model.Music, error) {
	result := m.db.Create(music)
	return music, result.Error
}

// FindById: where id = ? to find first
func (m *MusicDao) FindById(id uint) (*model.Music, error) {
	music := &model.Music{}
	result := m.db.Find(music, id)
	return music, result.Error
}

// FindByName: where name = ? to find many
func (m *MusicDao) FindByName(name string) (*model.Music, error) {

	music := &model.Music{}

	return music, nil
}
