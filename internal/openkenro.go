package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"stera/pkg"
	"strings"
)

// OpenKenro は堅ろう建物のファイルを開く
func OpenKenro() (kList []*KenroBuil, lk int) {
	fpk, er := os.Open("C:/data/kenro_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer fpk.Close()
	log.Println("fpk=", fpk)

	// 構造体のフィールド
	var id string
	var elv float64
	var story int
	var area string
	var bcr int
	var far int
	var cords [][]float64

	// 堅ろう建物データ（構造体）のスライスを作成する
	// kList := []*KenroBuil{}

	ks := bufio.NewScanner(fpk)

	for ks.Scan() {
		// 堅ろう建物ファイルを処理
		jStr := ks.Text()
		// 右端の「,」を削除，「,」がない行末でもエラーにならない
		jStr = strings.TrimRight(jStr, ",")

		// MultiPolygonをLineStringに置換する
		if strings.Contains(jStr, "[ [ [ [") {
			jStr = strings.Replace(jStr, "[ [ [ [", "[ [", 1)
			jStr = strings.Replace(jStr, "] ] ] ]", "] ]", 1)
		}
		// PolygonをLineStringに置換する
		if strings.Contains(jStr, "[ [ [") {
			jStr = strings.Replace(jStr, "[ [ [", "[ [", 1)
			jStr = strings.Replace(jStr, "] ] ]", "] ]", 1)
		}

		id, elv, story, area, bcr, far, cords = pkg.ParseJSON(jStr)
		// log.Println("IDデータ", id)
		// log.Println("標高データ", elv)
		// log.Println("座標データ", cords)
		// story = 3
		bf := 0 // TODO:
		kenro := KenroBuil{ID: id, Elv: elv, Story: story, Basement: bf, Area: area, Build: bcr, Floor: far, Cords: cords}
		kList = append(kList, &kenro)
		// log.Println("kenroList=", kList)
	}
	if ks.Err() != nil {
		log.Fatal(ks.Err())
	}

	// 堅ろう建物の配列の長さを取得する
	lk = (len(kList))
	fmt.Println("lk = ", lk)

	return kList, lk
}
