package main

import (
	"fmt"
	"online-bingo/backend/routes"
	"time"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	var hoge int

	fmt.Println(hoge)

	fmt.Println("Hello!")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin","content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
    // ルーターの設定
    // URLへのアクセスに対して静的ページを返す
    router.StaticFS("/", http.Dir("./public"))

	api_v1 := router.Group("/api/v1")
    {
		api_v1.POST("/getBingoCard", routes.GetBingoCard)
		api_v1.POST("/validBingoCard", routes.ValidBingoCard)
    }

	router.Run("localhost:8000")
}
