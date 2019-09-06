package main

import (
	"github.com/gin-gonic/gin"
	"log"
	s_file "s-file"
)

func main() {
	sfile := s_file.NewSFile("./files")
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.POST("/upload", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		log.Printf("upload file %v", fileHeader.Filename)
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		defer file.Close()
		content, key, err := sfile.SaveFile(c, fileHeader.Filename, file)
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"msg": content, "key": key, "size": fileHeader.Size})
		return
	})
	engine.POST("/delete/:key", func(c *gin.Context) {
		err := sfile.DeleteFile(c, c.Param("key"))
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"msg": "ok"})
		return
	})
	err := engine.Run(":80")
	log.Println(err)
}
