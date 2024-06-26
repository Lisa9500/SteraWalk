package pkg

import "log"

// CwRev は反時計回りを反時計回りに並び替える
func CwRev(num int, xy [][]float64, ext []float64, deg []float64) (revXY [][]float64,
	revExt []float64, revDeg []float64) {
	// log.Println("num", num)

	for i := num - 1; i >= 0; i-- {
		revXY = append(revXY, xy[i])
		revExt = append(revExt, ext[i])
		revDeg = append(revDeg, deg[i])
	}
	for i := 0; i < num; i++ {
		log.Println(i, "Y", revXY[i][0]) // Ctrl+/
		log.Println(i, "X", revXY[i][1]) // Ctrl+/
		log.Println(i, "外積", revExt[i])  // Ctrl+/
		log.Println(i, "内角", revDeg[i])  // Ctrl+/
	}
	return revXY, revExt, revDeg
}
