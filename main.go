package main

import (
	"github.com/joho/godotenv"
	postModel "github.com/pdh9523/gin-practice/internal/domain/post/model"
	userModel "github.com/pdh9523/gin-practice/internal/domain/user/model"
	"github.com/pdh9523/gin-practice/internal/infra/cache"
	"github.com/pdh9523/gin-practice/internal/infra/db"
	"github.com/pdh9523/gin-practice/internal/infra/redis"
	"github.com/pdh9523/gin-practice/internal/routes"
	"time"
)

func main() {
	// 환경 변수 로드
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}

	// DB 마운트
	db.Init()
	db.DB.AutoMigrate(&userModel.User{}, &postModel.Post{})

	cacheStore := cache.NewGoCacheStore(10*time.Minute, 30*time.Minute)

	redis.InitRedis()
	// 컨슈머 등록
	redis.StartEmailVerifiedConsumer()

	// 라우터 마운트
	r := routes.SetupRouter(cacheStore)
	r.Run(":8080")
}
