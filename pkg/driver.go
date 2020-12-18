package pkg


type Driver interface{
	MoveTo(x,y float64)
	RMoveTo(dx,dy float64)
	LineTo(x,y float64)
	RLineTo(rx,ry float64)
	HLine(x float64)
	RHLine(dx float64)
	VLine(y float64)
	RVLine(dy float64)
	ClosePath()
	Fill(string)
	Stroke(string)

	
 // 	C x1 y1, x2 y2, x y
 // 	c dx1 dy1, dx2 dy2, dx dy
 // 	S x2 y2, x y
 // (or)
 // 	s dx2 dy2, dx dy
 // 	Q x1 y1, x y
 // (or)
 // 	q dx1 dy1, dx dy
 // 	T x y
 // (or)
 // 	t dx dy

	////
	
	

}
