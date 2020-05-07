package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// File 文件
type File struct {
	gorm.Model
	UserID       uint   `gorm:"index"`
	FileName     string `gorm:"index;size:128"`
	Path         string `gorm:"size:128"`
	Hash         string `gorm:"size:128"`
	TimeStamp    time.Time
	IsUpload     bool
	UploadID     string `gorm:"size:64"`
	ManifestHash string `gorm:"size:64"`
}
