package controller

import (
	"Register-Login-Project/common"
	"Register-Login-Project/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileReceive(c *gin.Context) {
	//router := gin.Default()
	//router.MaxMultipartMemory = 8 << 20
	DB := common.GetDB()
	file, err := c.FormFile("file")
	//log.Println(file.Filename)
	filename := file.Filename
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  "上传失败",
		})
	}
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  "保存文件失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  file,
		"type": "上传",
	})
	//给前端返回数据
	newUpload := model.Upload{
		Name:    file.Filename,
		Address: "/" + filename,
	}
	DB.Create(&newUpload)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", file.Filename))
	//c.File("./" + file.Filename)
}
