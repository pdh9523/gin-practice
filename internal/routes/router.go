package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pdh9523/gin-practice/internal/domain/auth"
	"github.com/pdh9523/gin-practice/internal/domain/post"
	"github.com/pdh9523/gin-practice/internal/domain/user"
	"github.com/pdh9523/gin-practice/internal/infra/cache"
)

func SetupRouter(cacheStore cache.GlobalCacheStore) *gin.Engine {

	r := gin.Default()
	r.Group("/api/v1/")

	user.MountUserRoutes(r)
	post.MountPostRoutes(r)
	auth.MountAuthRoutes(r, cacheStore)

	return r
}
