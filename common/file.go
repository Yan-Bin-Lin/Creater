package common

import (
	"app/apperr"
	"app/log"
	"app/setting"
	"app/util/file"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

// write file and parse to html
func WriteFormFile(c *gin.Context, form *multipart.Form, fileName string) {
	fileHeader := form.File["content"][0]
	filePath := setting.Servers["main"].FilePath + "/" + form.Value["oid"][0] + "/" + fileName
	// check exist and create
	if err := file.Checkdir(filePath); err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "something error in write file", "something error in create folder")
	}
	go func() {
		if err := file.SaveMarkdown2Tmpl(fileHeader, filePath, fileName+".html"); err != nil {
			log.Warn(c, apperr.ErrPermissionDenied, err, "something error in write file", "something error in parse markdown")
		}
	}()
	if err := file.SaveFile(fileHeader, filePath, fileName+".md"); err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "something error in write file")
		return
	}
}