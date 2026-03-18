package core

/*
 * Interfaces In Go
 * ----------------
 * Interfaces are like blueprints of objects
 * it's a contract between the object and the code that uses the object
 * for e.g., below we have created an interface called Area every object
 * which is trying to calculate area using Area interface will have to
 * implement the GetArea() method.
 */

type Area interface {
	GetArea() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) GetArea() float64 {
	// area = pi * r^2
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) GetArea() float64 {
	// area = w * h
	return r.Width * r.Height
}
