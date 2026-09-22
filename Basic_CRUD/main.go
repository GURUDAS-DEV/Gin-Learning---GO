package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Details struct {
	name   string
	rollNo string
	dept   string
	cgpa   string
}

type DetailsDTO struct {
	Id     string `json:"id" form:"id"`
	Name   string `json:"name" form:"name" binding:"required"`
	RollNo string `json:"rollNo" form:"rollNo" binding:"required"`
	Dept   string `json:"dept" form:"dept" binding:"required"`
	Cgpa   string `json:"cgpa" form:"cgpa" binding:"required"`
}

func main() {
	router := gin.Default()

	m := make(map[string]Details)

	router.GET("/GetAll", func(ctx *gin.Context) {
		if len(m) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Nothing Found",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Nothing Found",
		})
	})

	router.POST("/createPost", func(ctx *gin.Context) {
		var details DetailsDTO

		err := ctx.ShouldBind(&details)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "Require all fields",
			})
			return
		}

		_, ok := m[details.Id]
		if ok == true {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "Already Exists",
			})
			return
		}

		value := Details{
			name:   details.Name,
			rollNo: details.RollNo,
			dept:   details.Dept,
			cgpa:   details.Cgpa,
		}
		m[details.Id] = value

		ctx.JSON(201, gin.H{
			"message": "Successfully created",
		})
	})

	router.GET("/getById/:Id", func(ctx *gin.Context) {
		id := ctx.Param("Id")
		id = strings.TrimSpace(id)

		if len(id) <= 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "id was not Found",
			})
			return
		}

		val, ok := m[id]
		if ok == false {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Entity was not Found",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Found",
			"val":     val,
		})
	})

	router.Run()
}
