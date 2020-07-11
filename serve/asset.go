package serve

import (
	"app/common"
	"app/log"
	"app/setting"
	"app/util/file"
	"github.com/gin-gonic/gin"
	"net/http"
)

var uploadParam = []string{"oid", "fileName"}

func UploadFile(c *gin.Context) {
	form, err := common.BindMultipartForm(c, uploadParam)
	log.Debug("", form)
	if err != nil {
		return
	}

	fileHeader := form.File["content"][0]
	filePath := setting.Servers["main"].FilePath + "/" + form.Value["oid"][0] + "/img";
	// check exist and create
	if err := file.Checkdir(filePath); err != nil {
		log.Warn(c, 1500001, err, "something error in write file", "something error in create folder")
	}
	if err := file.SaveFile(fileHeader, filePath, form.Value["fileName"][0]); err != nil {
		log.Warn(c, 1500001, err, "something error in write file")
		return
	}

	c.Status(http.StatusOK)
}

func GetFile(c *gin.Context) {
	c.File(setting.Servers["main"].FilePath + "/" + c.Param("oid") + "/img/" + c.Param("name"))
}