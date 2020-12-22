package pkg

type Driver interface {
	Init(attributes map[string]string)
	StartG(attributes map[string]string)
	EndG(attributes map[string]string)
	StartPath(attributes map[string]string)
	EndPath(attributes map[string]string)
	Close()
	CurveTo(x1, y1, x2, y2, x3, y3 float64)
	MoveTo(p *Point)
	RMoveTo(p *Point)
	LineTo(p []*Point)
	RLineTo(p []*Point)
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
