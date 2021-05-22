package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("few args:", len(os.Args))
		return
	}
	path := os.Args[1]

	if !Exists(path) {
		fmt.Println("no file:", path)
		return
	}

	list := CreatePlotData(path)

	// 図の生成
	p := plot.New()
	//label
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// 補助線
	p.Add(plotter.NewGrid())
	plot1, _ := plotter.NewScatter(list)
	plot1.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 55}

	//plot1,plot2をplot
	p.Add(plot1)

	// plot.pngに保存
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}

}
func CreatePlotData(path string) plotter.XYs {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	list := plotter.XYs{}
	reader := csv.NewReader(file)
	for idx := 0; idx < 100; idx++ {
		rand.Seed(time.Now().UnixNano())
		val, err := reader.Read()
		if err != nil {
			break
		}
		v, _ := strconv.Atoi(val[1])
		list = append(list, plotter.XY{X: (float64)(idx), Y: (float64)(v)})
	}

	return list
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
