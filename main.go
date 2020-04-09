package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Field struct {
	FieldType          string
	FieldName          string
	FieldFlags         float64
	FieldJustification string
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/dump-data-fields", func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
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
			c.Error(err)
			return
		}

		output := string(out[:])

		pdftkFields := strings.Split(output, "---\n")

		// Remove the empty string from the start of the fields list
		_, pdftkFields = pdftkFields[0], pdftkFields[1:]

		var fields []Field
		for _, pdftkField := range pdftkFields[0:] {
			fieldData := strings.Split(pdftkField, "\n")

			// Remove the empty string at the end of the field data list
			fieldData = fieldData[:len(fieldData)-1]

			fieldDataMap := make(map[string]string)
			for _, data := range fieldData {
				parts := strings.Split(data, ": ")
				keyPart := parts[0]
				valuePart := parts[1]
				fieldDataMap[keyPart] = valuePart
			}

			fieldFlags, err := strconv.ParseFloat(fieldDataMap["FieldFlags"], 64)

			if err != nil {
				c.Error(err)
				return
			}

			field := Field{
				FieldType:          fieldDataMap["FieldType"],
				FieldName:          fieldDataMap["FieldName"],
				FieldFlags:         fieldFlags,
				FieldJustification: fieldDataMap["FieldJustification"],
			}

			fields = append(fields, field)
		}

		c.JSON(200, fields)
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

	r.Run(":80")
}
