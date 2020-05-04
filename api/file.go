package api

import (
	"authen/eth"
	"authen/model"
	"authen/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

// GetFileList 获取已上传的文件列表
func GetFileList(c *gin.Context) {
	uid, ok := c.Get("user_id")
	if !ok {
		c.JSON(400, gin.H{
			"code": -1,
			"msg":  "not login",
		})
		return
	}
	userID := uid.(uint)
	var filelist []model.File
	err := model.DB.Model(&model.File{}).Where("user_id = ?", userID).Find(&filelist).Error
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, gin.H{
			"code":     0,
			"msg":      "success",
			"filelist": filelist,
		})
	}
}

// FileHashUpload 上传文件hash
func FileHashUpload(c *gin.Context) {
	var service service.FileHashUploadSvc
	err := c.BindJSON(&service)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "failed",
		})
	} else {
		uid, ok := c.Get("user_id")
		if !ok {
			c.JSON(400, gin.H{
				"code": -1,
				"msg":  "not login",
			})
			return
		}
		service.UserID = uid.(uint)
		_, err = service.UploadHash()
		if err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "failed",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
		})
	}
}

// FileUpload 文件上传
func FileUpload() gin.HandlerFunc {
	store := filestore.FileStore{
		Path: "/tmp",
	}
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/file/upload/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
		PreUploadCreateCallback: func(hook tusd.HookEvent) error {
			return nil
		},
	})
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	go func() {
		for {
			event := <-handler.CompleteUploads
			token, ok := event.Upload.MetaData["token"]
			fmt.Println("token is ", token)
			if !ok {
				continue
			}
			uid, err := model.GetUIDFromToken(token)
			if err != nil {
				fmt.Println("wrong token!")
				continue
			}
			var file model.File
			model.DB.Model(&model.File{}).Where(map[string]interface{}{
				"user_id":   uid,
				"file_name": event.Upload.MetaData["filename"],
			}).First(&file)
			file.IsUpload = true
			file.Path = "/tmp/" + event.Upload.ID
			file.UploadID = event.Upload.ID
			file.ManifestHash, err = eth.UploadETH(file.Path + ".info")
			if err != nil {
				fmt.Println("upload swarm error: ", err.Error())
			}
			model.DB.Save(&file)
			fmt.Printf("Upload %s finished, uri is %s\n", event.Upload.ID, event.HTTPRequest.URI)
		}
	}()
	return func(c *gin.Context) {
		http.StripPrefix("/file/upload/", handler).ServeHTTP(c.Writer, c.Request)
	}
}
