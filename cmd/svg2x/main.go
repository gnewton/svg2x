package main

import(
	"bufio"
	"regexp"
	"strings"
	"fmt"
	"github.com/JoshVarga/svgparser" // see github.com/catiepg/svg
	"log"
	"os"
	"github.com/gnewton/svg2x/api"
)

func main(){
	_ = new(Postscript)
	f, err := os.Open("../test/data/Res_Amazon-EC2_F1-Instance_48_Dark.svg")
	if err!=nil{
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)

	element, err := svgparser.Parse(reader, false)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v\n",element)
	printChildren(element)

	fmt.Printf("SVG width: %s\n", element.Attributes["width"])
	fmt.Printf("Circle fill: %s\n", element.Children[0].Attributes["fill"])

	// Output:
	// SVG width: 100
	// Circle fill: red
}

func printChildren(element *svgparser.Element){
	fmt.Println("---------------")
	for _,c:= range element.Children{
		fmt.Println("Name:", c.Name)
		fmt.Println("Content:", c.Name)
		//fmt.Println("Attributes", c.Attributes)
		for k, v := range c.Attributes {
			fmt.Println(k, "=", v)
			if k == "d"{
				pathParse(v)
			}
		}
		//fmt.Printf("Attributes: %+v\n",c.Attributes)
		//fmt.Printf("%+v\n",c)
		printChildren(c)
	}
}

func pathParse(s string){
	fmt.Println("*********************************")
	re := regexp.MustCompile(`[A-Za-z]`)
	s = re.ReplaceAllStringFunc(s,sep)
	parts := strings.Split(s,"|")
	fmt.Println("parts", parts)
	for i,p := range parts{
		fmt.Println(i,p)
	}
}
func sep(command string)string{
 	return "|" + command
}
