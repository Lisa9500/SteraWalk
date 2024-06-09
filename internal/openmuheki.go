package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"stera/pkg"
	"strings"
)

// OpenMuheki 無壁舎のファイルを開く
func OpenMuheki() (mList []*MuhekiBuil, lm int) {
	fpm, er := os.Open("C:/data/other_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer fpm.Close()

	// 構造体のフィールド
	var id string
	var elv float64
	var story int
	var area string
	var cords [][]float64

	// 無壁舎建物データ（構造体）のスライスを作成する
	// mList := []*MuhekiBuil{}

	ms := bufio.NewScanner(fpm)
	for ms.Scan() {

		// 無壁舎建物ファイルを処理
		jStr := ms.Text()
		// 右端の「,」を削除，「,」がない行末でもエラーにならない
		jStr = strings.TrimRight(jStr, ",")

		// MultiPolygonをLineStringに置換する
		if strings.Contains(jStr, "[ [ [ [") == true {
			jStr = strings.Replace(jStr, "[ [ [ [", "[ [", 1)
			jStr = strings.Replace(jStr, "] ] ] ]", "] ]", 1)
		}
		// PolygonをLineStringに置換する
		if strings.Contains(jStr, "[ [ [") == true {
			jStr = strings.Replace(jStr, "[ [ [", "[ [", 1)
			jStr = strings.Replace(jStr, "] ] ]", "] ]", 1)
		}

		id, elv, story, area, _, _, cords = pkg.ParseJSON(jStr)
		muheki := MuhekiBuil{ID: id, Elv: elv, Story: story, Area: area, Cords: cords}
		mList = append(mList, &muheki)
		// log.Println("IDデータ", id)
		// log.Println("標高データ", elv)
		// log.Println("座標データ", cords)
	}

	// 無壁舎建物の配列の長さを取得する
	lm = (len(mList))
	fmt.Println("lm = ", lm)

	return mList, lm
}
