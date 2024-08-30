package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"stera/internal"
	"stera/pkg"
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var (
	filename string
	demfile  string
	x_matrix [][]float64 // X座標
	y_matrix [][]float64 // Y座標
	z_matrix [][]float64 // Z座標
	x_len    float64
	y_len    float64
	x_dot    int
	y_dot    int
	x_max    float64
	x_min    float64
	y_max    float64
	y_min    float64
	z_max    float64
	z_min    float64
)

type MyMainWindow struct {
	*walk.MainWindow
	flabel      *walk.Label
	dlabel      *walk.Label
	d1_label    *walk.Label
	d2_label    *walk.Label
	d3_label    *walk.Label
	d4_label    *walk.Label
	d5_label    *walk.Label
	d6_label    *walk.Label
	d7_label    *walk.Label
	blabel      *walk.Label
	hlabel      *walk.Label
	klabel      *walk.Label
	olabel      *walk.Label
	le1         *walk.LineEdit
	le2         *walk.LineEdit
	le3         *walk.LineEdit
	le4         *walk.LineEdit
	le5         *walk.LineEdit
	le6         *walk.LineEdit
	le7         *walk.LineEdit
	le8         *walk.LineEdit
	le9         *walk.LineEdit
	le10        *walk.LineEdit
	t0label     *walk.Label
	t1label     *walk.Label
	t2label     *walk.Label
	edit7       *walk.TextEdit
	paintWidget *walk.CustomWidget

	progressBar *walk.ProgressBar

	path string
}

type CoorSys struct {
	Id   int
	Name string
}

