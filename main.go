package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

type Bubble struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func (Bubble) TableName() string {
	return "bubble"
}

func initMysql() (err error) {
	dsn := "root:12345678@tcp(47.101.207.93)/bubble?charset=utf8mb4"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	// 创建数据库链接
	if err := initMysql(); err != nil {
		panic(err.Error())
	}
	defer DB.Close()
	// 模型绑定
	DB.AutoMigrate(&Bubble{})
	// 建立
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static") // js,css文件的加载
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := router.Group("v1")

	// 添加
	v1Group.POST("/todo/", func(context *gin.Context) {
		// 前端发送数据
		// 1. 从前端获取传入数据
		var todo Bubble
		_ = context.BindJSON(&todo)
		// 2. 存入数据库
		if err := DB.Create(&todo).Error; err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"code":   "10000",
				"status": "success",
				"data":   todo,
			})
		}
	})
	// 删除
	v1Group.DELETE("/todo/:id", func(context *gin.Context) {

	})
	// 修改
	v1Group.PUT("/todo/:id", func(context *gin.Context) {

	})
	// 查看所有
	v1Group.GET("/", func(context *gin.Context) {

	})
	// 查看某个
	v1Group.GET("/:id", func(context *gin.Context) {

	})

	err := router.Run(":8081")
	if err != nil {
		panic(err.Error())
	}
}
