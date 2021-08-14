package main

import (
	"fmt"
	"online-bingo/backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	var hoge int

	fmt.Println(hoge)

	
	fmt.Println("Hello!")
	router := gin.Default()

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
