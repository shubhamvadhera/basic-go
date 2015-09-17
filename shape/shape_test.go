package shape

import "testing"

// For Circle area tests
type circAreaPair struct {
	c Circle
	a float64
}

var circAreaTests = []circAreaPair{
	{Circle{0, 0, 10}, 314.1592653589793},
	{Circle{0, 0, 5}, 78.53981633974483},
}

// For Circle perimeter tests
type circPeriPair struct {
	c Circle
	p float64
}

var circPeriTests = []circPeriPair{
	{Circle{0, 0, 3}, 18.84955592153876},
	{Circle{0, 0, 9}, 56.548667764616276},
}

// For rectangle area tests
type rectAreaPair struct {
	r Rectangle
	a float64
}

var rectAreaTests = []rectAreaPair{
	{Rectangle{0, 0, 10, 18}, 180},
	{Rectangle{0, 0, 1, 7}, 7},
}

// For rectangle perimeter tests
type rectPeriPair struct {
	r Rectangle
	p float64
}

var rectPeriTests = []rectPeriPair{
	{Rectangle{0, 0, 5, 15}, 40},
	{Rectangle{0, 0, 7, 7}, 28},
}

func TestAll(t *testing.T) {
	//Circle Area Tests
	for _, pair := range circAreaTests {
		v := pair.c.area()
		if v != pair.a {
			t.Error(
				"For Circle x=", pair.c.x, " y=", pair.c.y, " radius=", pair.c.r,
				"expected area", pair.a,
				"got", v,
			)
		}
	}
	//Cirle Perimeter Tests
	for _, pair := range circPeriTests {
		v := pair.c.perimeter()
		if v != pair.p {
			t.Error(
				"For Circle x=", pair.c.x, " y=", pair.c.y, " radius=", pair.c.r,
				"expected perimeter", pair.p,
				"got", v,
			)
		}
	}
	//Rectangle Area Tests
	for _, pair := range rectAreaTests {
		v := pair.r.area()
		if v != pair.a {
			t.Error(
				"For Rectangle x=", pair.r.x1, " y=", pair.r.y1, " side1=", pair.r.x2, " side2=", pair.r.y2,
				"expected area", pair.a,
				"got", v,
			)
		}
	}
	//Rectangle Perimeter Tests
	for _, pair := range rectPeriTests {
		v := pair.r.perimeter()
		if v != pair.p {
			t.Error(
				"For Rectangle x=", pair.r.x1, " y=", pair.r.y1, " side1=", pair.r.x2, " side2=", pair.r.y2,
				"expected perimeter", pair.p,
				"got", v,
			)
		}
	}
}