func main() {
	// ログファイルを新規作成，追記，書き込み専用，パーミッションは読むだけ
	file, err := os.OpenFile("main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// ログの出力先を変更
	log.SetOutput(file)

	mw := &MyMainWindow{}

	MW := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "３次元都市モデル生成プログラム",
		MinSize:  Size{240, 320},
		Size:     Size{600, 800},
		MenuItems: []MenuItem{
			Menu{
				Text: "ファイル",
				Items: []MenuItem{
					Action{
						Text: "開く",
						Shortcut: Shortcut{
							Key:       walk.KeyO,
							Modifiers: walk.ModControl,
						},
						OnTriggered: mw.menuOpen,
					},
					Action{
						Text: "ファイルを分割",
						Shortcut: Shortcut{
							Key:       walk.KeyS,
							Modifiers: walk.ModControl,
						},
						OnTriggered: mw.pbClicked,
					},
					Separator{},
					Action{
						Text: "終了",
						Shortcut: Shortcut{
							Key:       walk.KeyQ,
							Modifiers: walk.ModControl,
						},
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "処理",
				Items: []MenuItem{
					Menu{
						Text: "普通建物",
						Items: []MenuItem{
							Action{
								Text:        "四角形分割",
								OnTriggered: mw.hutsuBuild,
							},
						},
					},
					Action{
						Text:        "堅ろう建物",
						OnTriggered: mw.menuOpen,
					},
					Action{
						Text:        "無壁舎",
						OnTriggered: mw.menuOpen,
					},
					Separator{},
					Action{
						Text:        "３角形分割",
						OnTriggered: mw.trimeshDev,
					},
				},
			},
			Menu{
				Text: "建物モデル",
				Items: []MenuItem{
					Action{
						Text:        "COLLADA",
						OnTriggered: mw.BuilDAE,
					},
					Action{
						Text:        "IFC",
						OnTriggered: mw.LandXML, // TODO mw.BuilIFC
					},
					Action{
						Text:        "DXF file",
						OnTriggered: mw.LandDXF, // TODO mw.BuilDXF
					},
					Action{
						Text:        "OBJ file",
						OnTriggered: mw.LandDXF, // TODO mw.BuilOBJ
					},
				},
			},
			Menu{
				Text: "地形モデル",
				Items: []MenuItem{
					Action{
						Text: "開く",
						Shortcut: Shortcut{
							Key:       walk.KeyH,
							Modifiers: walk.ModControl,
						},
						OnTriggered: mw.menuOpen2,
					},
					Action{
						Text:        "DEMデータ抽出",
						OnTriggered: mw.demMake,
					},
					Menu{
						Text: "地形モデル作成",
						Items: []MenuItem{
							Action{
								Text:        "COLLADA",
								OnTriggered: mw.LandDAE,
							},
							Action{
								Text:        "LandXML",
								OnTriggered: mw.LandXML,
							},
							Action{
								Text:        "DXF file",
								OnTriggered: mw.LandDXF,
							},
						},
					},
				},
			},
		},
		Layout: VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					GroupBox{
						Layout: HBox{},
						Children: []Widget{
							Label{
								Font:     Font{PointSize: 9},
								AssignTo: &mw.flabel,
								Text:     "ファイル名",
							},
							LineEdit{
								AssignTo: &mw.le1,
							},
							PushButton{
								Text:      "開く",
								OnClicked: mw.menuOpen,
							},
						},
					},
					VSpacer{},
					GroupBox{
						Layout: VBox{},
						Children: []Widget{
							Label{
								Font:     Font{PointSize: 9},
								AssignTo: &mw.dlabel,
								Text:     "建築物データ",
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d1_label,
										Text:     "            データ数    ",
									},
									LineEdit{
										AssignTo: &mw.le2,
										MaxSize:  Size{100, 0},
									},
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d2_label,
										Text:     "    件    ",
									},
								},
							},
							PushButton{
								Text:      "ファイル分割",
								OnClicked: mw.pbClicked,
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.hlabel,
										Text:     "普 通 建 物",
									},
									LineEdit{
										AssignTo: &mw.le3,
										MaxSize:  Size{100, 0},
									},
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d2_label,
										Text:     "    件    ",
									},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.klabel,
										Text:     "堅ろう建物",
									},
									LineEdit{
										AssignTo: &mw.le4,
										MaxSize:  Size{100, 0},
									},
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d2_label,
										Text:     "    件    ",
									},
								},
							}, GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.olabel,
										Text:     "無  壁  舎",
									},
									LineEdit{
										AssignTo: &mw.le5,
										MaxSize:  Size{100, 0},
									},
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d2_label,
										Text:     "    件    ",
									},
								},
							},
						},
					},
					GroupBox{
						Layout: VBox{},
						Children: []Widget{
							Label{
								Font:     Font{PointSize: 9},
								AssignTo: &mw.blabel,
								Text:     "建物モデル作成",
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.t1label,
										Text:     "COLLADAファイル",
									},
									PushButton{
										Text:      "作成",
										OnClicked: mw.BuilDAE,
									},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.t1label,
										Text:     "IFCファイル",
									},
									PushButton{
										Text:      "作成",
										OnClicked: mw.LandXML, // TODO mw.BuilIFC
									},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.t1label,
										Text:     "DXFファイル",
									},
									PushButton{
										Text:      "作成",
										OnClicked: mw.LandDXF, // TODO mw.BuilDXF
									},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.t1label,
										Text:     "OBJファイル",
									},
									PushButton{
										Text:      "作成",
										OnClicked: mw.LandDXF, // TODO mw.BuilOBJ
									},
								},
							},
						},
					},
				},
			},
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					GroupBox{
						Layout: VBox{},
						Children: []Widget{
							Label{
								Font:     Font{PointSize: 9},
								AssignTo: &mw.olabel,
								Text:     "標高メッシュデータ",
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d1_label,
										Text:     "          データ数    ",
									},
									LineEdit{
										AssignTo: &mw.le6,
										MaxSize:  Size{100, 0},
									},
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d3_label,
										Text:     "    個    ",
									},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d4_label,
										Text:     "東西方向の大きさ    ",
									},
									LineEdit{
										AssignTo: &mw.le7,
										MaxSize:  Size{100, 0},
									},
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d4_label,
										Text:     "    ｍ    ",
									},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d5_label,
										Text:     "南北方向の大きさ    ",
									},
									LineEdit{
										AssignTo: &mw.le8,
										MaxSize:  Size{100, 0},
									},
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d4_label,
										Text:     "    ｍ    ",
									},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d6_label,
										Text:     "東西方向のグリッド数    ",
									},
									LineEdit{
										AssignTo: &mw.le9,
										MaxSize:  Size{100, 0},
									},
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d3_label,
										Text:     "    個    ",
									},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d7_label,
										Text:     "南北方向のグリッド数    ",
									},
									LineEdit{
										AssignTo: &mw.le10,
										MaxSize:  Size{100, 0},
									},
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.d3_label,
										Text:     "    個    ",
									},
								},
							},
							PushButton{
								Text:      "DEMデータ抽出",
								OnClicked: mw.demMake,
							},
						},
					},
					GroupBox{
						Layout: VBox{},
						Children: []Widget{
							Label{
								Font:     Font{PointSize: 9},
								AssignTo: &mw.t0label,
								Text:     "地形モデル作成",
							},
							// GroupBox{
							// 	Layout: HBox{},
							// 	Children: []Widget{
							// 		Label{
							// 			Text: "水平座標系",
							// 		},
							// 		ComboBox{
							// 			Value:         Bind("CoorSysID", SelRequired{}),
							// 			BindingMember: "Id",
							// 			DisplayMember: "Name",
							// 			Model:         CoordSystem,
							// 		},
							// 	},
							// },
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.t1label,
										Text:     "COLLADAファイル",
									},
									PushButton{
										Text:      "作成",
										OnClicked: mw.LandDAE,
									},
									// ProgressBar{
									// 	AssignTo:    &mw.progressBar,
									// 	MarqueeMode: true,
									// },
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.t1label,
										Text:     "LandXMLファイル",
									},
									PushButton{
										Text:      "作成",
										OnClicked: mw.LandXML,
									},
								},
							},
							GroupBox{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Font:     Font{PointSize: 9},
										AssignTo: &mw.t1label,
										Text:     "DXFファイル",
									},
									PushButton{
										Text:      "作成",
										OnClicked: mw.LandDXF,
									},
								},
							},
						},
					},
				},
			},
			// Composite{
			// 	Layout: VBox{},
			// 	Children: []Widget{
			// 		Label{
			// 			Font:     Font{PointSize: 9},
			// 			AssignTo: &mw.olabel,
			// 			Text:     "四角形分割結果",
			// 		},
			// 		CustomWidget{
			// 			AssignTo:            &mw.paintWidget,
			// 			ClearsBackground:    true,
			// 			InvalidatesOnResize: true,
			// 			Paint:               mw.drawStuff,
			// 		},
			// 		PushButton{
			// 			Text:      "描画",
			// 			OnClicked: mw.menuOpen,
			// 		},
			// 	},
			// },
		},
	}

	if _, err := MW.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (mw *MyMainWindow) menuOpen() {
	log.Println("FileMenu - Open Clicked")
	fn, _ := internal.Opfl()
	filename = fn
	// ファイル名
	basename := filepath.Base(fn)
	s := fmt.Sprintf(basename)
	mw.le1.SetText(string(s))
	// データ件数
	fl, er := os.Open(fn)
	_, l, _, _ := pkg.FileCount(fl)
	mw.le2.SetText(strconv.Itoa(l - 7))
	if er != nil {
		log.Fatal(er)
	}
	defer fl.Close()
}

