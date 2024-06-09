package pkg

import (
	"log"
	"strings"
)

// HeptaDiv は７角形を３つに分割して片流れ屋根を掛ける
func HeptaDiv(lrPtn []string, deg2 []float64, cord2 [][]float64, order map[string]int) (slice1 [][]float64,
	slice2 [][]float64, slice3 [][]float64, type1L, type2L, type3L string, story []int) {
	// L点が1つの場合と2つの場合で処理が異なる
	// L点を数える
	lrtxt := strings.Join(lrPtn, "")
	log.Println("lrtxt=", lrtxt)

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
			slice21 := cord2[L1+4:]
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
		// L点が2つの場合はL点の間のR点の数に応じて屋根の掛け方を変える
		if strings.Count(lrtxt, "L") == 2 {
			log.Println("lrtxt include Lx2")
			log.Println("order=", order)
			log.Println("order[\"L1\"]=", order["L1"])
			L1 := order["L1"]
			log.Println("order[\"L2\"]=", order["L2"])
			L2 := order["L2"]
			// L点に対向する最大角度のR点で四角形と５角形(四角形)に分割する
			log.Println("deg2=", deg2)
			if (L2-L1+7)%7 == 2 {
				var maxcnt int
				if L1 > L2 {
					if deg2[(L1-2+7)%7] > deg2[(L1-3+7)%7] {
						maxcnt = (L1 - 2 + 7) % 7
					} else {
						maxcnt = (L1 - 3 + 7) % 7
					}
				} else if L1 < L2 {
					if deg2[(L2+2)%7] > deg2[(L2+3)%7] {
						maxcnt = (L2 + 2) % 7
					} else {
						maxcnt = (L2 + 3) % 7
					}
				}
				log.Println("maxcnt=", maxcnt)

				if (maxcnt-L1+7)%7 == 4 {
					log.Println("四角形と５角形(四角形)-1")
					// 四角形・切妻屋根
					slice11 := cord2[L1 : L1+1]
					var slice12 [][]float64
					if maxcnt < 5 {
						slice12 = cord2[maxcnt : maxcnt+3]
					} else if maxcnt > 4 {
						slice12t1 := cord2[maxcnt:]
						slice12t2 := cord2[:(maxcnt+3)%7]
						slice12 = append(slice12, slice12t1...)
						slice12 = append(slice12, slice12t2...)
					}
					slice1 = append(slice1, slice11...)
					slice1 = append(slice1, slice12...)
					log.Println("slice1=", slice1)

					// 四角形・切妻屋根
					slice21 := cord2[maxcnt : maxcnt+1]
					var slice22 [][]float64
					if L1 < 6 {
						slice22 = cord2[L1 : L1+2]
					} else if L1 > 5 {
						slice22t1 := cord2[L1:]
						slice22t2 := cord2[:L2]
						slice22 = append(slice22, slice22t1...)
						slice22 = append(slice22, slice22t2...)
					}
					var slice23 [][]float64
					if L2 < 6 {
						slice23 = cord2[L2+1 : L2+2]
					} else if L2 > 5 {
						slice23 = cord2[:(L2+2)%7]
					}

					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, slice22...)
					slice2 = append(slice2, slice23...)
					log.Println("slice2=", slice2)
				}

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

				if (maxcnt-L1+7)%7 == 5 {
					log.Println("四角形と５角形(四角形)-3")
					// 四角形・切妻屋根
					slice11 := cord2[maxcnt : maxcnt+1]
					var slice12 [][]float64
					if L2 < 5 {
						slice12 = cord2[L2 : L2+3]
					} else if L2 > 4 {
						slice12t1 := cord2[L2:]
						slice12t2 := cord2[:(L2+3)%7]
						slice12 = append(slice12, slice12t1...)
						slice12 = append(slice12, slice12t2...)
					}
					slice1 = append(slice1, slice11...)
					slice1 = append(slice1, slice12...)
					log.Println("slice1=", slice1)

					// 四角形・切妻屋根
					slice21 := cord2[L2 : L2+1]
					var slice22 [][]float64
					if maxcnt < 6 {
						slice22 = cord2[maxcnt : maxcnt+2]
					} else if maxcnt > 5 {
						slice22t1 := cord2[maxcnt:]
						slice22t2 := cord2[:L1]
						slice22 = append(slice22, slice22t1...)
						slice22 = append(slice22, slice22t2...)
					}
					var slice23 [][]float64
					if L1 < 6 {
						slice23 = cord2[L1+1 : L1+2]
					} else if L1 > 5 {
						slice23t1 := cord2[L1:]
						slice23t2 := cord2[:(L1+2)%7]
						slice23 = append(slice23, slice23t1...)
						slice23 = append(slice23, slice23t2...)
					}

					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, slice22...)
					slice2 = append(slice2, slice23...)
					log.Println("slice2=", slice2)
				}

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
				slice31 := cord2[L1 : L1+1]
				slice32 := cord2[L2 : L2+1]
				slice33 := cord2[maxcnt : maxcnt+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				// 四角形slice1の内角をチェックする
				type1L = "kiri"
				_, degLst1, _ := TriVert(len(slice1), slice1)
				for _, v := range degLst1 {
					if v > 135 {
						type1L = "kata"
					}
				}

				// 四角形slice2の内角をチェックする
				type2L = "kiri"
				_, degLst2, _ := TriVert(len(slice2), slice2)
				for _, v := range degLst2 {
					if v > 135 {
						type2L = "kata"
					}
				}

				type3L = "flat"
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 1)

			} else if (L2-L1+7)%7 == 5 {
				var maxcnt int
				if L1 > L2 {
					if deg2[(L1+2)%7] > deg2[(L1+3)%7] {
						maxcnt = (L1 + 2) % 7
					} else {
						maxcnt = (L1 + 3) % 7
					}
				} else if L1 < L2 {
					if deg2[(L2-2+7)%7] > deg2[(L2-3+7)%7] {
						maxcnt = (L2 - 2 + 7) % 7
					} else {
						maxcnt = (L2 - 3 + 7) % 7
					}
				}
				log.Println("maxcnt=", maxcnt)

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

				if (maxcnt-L1+7)%7 == 2 {
					log.Println("四角形と５角形(四角形)-2")
					// 四角形・切妻屋根
					slice11 := cord2[L2 : L2+1]
					var slice12 [][]float64
					if maxcnt < 5 {
						slice12 = cord2[maxcnt : maxcnt+3]
					} else if maxcnt > 4 {
						slice12t1 := cord2[maxcnt:]
						slice12t2 := cord2[:(maxcnt+3)%7]
						slice12 = append(slice12, slice12t1...)
						slice12 = append(slice12, slice12t2...)
					}
					slice1 = append(slice1, slice11...)
					slice1 = append(slice1, slice12...)
					log.Println("slice1=", slice1)

					// 四角形・切妻屋根
					slice21 := cord2[maxcnt : maxcnt+1]
					var slice22 [][]float64
					if L2 < 6 {
						slice22 = cord2[L2 : L2+2]
					} else if L2 > 5 {
						slice22t1 := cord2[L2:]
						slice22t2 := cord2[:L1]
						slice22 = append(slice22, slice22t1...)
						slice22 = append(slice22, slice22t2...)
					}
					var slice23 [][]float64
					if L1 < 6 {
						slice23 = cord2[L1+1 : L1+2]
					} else if L1 > 5 {
						slice23 = cord2[:(L1+2)%7]
					}

					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, slice22...)
					slice2 = append(slice2, slice23...)
					log.Println("slice2=", slice2)
				}

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

				if (maxcnt-L1+7)%7 == 3 {
					log.Println("四角形と５角形(四角形)-4")
					// 四角形・切妻屋根
					slice11 := cord2[maxcnt : maxcnt+1]
					var slice12 [][]float64
					if L1 < 5 {
						slice12 = cord2[L1 : L1+3]
					} else if L1 > 4 {
						slice12t1 := cord2[L1:]
						slice12t2 := cord2[:(L1+3)%7]
						slice12 = append(slice12, slice12t1...)
						slice12 = append(slice12, slice12t2...)
					}
					slice1 = append(slice1, slice11...)
					slice1 = append(slice1, slice12...)
					log.Println("slice1=", slice1)

					// 四角形・切妻屋根
					slice21 := cord2[L1 : L1+1]
					var slice22 [][]float64
					if maxcnt < 6 {
						slice22 = cord2[maxcnt : maxcnt+2]
					} else if maxcnt > 5 {
						slice22t1 := cord2[maxcnt:]
						slice22t2 := cord2[:L2]
						slice22 = append(slice22, slice22t1...)
						slice22 = append(slice22, slice22t2...)
					}
					var slice23 [][]float64
					if L2 < 6 {
						slice23 = cord2[L2+1 : L2+2]
					} else if L2 > 5 {
						slice23 = cord2[:(L2+2)%7]
					}

					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, slice22...)
					slice2 = append(slice2, slice23...)
					log.Println("slice2=", slice2)
				}

				// var slice3 [][]float64
				slice31 := cord2[L1 : L1+1]
				slice32 := cord2[L2 : L2+1]
				slice33 := cord2[maxcnt : maxcnt+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				// 四角形slice1の内角をチェックする
				type1L = "kiri"
				_, degLst1, _ := TriVert(len(slice1), slice1)
				for _, v := range degLst1 {
					if v > 135 {
						type1L = "kata"
					}
				}

				// 四角形slice2の内角をチェックする
				type2L = "kiri"
				_, degLst2, _ := TriVert(len(slice2), slice2)
				for _, v := range degLst2 {
					if v > 135 {
						type2L = "kata"
					}
				}

				type3L = "flat"
				story = append(story, 1)
				story = append(story, 2)
				story = append(story, 1)

			} else if ((L1 < L2) && (L2-L1 == 3)) || ((L1 > L2) && (L1-L2 == 4)) {
				log.Println("二つの四角形に分割する-1")
				// 四角形・切妻屋根
				slice11 := cord2[L2 : L2+1]
				var slice12 [][]float64
				if L1 < 5 {
					slice12 = cord2[L1 : L1+3]
				} else if L1 > 4 {
					slice12t1 := cord2[L1:]
					slice12t2 := cord2[:(L1+3)%7]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				}
				slice1 = append(slice1, slice11...)
				slice1 = append(slice1, slice12...)
				log.Println("slice1=", slice1)

				// 四角形・切妻屋根
				if L2 < 3 {
					slice2 = cord2[L2 : L2+5]
				} else if L2 > 2 {
					slice21 := cord2[L2:]
					slice22 := cord2[:(L2+5)%7]
					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, slice22...)
				}

				// 作成した５角形にL点が含まれるとうまく屋根が掛からない
				// その場合は変形した四角形の切妻屋根とする
				n := len(slice2)
				ext, _, _ := TriVert(n, slice2)
				newslice := make([][]float64, 0, 5)
				for e, v := range ext {
					if v <= 0.0 {
						newslice = append(newslice, slice2[e])
					}
				}
				m := len(newslice)
				if m == 5 {
					newpenta, yane := PentaNode(deg2, newslice)
					slice2 = newpenta
					log.Println("slice2=", slice2)
					type2L = yane
				} else if m == 4 {
					slice2 = newslice
					type2L = "kiri"
				}
				log.Println("slice2=", slice2)

				// ダミー三角形
				midcnt := (L2 + 2) % 7

				slice31 := cord2[L1 : L1+1]
				slice32 := cord2[L2 : L2+1]
				slice33 := cord2[midcnt : midcnt+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				type3L = "flat"
				type1L = "kiri"
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 1)

			} else if ((L1 > L2) && (L1-L2 == 3)) || ((L1 < L2) && (L2-L1 == 4)) {
				log.Println("二つの四角形に分割する-2")
				// 四角形・切妻屋根
				slice11 := cord2[L1 : L1+1]
				var slice12 [][]float64
				if L2 < 5 {
					slice12 = cord2[L2 : L2+3]
				} else if L2 > 4 {
					slice12t1 := cord2[L2:]
					slice12t2 := cord2[:(L2+3)%7]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				}
				slice1 = append(slice1, slice11...)
				slice1 = append(slice1, slice12...)
				log.Println("slice1=", slice1)

				// 四角形・切妻屋根
				if L1 < 3 {
					slice2 = cord2[L1 : L1+5]
				} else if L1 > 2 {
					slice21 := cord2[L1:]
					slice22 := cord2[:(L1+5)%7]
					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, slice22...)
				}
				newpenta, _ := PentaNode(deg2, slice2)
				slice2 = newpenta
				log.Println("slice2=", slice2)

				// 作成した５角形にL点が含まれるとうまく屋根が掛からない
				// その場合は変形した四角形の切妻屋根とする
				n := len(slice2)
				ext, _, _ := TriVert(n, slice2)
				newslice := make([][]float64, 0, 5)
				for e, v := range ext {
					if v <= 0.0 {
						newslice = append(newslice, slice2[e])
					}
				}
				m := len(newslice)
				if m == 5 {
					newpenta, yane := PentaNode(deg2, newslice)
					slice2 = newpenta
					log.Println("slice2=", slice2)
					type2L = yane
				} else if m == 4 {
					slice2 = newslice
					type2L = "kiri"
				}
				log.Println("slice2=", slice2)

				// ダミー三角形
				midcnt := (L1 + 2) % 7

				slice31 := cord2[L1 : L1+1]
				slice32 := cord2[L2 : L2+1]
				slice33 := cord2[midcnt : midcnt+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				type3L = "flat"
				type1L = "kiri"
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 1)
			}

		} else if strings.Count(lrtxt, "L") == 1 {
			// L点が1つの場合はL点の反対側に流れる片流れ屋根を掛ける
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
			}
		}
	}
	log.Println("type1L=", type1L)
	log.Println("type2L=", type2L)
	log.Println("type3L=", type3L)

	return slice1, slice2, slice3, type1L, type2L, type3L, story
}
