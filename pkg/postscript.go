
package pkg

type Postscript struct{
	currentx,currenty float64
	haveCurrentPoint bool
}

func (p *Postscript) MoveTo(x,y float64){

}

func (p *Postscript) RMoveTo(dx,dy float64){

}

func (p *Postscript) LineTo(x,y float64){

}

func (p *Postscript) RLineTo(rx,ry float64){

}

func (p *Postscript) HLine(x float64){

}

func (p *Postscript) RHLine(dx float64){

}

func (p *Postscript) VLine(y float64){

}

func (p *Postscript) RVLine(dy float64){

}

func (p *Postscript) ClosePath(){

}

func (p *Postscript) Fill(string){

}

func (p *Postscript) Stroke(string){

}
