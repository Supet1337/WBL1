package point

type Point struct {
	x float64
	y float64
}

var x int

func PointNew(x, y float64) *Point {
	point := Point{
		x: x,
		y: y,
	}
	return &point
}
func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}
