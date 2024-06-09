package internal

import (
	"log"
	"math"
)

// Stcalc は建物階数から上面高さを設定する
func Stcalc(st, bcr, far int) (toph float64) {
	// 建物の階数が設定されていない場合は乱数で建物階数を設定する
	maxst := int(math.Round(float64(far)/float64(bcr) + 0.5))
	if st >= 3 {
		toph = float64(st) * 3.3
	} else {
		snum := RandStory()
		// 堅ろう建物の場合，全ての建物は３階建て以上とする
		if snum < 3 {
			snum = 3
		}
		if snum > maxst {
			snum = maxst
		}
		log.Println("snum=", snum)
		toph = float64(snum) * 3.3
	}

	return toph
}
