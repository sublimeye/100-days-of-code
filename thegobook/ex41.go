package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func Exercise43() {

}

func Exercise44() {
	fmt.Println("Exercise44")
	slice := []int{2, 3, 4, 5, 6, 7}
	// 5, 6, 7, 2, 3, 4
	// 6, 7, 2, 3, 4, 5
	newslice := rotate(slice, 0)
	fmt.Println(newslice, slice)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, shift int) []int {
	length := len(s)
	newslice := []int{}

	if shift > 0 {
		if shift > length {
			shift = length
		}

		newslice = append(s[shift:], s[:shift-1]...)
	} else {
		if shift < -length {
			shift = -length
		}
		newslice = append(s[length+shift:], s[:length+shift]...)
	}

	return newslice
}

/**
* Exercise 4.5
*
 */
func Exercise45() {
	// test := []string{"a", "b", "b", "b", "c", "d", "d", "d", "d", "d", "d", "d", "d", "e", "d", "d", "d", "d", "d", "d"}
	test2 := []string{}
	// a, b, c, d
	newtest := RemoveAdjacent(test2)
	fmt.Println("Exercise45", newtest)
}

// Find adjacent duplicates in a string slice
func RemoveAdjacent(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}
	// go through strings
	// if prev == current => remove prev || current

	out := strings[0:1] // zero-length slice of original

	// 0,1,2,3,4,5,6
	for i, j := 1, 1; i < len(strings); i++ {
		if strings[i] != out[j-1] {
			out = append(out, strings[i])
			j++
		}
	}

	return out
}

// Exercise 4.6
// Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace)
// in a UTF-8-encoded []byte slice into a single ASCII space.
func Exercise46() {

}

// Maps
func playingWithMaps() {
	var names []string
	ages := map[string]int{
		"zomb":      43,
		"abba":      15,
		"dabba":     13,
		"chika":     14,
		"hikka":     10,
		"somebooka": 200,
	}

	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	fmt.Printf("%v\n", names)
	fmt.Printf("%v\n", ages)

	if age, ok := ages["zomb"]; !ok {
		fmt.Println("zomba is not in ages map", age)
	} else {
		fmt.Printf("value %v key %v", age, "zomba")
	}
}

type person struct {
	name string
	age  int
}

func ComparingStructs() {
	nai := person{name: "Ian", age: 31}
	ian := person{name: "Ian", age: 44}
	// fmt.Printf("Is equal %v, %v, %v", ian == nai, nai, ian)

	persons := make(map[person]int)
	persons[nai] = 20
	fmt.Printf("Map Persons to Int %v\n", persons)
	persons[ian] = 10

	fmt.Printf("Map Persons to Int %v\n", persons)
}

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func StructEmbeddingAnonymousFields() {
	var w Wheel
	w.X = 10       // Shorthand for w.Circle.Point.X
	w.Y = 15       // Shorthand for w.Circle.Point.Y
	w.Radius = 555 // Shorthand for w.Circle.Radius
	w.Spokes = 1
	// You can not use struct literal like this: w = Wheel{X: 10, Y: 15, Radius: 555, Spokes: 1} won't work

	fmt.Println(w)

	w2 := Wheel{Circle{Point{10, 20}, 50}, 100}
	w3 := Wheel{
		Circle: Circle{
			Point:  Point{X: 10, Y: 50},
			Radius: 532,
		},
		Spokes: 2673,
	}
	fmt.Println(w2)
	fmt.Println(w3)
}

func test() {
	var movies = []Movie{
		{
			Title:  "Casablanca",
			Year:   1923,
			Color:  false,
			Actors: []string{"Cha", "Boo"},
		},
		{
			Title:  "Breaking Bad",
			Year:   2013,
			Color:  false,
			Actors: []string{"Aaron Paul", "Mr. White"},
		},
	}
	// fmt.Println(movies)

	jsonData, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatal("JSON FUCK")
	}

	// fmt.Printf("%s\n", jsonData)

	var theotherdata []Movie
	err = json.Unmarshal(jsonData, &theotherdata)
	if err != nil {
		log.Fatalln("JSON UNMARSHAL", err)
	}
	fmt.Printf("%v\n", theotherdata)
}

func main() {
	// Exercise43()
	// Exercise44()
	// Exercise45()
	// playingWithMaps()
	// ComparingStructs()
	// StructEmbeddingAnonymousFields()
	// test()

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %v %9.9s %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}
