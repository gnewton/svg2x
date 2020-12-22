package pkg

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Postscript struct {
	currentx, currenty float64
	haveCurrentPoint   bool
}

func (p *Postscript) Init(attributes map[string]string) {
	fmt.Println("%!PS-Adobe-3.0 EPSF-3.0")
	if box, ok := attributes["viewBox"]; ok {
		fmt.Println("%%BoundingBox: " + box)
		height := strings.Split(box, " ")[3]
		fmt.Println("10 10 scale")
		fmt.Println("0", height, "translate")
	}

}

func (p *Postscript) StartG(attributes map[string]string) {
	fmt.Println("%%StartG()")
	fmt.Println("gsave")
}

func (p *Postscript) EndG(attributes map[string]string) {
	fmt.Println("%% EndG()")
	fmt.Println("grestore")
}

func (p *Postscript) StartPath(attributes map[string]string) {
	fmt.Println("%% StartPath()")
	fmt.Println("newpath")
}

func (p *Postscript) EndPath(attributes map[string]string) {
	fmt.Println("%% EndPath()")
	if fill, ok := attributes["fill"]; ok {
		if fill != "none" {
			gsave()
			rgbColor(fill)
			if fillRule, ok := attributes["fill-rule"]; ok {
				if fillRule == "evenodd" {
					fmt.Println("eofill")
				} else {
					fmt.Println("fill")
				}
			} else {
				fmt.Println("fill")
			}
			grestore()
		}
	}

	if strokeColor, ok := attributes["stroke"]; ok {
		if strokeColor != "none" {
			gsave()
			rgbColor(strokeColor)
			stroke()
			grestore()
		}
	}
}
func grestore() {
	fmt.Println("grestore")
}

func stroke() {
	fmt.Println("stroke")
}

func gsave() {
	fmt.Println("gsave")
}
func rgbColor(hexColor string) {
	if hexColor != "none" && len(hexColor) > 1 && hexColor[0] == '#' {
		r, g, b, err := hex2rgb(hexColor)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%1.3f %1.3f %1.3f setrgbcolor\n", r, g, b)
	}
}

func hex2rgb(hexColor string) (r, g, b float32, err error) {
	if len(hexColor) != 7 {
		return r, g, b, errors.New("Not a hex string")
	}
	var tmp uint64

	hex := hexColor[1:3]
	tmp, err = strconv.ParseUint(hex, 16, 8)
	if err != nil {
		return r, g, b, err
	}
	r = float32(tmp) / 256.0

	hex = hexColor[3:5]
	tmp, err = strconv.ParseUint(hex, 16, 8)
	if err != nil {
		return r, g, b, err
	}
	g = float32(tmp) / 256.0

	hex = hexColor[5:7]
	tmp, err = strconv.ParseUint(hex, 16, 8)
	if err != nil {
		return r, g, b, err
	}
	b = float32(tmp) / 256.0

	return r, g, b, err
}

func (p *Postscript) Close() {
	fmt.Println("%% Close()")
	fmt.Println("showpage1")
}

func (p *Postscript) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	fmt.Println(x1, -y1, x2, -y2, x3, -y3, "curveto")
}

func (p *Postscript) MoveTo(pt *Point) {
	fmt.Println("")
	fmt.Println(pt.X, -pt.Y, "moveto")
}

func (p *Postscript) RMoveTo(pt *Point) {
	fmt.Println("")
	fmt.Println(pt.X, -pt.Y, "rmoveto")
}

func (p *Postscript) LineTo(ps []*Point) {
	for i, _ := range ps {
		fmt.Println(ps[i].X, -ps[i].Y, "lineto")
	}
}

func (p *Postscript) RLineTo(ps []*Point) {
	for i, _ := range ps {
		fmt.Println(ps[i].X, -ps[i].Y, "rlineto")
	}
}

func (p *Postscript) HLine(x float64) {

}

func (p *Postscript) RHLine(dx float64) {

}

func (p *Postscript) VLine(y float64) {

}

func (p *Postscript) RVLine(dy float64) {

}

func (p *Postscript) ClosePath() {
	fmt.Println("closepath")
	fmt.Println("")
}

func (p *Postscript) Fill(string) {

}

func (p *Postscript) Stroke(string) {

}
