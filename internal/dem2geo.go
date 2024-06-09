package internal

import (
	"bufio"
	"log"
	"math"
	"os"
	"stera/pkg"
	"strings"
)

type Zahyo struct {
	x float64
	y float64
	z float64
}

// TinMeshは標高メッシュデータから地形モデリングのための座標データを作成する
func TinMesh() (x_matrix, y_matrix, z_matrix [][]float64, x_len, y_len float64, x_dot, y_dot int64, x_max, x_min, y_max, y_min, z_max, z_min float64) {
	// 変数の宣言
	var kansan = 0.0254 // ｍをインチ換算
	var x_zahyo []float64
	var y_zahyo []float64
	var hyoukou []float64

	// データ数のカウント
	fl, er := os.Open("C:/data/dem_data.txt")
	_, l, _, _ := pkg.FileCount(fl)
	var counter = l
	// log.Println("counter=", counter)
	if er != nil {
		log.Fatal(er)
	}
	defer fl.Close()

	// 標高メッシュデータの読み込み
	fp, er := os.Open("C:/data/dem_data.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer fp.Close()
	// 一行ずつデータを読み込む
	scanner := bufio.NewScanner(fp)
	zahyo := []Zahyo{}
	var lines = 0
	for scanner.Scan() {
		// ここで一行ずつ処理
		gStr := scanner.Text()
		// 右端の「,」を削除，「,」がない行末でもエラーにならない
		gStr = strings.TrimRight(gStr, ",")
		// GeoJson構造体の変数stcDataを宣言
		geo := pkg.ParseDEM(gStr)
		// log.Println("geo=", geo)

		// データを構造体のスライスに読み込む
		xyz := Zahyo{x: geo[0], y: geo[1], z: geo[2]}
		zahyo = append(zahyo, xyz)
		lines += 1
	}
	// log.Println("lines=", lines)

	// 各座標データを各スライスに読み込む
	for i := 0; i < lines; i++ {
		x_zahyo = append(x_zahyo, zahyo[i].x/kansan) // X座標の読み込み
		// log.Println("x_zahyo=", x_zahyo[i])
		y_zahyo = append(y_zahyo, zahyo[i].y/kansan) // Y座標の読み込み
		// log.Println("y_zahyo=", y_zahyo[i])
		hyoukou = append(hyoukou, zahyo[i].z/kansan) // 地盤高データの読み込み
		// log.Println("hyoukou=", hyoukou[i])
	}

	// 配列（マトリックス）の大きさを決める処理
	x_max = x_zahyo[0] // X座標の最大値の初期化
	y_max = y_zahyo[0] // Y座標の最大値の初期化
	x_min = x_zahyo[0] // X座標の最小値の初期化
	y_min = y_zahyo[0] // Y座標の最小値の初期化

	// X座標・Y座標の最大値と最小値を求める
	// 標高データの最大値と最小値を求める
	for i := 0; i < counter; i++ {
		if x_max < x_zahyo[i] {
			x_max = x_zahyo[i] // X座標の最大値の更新
		}
		if y_max < y_zahyo[i] {
			y_max = y_zahyo[i] // Y座標の最大値の更新
		}
		if x_min > x_zahyo[i] {
			x_min = x_zahyo[i] // X座標の最小値の更新
		}
		if y_min > y_zahyo[i] {
			y_min = y_zahyo[i] // Y座標の最小値の更新
		}
		if z_max < hyoukou[i] {
			z_max = hyoukou[i] // Z座標の最大値の更新
		}
		if z_min > hyoukou[i] {
			z_min = hyoukou[i] // Z座標の最小値の更新
		}
	}
	// log.Println("x_max=", x_max)
	// log.Println("y_max=", y_max)
	// log.Println("x_min=", x_min)
	// log.Println("y_min=", y_min)

	// マトリックスの大きさを決めるためにX(東西)方向とY(南北)方向の大きさを求める
	x_len = math.Abs(x_max - x_min) // X(東西)方向の幅
	y_len = math.Abs(y_max - y_min) // Y(南北)方向の高さ

	// var y_count = 0       // Y(南北)方向マス数

	// X(東西)方向のマス目の大きさ(最小値)を求める
	// X座標の間隔は上(北)の方は狭く，下(南)は広く，同じ行(段)は等間隔になっている
	// X(東西)方向のマス目の大きさ
	x_dist := math.Abs(x_zahyo[1] - x_zahyo[0])

	var x_dist_temp float64
	for i := 2; i < counter; i++ {
		// X座標値は東／右に行くほど値が大きくなる
		// X座標値が小さくなっている個所で改行（段）されたと判断できる
		if x_zahyo[i] > x_zahyo[i-1] {
			// X(東西)方向のマス目の計算過程での大きさ
			x_dist_temp = math.Abs(x_zahyo[i] - x_zahyo[i-1])
			if x_dist > x_dist_temp {
				x_dist = x_dist_temp
			}
		}
	}
	log.Println("x_dist=", x_dist)

	// X(東西)方向のデータ数を求める
	// X(東西)方向マス数最大値の初期化
	x_countMax := int64(math.Round(x_len / x_dist))
	log.Println("x_countMax=", x_countMax)
	// X(東西)方向の頂点データ数
	x_dot = x_countMax + 1
	log.Println("x_dot=", x_dot)

	// Y(南北)方向のマス目の大きさを求める
	var y_zahyo_accu = 0.0
	var y_cnt_tmp = 0

	// 各行（段）のY座標平均値を格納する配列の準備
	var y_zahyo_arr []float64
	var row_cnt = 0
	// var row_cnt = 1
	// for i := 1; i < (counter - 1); i++ {
	for i := 1; i < counter; i++ {
		// Y座標値を累積して平均値を求める
		// Y座標値は右肩上がり（東／右に行くほど値が大きくなる）
		if x_zahyo[i] > x_zahyo[i-1] {
			y_zahyo_accu = y_zahyo_accu + y_zahyo[i-1]
			y_cnt_tmp = y_cnt_tmp + 1
		} else {
			y_zahyo_accu = y_zahyo_accu + y_zahyo[i-1]
			y_cnt_tmp = y_cnt_tmp + 1
			y_zahyo_arr = append(y_zahyo_arr, y_zahyo_accu/float64(y_cnt_tmp))
			// Y座標平均値用配列にデータを追加したのでカウンターを１増やす
			row_cnt = row_cnt + 1
			y_zahyo_accu = 0.0
			y_cnt_tmp = 0
		}
		// if i == (counter - 1) {
		if i == counter {
			y_zahyo_accu = y_zahyo_accu + y_zahyo[i]
			y_cnt_tmp = y_cnt_tmp + 1
			y_zahyo_arr = append(y_zahyo_arr, y_zahyo_accu/float64(y_cnt_tmp))
		}
	}

	var y_dist_temp = 0.0 // Y(南北)方向の計算過程での高さ
	var y_dist_cnt = 0    // Y(南北)方向の計算過程での累積回数
	var y_diff_temp float64
	for j := 1; j < row_cnt; j++ {
		// Y(南北)方向のマス目の計算過程での大きさ
		y_diff_temp = math.Abs(y_zahyo_arr[j] - y_zahyo_arr[j-1])
		// log.Println("y_diff_temp=", y_diff_temp)
		// 空白となっている行（段）は除外する
		// Y座標の間隔がX座標の間隔の1.5倍未満の場合は空白の行（段）ではないと判断する
		if y_diff_temp < x_dist*1.5 {
			y_dist_temp = y_dist_temp + y_diff_temp
			y_dist_cnt = y_dist_cnt + 1
		}
	}

	// Y(南北)方向のマス目の大きさ
	y_dist := y_dist_temp / float64(y_dist_cnt)
	// log.Println("y_dist=", y_dist)

	// マトリックスのY(南北)方向の大きさをY座標の傾きの分だけ修正する
	var y_zahyo_ini = 0.0
	// var y_zahyo_accu2 = 0.0
	var y_cnt_tmp2 = 0
	var y_tilt_dif = 0.0
	var y_tilt = 0.0

	// Y座標値は右肩上がり（東／右に行くほど値が大きくなる）
	for i := 1; i < counter; i++ {
		// X座標値が小さくなっている個所で改行（段）されたと判断できる
		if x_zahyo[i] < x_zahyo[i-1] {
			// １番始めの行（段）で改行された場合
			if y_cnt_tmp2 == 0 {
				y_zahyo_ini = y_zahyo[0]
				y_tilt = y_zahyo[i-1] - y_zahyo_ini
				y_zahyo_ini = y_zahyo[i]
				y_cnt_tmp2 = 1
			} else {
				// Y座標値の差分の最大値を求める
				y_tilt_dif = y_zahyo[i-1] - y_zahyo_ini
				if y_tilt < y_tilt_dif {
					y_tilt = y_tilt_dif
				}
				y_zahyo_ini = y_zahyo[i]
				y_cnt_tmp2 = y_cnt_tmp2 + 1
			}
		}
	}
	// log.Println("y_cnt_tmp2=", y_cnt_tmp2)

	// y_tilt := y_zahyo_accu2 / (float64(y_cnt_tmp2) - 1)
	log.Println("y_tilt=", y_tilt)
	y_len = y_len - y_tilt
	// log.Println("y_len=", y_len)

	// Y(南北)方向のデータ数を求める
	y_countmax := math.Round(y_len / y_dist)
	// log.Println("y_countmax=", y_countmax)
	// Y(南北)方向の頂点データ数
	y_dot = int64(y_countmax) + 1
	// log.Println("y_dot=", y_dot)

	// マトリックスの大きさ（２次元配列）を決める
	// var x_matrix [][]float64 // X座標
	x_matrix = make([][]float64, y_dot)
	for k := int64(0); k < y_dot; k++ {
		x_matrix[k] = make([]float64, x_dot)
	}
	// var y_matrix [][]float64 // Y座標
	y_matrix = make([][]float64, y_dot)
	for k := int64(0); k < y_dot; k++ {
		y_matrix[k] = make([]float64, x_dot)
	}
	// var z_matrix [][]float64 // Z座標
	z_matrix = make([][]float64, y_dot)
	for k := int64(0); k < y_dot; k++ {
		z_matrix[k] = make([]float64, x_dot)
	}
	// log.Println("x_matrix=", x_matrix)
	// log.Println("y_matrix=", y_matrix)
	// log.Println("z_matrix=", z_matrix)

	// マトリックスにデータを割り付ける
	col_num := int64(0) // 桁の番号，左端が１
	row_num := int64(1) // 行（段）の番号，下端が１
	// row_num := int64(0) // 行（段）の番号，下端が１
	blunk := int64(0)   // 空白部分の桁数，初期値は０
	var x_row []float64 // X座標の配列
	var y_row []float64 // Y座標の配列
	var z_row []float64 // Z座標の配列

	// Y座標値のデータは南から北（下から上）へ並ぶ
	// X座標値のデータは西から東（左から右）へ並ぶ
	for i := 0; i < counter; i++ {
		if i == 0 {
			// １番目のデータ
			// 各行（段）の１番目のデータの配列（配置場所）を決める
			// 左端からの距離は許容範囲内か？
			if (x_zahyo[i] - x_min) < x_dist*0.8 {
				x_row = append(x_row, x_zahyo[i])
				y_row = append(y_row, y_zahyo[i])
				z_row = append(z_row, hyoukou[i])
				// log.Println("z_row=", z_row)
				// log.Println("hyoukou=", hyoukou[i])
				col_num = col_num + 1
				// log.Println("i=", i)
			} else {
				// 左端との間に空白データがある
				// 空白部分の桁数を求める
				blunk = int64(math.Round((x_zahyo[i] - x_min) / x_dist))
				// log.Println("blunk(1)=", blunk)
				// 空白部分を標高０としてダミーで埋める
				for j := int64(0); j < blunk; j++ {
					x_row = append(x_row, x_min+x_dist*float64(j))
					y_row = append(y_row, y_zahyo[i])
					z_row = append(z_row, z_min)
					col_num = col_num + 1
				}
				// col_num = col_num + blunk
				// 空白データの後ろにデータを追加する
				x_row = append(x_row, x_zahyo[i])
				y_row = append(y_row, y_zahyo[i])
				z_row = append(z_row, hyoukou[i])
				col_num = col_num + 1
				// log.Println("i=", i)
			}
			// log.Println("row_num=", row_num)
			// log.Println("col_num=", col_num)

		} else if i > 0 && (x_zahyo[i] > x_zahyo[i-1]) {
			// 各行（段）の２番目以降のデータの配列（配置場所）を決める
			// １番目のデータからの距離は許容範囲内か？
			if (x_zahyo[i] - x_zahyo[i-1]) < x_dist*1.1 {
				x_row = append(x_row, x_zahyo[i])
				y_row = append(y_row, y_zahyo[i])
				z_row = append(z_row, hyoukou[i])
				col_num = col_num + 1
				// log.Println("i=", i)
			} else {
				blunk = int64(math.Round(x_zahyo[i]-x_zahyo[i-1]) / x_dist)
				// log.Println("x_zahyo[i]=", x_zahyo[i])
				// log.Println("x_zahyo[i-1]=", x_zahyo[i-1])
				// log.Println("blunk(2)=", blunk)
				// 空白部分を標高０としてダミーで埋める
				for j := int64(1); j < blunk; j++ {
					x_row = append(x_row, x_zahyo[i-1]+x_dist*float64(j))
					y_row = append(y_row, y_zahyo[i])
					z_row = append(z_row, z_min)
					col_num = col_num + 1
				}
				// col_num = col_num + blunk - 1
				// 空白データの後ろにデータを追加する
				x_row = append(x_row, x_zahyo[i])
				y_row = append(y_row, y_zahyo[i])
				z_row = append(z_row, hyoukou[i])
				col_num = col_num + 1
				// log.Println("i=", i)
			}
			// log.Println("row_num=", row_num)
			// log.Println("col_num=", col_num)
		}

		// 	改行されるまでデータ並びを配列に追加する
		if i > 0 && x_zahyo[i] < x_zahyo[i-1] {
			// 改行されたので配列をマトリックスに追加する
			// 改行される前に空白部分がないか調べる
			if col_num < x_dot {
				blunk = x_dot - col_num
				// log.Println("blunk(3)=", blunk)

				// 空白部分を標高０としてダミーで埋める
				for j := int64(1); j <= blunk; j++ {
					x_row = append(x_row, x_zahyo[i-1]+x_dist*float64(j))
					y_row = append(y_row, y_zahyo[i-1])
					z_row = append(z_row, z_min)
					col_num = col_num + 1
				}
				// col_num = col_num + blunk
			}
			// log.Println("col_num=", col_num)
			pre_yzahyo := y_row[0] // １つ前の行（段）の１番目のデータのY座標
			// log.Println("pre_yzahyo=", pre_yzahyo)

			// x_matrix = append(x_matrix, x_row)
			for x := int64(0); x < x_dot; x++ {
				x_matrix[row_num-1][x] = x_row[x]
			}
			// log.Println("x_matrix=", x_matrix)
			// log.Println("x_row=", len(x_row))
			// log.Println(x_row)
			x_row = x_row[:0]
			// y_matrix = append(y_matrix, y_row)
			for y := int64(0); y < x_dot; y++ {
				y_matrix[row_num-1][y] = y_row[y]
			}
			// log.Println("y_matrix=", y_matrix)
			// log.Println("y_row=", len(y_row))
			// log.Println(y_row)
			y_row = y_row[:0]
			// z_matrix = append(z_matrix, z_row)
			for z := int64(0); z < x_dot; z++ {
				z_matrix[row_num-1][z] = z_row[z]
			}
			// log.Println("z_matrix=", z_matrix)
			// log.Println("z_row=", len(z_row))
			// log.Println(z_row)
			z_row = z_row[:0]

			// 改行されているので行(row)番号が変更になる
			row_num = row_num + 1

			// 帯状空白部分がある場合は標高０としてダミーで埋める
			if (y_zahyo[i] - pre_yzahyo) > y_dist*1.8 {
				y_blk := int64(math.Round((y_zahyo[i] - pre_yzahyo) / y_dist))
				for k := int64(1); k < y_blk; k++ {
					for j := int64(0); j < x_dot; j++ {
						x_row = append(x_row, x_min+x_dist*float64(j))
						y_row = append(y_row, y_zahyo[i]+y_dist*float64(k))
						z_row = append(z_row, z_min)
					}
				}
				// x_matrix = append(x_matrix, x_row)
				for x := int64(0); x < x_dot; x++ {
					x_matrix[row_num-1][x] = x_row[x]
				}
				// log.Println("x_matrix=", x_matrix)
				// log.Println("x_row=", len(x_row))
				// log.Println(x_row)
				x_row = x_row[:0]
				y_matrix = append(y_matrix, y_row)
				for y := int64(0); y < x_dot; y++ {
					y_matrix[row_num-1][y] = y_row[y]
				}
				// log.Println("y_matrix=", y_matrix)
				// log.Println("y_row=", len(y_row))
				// log.Println(y_row)
				y_row = y_row[:0]
				// z_matrix = append(z_matrix, z_row)
				for z := int64(0); z < x_dot; z++ {
					z_matrix[row_num-1][z] = z_row[z]
				}
				// log.Println("z_matrix=", z_matrix)
				// log.Println("z_row=", len(z_row))
				// log.Println(z_row)
				z_row = z_row[:0]

				// 改行されているので行(row)番号が変更になる
				row_num = row_num + y_blk - 1
			}
			// 改行されているので桁(col)番号を１に戻す
			col_num = 0

			// X座標が東から西へ折り返した最初のデータ
			// 各行（段）の１番目のデータの配列（配置場所）を決める
			// 左端からの距離は許容範囲内か？
			if (x_zahyo[i] - x_min) < x_dist*0.8 {
				x_row = append(x_row, x_zahyo[i])
				y_row = append(y_row, y_zahyo[i])
				z_row = append(z_row, hyoukou[i])
				col_num = col_num + 1
				// log.Println("i=", i)
			} else {
				// 左端との間に空白データがある
				// 空白部分の桁数を求める
				blunk = int64(math.Round((x_zahyo[i] - x_min) / x_dist))
				// log.Println("blunk(1)=", blunk)
				// 空白部分を標高０としてダミーで埋める
				for j := int64(0); j < blunk; j++ {
					x_row = append(x_row, x_min+x_dist*float64(j))
					y_row = append(y_row, y_zahyo[i])
					z_row = append(z_row, z_min)
					col_num = col_num + 1
				}
				// col_num = col_num + blunk
				// 空白データの後ろにデータを追加する
				x_row = append(x_row, x_zahyo[i])
				y_row = append(y_row, y_zahyo[i])
				z_row = append(z_row, hyoukou[i])
				col_num = col_num + 1
				// log.Println("i=", i)
			}
			// log.Println("row_num=", row_num)
			// log.Println("col_num=", col_num)
		}

		// 全ての座標データが追加されたがrow_numがy_dotに到達していない場合
		if i == (counter - 1) {
			// 最後の行（段）のデータが右端（東端）に到達していない
			if row_num <= y_dot {
				if col_num < x_dot {
					blunk = x_dot - col_num
					// log.Println("blunk(3)=", blunk)

					// 空白部分を標高０としてダミーで埋める
					for j := int64(0); j < blunk; j++ {
						x_row = append(x_row, x_zahyo[i]+x_dist*float64(j))
						y_row = append(y_row, y_zahyo[i])
						z_row = append(z_row, z_min)
					}
					col_num = col_num + blunk
					// log.Println("col_num=", col_num)
					// x_matrix = append(x_matrix, x_row)
					for x := int64(0); x < x_dot; x++ {
						x_matrix[row_num-1][x] = x_row[x]
					}
					// log.Println("x_matrix=", x_matrix)
					// log.Println("x_row=", len(x_row))
					// log.Println(x_row)
					x_row = x_row[:0]
					y_matrix = append(y_matrix, y_row)
					for y := int64(0); y < x_dot; y++ {
						y_matrix[row_num-1][y] = y_row[y]
					}
					// log.Println("y_matrix=", y_matrix)
					// log.Println("y_row=", len(y_row))
					// log.Println(y_row)
					y_row = y_row[:0]
					// z_matrix = append(z_matrix, z_row)
					for z := int64(0); z < x_dot; z++ {
						z_matrix[row_num-1][z] = z_row[z]
					}
					// log.Println("z_matrix=", z_matrix)
					// log.Println("z_row=", len(z_row))
					// log.Println(z_row)
					z_row = z_row[:0]
				}

				// 最後の行（段）が上端（北端）に到達していない
				y_rem := y_dot - row_num
				for i := int64(0); i < y_rem; i++ {
					for j := int64(0); j < x_dot; j++ {
						x_row = append(x_row, x_min+x_dist*float64(j))
						y_row = append(y_row, y_zahyo[i]+y_dist*float64(i))
						z_row = append(z_row, z_min)
					}
					row_num = row_num + 1
					// log.Println("row_num=", row_num)
					// x_matrix = append(x_matrix, x_row)
					for x := int64(0); x < x_dot; x++ {
						x_matrix[row_num-1][x] = x_row[x]
					}
					// log.Println("x_row=", len(x_row))
					// log.Println(x_row)
					// x_row = x_row[:0]
					// y_matrix = append(y_matrix, y_row)
					for y := int64(0); y < x_dot; y++ {
						y_matrix[row_num-1][y] = y_row[y]
					}
					// log.Println("y_matrix=", y_matrix)
					// log.Println("y_row=", len(y_row))
					// log.Println(y_row)
					y_row = y_row[:0]
					// z_matrix = append(z_matrix, z_row)
					for z := int64(0); z < x_dot; z++ {
						z_matrix[row_num-1][z] = z_row[z]
					}
					// log.Println("z_matrix=", z_matrix)
					// log.Println("z_row=", len(z_row))
					// log.Println(z_row)
					z_row = z_row[:0]
				}
			} else {
				// 最後の行（段）の最後のデータが右端（東端）に到達している
				// x_matrix = append(x_matrix, x_row)
				for x := int64(0); x < x_dot; x++ {
					x_matrix[row_num-1][x] = x_row[x]
				}
				// log.Println("x_matrix=", x_matrix)
				// log.Println("x_row=", len(x_row))
				// log.Println(x_row)
				x_row = x_row[:0]
				y_matrix = append(y_matrix, y_row)
				for y := int64(0); y < x_dot; y++ {
					y_matrix[row_num-1][y] = y_row[y]
				}
				// log.Println("y_matrix=", y_matrix)
				// log.Println("y_row=", len(y_row))
				// log.Println(y_row)
				y_row = y_row[:0]
				// z_matrix = append(z_matrix, z_row)
				for z := int64(0); z < x_dot; z++ {
					z_matrix[row_num-1][z] = z_row[z]
				}
				// log.Println("z_matrix=", z_matrix)
				// log.Println("z_row=", len(z_row))
				// log.Println(z_row)
				z_row = z_row[:0]
			}
		}
	}
	return x_matrix, y_matrix, z_matrix, x_len, y_len, x_dot, y_dot, x_max, x_min, y_max, y_min, z_max, z_min
}
