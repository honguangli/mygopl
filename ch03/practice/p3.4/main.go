// 练习 3.4： 参考1.7节Lissajous例子的函数，构造一个web服务器，
// 用于计算函数曲面然后返回SVG数据给客户端。服务器必须设置Content-Type头部：
// w.Header().Set("Content-Type", "image/svg+xml")
//（这一步在Lissajous例子中不是必须的，因为服务器使用标准的PNG图像格式，
// 可以根据前面的512个字节自动输出对应的头部。）
// 允许客户端通过HTTP请求参数设置高度、宽度和颜色等参数。
package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var (
	width, height = 600, 320                         // canvas size in  pixels
	xyscale       = float64(width) / 2 / xyrange     // pixels per x or y unit
	zscale        = float64(height) * 0.4            // pixels per z unit
	sin30, cos30  = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
	color         = "grey"
)

func main() {
	http.HandleFunc("/", surface)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("http error: %s", err.Error())
	}
}

func surface(w http.ResponseWriter, r *http.Request) {
	var bf bytes.Buffer

	if err := r.ParseForm(); err == nil {
		if r.FormValue("width") != "" {
			width, err = strconv.Atoi(r.FormValue("width"))
			xyscale = float64(width) / 2 / xyrange
		}
		if r.FormValue("height") != "" {
			height, err = strconv.Atoi(r.FormValue("height"))
			zscale = float64(height) * 0.4
		}
		if r.FormValue("color") != "" {
			color = r.FormValue("color")
		}
	}
	bf.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: #%s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			bf.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	bf.WriteString(fmt.Sprintf("</svg>"))

	w.Header().Set("Content-Type", "image/svg+xml")

	if _, err := w.Write(bf.Bytes()); err != nil {
		log.Fatalf("output error: %s", err.Error())
	}
}

func corner(i, j int) (sx float64, sy float64, ok bool) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, false
	}

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx = float64(width)/2 + (x-y)*cos30*xyscale
	sy = float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (z float64, ok bool) {
	r := math.Hypot(x, y) // distance from (0, 0)
	z = math.Sin(r) / r
	if math.IsNaN(z) {
		return 0, false
	}
	return z, true
}
