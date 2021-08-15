package general

import (
	"context"
	"errors"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func ByteArray2HexString(array [32]byte) string{
	id_hash_str := ""
	for i := range array {
		id_hash_str += fmt.Sprintf("%02x", array[i])
	}
	return id_hash_str
}

func InitFirestore() firestore.Client {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return *client
}

func Matrix2Array(matrix [][]int) []int{
	var array []int
	if len(matrix) == 0{
		return array
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			array = append(array, matrix[i][j])
		}
	}
	return array
}

func Array2Matrix55(array []int, size int)( [][]int,error){
	initial_num := -1;
	matrix := make([][]int, size)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = []int{initial_num, initial_num, initial_num, initial_num, initial_num}
	}
	if len(array) != size*size {
		return nil,errors.New(fmt.Sprintf("Array2Matrix55/配列サイズが正しくありません Length:%d",len(array)))
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] =  array[i*size+j]
		}
	}
	return matrix,nil
}