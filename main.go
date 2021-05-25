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
	// 図の生成
	p := plot.New()
	//label
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	for i := 1; i < len(os.Args); i++ {

		path := os.Args[i]

		if !Exists(path) {
			fmt.Println("no file:", path)
			continue
		}

		list := CreatePlotData(path)
		// 補助線
		p.Add(plotter.NewGrid())
		plot, _ := plotter.NewScatter(list)
		plot.GlyphStyle.Color = GetColor(i)

		//plotをplot
		p.Add(plot)
	}

	// plot.pngに保存
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}

}
func GetColor(idx int) color.RGBA {
	var c color.RGBA
	switch idx {
	case 0:
		c = color.RGBA{R: 255, G: 128, B: 0, A: 255}
	case 1:
		c = color.RGBA{R: 255, B: 0, A: 255}
	case 2:
		c = color.RGBA{R: 0, B: 255, A: 255}
	case 3:
		c = color.RGBA{R: 0, B: 0, A: 255}
	}
	return c

}
func CreatePlotData(path string) plotter.XYs {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	list := plotter.XYs{}
	reader := csv.NewReader(file)
	idx := 1
	for {
		rand.Seed(time.Now().UnixNano())
		val, err := reader.Read()
		if err != nil {
			break
		}
		v, _ := strconv.ParseFloat(val[0], 64)
		list = append(list, plotter.XY{X: (float64)(idx), Y: v})
		idx += 1
	}

	return list
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