func (mw *MyMainWindow) menuOpen2() {
	log.Println("FileMenu - DEM Open Clicked")
	fn2, _ := internal.Opfl()
	demfile = fn2
	internal.MakeDem(demfile)

	// データ数のカウント
	fl, er := os.Open("C:/data/dem_data.txt")
	_, l, _, _ := pkg.FileCount(fl)
	mw.le6.SetText(strconv.Itoa(l))
	if er != nil {
		log.Fatal(er)
	}
	defer fl.Close()
}

func (mw *MyMainWindow) demMake() {
	log.Println("FileMenu - DEM Make Clicked")
	x_matrix, y_matrix, z_matrix, x_len, y_len, x_dot, y_dot, x_max, x_min, y_max, y_min, z_max, z_min = internal.TinMesh()
	mw.le7.SetText(strconv.FormatFloat(math.Round(x_len*0.0254), 'f', -1, 64))
	mw.le8.SetText(strconv.FormatFloat(math.Round(y_len*0.0254), 'f', -1, 64))
	mw.le9.SetText(strconv.Itoa(x_dot - 1))
	mw.le10.SetText(strconv.Itoa(y_dot - 1))
}

func (mw *MyMainWindow) pbClicked() {
	log.Println("Button Clicked")
	internal.DivideLine(filename)
	// wfl1, wfl2, wfl3 := internal.DispResult()

	wfl1, er := os.Open("C:/data/hutsu_list.txt")
	_, l1, _, _ := pkg.FileCount(wfl1)
	mw.le3.SetText(strconv.Itoa(l1))
	if er != nil {
		log.Fatal(er)
	}
	defer wfl1.Close()

	wfl2, er := os.Open("C:/data/kenro_list.txt")
	_, l2, _, _ := pkg.FileCount(wfl2)
	mw.le4.SetText(strconv.Itoa(l2))
	if er != nil {
		log.Fatal(er)
	}
	defer wfl2.Close()

	wfl3, er := os.Open("C:/data/other_list.txt")
	_, l3, _, _ := pkg.FileCount(wfl3)
	mw.le5.SetText(strconv.Itoa(l3))
	if er != nil {
		log.Fatal(er)
	}
	defer wfl3.Close()
}

