package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/fill-pdf", func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
			c.Error(err)
			return
		}

		jsonString := c.PostForm("json")
		dynamic := make(map[string]interface{})
		json.Unmarshal([]byte(jsonString), &dynamic)

		tempdir := os.TempDir()
		tempfile := tempdir + "/temp.pdf"
		tempfilefilled := tempdir + "/temp.pdf"

		err = c.SaveUploadedFile(file, tempfile)

		if err != nil {
			c.Error(err)
			return
		}

		err = Fill(dynamic, tempfile, tempfilefilled, true)

		if err != nil {
			c.Error(err)
			return
		}

		c.File(tempfilefilled)
	})

	r.Run(":80")
}
