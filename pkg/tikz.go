package pkg

import (
	//	"errors"
	"fmt"
	//	"log"
	//"strconv"
	"strings"
)

type Tikz struct {
}

func (p *Tikz) Init(attributes map[string]string) {
	fmt.Println("%!PS-Adobe-3.0 EPSF-3.0")
	if box, ok := attributes["viewBox"]; ok {
		fmt.Println("%%BoundingBox: " + box)
		height := strings.Split(box, " ")[3]
		fmt.Println("10 10 scale")
		fmt.Println("0", height, "translate")
	}

}

func (p *Tikz) StartG(attributes map[string]string) {
	fmt.Println("%%StartG()")
	fmt.Println("gsave")
}

func (p *Tikz) EndG(attributes map[string]string) {
	fmt.Println("%% EndG()")
	fmt.Println("grestore")
}

func (p *Tikz) StartPath(attributes map[string]string) {
	fmt.Println("%% StartPath()")
	fmt.Println("newpath")
}

func (p *Tikz) EndPath(attributes map[string]string) {
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

func (p *Tikz) Close() {
	fmt.Println("%% Close()")
	fmt.Println("showpage1")
}

func (p *Tikz) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	fmt.Println(x1, -y1, x2, -y2, x3, -y3, "curveto")
}

func (p *Tikz) MoveTo(x, y float64) {
	fmt.Println("")
	fmt.Println(x, -y, "moveto")
}

func (p *Tikz) RMoveTo(dx, dy float64) {

}

func (p *Tikz) LineTo(x, y float64) {
	fmt.Println(x, -y, "lineto")
}

func (p *Tikz) RLineTo(rx, ry float64) {

}

func (p *Tikz) HLine(x float64) {

}

func (p *Tikz) RHLine(dx float64) {

}

func (p *Tikz) VLine(y float64) {

}

func (p *Tikz) RVLine(dy float64) {

}

func (p *Tikz) ClosePath() {
	fmt.Println("closepath")
	fmt.Println("")
}

func (p *Tikz) Fill(string) {

}

func (p *Tikz) Stroke(string) {

}
