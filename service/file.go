package service

import (
	"authen/model"
	"time"
)

// FileHashUploadSvc 上传文件hash
type FileHashUploadSvc struct {
	UserID    uint   `json:"-"`
	FileName  string `json:"filename"`
	Hash      string `json:"hash"`
	Timestamp int64  `json:"timestamp"`
}

// UploadHash 上传文件hash
func (service *FileHashUploadSvc) UploadHash() (model.File, error) {
	var file model.File
	count := 0
	model.DB.Model(&model.File{}).Where("user_id = ? and file_name = ?", service.UserID, service.FileName).Count(&count)
	if count != 0 {
		model.DB.Model(&model.File{}).Where("user_id = ? and file_name = ?", service.UserID, service.FileName).First(&file)
		file.Hash = service.Hash
		file.TimeStamp = time.Unix(service.Timestamp, 0)
		err := model.DB.Save(&file).Error
		return file, err
	}
	file.UserID = service.UserID
	file.FileName = service.FileName
	file.Hash = service.Hash
	file.TimeStamp = time.Unix(service.Timestamp, 0)
	err := model.DB.Create(&file).Error
	return file, err
}