func (mw *MyMainWindow) hutsuBuild() {
	log.Println("FileMenu - HutsuBuildings Clicked")
	internal.SquarePoly()
}

func (mw *MyMainWindow) trimeshDev() {
	log.Println("FileMenu - TriangleMeshDevide Clicked")
	internal.TriMeshDiv()
}

func (mw *MyMainWindow) BuilDAE() {
	log.Println("FileMenu - COLLADA_Building Clicked")
	internal.BuildDAE()
	// internal.TriMeshDiv()
}

func (mw *MyMainWindow) LandDAE() {
	log.Println("FileMenu - COLLADA_Terrain Clicked")
	internal.TerrDAE(x_matrix, y_matrix, z_matrix, x_dot, y_dot)
	// internal.TerrDAE()
}

func (mw *MyMainWindow) LandXML() {
	log.Println("FileMenu - LandXML_Terrain Clicked")
	internal.TerrXML(x_matrix, y_matrix, z_matrix, x_dot, y_dot, z_max, z_min)
	// internal.TerrXML()
}

func (mw *MyMainWindow) LandDXF() {
	log.Println("FileMenu - LandDXF_Terrain Clicked")
	internal.TerrDXF(x_matrix, y_matrix, z_matrix, x_dot, y_dot, x_max, x_min, y_max, y_min, z_max, z_min)
	// internal.TerrXML()
}

// func CoordSystem() []*CoorSys {
// 	return []*CoorSys{
// 		{1, "第Ⅰ系"},
// 		{2, "第Ⅱ系"},
// 		{3, "第Ⅲ系"},
// 		{4, "第Ⅳ系"},
// 		{5, "第Ⅴ系"},
// 		{6, "第Ⅵ系"},
// 		{7, "第Ⅶ系"},
// 		{8, "第Ⅷ系"},
// 		{9, "第Ⅸ系"},
// 		{10, "第Ⅹ系"},
// 		{11, "第ⅩⅠ系"},
// 		{12, "第ⅩⅡ系"},
// 		{13, "第ⅩⅢ系"},
// 		{14, "第ⅩⅣ系"},
// 		{15, "第ⅩⅤ系"},
// 		{16, "第ⅩⅥ系"},
// 		{17, "第ⅩⅦ系"},
// 		{18, "第ⅩⅧ系"},
// 		{19, "第ⅩⅨ系"},
// 	}
// }

