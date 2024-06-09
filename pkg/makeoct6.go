package pkg

import (
	"log"
	"math"
)

// MakeOct6 は凸型2の８角形を３つの４角形に分割する
func MakeOct6(XY [][]float64, order map[string]int) (cord [][]float64,
	rect1List [][]float64, rect2List [][]float64, rect3List [][]float64,
	story []int, yane []string) {
	// octT := "凸型2"
	nodOct := len(XY)
	if nodOct != 8 {
		// TODO:8頂点でない多角形は，三角メッシュ分割
		return
	}

	num1 := order["L1"]
	// 直交する辺は．L1点と1つ前の点で結ばれる線分
	// 直交する辺の座標ペア
	chokuCord1a := make([][]float64, 2)
	num1P1a := (num1 - 1 + nodOct) % nodOct
	chokuCord1a[0] = XY[num1]
	chokuCord1a[1] = XY[num1P1a]
	// L1点と1つ前の点の間の距離
	dl1p1a := DistVerts(XY[num1][0], XY[num1][1], XY[num1P1a][0], XY[num1P1a][1])
	// 対向する辺は，L1点から２つ目と３つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoCord1a := make([][]float64, 2)
	num1P2a := (num1 + 2) % nodOct
	taikoCord1a[0] = XY[num1P2a]
	num1P3a := (num1 + 3) % nodOct
	taikoCord1a[1] = XY[num1P3a]
	// 直交する直線1aと対向する辺との直交条件を確認する
	int1aX, int1aY, theta := OrthoAngle(chokuCord1a, taikoCord1a)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta < 45 || theta > 135 {
		log.Println("theta=", theta)
		// return
	}
	// もう一方の直交する辺は．L点と次の点で結ばれる線分
	// 直交する辺の座標ペア
	chokuCord1b := make([][]float64, 2)
	num1N1b := (num1 + 1) % nodOct
	chokuCord1b[0] = XY[num1]
	chokuCord1b[1] = XY[num1N1b]
	// L1点と次の点の間の距離
	dl1n1b := DistVerts(XY[num1][0], XY[num1][1], XY[num1N1b][0], XY[num1N1b][1])
	// 対向する辺は，L1点から５つ目と６つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoCord1b := make([][]float64, 2)
	num1N5b := (num1 + 5) % nodOct
	taikoCord1b[0] = XY[num1N5b]
	num1N6b := (num1 + 6) % nodOct
	taikoCord1b[1] = XY[num1N6b]
	// 直交する直線1bと対向する辺との直交条件を確認する
	int1bX, int1bY, theta := OrthoAngle(chokuCord1b, taikoCord1b)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta < 45 || theta > 135 {
		log.Println("theta=", theta)
		// return
	}

	num2 := order["L2"]
	// ２つ目の直交する辺は．L2点と1つ前の点で結ばれる線分
	//  直交する辺の座標ペア
	chokuXYa := make([][]float64, 2)
	num2P1a := (num2 - 1 + nodOct) % nodOct
	chokuXYa[0] = XY[num2]
	chokuXYa[1] = XY[num2P1a]
	// L2点と1つ前の点の間の距離
	dl2p1a := DistVerts(XY[num2][0], XY[num2][1], XY[num2P1a][0], XY[num2P1a][1])
	// 対向する辺は，L2点から４つ目と５つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoXYa := make([][]float64, 2)
	num2P2a := (num2 + 4) % nodOct
	taikoXYa[0] = XY[num2P2a]
	num2P3a := (num2 + 5) % nodOct
	taikoXYa[1] = XY[num2P3a]
	// 直交する直線2aと対向する辺との直交条件を確認する
	int2aX, int2aY, theta2 := OrthoAngle(chokuXYa, taikoXYa)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta2 < 45 || theta2 > 135 {
		log.Println("theta=", theta)
		// return
	}
	// もう一方の直交する辺は．L2点と次の点で結ばれる線分
	//  直交する辺の座標ペア
	chokuXYb := make([][]float64, 2)
	num2N1b := (num2 + 1) % nodOct
	chokuXYb[0] = XY[num2]
	chokuXYb[1] = XY[num2N1b]
	// L1点と次の点の間の距離
	dl2n1b := DistVerts(XY[num2][0], XY[num2][1], XY[num2N1b][0], XY[num2N1b][1])
	// 対向する辺は，L2点から５つ目と６つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoXYb := make([][]float64, 2)
	num2N5b := (num2 + 5) % nodOct
	taikoXYb[0] = XY[num2N5b]
	num2N6b := (num2 + 6) % nodOct
	taikoXYb[1] = XY[num2N6b]
	// 直交する直線2bと対向する辺との直交条件を確認する
	int2bX, int2bY, theta2 := OrthoAngle(chokuXYb, taikoXYb)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta2 < 45 || theta2 > 135 {
		log.Println("theta=", theta)
		// return
	}

	// 四角形の頂点のリストを３つ用意する．
	var rect1name []string
	var rect2name []string
	var rect3name []string

	// L点から対向する二辺までの距離を比較する
	// 交点1aまでの距離
	divLine1a := DistVerts(XY[num1][0], XY[num1][1], int1aX, int1aY)
	log.Println("divLine1a=", divLine1a)
	// 交点1bが L2-R5上にあるかチェックする
	t1b := PosLine(taikoCord1b[1][0], int1bX, taikoCord1b[1][1], int1bY, taikoCord1b[0][1], taikoCord1b[0][0])
	log.Println("t1b=", t1b)

	var divLine1b float64
	// if (taikoCord1b[0][0] < int1bX && int1bX < taikoCord1b[1][0]) || (taikoCord1b[0][0] > int1bX && int1bX > taikoCord1b[1][0]) {
	if (taikoCord1b[0][0] < int1bX && int1bX < taikoCord1b[1][0]) || (taikoCord1b[0][0] > int1bX && int1bX > taikoCord1b[1][0]) {
		if (taikoCord1b[0][1] < int1bY && int1bY < taikoCord1b[1][1]) || (taikoCord1b[0][1] > int1bY && int1bY > taikoCord1b[1][1]) {
			if (int1bY*(taikoCord1b[0][0]-taikoCord1b[1][0]))+(taikoCord1b[0][1]*(taikoCord1b[1][0]-int1bX))+(taikoCord1b[1][1]*(int1bX-taikoCord1b[0][0])) == 0 {
				// 交点1bまでの距離
				divLine1b = DistVerts(XY[num1][0], XY[num1][1], int1bX, int1bY)
				log.Println("divLine1b=", divLine1b)
			}
		}
	} else {
		// f_inf = float('inf')
		divLine1b = math.Inf(1)
		log.Println("divLine1b=", divLine1b)
	}
	// 交点2aが R6-L1上にあるかチェックする
	t2a := PosLine(taikoXYa[1][0], int2aX, taikoXYa[1][1], int2aY, taikoXYa[0][1], taikoXYa[0][0])
	log.Println("t2a=", t2a)

	var divLine2a float64
	// if (taikoXYa[0][0] < int2aX && int2aX < taikoXYa[1][0]) || (taikoXYa[0][0] > int2aX && int2aX > taikoXYa[1][0]) {
	if (taikoXYa[0][0] < int2aX && int2aX < taikoXYa[1][0]) || (taikoXYa[0][0] > int2aX && int2aX > taikoXYa[1][0]) {
		if (taikoXYa[0][1] < int2aY && int2aY < taikoXYa[1][1]) || (taikoXYa[0][1] > int2aY && int2aY > taikoXYa[1][1]) {
			if (int2aY*(taikoXYa[0][0]-taikoXYa[1][0]))+(taikoXYa[0][1]*(taikoXYa[1][0]-int2aX))+(taikoXYa[1][1]*(int2aX-taikoXYa[0][0])) == 0 {
				// 交点2aまでの距離
				divLine2a = DistVerts(XY[num2][0], XY[num2][1], int2aX, int2aY)
				log.Println("divLine2a=", divLine2a)
			}
		}
	} else {
		// f_inf = float('inf')
		divLine2a = math.Inf(1)
		log.Println("divLine2a=", divLine2a)
	}
	// 交点2bまでの距離
	divLine2b := DistVerts(XY[num2][0], XY[num2][1], int2bX, int2bY)
	log.Println("divLine2b=", divLine2b)

	// 距離の短い方の線分を分割線とする
	// 距離が無限大の分割線では四角形を分割できない
	if divLine1b == math.Inf(1) {
		log.Println("分割線はdivLine1a")
		// 分割点はD1a点（交点１）
		d1 := []float64{int1aX, int1aY}
		// 座標値のリストにD1点の座標値を追加する
		XY = append(XY, d1)
		// 頂点並びの辞書に分割点を追加する
		d1num := nodOct
		order["D1"] = d1num

		// 四角形D1-L1-R1-R2
		rect1name = []string{"D1", "L1", "R1", "R2"}
		story = append(story, 1)
		if dl1n1b > 1.8 {
			yane = append(yane, "kiri")
		} else {
			yane = append(yane, "kata")
		}

		// 距離の短い方の線分を分割線とする
		if divLine2a < divLine2b {
			log.Println("分割線はdivLine2a")
			// 分割点はD2a点（交点2）
			d2 := []float64{int2aX, int2aY}
			// 座標値のリストにD2点の座標値を追加する
			XY = append(XY, d2)
			// 頂点並びの辞書に分割点を追加する
			d2num := nodOct + 1
			order["D2"] = d2num

			// 四角形D2-L2-R5-R6
			rect2name = []string{"D2", "L2", "R5", "R6"}
			story = append(story, 1)
			if dl2n1b > 1.8 {
				yane = append(yane, "kiri")
			} else {
				yane = append(yane, "kata")
			}
			// 四角形R4-D2-D1-R3
			rect3name = []string{"R4", "D2", "D1", "R3"}
			story = append(story, 2)
			yane = append(yane, "yose")
		} else if divLine2a > divLine2b {
			log.Println("分割線はdivLine2b")
			// 分割点はD2b点（交点2）
			d2 := []float64{int2bX, int2bY}
			// 座標値のリストにD2点の座標値を追加する
			XY = append(XY, d2)
			// 頂点並びの辞書に分割点を追加する
			d2num := nodOct + 1
			order["D2"] = d2num

			// 四角形L2-D2-R3-R4
			rect2name = []string{"L2", "D2", "R3", "R4"}
			story = append(story, 1)
			if dl2p1a > 1.8 {
				yane = append(yane, "kiri")
			} else {
				yane = append(yane, "kata")
			}
			// 四角形D2-R5-R6-D1
			rect3name = []string{"D2", "R5", "R6", "D1"}
			story = append(story, 2)
			yane = append(yane, "kiri")
		}
	} else if divLine2a == math.Inf(1) {
		log.Println("分割線はdivLine2b")
		// 分割点はD2b点（交点２）
		d2 := []float64{int2bX, int2bY}
		// 座標値のリストにD1点の座標値を追加する
		XY = append(XY, d2)
		print(XY)
		// 頂点並びの辞書に分割点を追加する
		d2num := nodOct
		order["D2"] = d2num

		// 四角形L2-D2-R3-R4
		rect1name = []string{"L2", "D2", "R3", "R4"}
		story = append(story, 1)
		if dl2p1a > 1.8 {
			yane = append(yane, "kiri")
		} else {
			yane = append(yane, "kata")
		}

		// 距離の短い方の線分を分割線とする
		if divLine1a < divLine1b {
			log.Println("分割線はdivLine1a")
			// 分割点はD1a点（交点１）
			d1 := []float64{int1aX, int1aY}
			// 座標値のリストにD1点の座標値を追加する
			XY = append(XY, d1)
			// 頂点並びの辞書に分割点を追加する
			d1num := nodOct + 1
			order["D1"] = d1num

			// 四角形D1-L1-R1-R2
			rect2name = []string{"D1", "L1", "R1", "R2"}
			story = append(story, 1)
			if dl1n1b > 1.8 {
				yane = append(yane, "kiri")
			} else {
				yane = append(yane, "kata")
			}
			// 四角形D2-R5-R6-D1
			rect3name = []string{"D2", "R5", "R6", "D1"}
			story = append(story, 2)
			yane = append(yane, "kiri")
		} else if divLine1a > divLine1b {
			log.Println("分割線はdivLine1b")
			// 分割点はD1b点（交点１）
			d1 := []float64{int1bX, int1bY}
			// 座標値のリストにD1点の座標値を追加する
			XY = append(XY, d1)
			// 頂点並びの辞書に分割点を追加する
			d1num := nodOct + 1
			order["D1"] = d1num

			// 四角形L1-D1-R5-R6
			rect2name = []string{"L1", "D1", "R5", "R6"}
			story = append(story, 1)
			if dl1p1a > 1.8 {
				yane = append(yane, "kiri")
			} else {
				yane = append(yane, "kata")
			}
			// 四角形D2-D1-L1-R2
			rect3name = []string{"D2", "D1", "R1", "R2"}
			story = append(story, 2)
			yane = append(yane, "yose")
		}
	}

	// 辞書の中身に従ってリストの座標データで四角形を作る
	rect1List = MakeRectList(XY, order, rect1name)
	log.Println("rect1List=", rect1List)
	rect2List = MakeRectList(XY, order, rect2name)
	log.Println("rect2List=", rect2List)
	rect3List = MakeRectList(XY, order, rect3name)
	log.Println("rect3List=", rect3List)

	cord = XY
	return cord, rect1List, rect2List, rect3List, story, yane
}
