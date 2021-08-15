package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"log"
	g "online-bingo/backend/routes/general"
	"strconv"
	"time"
)

func GetWinNumbers() (map[string]bool, error) {

	ctx := context.Background()
	firestoreClient := g.InitFirestore()
	dsnap, err := firestoreClient.Collection("games").Doc("active").Get(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred(GetWinNumbers): %s", err)
		return nil, err
	}
	defer firestoreClient.Close()

	data := dsnap.Data()
	if data["numbers"] == nil {
		fmt.Println("nill!!")
		return nil, errors.New("開始前です")
	}
	numbers := make(map[string]bool)
	for i := 0; i < 75; i++ {
		numbers[strconv.Itoa(i)] = data["numbers"].(map[string]interface{})[strconv.Itoa(i)].(bool)
	}

	return numbers, nil
}

func GetBingoMatrix(studentId string) ([][]int, error) {
	type Body struct {
		last_modified time.Time
		status        string
		bingo_array   []int
	}

	ctx := context.Background()
	firestoreClient := g.InitFirestore()
	dsnap, err := firestoreClient.Collection("cards").Doc(studentId).Get(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred(GetBingoMatrix): %s", err)
		return nil, err
	}
	defer firestoreClient.Close()

	var data Body
	dsnap.DataTo(&data)
	data2 := dsnap.Data()
	
	if data2["status"].(string) == "bingo"{
		err = errors.New("すでにビンゴ登録が完了しています")
		return nil, err;
	}

	var array []int
	for i := range data2["bingo_array"].([]interface{}) {
		tmp := data2["bingo_array"].([]interface{})[i].(int64)
		array = append(array, int(tmp))
	}

	bingo_matrix, err := g.Array2Matrix55(array, 5)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return bingo_matrix, nil
}

func AddWinUser(studentId string, studentIdHash string, displayName string) error {
	ctx := context.Background()
	firestoreClient := g.InitFirestore()
	_, err := firestoreClient.Collection("games").Doc("active").Collection("winners").Doc(studentIdHash).Set(ctx, map[string]interface{}{
		"studentId":      studentId,
		"studentId_hash": studentIdHash,
		"display_name":   displayName,
		"timestamp":      firestore.ServerTimestamp,
		"prize_number":   -1,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}

	_, err = firestoreClient.Collection("cards").Doc(studentIdHash).Set(ctx, map[string]interface{}{
		"status":        "bingo",
		"last_modified": firestore.ServerTimestamp,
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

// func AddWinNumber(number int) error {
// 	// DBへ保存
// 	ctx := context.Background()
// 	firestoreClient := g.InitFirestore()
// 	ref := firestoreClient.Collection("win_numbers").Doc("play")
// 	err := firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
// 		doc, err := tx.Get(ref) // tx.Get, NOT ref.Get!
// 		if err != nil {
// 			return err
// 		}
// 		pop, err := doc.DataAt("numbers")
// 		if err != nil {
// 			return err
// 		}
// 		var array []int
// 		for i := range pop.([]interface{}) {
// 			tmp := pop.([]interface{})[i].(int64)
// 			array = append(array, int(tmp))
// 		}
// 		array = append(array, number)
// 		return tx.Set(ref, map[string]interface{}{
// 			"numbers": array,
// 		}, firestore.MergeAll)
// 	})
// 	if err != nil {
// 		// Handle any errors appropriately in this section.
// 		log.Printf("An error has occurred: %s", err)
// 		return err
// 	}
// 	return nil

// }