func (mw *MyMainWindow) drawStuff(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	bmp, err := createBitmap()
	if err != nil {
		return err
	}
	defer bmp.Dispose()

	bounds := mw.paintWidget.ClientBounds()

	rectpen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(255, 0, 0))
	if err != nil {
		return err
	}
	defer rectpen.Dispose()

	if err := canvas.DrawRectanglePixels(rectpen, bounds); err != nil {
		return err
	}

	// ellipseBrush, err := walk.NewHatchBrush(walk.RGB(0, 255, 0), walk.HatchCross)
	// if err != nil {
	// 	return err
	// }
	// defer ellipseBrush.Dispose()

	// if err := canvas.FillEllipsePixels(ellipseBrush, bounds); err != nil {
	// 	return err
	// }

	linesBrush, err := walk.NewSolidColorBrush(walk.RGB(0, 0, 255))
	if err != nil {
		return err
	}
	defer linesBrush.Dispose()

	linesPen, err := walk.NewGeometricPen(walk.PenSolid, 2, linesBrush)
	if err != nil {
		return err
	}
	defer linesPen.Dispose()

	// if err := canvas.DrawLinePixels(linesPen, walk.Point{bounds.X, bounds.Y}, walk.Point{bounds.Width, bounds.Height}); err != nil {
	// 	return err
	// }
	// if err := canvas.DrawLinePixels(linesPen, walk.Point{bounds.X, bounds.Height}, walk.Point{bounds.Width, bounds.Y}); err != nil {
	// 	return err
	// }
	// fmt.Println(bounds.X, bounds.Y)
	// fmt.Println(bounds.Width, bounds.Height)

	points := make([]walk.Point, 10)
	dx := bounds.Width / (len(points) - 1)
	for i := range points {
		points[i].X = i * dx
		points[i].Y = int(float64(bounds.Height) / math.Pow(float64(bounds.Width/2), 2) * math.Pow(float64(i*dx-bounds.Width/2), 2))
	}
	if err := canvas.DrawPolylinePixels(linesPen, points); err != nil {
		return err
	}

	// bmpsize := bmp.Size()
	// if err := canvas.DrawImagePixels(bmp, walk.Point{(bounds.Width - bmpsize.Width) / 2, (bounds.Height - bmpsize.Height) / 2}); err != nil {
	// 	return err
	// }

	return nil
}

func createBitmap() (*walk.Bitmap, error) {
	bounds := walk.Rectangle{Width: 200, Height: 200}

	bmp, err := walk.NewBitmapForDPI(bounds.Size(), 32)
	if err != nil {
		return nil, err
	}

	// succeeded := false
	// defer func() {
	// 	if !succeeded {
	// 		bmp.Dispose()
	// 	}
	// }()

	// canvas, err := walk.NewCanvasFromImage(bmp)
	// if err != nil {
	// 	return nil, err
	// }
	// defer canvas.Dispose()

	// brushBmp, err := walk.NewBitmapFromFileForDPI("../img/plus.ping", 32)
	// if err != nil {
	// 	return nil, err
	// }
	// defer brushBmp.Dispose()

	// brush, err := walk.NewBitmapBrush(brushBmp)
	// if err != nil {
	// 	return nil, err
	// }
	// defer brush.Dispose()

	// if err := canvas.FillRectanglePixels(brush, bounds); err != nil {
	// 	return nil, err
	// }

	// font, err := walk.NewFont("Times New Roman", 40, walk.FontBold|walk.FontItalic)
	// if err != nil {
	// 	return nil, err
	// }
	// defer font.Dispose()

	// if err := canvas.DrawTextPixels("Walk Drawing Example", font, walk.RGB(0, 0, 0), bounds, walk.TextWordbreak); err != nil {
	// 	return nil, err
	// }

	// succeeded = true

	return bmp, nil
}
