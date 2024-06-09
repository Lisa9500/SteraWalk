package pkg

import "log"

// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
func PentaNode(deg []float64, cord [][]float64) (cord5 [][]float64, yane string) {
	// ５角形の各頂点の内角を確認する
	var wa []int
	for d := range deg {
		if deg[d] > 115 {
			wa = append(wa, d)
		}
	}
	log.Println("wa=", wa)

	maxcnt := 0
	// 内角が広角の頂点数に応じて屋根を掛ける
	if len(wa) == 1 {
		// ５角形に変形の寄棟屋根を掛ける
		degmax := 0.0
		for i := range deg {
			if deg[i] > degmax {
				degmax = deg[i]
				maxcnt = i
			}
		}
		yane = "penta"
	} else if len(wa) == 2 {
		maxcnt = wa[0]
		yane = "kiri5"
	}
	log.Println("maxcnt=", maxcnt)

	if maxcnt < 2 {
		slice1 := cord[(maxcnt + 3):5]
		slice2 := cord[0:(maxcnt + 3)]
		cord5 = append(cord5, slice1...)
		cord5 = append(cord5, slice2...)
	} else if maxcnt > 2 {
		slice1 := cord[(maxcnt - 2):5]
		slice2 := cord[0:(maxcnt - 2)]
		cord5 = append(cord5, slice1...)
		cord5 = append(cord5, slice2...)
	} else {
		cord5 = cord
	}

	log.Println("cord5=", cord5)

	return cord5, yane
}
