package internal

// AreaRGB は用途地域ごとに着色する色のRGB値を設定する
func AreaRGB(num int) (rgb []float64) {
	if num == 1 {
		rgb = append(rgb, 0)
		rgb = append(rgb, 0.647)
		rgb = append(rgb, 0.408)
		rgb = append(rgb, 1)
	} else if num == 2 {
		rgb = append(rgb, 0.467)
		rgb = append(rgb, 0.718)
		rgb = append(rgb, 0.620)
		rgb = append(rgb, 1)
	} else if num == 3 {
		rgb = append(rgb, 0.314)
		rgb = append(rgb, 0.686)
		rgb = append(rgb, 0.424)
		rgb = append(rgb, 1)
	} else if num == 4 {
		rgb = append(rgb, 0.745)
		rgb = append(rgb, 0.804)
		rgb = append(rgb, 0)
		rgb = append(rgb, 1)
	} else if num == 5 {
		rgb = append(rgb, 0)
		rgb = append(rgb, 0.937)
		rgb = append(rgb, 0.267)
		rgb = append(rgb, 1)
	} else if num == 6 {
		rgb = append(rgb, 0.976)
		rgb = append(rgb, 0.698)
		rgb = append(rgb, 0)
		rgb = append(rgb, 1)
	} else if num == 7 {
		rgb = append(rgb, 0.933)
		rgb = append(rgb, 0.498)
		rgb = append(rgb, 0)
		rgb = append(rgb, 1)
	} else if num == 8 {
		rgb = append(rgb, 0.941)
		rgb = append(rgb, 0.569)
		rgb = append(rgb, 0.604)
		rgb = append(rgb, 1)
	} else if num == 9 {
		rgb = append(rgb, 0.910)
		rgb = append(rgb, 0.345)
		rgb = append(rgb, 0.522)
		rgb = append(rgb, 1)
	} else if num == 10 {
		rgb = append(rgb, 0.820)
		rgb = append(rgb, 0.741)
		rgb = append(rgb, 0.851)
		rgb = append(rgb, 1)
	} else if num == 11 {
		rgb = append(rgb, 0.749)
		rgb = append(rgb, 0.886)
		rgb = append(rgb, 0.906)
		rgb = append(rgb, 1)
	} else if num == 12 {
		rgb = append(rgb, 0.314)
		rgb = append(rgb, 0.420)
		rgb = append(rgb, 0.678)
		rgb = append(rgb, 1)
	}

	return rgb
}
