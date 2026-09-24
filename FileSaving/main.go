package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "File upload unsuccessfull!",
			})
			return
		}

		log.Println(file.Filename)

		dst := filepath.Join("./files/", filepath.Base(file.Filename))
		ctx.SaveUploadedFile(file, dst)
		ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})

	router.POST("/uploadMultiple", func(ctx *gin.Context) {
		// Multipart form
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			dst := filepath.Join("./files/", filepath.Base(file.Filename))
			ctx.SaveUploadedFile(file, dst)
		}
		ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run()
}
