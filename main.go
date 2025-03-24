package main

import (
	"github.com/joho/godotenv"
	"github.com/pdh9523/gin-practice/internal/db"
	"github.com/pdh9523/gin-practice/internal/model"
	"github.com/pdh9523/gin-practice/internal/routes"
)

func main() {
	// 환경 변수 로드
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}

	// DB 마운트
	db.Init()
	db.DB.AutoMigrate(&model.Post{})

	// 라우터 마운트
	r := routes.SetupRouter()
	r.Run(":8080")
}
