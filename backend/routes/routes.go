package routes

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"time"

	"online-bingo/backend/routes/bingoCard"
	"online-bingo/backend/routes/db"
	g "online-bingo/backend/routes/general"

	"github.com/gin-gonic/gin"
)


func GetBingoCard(ctx *gin.Context) {
	type card struct {
		StudentId string `json:"student_id"`
	}
	// fmt.Println(ctx.Params)
	// param := new(card)
	var param card
	ctx.Bind(&param)
	fmt.Printf("param: %v\n", param)
	if param.StudentId == "" {
		fmt.Println("400:パラメーターが未入力")
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"matrix": "", "message": "パラメーターが未入力"})
		return
	}
	// id := ctx.PostForm("student_id")
	// fmt.Println()
	id_hash := sha256.Sum256([]byte(param.StudentId))
	// showHashArray(id_hash)
	bingo_matrix := bingoCard.GetBingoMatrix(id_hash)
	// fmt.Println(bingo_matrix)

	// 学籍番号のハッシュを文字列化
	id_hash_str := g.ByteArray2HexString(id_hash)
	fmt.Println(id_hash_str)
	// DBへ保存
	firestoreClient := g.InitFirestore()
	_, err := firestoreClient.Collection("cards").Doc(id_hash_str).Set(ctx, map[string]interface{}{
		"student_id":    id_hash,
		"status":        "play",
		"last_modified": time.Now(),
		"bingo_array":   g.Matrix2Array(bingo_matrix),
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": fmt.Sprintf("%S", err)})
		return
	}
	defer firestoreClient.Close()
	ctx.JSON(http.StatusOK, map[string]interface{}{"matrix": bingo_matrix, "student_id_hash": id_hash_str})
}

func ValidBingoCard(ctx *gin.Context) {
	type Params struct {
		StudentIdHash string `json:"student_id_hash"`
		StudentId string `json:"student_id"`
		Array     []int  `json:"array"`
		DisplauName string `json:"display_name"`
	}
	var body Params
	err := ctx.Bind(&body)
	print("body.StudentIdHash:")
	fmt.Println(body.StudentIdHash)
	print("body.Matrix:")
	fmt.Println(body.Array)
	if err != nil {
		fmt.Printf("ValidBingoCard bind/%s", err)
		ctx.JSON(http.StatusOK, map[string]interface{}{"result": "Rejected", "message": "パラメーター受信失敗"})
		return
	}
	requestedmatrix, err := g.Array2Matrix55(body.Array, 5)
	if err != nil {
		fmt.Printf("ValidBingoCard Array2Matrix55/%s", err)
		ctx.JSON(http.StatusOK, map[string]interface{}{"result": "Rejected", "message": err})
		return
	}
	//DBに保存済みのビンゴ行列を取得
	//TODO:DB取得エラーなのか、ビンゴカードが見つからなかったのかを判別する
	verifyMatrix, err := db.GetBingoMatrix(body.StudentIdHash)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{"result": "Rejected", "message": "ビンゴカード未発行かも"})
		return
	}
	//DB保存済みのビンゴ行列とユーザーから届いた行列が一致しているかチェック（不正改ざんチェック）
	for i := 0; i < len(requestedmatrix); i++ {
		for j := 0; j < len(requestedmatrix[0]); j++ {
			if verifyMatrix[i][j] != requestedmatrix[i][j] {
				ctx.JSON(http.StatusOK, map[string]interface{}{"result": "Rejected", "message": "送信されたビンゴ行列が不正です"})
				return
			}
		}
	}
	// ビンゴしているかチェック
	bingoCount, err := bingoCard.GetBingoCount(requestedmatrix)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"result": "Failed", "message": "Dbエラー/出た数字の取得に失敗"})
		return
	}
	if bingoCount <= 0 {
		ctx.JSON(http.StatusOK, map[string]interface{}{"result": "Rejected", "message": "ビンゴしていません"})
		return
	}
	// ビンゴ認定
	err = db.AddWinUser(body.StudentId,body.StudentIdHash,body.DisplauName)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"result": "Failed", "message": "Dbエラー/ビンゴ者リストへの登録失敗"})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"result": "Passed", "message": ""})
	
}

// func AddWinNumber(ctx *gin.Context){
// 	type Params struct {
// 		Number int `json:"number"`
// 		Token string `json:"token"`
// 	}
// 	var body Params
// 	err := ctx.Bind(&body)
// 	fmt.Printf("param: %v\n", body)
// 	if err != nil{
// 		fmt.Printf("AddWinNumber bind/%s", err)
// 		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"number": body.Number, "message": "パラメーター受信失敗"})
// 		return
// 	}
// 	fmt.Printf("AddWinNumber/number:%d\n",body.Number)
// 	err = db.AddWinNumber(body.Number)

// 	if err != nil{
// 		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"number": body.Number, "message": "DB書き込み失敗"})
// 		return 
// 	}

// 	ctx.JSON(http.StatusOK, map[string]interface{}{"number": body.Number, "message": ""})
// }
