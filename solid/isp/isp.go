package isp

// // This disobeys ISP because the client is depending on a method
// // it does not need.
// type shape interface {
// 	area() float64
// 	volume() float64
// }

// Fixed to obey ISP
type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

type square struct {
	sideLen float64
}

func (s square) area() float64 {
	return s.sideLen * s.sideLen
}

// // This disobeys ISP because the client is depending on a method
// // it does not need.
// func (s square) volume() float64 {
// 	return 0
// }

// type cube struct {
// 	sideLen float64
// }

// Fix to support ISP by extending the cube struct from the square shape
type cube struct {
	square
}

func (c cube) volume() float64 {
	return c.sideLen * c.sideLen * c.sideLen
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

// Fix to support ISP by using the object interface for volume + area calculations
func areaVolumeSum(objects ...object) float64 {
	var sum float64
	for _, ob := range objects {
		sum += ob.volume() + ob.area()
	}
	return sum
}
