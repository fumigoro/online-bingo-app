package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"online-bingo/backend/routes"
	"os"
	"time"
)

func main() {

	err := godotenv.Load(fmt.Sprintf("env/%s.env",os.Getenv("GO_ENV")))
	var apiUrl string
	var baseUrl string
	//もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
		apiUrl = "localhost:8000"
		baseUrl = "http://localhost:3000"
	} else {
		apiUrl = os.Getenv("API_HOST")+":"+os.Getenv("API_PORT")
		baseUrl = os.Getenv("BASE_HOST")+":"+os.Getenv("BASE_PORT")
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{baseUrl},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "content-type"},
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
	fmt.Println(apiUrl)
	router.Run(apiUrl)
}
