package main

import "fmt"
import "math/rand"
import "sort"

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	//swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

func removeAtIndex2(source []int, index int) []int {
	temp := source[0]
	source[0] = source[index]
	source[index] = temp
	return source[1:]
}

//map in struct
type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

func main() {
	var wutever [10]int
	wutever[0] = 1

	scores := [3]int{1, 2, 3}
	fmt.Println("length of scores: ", len(scores))

	//slices
	//sscores := []int{1, 5, 8, 3, 6}
	//sscores2 := make([]int, 10)

	///
	tes := make([]int, 0, 5)
	c := cap(tes)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		tes = append(tes, i)
		//captes := cap(tes)
		//fmt.Println("debug i, captes: ", i, captes)
		if cap(tes) != c {
			c = cap(tes)
			fmt.Println("c: ", c)
		}
	}

	///
	zz := make([]int, 5)
	zz = append(zz, 9332)
	fmt.Println(zz)

	///
	aa := []int{1, 2, 3, 4, 5}
	aa = removeAtIndex(aa, 2)
	fmt.Println("aa: ", aa)

	ab := []int{1, 2, 3, 4, 5}
	ab = removeAtIndex2(ab, 2)
	fmt.Println("ab: ", ab)

	//copy
	cc := make([]int, 20)
	for i := 0; i < 20; i++ {
		cc[i] = int(rand.Int31n(999))
	}
	sort.Ints(cc)
	fmt.Println(cc)
	worst := make([]int, 5)
	copy(worst, cc[:3])
	fmt.Println(worst)
	worst2 := make([]int, 5)
	copy(worst2[2:4], cc[:5])
	fmt.Println(worst2)

	//maps
	lookup := make(map[string]int)
	lookup["goku"] = 9000
	lookup["gohan"] = 7000
	lookup["vegeta"] = 8999
	lookup["raditz"] = 8000
	lookup["napa"] = 7999
	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)
	//get key total number
	total := len(lookup)
	fmt.Println(total)
	//delete value by its key
	delete(lookup, "goku")

	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}

	goku.Friends["Vegeta"] = &Saiyan{
		Name:    "Vegeta",
		Friends: nil,
	}
	name1, existkah := goku.Friends["Vegeta"]
	fmt.Println(name1.Name, existkah)

	for key, val1 := range lookup {
		fmt.Println("key, value: ", key, val1)
	}
}
