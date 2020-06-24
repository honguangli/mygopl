// 练习 1.12： 修改Lissajour服务，从URL读取变量，
// 比如你可以访问 http://localhost:8000/?cycles=20 这个URL，
// 这样访问可以将程序里的cycles默认的5修改为20。
// 字符串转换为数字可以调用strconv.Atoi函数。
// 你可以在godoc里查看strconv.Atoi的详细说明。
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	var opt = option{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
	if cycles, err := strconv.Atoi(r.FormValue("cycles")); err == nil {
		opt.cycles = cycles
	}
	if res, err := strconv.ParseFloat(r.FormValue("res"), 64); err == nil {
		opt.res = res
	}
	if size, err := strconv.Atoi(r.FormValue("size")); err == nil {
		opt.size = size
	}
	if nframes, err := strconv.Atoi(r.FormValue("nframes")); err == nil {
		opt.nframes = nframes
	}
	if delay, err := strconv.Atoi(r.FormValue("delay")); err == nil {
		log.Println(delay)
		opt.delay = delay
	}
	lissajous(opt, w)
}

type option struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func lissajous(opt option, out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of  y oscillator
	anim := gif.GIF{LoopCount: opt.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < opt.nframes; i++ {
		rect := image.Rect(0, 0, 2*opt.size+1, 2*opt.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(opt.cycles)*2*math.Pi; t += opt.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(opt.size+int(x*float64(opt.size)+0.5), opt.size+int(y*float64(opt.size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, opt.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
