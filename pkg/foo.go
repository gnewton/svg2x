package pkg

import(
	"github.com/JoshVarga/svgparser" // see github.com/catiepg/svg
)

func f(){
	_,_ = svgparser.Parse(nil, false)
}
