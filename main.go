package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/dump-data-fields", func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
			println("No file")
			c.Error(err)
			return
		}

		// Create the tempdir and tempfile for pdftk to operate on
		tempdir := os.TempDir()
		tempfile := tempdir + "/temp.pdf"

		err = c.SaveUploadedFile(file, tempfile)

		args := []string{
			tempfile,
			"dump_data_fields",
		}

		out, err := exec.Command("pdftk", args...).Output()

		if err != nil {
			println("pdftk error: " + err.Error())
		}

		output := string(out[:])

		fields := strings.Split(output, "")

		res := []string{"foo", "bar", string(output)}
		c.JSON(200, res)
	})

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
	port := ":9092"
	fmt.Print("Running on port " + port)
	r.Run(port)
}
