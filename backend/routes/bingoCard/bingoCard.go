package bingoCard

import (
	"fmt"
	"online-bingo/backend/routes/db"
	"strconv"
)

func GetBingoMatrix(id_hash [32]byte) [][]int {
	initial_num := -1
	bingo_matrix := make([][]int, 5)
	for i := 0; i < len(bingo_matrix); i++ {
		bingo_matrix[i] = []int{initial_num, initial_num, initial_num, initial_num, initial_num}
	}

	for i := 0; i < len(bingo_matrix); i++ {
		for j := 0; j < len(bingo_matrix[0]); j++ {
			current_num := int(id_hash[i*5+j])%15 + 1 + j*15
			// 重複チェック
			for k := 0; isDuplicate(current_num, bingo_matrix); k++ {
				// fmt.Printf("重複：%d,%d<-%d\n", k, ((current_num+7)%15)+1+j*15, current_num)
				current_num = ((current_num + 7) % 15) + 1 + j*15
			}
			bingo_matrix[i][j] = current_num
		}
	}
	//中心のマスをフリーにする
	bingo_matrix[2][2] = 0
	return bingo_matrix
}

func isDuplicate(check_number int, bingo_matrix [][]int) bool {
	for i := 0; i < len(bingo_matrix); i++ {
		for j := 0; j < len(bingo_matrix[0]); j++ {
			if check_number == bingo_matrix[i][j] {
				return true
			}
		}
	}
	return false
}

func GetBingoCount(bingo_matrix [][]int) (int, error) {
	bingoCount := 0
	win_numbers, err := db.GetWinNumbers()
	// 出た数字の取得に失敗した場合
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	checkMatrix := make([][]int, 5)
	for i := 0; i < 5; i++ {
		checkMatrix[i] = make([]int, 5)
	}

	// 穴をあける（各数字が出たかどうかチェック）
	for i := 0; i < len(bingo_matrix); i++ {
		for j := 0; j < len(bingo_matrix[0]); j++ {
			if win_numbers[strconv.Itoa(bingo_matrix[i][j])] == true {
				checkMatrix[i][j] = 1
			}
		}
	}
	fmt.Printf("%v\n", checkMatrix)
	/*
		横5行＋縦5列+対角＼+対角／
		それぞれの穴の数を入れる配列
	*/
	var bingoSum [5 + 5 + 2]int
	for i := range bingoSum {
		bingoSum[i] = 0
	}
	// 空いた穴の数をカウント
	for i := 0; i < len(checkMatrix); i++ {
		for j := 0; j < len(checkMatrix[0]); j++ {
			//縦
			bingoSum[i] += checkMatrix[i][j]
			//横
			bingoSum[j+5] += checkMatrix[i][j]
			//対角＼
			if i == j {
				bingoSum[10] += checkMatrix[i][j]
			}
			//対角／
			if i+j == 5-1 {
				bingoSum[11] += checkMatrix[i][j]
			}
		}
	}
	// fmt.Print("bingo_matrix:")
	// fmt.Println(bingo_matrix)
	// fmt.Print("checkMatrix:")
	// fmt.Println(checkMatrix)
	fmt.Print("bingoSum:")
	fmt.Println(bingoSum)
	for i := range bingoSum {
		if bingoSum[i] == 5 {
			bingoCount++
		}
	}

	return bingoCount, nil

}
