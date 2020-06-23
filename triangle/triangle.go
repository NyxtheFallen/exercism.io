//Package triangle provides tools for identifying triangles.
package triangle

import (
	"math"
	"sort"
)

//Kind is a wrapper around int, used to describe the kind of a triangle.
type Kind int

//This group of constants describe the kind of a triangle.
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

//KindFromSides accepts three floats--the sides of a triangle--and returns whether that triangle is scalene, isoceles, equilateral, or broken.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	var sides sort.Float64Slice = []float64{a, b, c}
	sides.Sort()
	var isInvalid bool
	for _, val := range sides {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			isInvalid = true
		}
	}
	//Because the slice is sorted, we can skip some cases.
	if isInvalid || sides[2] > sides[1]+sides[0] || sides[2] <= 0 {
		k = NaT
	} else if sides[2] == sides[0] {
		k = Equ
	} else if sides[2] == sides[1] || sides[1] == sides[0] {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
