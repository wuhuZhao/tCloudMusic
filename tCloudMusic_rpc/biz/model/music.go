package model

type Music struct {
	ID          uint
	MusicName   string
	MusicAuthor string
	MusicPath   string
	CommentId   uint
}

func (Music) TableName() string {
	return "music"
}
