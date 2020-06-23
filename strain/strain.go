//Package strain provides methods for filtering different types of lists.
package strain

//Ints is a wrapper around []int.
type Ints []int

//Lists is a slice of slices of ints.
type Lists [][]int

//Strings is a slice of strings.
type Strings []string

//Keep accepts a predicate argument and returns the slice it was called on, minus the records for which
//the predicate returned false.
func (i Ints) Keep(filter func(int) bool) Ints {
	var newInts Ints
	for _, num := range i {
		if filter(num) {
			newInts = append(newInts, num)
		}
	}
	return newInts
}

//Discard accepts a predicate argument and returns the slice it was called on, minus the records for which
//the predicate returned true.
func (i Ints) Discard(filter func(int) bool) Ints {
	var newInts Ints
	for _, num := range i {
		if !filter(num) {
			newInts = append(newInts, num)
		}
	}
	return newInts
}

//Keep accepts a predicate argument and returns the slice it was called on, minus the records for which
//the predicate returned false.
func (l Lists) Keep(filter func([]int) bool) Lists {
	var newLists Lists
	for _, list := range l {
		if filter(list) {
			newLists = append(newLists, list)
		}
	}
	return newLists
}

//Keep accepts a predicate argument and returns the slice it was called on, minus the records for which
//the predicate returned false.
func (s Strings) Keep(filter func(string) bool) Strings {
	var newStrings Strings
	for _, num := range s {
		if filter(num) {
			newStrings = append(newStrings, num)
		}
	}
	return newStrings
}
