package main

import (
	"bufio"
	//"errors"
	"fmt"
	"github.com/JoshVarga/svgparser" // see github.com/catiepg/svg
	"github.com/gnewton/svg2x/pkg"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	driver := new(pkg.Postscript)
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)

	element, err := svgparser.Parse(reader, false)
	if err != nil {
		log.Fatal(err)
	}

	// element=svg
	attributes := make(map[string]string)

	driver.Init(element.Attributes)

	setAttributes(attributes, element.Attributes)
	err = renderChildren(driver, element, attributes)
	if err != nil {
		log.Fatal(err)
	}
	if true {
		return
	}

	printChildren(element)

	fmt.Printf("%% SVG width: %s\n", element.Attributes["width"])
	fmt.Printf("%% Circle fill: %s\n", element.Children[0].Attributes["fill"])

	// Output:
	// SVG width: 100
	// Circle fill: red
}

func renderChildren(driver pkg.Driver, element *svgparser.Element, attributes map[string]string) error {

	for _, c := range element.Children {
		fmt.Printf("\n%% Name=%s\n", c.Name)
		fmt.Printf("%% %+v\n", c)
		xatts := make(map[string]string)
		// Inherit parent's attributes
		setAttributes(xatts, attributes)
		// Then overwrite them
		setAttributes(xatts, c.Attributes)

		if c.Name == "g" {
			//driver.StartG(c.Attributes)
			driver.StartG(xatts)
		} else {
			if c.Name == "path" {
				//driver.StartPath(c.Attributes)
				driver.StartPath(xatts)
				//for k, v := range c.Attributes {
				for k, v := range xatts {
					fmt.Println("%%", k, "=", v)
					if k == "d" {
						err := pathParse(driver, v)
						if err != nil {
							return err
						}
					}
				}
			}
		}
		err := renderChildren(driver, c, xatts)
		if err != nil {
			return err
		}
		if c.Name == "g" {
			//driver.EndG(c.Attributes)
			driver.EndG(xatts)
		} else {
			if c.Name == "path" {
				//driver.EndPath(c.Attributes)
				driver.EndPath(xatts)
			}
		}
	}
	return nil
}

func printChildren(element *svgparser.Element) {
	fmt.Println("%%---------------")
	for _, c := range element.Children {
		fmt.Println("%% Name:", c.Name)
		fmt.Println("%% Content:", c.Name)
		//fmt.Println("Attributes", c.Attributes)
		for k, v := range c.Attributes {
			fmt.Println("%%", k, "=", v)
			if k == "d" {
				pathParse2(v)
			}
		}
		//fmt.Printf("Attributes: %+v\n",c.Attributes)
		//fmt.Printf("%+v\n",c)
		printChildren(c)
	}

}

func pathParse(driver pkg.Driver, s string) error {
	re := regexp.MustCompile(`[A-Za-z]`)
	s = re.ReplaceAllStringFunc(s, sep)
	//log.Println(s)
	parts := strings.Split(s, "|")

	//log.Println(parts)
	for i, p := range parts {

		fmt.Println("%%%%%%%%", i, p)
		p = strings.TrimSpace(p)
		if len(p) == 0 {
			fmt.Println("%%%%% ---- trimspace=0", p)
			continue
		}
		//log.Println(p)
		switch p[0] {
		case 'M':
			numbers, err := pkg.ParseNumbers(p)
			if len(numbers)%2 != 0 {
				return fmt.Errorf("Wrong number of numbers: %d  results: %v  original: %s", len(numbers), numbers, p)
			}
			if err != nil {
				return err
			}
			driver.MoveTo(list2Points(numbers)[0])
		case 'L':
			numbers, err := pkg.ParseNumbers(p)
			if len(numbers)%2 != 0 {
				return fmt.Errorf("Wrong number of numbers: %d  results: %v  original: %s", len(numbers), numbers, p)
			}
			if err != nil {
				return err
			}
			driver.LineTo(list2Points(numbers))
		case 'l':
			numbers, err := pkg.ParseNumbers(p)
			if len(numbers)%2 != 0 {
				return fmt.Errorf("Wrong number of numbers: %d  results: %v  original: %s", len(numbers), numbers, p)
			}
			if err != nil {
				return err
			}
			//driver.RLineTo(x, y)
			log.Println(p, numbers)
			driver.RLineTo(list2Points(numbers))
		case 'C':
			x1, y1, x2, y2, x3, y3, err := convert6(p)
			if err != nil {
				return err
			}
			driver.CurveTo(x1, y1, x2, y2, x3, y3)
		case 'Z':
			driver.ClosePath()
		}
	}
	return nil
}

func pathParse2(s string) {
	fmt.Println("%% *********************************")
	re := regexp.MustCompile(`[A-Za-z]`)
	s = re.ReplaceAllStringFunc(s, sep)
	parts := strings.Split(s, "|")
	fmt.Println("%% parts", parts)
	for i, p := range parts {
		fmt.Println("%%%", i, p)
	}
}
func sep(command string) string {
	return "|" + command
}

func list2Points(num []float64) []*pkg.Point {

	pts := make([]*pkg.Point, len(num)/2)
	counter := 0
	for i := 0; i < len(num); i += 2 {
		pt := new(pkg.Point)
		pt.X = num[i]
		pt.Y = num[i+1]
		pts[counter] = pt
		counter++
	}
	return pts
}

// func convertPairsM(s string) (x, y []float64, err error) {
// 	s = s[1:len(s)]
// 	x = make([]float64, 0)
// 	y = make([]float64, 0)
// 	parts := strings.Split(s, ",")
// 	fmt.Println("%%%%num parts=", len(parts), parts)
// 	x, err = strconv.ParseFloat(parts[0], 64)
// 	if err != nil {
// 		log.Println(parts)
// 		log.Println(err)
// 		return x, y, err
// 	}
// 	y, err = strconv.ParseFloat(parts[1], 64)
// 	return x, y, err

// }

func convert6(s string) (x1, y1, x2, y2, x3, y3 float64, err error) {
	s = s[1:len(s)]
	parts := strings.Split(s, " ")
	first := strings.Split(parts[0], ",")
	x1, err = strconv.ParseFloat(first[0], 64)
	if err != nil {
		return x1, y1, x2, y2, x3, y3, err
	}
	y1, err = strconv.ParseFloat(first[1], 64)
	if err != nil {
		return x1, y1, x2, y2, x3, y3, err
	}

	second := strings.Split(parts[1], ",")
	x2, err = strconv.ParseFloat(second[0], 64)
	if err != nil {
		return x1, y1, x2, y2, x3, y3, err
	}
	y2, err = strconv.ParseFloat(second[1], 64)
	if err != nil {
		return x1, y1, x2, y2, x3, y3, err
	}

	third := strings.Split(parts[2], ",")
	x3, err = strconv.ParseFloat(third[0], 64)
	if err != nil {
		return x1, y1, x2, y2, x3, y3, err
	}
	y3, err = strconv.ParseFloat(third[1], 64)
	return x1, y1, x2, y2, x3, y3, err

}

func fix(s string) string {
	return strings.ReplaceAll(s, ",", " ")
}

func setAttributes(dest, source map[string]string) {
	for k, v := range source {
		dest[k] = v
	}
}
