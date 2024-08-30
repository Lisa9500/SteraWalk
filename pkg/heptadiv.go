package pkg

import (
	"log"
	"math"
	"strings"
)

// HeptaDiv は７角形を３つに分割して片流れ屋根を掛ける
func HeptaDiv(lrPtn []string, deg2 []float64, cord2 [][]float64, order map[string]int) (slice1 [][]float64,
	slice2 [][]float64, slice3 [][]float64, type1L, type2L, type3L string, story []int, chk bool) {
	// L点が1つの場合と2つの場合で処理が異なる
	// L点を数える
	lrtxt := strings.Join(lrPtn, "")
	log.Println("lrtxt=", lrtxt)
	var roof5 string

	// ７角形を３つに分割
	// L点が連続する場合はL点の側に流れる片流れ屋根を掛ける
	if strings.Contains(lrtxt, "LL") {
		log.Println("lrtxt include LL")
		log.Println("order=", order)
		log.Println("order[\"L1\"]=", order["L1"])
		L1 := order["L1"]
		log.Println("order[\"L2\"]=", order["L2"])
		L2 := order["L2"]

		var slice1t [][]float64
		if L2 < 4 {
			slice1t = cord2[L2 : L2+4]
		} else if L2 > 3 {
			slice11 := cord2[L2:]
			slice12 := cord2[:L2-3]
			slice1t = append(slice1t, slice11...)
			slice1t = append(slice1t, slice12...)
		}
		slice1t1 := slice1t[2:]
		slice1t2 := slice1t[:2]
		slice1 = append(slice1, slice1t1...)
		slice1 = append(slice1, slice1t2...)
		log.Println("slice1=", slice1)

		if L1 > 2 {
			slice2 = cord2[L1-3 : L1+1]
			log.Println("slice2=", slice2)
		} else if L1 < 4 {
			slice21 := cord2[(L1+4)%7:]
			slice22 := cord2[:L1+1]
			slice2 = append(slice2, slice21...)
			slice2 = append(slice2, slice22...)
			log.Println("slice2=", slice2)
		}

		slice31 := cord2[(L2+3)%7 : (L2+3)%7+1]
		slice32 := cord2[L1 : L1+1]
		slice33 := cord2[L2 : L2+1]
		slice3 = append(slice3, slice31...)
		slice3 = append(slice3, slice32...)
		slice3 = append(slice3, slice33...)
		log.Println("slice3=", slice3)

		type1L = "kata1"
		type2L = "kata2"
		type3L = "heptri"
		for i := 0; i < 3; i++ {
			story = append(story, 2)
		}

	} else if strings.Contains(lrtxt, "L") {
		chk = true
		// L点が2つの場合はL点の間のR点の数に応じて屋根の掛け方を変える
		if strings.Count(lrtxt, "L") == 2 {
			log.Println("lrtxt include Lx2")
			log.Println("order=", order)
			log.Println("order[\"L1\"]=", order["L1"])
			L1 := order["L1"]
			log.Println("order[\"L2\"]=", order["L2"])
			L2 := order["L2"]

			// ７角形を３つに分割
			log.Println("deg2=", deg2)
			if L1 == 0 && L2 == 6 {
				// L点が連続する場合はL点の側に流れる片流れ屋根を掛ける
				var slice1t [][]float64
				if L1 < 4 {
					slice1t = cord2[L1 : L1+4]
				} else if L1 > 3 {
					slice11 := cord2[L1:]
					slice12 := cord2[:L1-3]
					slice1t = append(slice1t, slice11...)
					slice1t = append(slice1t, slice12...)
				}
				slice1t1 := slice1t[2:]
				slice1t2 := slice1t[:2]
				slice1 = append(slice1, slice1t1...)
				slice1 = append(slice1, slice1t2...)
				log.Println("slice1=", slice1)

				if L2 > 2 {
					slice2 = cord2[L2-3 : L2+1]
					log.Println("slice2=", slice2)
				} else if L2 < 4 {
					slice21 := cord2[(L2+4)%7:]
					slice22 := cord2[:L2+1]
					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, slice22...)
					log.Println("slice2=", slice2)
				}

				slice31 := cord2[(L1+3)%7 : (L1+3)%7+1]
				slice32 := cord2[L2 : L2+1]
				slice33 := cord2[L1 : L1+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				type1L = "kata1"
				type2L = "kata2"
				type3L = "heptri"
				for i := 0; i < 3; i++ {
					story = append(story, 2)
				}

			} else if (L2-L1+7)%7 == 2 {
				// L点が２つ離れた王冠型−１
				log.Println("２つの三角形と１つの５角形に分割する-1")

				// var maxcnt int
				// if L1 > L2 {
				// 	if deg2[(L1-2+7)%7] > deg2[(L1-3+7)%7] {
				// 		maxcnt = (L1 - 2 + 7) % 7
				// 	} else {
				// 		maxcnt = (L1 - 3 + 7) % 7
				// 	}
				// } else if L1 < L2 {
				// 	if deg2[(L2+2)%7] > deg2[(L2+3)%7] {
				// 		maxcnt = (L2 + 2) % 7
				// 	} else {
				// 		maxcnt = (L2 + 3) % 7
				// 	}
				// }
				// log.Println("maxcnt=", maxcnt)

				// 三角形・片流れ屋根
				// var slice1 [][]float64
				slice11 := cord2[L1 : L1+1]
				var slice12 [][]float64
				if L1 < 1 {
					slice12 = cord2[L1+5:]
				} else if L1 == 1 {
					slice12t1 := cord2[L1+5:]
					slice12t2 := cord2[:(L1+7)%7]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				} else if L1 > 1 {
					slice12 = cord2[(L1+5)%7 : (L1+7)%7]
				}
				slice1 = append(slice1, slice11...)
				slice1 = append(slice1, slice12...)
				log.Println("slice1=", slice1)

				// 三角形・片流れ屋根
				var slice21 [][]float64
				if L2 < 5 {
					slice21 = cord2[L2+2 : L2+3]
				} else if L2 > 4 {
					slice21 = cord2[(L2+2)%7 : (L2+3)%7]
				}
				var slice22 [][]float64
				if L2 < 6 {
					slice22 = cord2[L2 : L2+2]
				} else if L2 == 6 {
					slice22t1 := cord2[L2:]
					slice22t2 := cord2[:(L2+2)%7]
					slice22 = append(slice22, slice22t1...)
					slice22 = append(slice22, slice22t2...)
				}
				slice2 = append(slice2, slice21...)
				slice2 = append(slice2, slice22...)
				log.Println("slice2=", slice2)

				// ５角形
				// var slice3 [][]float64
				var slice31 [][]float64
				if L1 < 5 {
					slice31 = cord2[L1 : L1+3]
				} else if L1 == 5 {
					slice31t1 := cord2[L1 : L1+2]
					slice31t2 := cord2[(L1+2)%7]
					slice31 = append(slice31, slice31t1...)
					slice31 = append(slice31, slice31t2)
				} else if L1 > 5 {
					slice31t1 := cord2[L1]
					slice31t2 := cord2[:(L1+3)%7]
					slice31 = append(slice31, slice31t1)
					slice31 = append(slice31, slice31t2...)
				}
				var slice32 [][]float64
				if L1 < 2 {
					slice32 = cord2[(L1+4)%7 : (L1+6)%7]
				} else if L1 == 2 {
					slice32t1 := cord2[(L1+4)%7:]
					slice32t2 := cord2[:(L1+6)%7]
					slice32 = append(slice32, slice32t1...)
					slice32 = append(slice32, slice32t2...)
				} else if L1 > 2 {
					slice32 = cord2[(L1+4)%7 : (L1+6)%7]
				}
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				log.Println("slice3=", slice3)

				deg5 := PentaDeg(slice3)
				slice3, roof5 = PentaNode(deg5, slice3)

				// slice22t2 := cord2[:(L2+7)%7]
				// slice12 = append(slice12, slice12t1...)
				// slice12 = append(slice12, slice12t2...)
				// slice1 = append(slice1, slice11...)
				// slice1 = append(slice1, slice12...)
				// log.Println("slice1=", slice1)

				// if (maxcnt-L1+7)%7 == 4 {
				// 	log.Println("四角形と５角形(四角形)-1")
				// 	// 四角形・切妻屋根
				// 	slice11 := cord2[L1 : L1+1]
				// 	var slice12 [][]float64
				// 	if maxcnt < 5 {
				// 		slice12 = cord2[maxcnt : maxcnt+3]
				// 	} else if maxcnt > 4 {
				// 		slice12t1 := cord2[maxcnt:]
				// 		slice12t2 := cord2[:(maxcnt+3)%7]
				// 		slice12 = append(slice12, slice12t1...)
				// 		slice12 = append(slice12, slice12t2...)
				// 	}
				// 	slice1 = append(slice1, slice11...)
				// 	slice1 = append(slice1, slice12...)
				// 	log.Println("slice1=", slice1)

				// 	// 四角形・切妻屋根
				// 	slice21 := cord2[maxcnt : maxcnt+1]
				// 	var slice22 [][]float64
				// 	if L1 < 6 {
				// 		slice22 = cord2[L1 : L1+2]
				// 	} else if L1 > 5 {
				// 		slice22t1 := cord2[L1:]
				// 		slice22t2 := cord2[:L2]
				// 		slice22 = append(slice22, slice22t1...)
				// 		slice22 = append(slice22, slice22t2...)
				// 	}
				// 	var slice23 [][]float64
				// 	if L2 < 6 {
				// 		slice23 = cord2[L2+1 : L2+2]
				// 	} else if L2 > 5 {
				// 		slice23 = cord2[:(L2+2)%7]
				// 	}

				// 	slice2 = append(slice2, slice21...)
				// 	slice2 = append(slice2, slice22...)
				// 	slice2 = append(slice2, slice23...)
				// 	log.Println("slice2=", slice2)
				// }

				// if (maxcnt-L1+7)%7 == 2 {
				// 	log.Println("四角形と５角形(四角形)-2")
				// 	// 四角形・切妻屋根
				// 	slice11 := cord2[L2 : L2+1]
				// 	var slice12 [][]float64
				// 	if maxcnt < 5 {
				// 		slice12 = cord2[maxcnt : maxcnt+3]
				// 	} else if maxcnt > 4 {
				// 		slice12t1 := cord2[maxcnt:]
				// 		slice12t2 := cord2[:(maxcnt+3)%7]
				// 		slice12 = append(slice12, slice12t1...)
				// 		slice12 = append(slice12, slice12t2...)
				// 	}
				// 	slice1 = append(slice1, slice11...)
				// 	slice1 = append(slice1, slice12...)
				// 	log.Println("slice1=", slice1)

				// 	// 四角形・切妻屋根
				// 	slice21 := cord2[maxcnt : maxcnt+1]
				// 	var slice22 [][]float64
				// 	if L2 < 6 {
				// 		slice22 = cord2[L2 : L2+2]
				// 	} else if L2 > 5 {
				// 		slice22t1 := cord2[L2:]
				// 		slice22t2 := cord2[:L1]
				// 		slice22 = append(slice22, slice22t1...)
				// 		slice22 = append(slice22, slice22t2...)
				// 	}
				// 	var slice23 [][]float64
				// 	if L1 < 6 {
				// 		slice23 = cord2[L1+1 : L1+2]
				// 	} else if L1 > 5 {
				// 		slice23 = cord2[:(L1+2)%7]
				// 	}

				// 	slice2 = append(slice2, slice21...)
				// 	slice2 = append(slice2, slice22...)
				// 	slice2 = append(slice2, slice23...)
				// 	log.Println("slice2=", slice2)
				// }

				// if (maxcnt-L1+7)%7 == 5 {
				// 	log.Println("四角形と５角形(四角形)-3")
				// 	// 四角形・切妻屋根
				// 	slice11 := cord2[maxcnt : maxcnt+1]
				// 	var slice12 [][]float64
				// 	if L2 < 5 {
				// 		slice12 = cord2[L2 : L2+3]
				// 	} else if L2 > 4 {
				// 		slice12t1 := cord2[L2:]
				// 		slice12t2 := cord2[:(L2+3)%7]
				// 		slice12 = append(slice12, slice12t1...)
				// 		slice12 = append(slice12, slice12t2...)
				// 	}
				// 	slice1 = append(slice1, slice11...)
				// 	slice1 = append(slice1, slice12...)
				// 	log.Println("slice1=", slice1)

				// 	// 四角形・切妻屋根
				// 	slice21 := cord2[L2 : L2+1]
				// 	var slice22 [][]float64
				// 	if maxcnt < 6 {
				// 		slice22 = cord2[maxcnt : maxcnt+2]
				// 	} else if maxcnt > 5 {
				// 		slice22t1 := cord2[maxcnt:]
				// 		slice22t2 := cord2[:L1]
				// 		slice22 = append(slice22, slice22t1...)
				// 		slice22 = append(slice22, slice22t2...)
				// 	}
				// 	var slice23 [][]float64
				// 	if L1 < 6 {
				// 		slice23 = cord2[L1+1 : L1+2]
				// 	} else if L1 > 5 {
				// 		slice23t1 := cord2[L1:]
				// 		slice23t2 := cord2[:(L1+2)%7]
				// 		slice23 = append(slice23, slice23t1...)
				// 		slice23 = append(slice23, slice23t2...)
				// 	}

				// 	slice2 = append(slice2, slice21...)
				// 	slice2 = append(slice2, slice22...)
				// 	slice2 = append(slice2, slice23...)
				// 	log.Println("slice2=", slice2)
				// }

				// if (maxcnt-L1+7)%7 == 3 {
				// 	log.Println("四角形と５角形(四角形)-4")
				// 	// 四角形・切妻屋根
				// 	slice11 := cord2[maxcnt : maxcnt+1]
				// 	var slice12 [][]float64
				// 	if L1 < 5 {
				// 		slice12 = cord2[L1 : L1+3]
				// 	} else if L1 > 4 {
				// 		slice12t1 := cord2[L1:]
				// 		slice12t2 := cord2[:(L1+3)%7]
				// 		slice12 = append(slice12, slice12t1...)
				// 		slice12 = append(slice12, slice12t2...)
				// 	}
				// 	slice1 = append(slice1, slice11...)
				// 	slice1 = append(slice1, slice12...)
				// 	log.Println("slice1=", slice1)

				// 	// 四角形・切妻屋根
				// 	slice21 := cord2[L1 : L1+1]
				// 	var slice22 [][]float64
				// 	if maxcnt < 6 {
				// 		slice22 = cord2[maxcnt : maxcnt+2]
				// 	} else if maxcnt > 5 {
				// 		slice22t1 := cord2[maxcnt:]
				// 		slice22t2 := cord2[:L2]
				// 		slice22 = append(slice22, slice22t1...)
				// 		slice22 = append(slice22, slice22t2...)
				// 	}
				// 	var slice23 [][]float64
				// 	if L2 < 6 {
				// 		slice23 = cord2[L2+1 : L2+2]
				// 	} else if L2 > 5 {
				// 		slice23 = cord2[:(L2+2)%7]
				// 	}

				// 	slice2 = append(slice2, slice21...)
				// 	slice2 = append(slice2, slice22...)
				// 	slice2 = append(slice2, slice23...)
				// 	log.Println("slice2=", slice2)
				// }

				// // var slice3 [][]float64
				// slice31 := cord2[L1 : L1+1]
				// slice32 := cord2[L2 : L2+1]
				// slice33 := cord2[maxcnt : maxcnt+1]
				// slice3 = append(slice3, slice31...)
				// slice3 = append(slice3, slice32...)
				// slice3 = append(slice3, slice33...)
				// log.Println("slice3=", slice3)

				// // 四角形slice1の内角をチェックする
				// type1L = "kiri"
				// _, degLst1, _ := TriVert(len(slice1), slice1)
				// for _, v := range degLst1 {
				// 	if v > 135 {
				// 		type1L = "kata"
				// 	}
				// }
				// for _, v := range degLst1 {
				// 	if v < 45 {
				// 		type1L = "flat"
				// 	}
				// }

				// // 四角形slice2の内角をチェックする
				// type2L = "kiri"
				// _, degLst2, _ := TriVert(len(slice2), slice2)
				// for _, v := range degLst2 {
				// 	if v > 135 {
				// 		type2L = "kata"
				// 	}
				// }
				// for _, v := range degLst2 {
				// 	if v < 45 {
				// 		type1L = "flat"
				// 	}
				// }

				type1L = "tri"
				type2L = "tri"
				type3L = roof5
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 2)

			} else if (L2-L1+7)%7 == 5 {
				// L点が２つ離れた王冠型−2
				log.Println("２つの三角形と１つの５角形に分割する-2")

				// var maxcnt int
				// if L1 > L2 {
				// 	if deg2[(L1+2)%7] > deg2[(L1+3)%7] {
				// 		maxcnt = (L1 + 2) % 7
				// 	} else {
				// 		maxcnt = (L1 + 3) % 7
				// 	}
				// } else if L1 < L2 {
				// 	if deg2[(L2-2+7)%7] > deg2[(L2-3+7)%7] {
				// 		maxcnt = (L2 - 2 + 7) % 7
				// 	} else {
				// 		maxcnt = (L2 - 3 + 7) % 7
				// 	}
				// }
				// log.Println("maxcnt=", maxcnt)

				// 三角形・片流れ屋根
				// var slice1 [][]float64
				slice11 := cord2[L2 : L2+1]
				var slice12 [][]float64
				if L2 < 1 {
					slice12 = cord2[L2+5:]
				} else if L2 == 1 {
					slice12t1 := cord2[L2+5:]
					slice12t2 := cord2[:(L2+7)%7]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				} else if L1 > 1 {
					slice12 = cord2[(L2+5)%7 : (L2+7)%7]
				}
				slice1 = append(slice1, slice11...)
				slice1 = append(slice1, slice12...)
				log.Println("slice1=", slice1)

				// 三角形・片流れ屋根
				var slice21 [][]float64
				if L1 < 5 {
					slice21 = cord2[L1+2 : L1+3]
				} else if L1 > 4 {
					slice21 = cord2[(L1+2)%7 : (L1+3)%7]
				}
				var slice22 [][]float64
				if L1 < 6 {
					slice22 = cord2[L1 : L1+2]
				} else if L1 == 6 {
					slice22t1 := cord2[L1:]
					slice22t2 := cord2[:(L1+2)%7]
					slice22 = append(slice22, slice22t1...)
					slice22 = append(slice22, slice22t2...)
				}
				slice2 = append(slice2, slice21...)
				slice2 = append(slice2, slice22...)
				log.Println("slice2=", slice2)

				// ５角形
				// var slice3 [][]float64
				var slice31 [][]float64
				if L2 < 5 {
					slice31 = cord2[L2 : L2+3]
				} else if L2 == 5 {
					slice31t1 := cord2[L2 : L2+2]
					slice31t2 := cord2[(L2+2)%7]
					slice31 = append(slice31, slice31t1...)
					slice31 = append(slice31, slice31t2)
				} else if L2 > 5 {
					slice31t1 := cord2[L2]
					slice31t2 := cord2[:(L2+3)%7]
					slice31 = append(slice31, slice31t1)
					slice31 = append(slice31, slice31t2...)
				}
				var slice32 [][]float64
				if L2 < 2 {
					slice32 = cord2[(L2+4)%7 : (L2+6)%7]
				} else if L2 == 2 {
					slice32t1 := cord2[(L2+4)%7:]
					slice32t2 := cord2[:(L2+6)%7]
					slice32 = append(slice32, slice32t1...)
					slice32 = append(slice32, slice32t2...)
				} else if L2 > 2 {
					slice32 = cord2[(L2+4)%7 : (L2+6)%7]
				}
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				log.Println("slice3=", slice3)

				deg5 := PentaDeg(slice3)
				slice3, roof5 = PentaNode(deg5, slice3)

				// if (maxcnt-L1+7)%7 == 4 {
				// 	log.Println("四角形と５角形(四角形)-1")
				// 	// 四角形・切妻屋根
				// 	slice11 := cord2[L1 : L1+1]
				// 	var slice12 [][]float64
				// 	if maxcnt < 5 {
				// 		slice12 = cord2[maxcnt : maxcnt+3]
				// 	} else if maxcnt > 4 {
				// 		slice12t1 := cord2[maxcnt:]
				// 		slice12t2 := cord2[:(maxcnt+3)%7]
				// 		slice12 = append(slice12, slice12t1...)
				// 		slice12 = append(slice12, slice12t2...)
				// 	}
				// 	slice1 = append(slice1, slice11...)
				// 	slice1 = append(slice1, slice12...)
				// 	log.Println("slice1=", slice1)

				// 	// 四角形・切妻屋根
				// 	slice21 := cord2[maxcnt : maxcnt+1]
				// 	var slice22 [][]float64
				// 	if L1 < 6 {
				// 		slice22 = cord2[L1 : L1+2]
				// 	} else if L1 > 5 {
				// 		slice22t1 := cord2[L1:]
				// 		slice22t2 := cord2[:L2]
				// 		slice22 = append(slice22, slice22t1...)
				// 		slice22 = append(slice22, slice22t2...)
				// 	}
				// 	var slice23 [][]float64
				// 	if L2 < 6 {
				// 		slice23 = cord2[L2+1 : L2+2]
				// 	} else if L2 > 5 {
				// 		slice23 = cord2[:(L2+2)%7]
				// 	}

				// 	slice2 = append(slice2, slice21...)
				// 	slice2 = append(slice2, slice22...)
				// 	slice2 = append(slice2, slice23...)
				// 	log.Println("slice2=", slice2)
				// }

				// if (maxcnt-L1+7)%7 == 2 {
				// 	log.Println("四角形と５角形(四角形)-2")
				// 	// 四角形・切妻屋根
				// 	slice11 := cord2[L2 : L2+1]
				// 	var slice12 [][]float64
				// 	if maxcnt < 5 {
				// 		slice12 = cord2[maxcnt : maxcnt+3]
				// 	} else if maxcnt > 4 {
				// 		slice12t1 := cord2[maxcnt:]
				// 		slice12t2 := cord2[:(maxcnt+3)%7]
				// 		slice12 = append(slice12, slice12t1...)
				// 		slice12 = append(slice12, slice12t2...)
				// 	}
				// 	slice1 = append(slice1, slice11...)
				// 	slice1 = append(slice1, slice12...)
				// 	log.Println("slice1=", slice1)

				// 	// 四角形・切妻屋根
				// 	slice21 := cord2[maxcnt : maxcnt+1]
				// 	var slice22 [][]float64
				// 	if L2 < 6 {
				// 		slice22 = cord2[L2 : L2+2]
				// 	} else if L2 > 5 {
				// 		slice22t1 := cord2[L2:]
				// 		slice22t2 := cord2[:L1]
				// 		slice22 = append(slice22, slice22t1...)
				// 		slice22 = append(slice22, slice22t2...)
				// 	}
				// 	var slice23 [][]float64
				// 	if L1 < 6 {
				// 		slice23 = cord2[L1+1 : L1+2]
				// 	} else if L1 > 5 {
				// 		slice23 = cord2[:(L1+2)%7]
				// 	}

				// 	slice2 = append(slice2, slice21...)
				// 	slice2 = append(slice2, slice22...)
				// 	slice2 = append(slice2, slice23...)
				// 	log.Println("slice2=", slice2)
				// }

				// if (maxcnt-L1+7)%7 == 5 {
				// 	log.Println("四角形と５角形(四角形)-3")
				// 	// 四角形・切妻屋根
				// 	slice11 := cord2[maxcnt : maxcnt+1]
				// 	var slice12 [][]float64
				// 	if L2 < 5 {
				// 		slice12 = cord2[L2 : L2+3]
				// 	} else if L2 > 4 {
				// 		slice12t1 := cord2[L2:]
				// 		slice12t2 := cord2[:(L2+3)%7]
				// 		slice12 = append(slice12, slice12t1...)
				// 		slice12 = append(slice12, slice12t2...)
				// 	}
				// 	slice1 = append(slice1, slice11...)
				// 	slice1 = append(slice1, slice12...)
				// 	log.Println("slice1=", slice1)

				// 	// 四角形・切妻屋根
				// 	slice21 := cord2[L2 : L2+1]
				// 	var slice22 [][]float64
				// 	if maxcnt < 6 {
				// 		slice22 = cord2[maxcnt : maxcnt+2]
				// 	} else if maxcnt > 5 {
				// 		slice22t1 := cord2[maxcnt:]
				// 		slice22t2 := cord2[:L1]
				// 		slice22 = append(slice22, slice22t1...)
				// 		slice22 = append(slice22, slice22t2...)
				// 	}
				// 	var slice23 [][]float64
				// 	if L1 < 6 {
				// 		slice23 = cord2[L1+1 : L1+2]
				// 	} else if L1 > 5 {
				// 		slice23t1 := cord2[L1:]
				// 		slice23t2 := cord2[:(L1+2)%7]
				// 		slice23 = append(slice23, slice23t1...)
				// 		slice23 = append(slice23, slice23t2...)
				// 	}

				// 	slice2 = append(slice2, slice21...)
				// 	slice2 = append(slice2, slice22...)
				// 	slice2 = append(slice2, slice23...)
				// 	log.Println("slice2=", slice2)
				// }

				// if (maxcnt-L1+7)%7 == 3 {
				// 	log.Println("四角形と５角形(四角形)-4")
				// 	// 四角形・切妻屋根
				// 	slice11 := cord2[maxcnt : maxcnt+1]
				// 	var slice12 [][]float64
				// 	if L1 < 5 {
				// 		slice12 = cord2[L1 : L1+3]
				// 	} else if L1 > 4 {
				// 		slice12t1 := cord2[L1:]
				// 		slice12t2 := cord2[:(L1+3)%7]
				// 		slice12 = append(slice12, slice12t1...)
				// 		slice12 = append(slice12, slice12t2...)
				// 	}
				// 	slice1 = append(slice1, slice11...)
				// 	slice1 = append(slice1, slice12...)
				// 	log.Println("slice1=", slice1)

				// 	// 四角形・切妻屋根
				// 	slice21 := cord2[L1 : L1+1]
				// 	var slice22 [][]float64
				// 	if maxcnt < 6 {
				// 		slice22 = cord2[maxcnt : maxcnt+2]
				// 	} else if maxcnt > 5 {
				// 		slice22t1 := cord2[maxcnt:]
				// 		slice22t2 := cord2[:L2]
				// 		slice22 = append(slice22, slice22t1...)
				// 		slice22 = append(slice22, slice22t2...)
				// 	}
				// 	var slice23 [][]float64
				// 	if L2 < 6 {
				// 		slice23 = cord2[L2+1 : L2+2]
				// 	} else if L2 > 5 {
				// 		slice23 = cord2[:(L2+2)%7]
				// 	}

				// 	slice2 = append(slice2, slice21...)
				// 	slice2 = append(slice2, slice22...)
				// 	slice2 = append(slice2, slice23...)
				// 	log.Println("slice2=", slice2)
				// }

				// var slice3 [][]float64
				// slice31 := cord2[L1 : L1+1]
				// slice32 := cord2[L2 : L2+1]
				// slice33 := cord2[maxcnt : maxcnt+1]
				// slice3 = append(slice3, slice31...)
				// slice3 = append(slice3, slice32...)
				// slice3 = append(slice3, slice33...)
				// log.Println("slice3=", slice3)

				// // 四角形slice1の内角をチェックする
				// type1L = "kiri"
				// _, degLst1, _ := TriVert(len(slice1), slice1)
				// for _, v := range degLst1 {
				// 	if v > 135 {
				// 		type1L = "kata"
				// 	}
				// }
				// for _, v := range degLst1 {
				// 	if v < 45 {
				// 		type1L = "flat"
				// 	}
				// }

				// // 四角形slice2の内角をチェックする
				// type2L = "kiri"
				// _, degLst2, _ := TriVert(len(slice2), slice2)
				// for _, v := range degLst2 {
				// 	if v > 135 {
				// 		type2L = "kata"
				// 	}
				// }
				// for _, v := range degLst2 {
				// 	if v < 45 {
				// 		type1L = "flat"
				// 	}
				// }

				type1L = "tri"
				type2L = "tri"
				type3L = roof5
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 2)

			} else if ((L1 < L2) && (L2-L1 == 3)) || ((L1 > L2) && (L1-L2 == 4)) {
				// L点が３つ離れたイカ型−１
				log.Println("１つの５角形と１つの四角形に分割する-1")
				// // 四角形・切妻屋根
				// slice11 := cord2[L2 : L2+1]
				// var slice12 [][]float64
				// if L1 < 5 {
				// 	slice12 = cord2[L1 : L1+3]
				// } else if L1 > 4 {
				// 	slice12t1 := cord2[L1:]
				// 	slice12t2 := cord2[:(L1+3)%7]
				// 	slice12 = append(slice12, slice12t1...)
				// 	slice12 = append(slice12, slice12t2...)
				// }
				// slice1 = append(slice1, slice11...)
				// slice1 = append(slice1, slice12...)
				// log.Println("slice1=", slice1)

				// ５角形
				// var slice1 [][]float64
				if L2 < 3 {
					slice1 = cord2[L2 : L2+5]
				} else if L2 > 2 {
					slice11 := cord2[L2:]
					slice12 := cord2[:L2-2]
					slice1 = append(slice1, slice11...)
					slice1 = append(slice1, slice12...)
				}
				log.Println("slice1=", slice1)

				deg5 := PentaDeg(slice1)
				slice1, roof5 = PentaNode(deg5, slice1)

				// 四角形・切妻屋根
				// var slice2 [][]float64
				slice21 := cord2[L2 : L2+1]
				var slice22 [][]float64
				if L2 < 1 {
					slice22 = cord2[L2+4:]
				} else if L2 > 0 && L2 < 3 {
					slice22t1 := cord2[L2+4:]
					slice22t2 := cord2[:L2]
					slice22 = append(slice22, slice22t1...)
					slice22 = append(slice22, slice22t2...)
				} else if L2 > 2 {
					slice22 = cord2[L2-3 : L2]
				}
				slice2 = append(slice2, slice21...)
				slice2 = append(slice2, slice22...)
				log.Println("slice2=", slice2)

				// // 作成した５角形にL点が含まれるとうまく屋根が掛からない
				// // その場合は変形した四角形の切妻屋根とする
				// n := len(slice2)
				// ext, _, _ := TriVert(n, slice2)
				// newslice := make([][]float64, 0, 5)
				// for e, v := range ext {
				// 	if v <= 0.0 {
				// 		newslice = append(newslice, slice2[e])
				// 	}
				// }
				// m := len(newslice)
				// if m == 5 {
				// 	newpenta, yane := PentaNode(deg2, newslice)
				// 	slice2 = newpenta
				// 	log.Println("slice2=", slice2)
				// 	type2L = yane
				// } else if m == 4 {
				// 	slice2 = newslice
				// 	type2L = "kiri"
				// }
				// log.Println("slice2=", slice2)

				// ダミー三角形
				midcnt := (L2 + 2) % 7

				slice31 := cord2[L1 : L1+1]
				slice32 := cord2[midcnt : midcnt+1]
				slice33 := cord2[L2 : L2+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				type1L = roof5
				type2L = "kiri"
				type3L = "flat"
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 1)

			} else if ((L1 > L2) && (L1-L2 == 3)) || ((L1 < L2) && (L2-L1 == 4)) {
				// L点が３つ離れたイカ型−１
				log.Println("１つの５角形と１つの四角形に分割する-2")
				// // 四角形・切妻屋根
				// slice11 := cord2[L1 : L1+1]
				// var slice12 [][]float64
				// if L2 < 5 {
				// 	slice12 = cord2[L2 : L2+3]
				// } else if L2 > 4 {
				// 	slice12t1 := cord2[L2:]
				// 	slice12t2 := cord2[:(L2+3)%7]
				// 	slice12 = append(slice12, slice12t1...)
				// 	slice12 = append(slice12, slice12t2...)
				// }
				// slice1 = append(slice1, slice11...)
				// slice1 = append(slice1, slice12...)
				// log.Println("slice1=", slice1)

				// ５角形
				// var slice1 [][]float64
				if L1 < 3 {
					slice1 = cord2[L1 : L1+5]
				} else if L1 > 2 {
					slice11 := cord2[L1:]
					slice12 := cord2[:L1-2]
					slice1 = append(slice1, slice11...)
					slice1 = append(slice1, slice12...)
				}
				log.Println("slice1=", slice1)

				deg5 := PentaDeg(slice1)
				slice1, roof5 = PentaNode(deg5, slice1)

				// 四角形・切妻屋根
				// var slice2 [][]float64
				slice21 := cord2[L1 : L1+1]
				var slice22 [][]float64
				if L1 < 1 {
					slice22 = cord2[L1+4:]
				} else if L1 > 0 && L1 < 3 {
					slice22t1 := cord2[L1+4:]
					slice22t2 := cord2[:L1]
					slice22 = append(slice22, slice22t1...)
					slice22 = append(slice22, slice22t2...)
				} else if L1 > 2 {
					slice22 = cord2[L1-3 : L1]
				}
				slice2 = append(slice2, slice21...)
				slice2 = append(slice2, slice22...)
				log.Println("slice2=", slice2)

				// // 作成した５角形にL点が含まれるとうまく屋根が掛からない
				// // その場合は変形した四角形の切妻屋根とする
				// n := len(slice2)
				// ext, _, _ := TriVert(n, slice2)
				// newslice := make([][]float64, 0, 5)
				// for e, v := range ext {
				// 	if v <= 0.0 {
				// 		newslice = append(newslice, slice2[e])
				// 	}
				// }
				// m := len(newslice)
				// if m == 5 {
				// 	newpenta, yane := PentaNode(deg2, newslice)
				// 	slice2 = newpenta
				// 	log.Println("slice2=", slice2)
				// 	type2L = yane
				// } else if m == 4 {
				// 	slice2 = newslice
				// 	type2L = "kiri"
				// }
				// log.Println("slice2=", slice2)

				// ダミー三角形
				midcnt := (L1 + 2) % 7

				slice31 := cord2[L1 : L1+1]
				slice32 := cord2[midcnt : midcnt+1]
				slice33 := cord2[L2 : L2+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				type1L = roof5
				type2L = "kiri"
				type3L = "flat"
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 1)
			}

		} else if strings.Count(lrtxt, "L") == 1 {
			// L1点から伸ばした線が対向する辺に直交する場合
			log.Println("直交する点で四角形と５角形に分割する")
			// L1点の頂点番号を確認する
			var num int
			for LRkey := range order {
				log.Println("LRkey=", LRkey) // Ctrl+/
				if LRkey == "L1" {
					num = order[LRkey]      // 頂点番号
					log.Println("num", num) // Ctrl+/
				}
			}
			// L1点のX座標
			p := cord2[num][0]
			// L1点のY座標
			q := cord2[num][1]

			// 対向する辺は，L点から２つ目と３つ目の点で結ばれる線分
			// 対向する辺２−３の座標ペア
			taikoCord1 := make([][]float64, 2)
			numP2 := (num + 2) % 7
			taikoCord1[0] = cord2[numP2]
			numP3 := (num + 3) % 7
			taikoCord1[1] = cord2[numP3]
			// 対向する辺の直線の方程式
			line1 := LineEquat(taikoCord1[0][0], taikoCord1[0][1], taikoCord1[1][0], taikoCord1[1][1])
			a1 := line1["m"]
			b1 := line1["n"]
			// L1点から対向する辺に下ろした垂線の交点の座標
			x1 := (a1*(q-b1) + p) / (math.Pow(a1, 2) + 1)
			y1 := (a1*(a1*(q-b1)+p))/(math.Pow(a1, 2)+1) + b1
			D1 := []float64{x1, y1}
			log.Println("D1=", D1)
			// 垂線の交点が辺の上にあるかどうかチェックする
			chk1 := PointonLine(taikoCord1[0][0], taikoCord1[0][1], taikoCord1[1][0], taikoCord1[1][1], x1, y1)
			// perpen1 := [][]float64{{p, q}, {x1, y1}}
			// chk1 := PosLine2(taikoCord1, perpen1)
			// chk1 := PosLine(taikoCord1[1][0], taikoCord1[0][0], taikoCord1[1][1], taikoCord1[0][1], y1, x1)
			log.Println("chk1=", chk1)

			// もう一方の対向する辺は，L点から４つ目と５つ目の点で結ばれる線分
			// 対向する辺4-5の座標ペア
			taikoCord2 := make([][]float64, 2)
			numN4 := (num + 4) % 7
			taikoCord2[0] = cord2[numN4]
			numN5 := (num + 5) % 7
			taikoCord2[1] = cord2[numN5]
			// 対向する辺の直線の方程式
			line2 := LineEquat(taikoCord2[0][0], taikoCord2[0][1], taikoCord2[1][0], taikoCord2[1][1])
			a2 := line2["m"]
			b2 := line2["n"]
			// L1点から対向する辺に下ろした垂線の交点の座標
			x2 := (a2*(q-b2) + p) / (math.Pow(a2, 2) + 1)
			y2 := (a2*(a2*(q-b2)+p))/(math.Pow(a2, 2)+1) + b2
			D2 := []float64{x2, y2}
			log.Println("D2=", D2)
			// 垂線の交点が辺の上にあるかどうかチェックする
			chk2 := PointonLine(taikoCord2[0][0], taikoCord2[0][1], taikoCord2[1][0], taikoCord2[1][1], x2, y2)
			// perpen2 := [][]float64{{p, q}, {x2, y2}}
			// chk2 := PosLine2(taikoCord2, perpen2)
			// chk2 := PosLine(taikoCord2[1][0], taikoCord2[0][0], taikoCord2[1][1], taikoCord2[0][1], y2, x2)
			log.Println("chk2=", chk2)

			// さらにもう一方つ対向する辺は，L点から３つ目と４つ目の点で結ばれる線分
			// 対向する辺3-4の座標ペア
			taikoCord3 := make([][]float64, 2)
			numC3 := (num + 3) % 7
			taikoCord3[0] = cord2[numC3]
			numC4 := (num + 4) % 7
			taikoCord3[1] = cord2[numC4]
			// 対向する辺の直線の方程式
			line3 := LineEquat(taikoCord3[0][0], taikoCord3[0][1], taikoCord3[1][0], taikoCord3[1][1])
			a3 := line3["m"]
			b3 := line3["n"]
			// L1点から対向する辺に下ろした垂線の交点の座標
			x3 := (a3*(q-b3) + p) / (math.Pow(a3, 2) + 1)
			y3 := (a3*(a3*(q-b3)+p))/(math.Pow(a3, 2)+1) + b3
			D3 := []float64{x3, y3}
			log.Println("D3=", D3)
			// 垂線の交点が辺の上にあるかどうかチェックする
			chk3 := PointonLine(taikoCord3[0][0], taikoCord3[0][1], taikoCord3[1][0], taikoCord3[1][1], x3, y3)
			// perpen3 := [][]float64{{p, q}, {x3, y3}}
			// chk3 := PosLine2(taikoCord3, perpen3)
			// chk2 := PosLine(taikoCord2[1][0], taikoCord2[0][0], taikoCord2[1][1], taikoCord2[0][1], y2, x2)
			log.Println("chk3=", chk3)

			// 四角形aを分割する
			// var slice1 [][]float64
			if chk1 == true {
				log.Println("四角形１を分割する")
				slice1 = append(slice1, D1)
				var slice12 [][]float64
				if num < 5 {
					slice12 = cord2[num : num+3]
				} else if num > 4 {
					slice12t1 := cord2[num:]
					slice12t2 := cord2[:(num-4)%7]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				}
				slice1 = append(slice1, slice12...)
				type1L = "kiri"
				story = append(story, 2)

				if chk2 == true {
					log.Println("四角形２を分割する")
					slice21 := cord2[num : num+1]
					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, D2)
					var slice22 [][]float64
					if num < 1 {
						slice22 = cord2[(num+5)%7 : (num+7)%7]
					} else if num == 1 {
						slice22t1 := cord2[(num+5)%7:]
						slice22t2 := cord2[:(num+7)%7]
						slice22 = append(slice22, slice22t1...)
						slice22 = append(slice22, slice22t2...)
					} else if num > 1 {
						slice22 = cord2[(num-2+7)%7 : num]
					}
					slice2 = append(slice2, slice22...)
					type2L = "kiri"
					story = append(story, 2)

					log.Println("5角形を分割する")
					slice31 := cord2[num : num+1]
					slice3 = append(slice3, slice31...)
					slice3 = append(slice3, D1)
					var slice32 [][]float64
					if num < 3 {
						slice32 = cord2[num+3 : num+5]
						slice3 = append(slice3, slice32...)
					} else if num == 3 {
						slice32t1 := cord2[num+3:]
						slice32t2 := cord2[:(num+5)%7]
						slice32 = append(slice32, slice32t1...)
						slice32 = append(slice32, slice32t2...)
						slice3 = append(slice3, slice32...)
					} else if num > 3 {
						slice32 = cord2[num-4 : num-2]
						slice3 = append(slice3, slice32...)
					}
					slice3 = append(slice3, D2)
					log.Println("slice3=", slice3)

					deg5 := PentaDeg(slice3)
					slice3, roof5 = PentaNode(deg5, slice3)
					type3L = roof5
					story = append(story, 2)

				} else if chk3 == true {
					log.Println("四角形３を分割する")
					slice21 := cord2[num : num+1]
					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, D1)
					var slice22 [][]float64
					if num < 3 {
						slice22 = cord2[num+3 : num+5]
						slice2 = append(slice2, slice22...)
					} else if num == 3 {
						slice22t1 := cord2[num+3:]
						slice22t2 := cord2[:(num+5)%7]
						slice22 = append(slice22, slice22t1...)
						slice22 = append(slice22, slice22t2...)
						slice2 = append(slice2, slice22...)
					} else if num > 3 {
						slice22 = cord2[num-4 : num-2]
						slice2 = append(slice2, slice22...)
					}
					type2L = "kiri"
					story = append(story, 2)

					log.Println("5角形を分割する")
					slice31 := cord2[num : num+1]
					slice3 = append(slice3, slice31...)
					slice3 = append(slice3, D3)
					var slice32 [][]float64
					if num < 1 {
						slice32 = cord2[num+4 : num+7]
						slice3 = append(slice3, slice32...)
					} else if num > 0 && num < 3 {
						slice32t1 := cord2[num+4:]
						slice32t2 := cord2[:num]
						slice32 = append(slice32, slice32t1...)
						slice32 = append(slice32, slice32t2...)
						slice3 = append(slice3, slice32...)
					} else if num > 2 {
						slice32 = cord2[num-3 : num]
						slice3 = append(slice3, slice32...)
					}
					slice3 = append(slice3, D2)

					deg5 := PentaDeg(slice3)
					slice3, roof5 = PentaNode(deg5, slice3)
					type3L = roof5
					story = append(story, 2)
				} else {
					chk = false
				}
			} else if chk2 == true {
				log.Println("四角形２を分割する")
				slice11 := cord2[num : num+1]
				slice1 = append(slice1, slice11...)
				slice1 = append(slice1, D2)
				var slice12 [][]float64
				if num < 1 {
					slice12 = cord2[num+5 : num+7]
				} else if num == 1 {
					slice12t1 := cord2[num+5:]
					slice12t2 := cord2[:num]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				} else if num > 1 {
					slice12 = cord2[(num+5)%7 : num]
				}
				slice1 = append(slice1, slice12...)
				type1L = "kiri"
				story = append(story, 2)

				if chk3 == true {
					log.Println("5角形を分割する")
					slice21 := cord2[num : num+1]
					slice2 = append(slice2, slice21...)
					var slice22 [][]float64
					if num < 4 {
						slice22 = cord2[num+1 : num+4]
						slice2 = append(slice2, slice22...)
					} else if num > 3 && num < 6 {
						slice22t1 := cord2[num+1:]
						slice22t2 := cord2[:(num+4)%7]
						slice22 = append(slice22, slice22t1...)
						slice22 = append(slice22, slice22t2...)
						slice2 = append(slice2, slice22...)
					} else if num > 5 {
						slice22 = cord2[(num+1)%7 : (num+4)%7]
						slice2 = append(slice2, slice22...)
					}
					slice2 = append(slice2, D2)

					deg5 := PentaDeg(slice2)
					slice2, roof5 = PentaNode(deg5, slice2)
					type2L = roof5
					story = append(story, 2)

					log.Println("四角形４を分割する")
					slice31 := cord2[num : num+1]
					slice3 = append(slice3, slice31...)
					slice3 = append(slice3, D3)
					var slice32 [][]float64
					if num < 3 {
						slice32 = cord2[num+4 : num+5]
						slice3 = append(slice3, slice32...)
					} else if num > 2 {
						slice32 = cord2[(num+4)%7 : (num+5)%7]
						slice3 = append(slice3, slice32...)
					}
					slice3 = append(slice3, D2)
					type3L = "flat"
					story = append(story, 2)
				} else {
					chk = false
				}
			} else {
				chk = false
			}
			//  else if chk1 < 0 {
			// 	log.Println("四角形１を分割しない")
			// 	if num < 4 {
			// 		slice1 = cord2[num : num+4]
			// 	} else if num > 3 {
			// 		slice11 := cord2[num:]
			// 		slice12 := cord2[:num-3]
			// 		slice1 = append(slice1, slice11...)
			// 		slice1 = append(slice1, slice12...)
			// 	}
			// 	type1L = "flat"
			// 	story = append(story, 2)
			// }

			// 四角形bを分割する
			// var slice2 [][]float64
			// if chk2 > 0 {
			// 	log.Println("四角形２を分割する")
			// 	slice21 := cord2[num : num+1]
			// 	slice2 = append(slice2, slice21...)
			// 	slice2 = append(slice2, D2)
			// 	var slice22 [][]float64
			// 	if num < 1 {
			// 		slice22 = cord2[(num+5)%7 : (num+7)%7]
			// 	} else if num == 1 {
			// 		slice22t1 := cord2[(num+5)%7:]
			// 		slice22t2 := cord2[:(num+7)%7]
			// 		slice22 = append(slice22, slice22t1...)
			// 		slice22 = append(slice22, slice22t2...)
			// 	} else if num > 1 {
			// 		slice22 = cord2[(num-2+7)%7 : num]
			// 	}
			// 	slice2 = append(slice2, slice22...)
			// 	type2L = "kiri"
			// 	story = append(story, 2)
			// }
			//  else if chk2 < 0 {
			// 	log.Println("四角形２を分割しない")
			// 	slice21 := cord2[num : num+1]
			// 	slice2 = append(slice2, slice21...)
			// 	var slice22 [][]float64
			// 	if num < 1 {
			// 		slice22 = cord2[num+4:]
			// 	} else if num > 0 && num < 3 {
			// 		slice22t1 := cord2[num+4 : (num + 6)]
			// 		slice22t2 := cord2[:num]
			// 		slice22 = append(slice22, slice22t1...)
			// 		slice22 = append(slice22, slice22t2...)
			// 	} else if num > 2 {
			// 		slice22 = cord2[num-3 : num]
			// 	}
			// 	slice2 = append(slice2, slice22...)
			// 	type2L = "flat"
			// 	story = append(story, 2)
			// }

			// ５角形を分割する
			// var slice3 [][]float64
			// if chk1 < 0 && chk2 < 0 {
			// 	log.Println("5角形を分割する")
			// 	slice31 := cord2[num : num+1]
			// 	slice3 = append(slice3, slice31...)
			// 	slice3 = append(slice3, D1)
			// 	var slice32 [][]float64
			// 	if num < 3 {
			// 		slice32 = cord2[num+3 : num+5]
			// 		slice3 = append(slice3, slice32...)
			// 	} else if num == 3 {
			// 		slice32t1 := cord2[num+3:]
			// 		slice32t2 := cord2[:(num+5)%7]
			// 		slice32 = append(slice32, slice32t1...)
			// 		slice32 = append(slice32, slice32t2...)
			// 		slice3 = append(slice3, slice32...)
			// 	} else if num > 3 {
			// 		slice32 = cord2[num-4 : num-2]
			// 		slice3 = append(slice3, slice32...)
			// 	}
			// 	slice3 = append(slice3, D2)
			// 	type3L = "penta"
			// 	story = append(story, 2)

			// } else if chk1 > 0 && chk2 < 0 {
			// 	log.Println("四角形３を分割する")
			// 	slice31 := cord2[num : num+1]
			// 	slice3 = append(slice3, slice31...)
			// 	var slice32 [][]float64
			// 	if num < 3 {
			// 		slice32 = cord2[num+3 : num+5]
			// 		slice3 = append(slice3, slice32...)
			// 	} else if num == 3 {
			// 		slice32t1 := cord2[num+3:]
			// 		slice32t2 := cord2[:(num+5)%7]
			// 		slice32 = append(slice32, slice32t1...)
			// 		slice32 = append(slice32, slice32t2...)
			// 		slice3 = append(slice3, slice32...)
			// 	} else if num > 3 {
			// 		slice32 = cord2[num-4 : num-2]
			// 		slice3 = append(slice3, slice32...)
			// 	}
			// 	slice3 = append(slice3, D2)
			// 	type3L = "flat"
			// 	story = append(story, 2)

			// } else if chk1 < 0 && chk2 > 0 {
			// 	log.Println("四角形４を分割する")
			// 	slice31 := cord2[num : num+1]
			// 	slice3 = append(slice3, slice31...)
			// 	slice3 = append(slice3, D1)
			// 	var slice32 [][]float64
			// 	if num < 3 {
			// 		slice32 = cord2[num+3 : num+5]
			// 		slice3 = append(slice3, slice32...)
			// 	} else if num == 3 {
			// 		slice32t1 := cord2[num+3:]
			// 		slice32t2 := cord2[:(num+5)%7]
			// 		slice32 = append(slice32, slice32t1...)
			// 		slice32 = append(slice32, slice32t2...)
			// 		slice3 = append(slice3, slice32...)
			// 	} else if num > 3 {
			// 		slice32 = cord2[num-4 : num-2]
			// 		slice3 = append(slice3, slice32...)
			// 	}
			// 	type3L = "flat"
			// 	story = append(story, 2)

			// } else if chk1 > 0 && chk2 > 0 {
			// 	log.Println("三角形を分割する")
			// 	slice31 := cord2[num : num+1]
			// 	slice3 = append(slice3, slice31...)
			// 	var slice32 [][]float64
			// 	if num < 3 {
			// 		slice32 = cord2[num+3 : num+5]
			// 		slice3 = append(slice3, slice32...)
			// 	} else if num == 3 {
			// 		slice32t1 := cord2[num+3:]
			// 		slice32t2 := cord2[:(num+5)%7]
			// 		slice32 = append(slice32, slice32t1...)
			// 		slice32 = append(slice32, slice32t2...)
			// 		slice3 = append(slice3, slice32...)
			// 	} else if num > 3 {
			// 		slice32 = cord2[num-4 : num-2]
			// 		slice3 = append(slice3, slice32...)
			// 	}
			// 	type3L = "flat"
			// 	story = append(story, 2)
			// }

			/* // L点から伸ばした線が対向する辺に直交する場合
			// 直行する点で四角形と５角形に分割する
			var num int
			nodHex := len(cord2)
			// L点の直交条件．対向する辺との交点の角度制限を確認する．
			var int1stX float64
			var int1stY float64
			var int2ndX float64
			var int2ndY float64
			// 交点が対向する辺の上にあるか確認する
			var lnincl1 bool
			var lnincl2 bool
			// L1点の頂点番号を確認する
			for LRkey := range order {
				log.Println("LRkey=", LRkey) // Ctrl+/
				if LRkey == "L1" {
					num = order[LRkey]      // 頂点番号
					log.Println("num", num) // Ctrl+/

					// 直交する辺は．L点と1つ前の点で結ばれる線分
					// 直交する辺の座標ペア
					chokuCord1 := make([][]float64, 2)
					numP1 := (num - 1 + nodHex) % nodHex
					chokuCord1[0] = cord2[num]
					chokuCord1[1] = cord2[numP1]
					// 対向する辺は，L点から２つ目と３つ目の点で結ばれる線分
					// 対向する辺の座標ペア
					taikoCord1 := make([][]float64, 2)
					numP2 := (num + 2) % nodHex
					taikoCord1[0] = cord2[numP2]
					numP3 := (num + 3) % nodHex
					taikoCord1[1] = cord2[numP3]
					// 直交する直線aと対向する辺との直交条件を確認する
					intX, intY, theta := OrthoAngle(chokuCord1, taikoCord1) // OrthoAngleの戻り値X･Yが逆転
					int1stX = intX
					log.Println("int1stX=", int1stX) // Ctrl+/
					int1stY = intY
					log.Println("int1stY=", int1stY) // Ctrl+/
					// 交差角度が制限範囲内でない場合は処理を中断する
					log.Println("theta=", theta) // Ctrl+/
					if theta < 45 || theta > 135 {
						// TODO:折れ曲がりの切妻屋根
						return
					}
					lnchk1 := PosLine2(chokuCord1, taikoCord1)
					if lnchk1 < 0 {
						lnincl1 = true
					}

					// もう一方の直交する辺は．L点と1つ次の点で結ばれる線分
					// 直交する辺の座標ペア
					chokuCord2 := make([][]float64, 2)
					numN1 := (num + 1) % nodHex
					chokuCord2[0] = cord2[num]
					chokuCord2[1] = cord2[numN1]
					// もう一方の対向する辺は，L点から３つ目と４つ目の点で結ばれる線分
					// 対向する辺の座標ペア
					taikoCord2 := make([][]float64, 2)
					numN3 := (num + 3) % nodHex
					taikoCord2[0] = cord2[numN3]
					numN4 := (num + 4) % nodHex
					taikoCord2[1] = cord2[numN4]
					// 直交する直線bと対向する辺との直交条件を確認する
					int2X, int2Y, theta2 := OrthoAngle(chokuCord2, taikoCord2) // OrthoAngleの戻り値X･Yが逆転
					int2ndX = int2X
					log.Println("int2ndX=", int2ndX) // Ctrl+/
					int2ndY = int2Y
					log.Println("int2ndY=", int2ndY) // Ctrl+/
					// 交差角度が制限範囲内でない場合は処理を中断する
					log.Println("theta2=", theta2) // Ctrl+/
					if theta < 45 || theta > 135 {
						// TODO:折れ曲がりの切妻屋根
						return
					}
					lnchk2 := PosLine2(chokuCord2, taikoCord2)
					if lnchk2 < 0 {
						lnincl2 = true
					}

				}
				// L点から対向する二辺までの距離を比較する
				// L点の座標
				log.Println("X座標", cord2[num][0]) // Ctrl+/
				log.Println("Y座標", cord2[num][1]) // Ctrl+/
				// 交点１までの距離
				divLa := DistVerts(cord2[num][0], cord2[num][1], int1stX, int1stY)
				log.Println("divLa=", divLa)
				// 交点２までの距離
				divLb := DistVerts(cord2[num][0], cord2[num][1], int2ndX, int2ndY)
				log.Println("divLb=", divLb)
				var D1splt float64
				var D2splt float64

				if lnincl1 {
					// 交点１から隣り合うR点までの最短距離を求める
					// D1点（交点１）と１つ前のR2点までの距離
					D1toR2 := DistVerts(int1stX, int1stY, cord2[(num+2)%6][0], cord2[(num+2)%6][1])
					// D1点（交点１）と１つ後ろのR3点までの距離
					D1toR3 := DistVerts(int1stX, int1stY, cord2[(num+3)%6][0], cord2[(num+3)%6][1])
					if D1toR2 < D1toR3 {
						D1splt = D1toR2
					} else {
						D1splt = D1toR3
					}
				}
				if lnincl2 {
					// 交点２から隣り合うR点までの最短距離を求める
					// D2点（交点２）と１つ前のR3点までの距離
					D2toR3 := DistVerts(int2ndX, int2ndY, cord2[(num+3)%6][0], cord2[(num+3)%6][1])
					// D2点（交点２）と１つ後ろのR4点までの距離
					D2toR4 := DistVerts(int2ndX, int2ndY, cord2[(num+4)%6][0], cord2[(num+4)%6][1])
					if D2toR3 < D2toR4 {
						D2splt = D2toR3
					} else {
						D2splt = D2toR4
					}
				}

				// 交点からR点までの最短距離を比較する
				if (D1splt > D2splt) && D1splt >= 1.0 {
					if divLa >= 1.0 {
						// 分割点はD1点（交点１）
						d1 := []float64{int1stX, int1stY}
						log.Println("d1=", d1)
						// 座標値のリストにD1点の座標値を追加する
						cord2 = append(cord2, d1)
						log.Println(cord2) // Ctrl+/
						// 頂点並びの辞書に分割点を追加する
						d1Num := nodHex
						order["D1"] = d1Num
						log.Println("line_a", order) // Ctrl+/

						// 四角形D1-L1-R1-R2(1-2頂点，3-4頂点が妻面)
						rect1name = []string{"D1", "L1", "R1", "R2"}
						// 四角形R5-D1-R3-R4(1-2頂点，3-4頂点が妻面)
						rect2name = []string{"R5", "D1", "R3", "R4"}
					}
				} else if (D1splt < D2splt) && D2splt >= 1.0 {
					if divLb >= 1.0 {
						// 分割点はD2点（交点２）
						d2 := []float64{int2ndX, int2ndY}
						log.Println("d2=", d2)
						// 座標値のリストにD2点の座標値を追加する
						cord2 = append(cord2, d2)
						log.Println(cord2) // Ctrl+/
						// 頂点並びの辞書に分割点を追加する
						d2Num := nodHex
						order["D2"] = d2Num
						log.Println("line_b", order) // Ctrl+/

						// 四角形D2-R1-R2-R3(1-2頂点，3-4頂点が妻面)
						rect1name = []string{"D2", "R1", "R2", "R3"}
						// 四角形L1-D2-R4-R5(1-2頂点，3-4頂点が妻面)
						rect2name = []string{"L1", "D2", "R4", "R5"}
					}
				}
			} */

			/* // L点が1つの場合はL点の反対側に流れる片流れ屋根を掛ける
			log.Println("lrtxt include L")
			log.Println("order=", order)
			log.Println("order[\"L1\"]=", order["L1"])
			L1 := order["L1"]

			var slice1t [][]float64
			if L1 < 3 {
				slice11 := cord2[L1+4:]
				slice12 := cord2[:L1+1]
				slice1t = append(slice1t, slice11...)
				slice1t = append(slice1t, slice12...)
			} else if L1 > 2 {
				slice1t = cord2[(L1+4)%7 : L1+1]
			}
			slice1t1 := slice1t[2:]
			slice1t2 := slice1t[:2]
			slice1 = append(slice1, slice1t1...)
			slice1 = append(slice1, slice1t2...)
			log.Println("slice1=", slice1)

			if L1 < 4 {
				slice2 = cord2[L1 : L1+4]
			} else if L1 > 3 {
				slice21 := cord2[L1:]
				slice22 := cord2[:(L1+4)%7]
				slice2 = append(slice2, slice21...)
				slice2 = append(slice2, slice22...)
			}
			log.Println("slice2=", slice2)

			slice31 := cord2[L1 : L1+1]
			slice32 := cord2[(L1+3)%7 : (L1+3)%7+1]
			slice33 := cord2[(L1+4)%7 : (L1+4)%7+1]
			slice3 = append(slice3, slice31...)
			slice3 = append(slice3, slice32...)
			slice3 = append(slice3, slice33...)
			log.Println("slice3=", slice3)

			type1L = "kata1"
			type2L = "kata2"
			type3L = "heptri"
			for i := 0; i < 3; i++ {
				story = append(story, 2)
			} */
		}

	} else if !strings.Contains(lrtxt, "L") {
		chk = false
	}

	log.Println("type1L=", type1L)
	log.Println("type2L=", type2L)
	log.Println("type3L=", type3L)

	return slice1, slice2, slice3, type1L, type2L, type3L, story, chk
}
