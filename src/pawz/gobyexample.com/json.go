package main

import "encoding/json"
import "fmt"
import "os"

var f = fmt.Println

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	boolB, _ := json.Marshal(true)
	f(string(boolB))

	intB, _ := json.Marshal(1)
	f(string(intB))

	fltB, _ := json.Marshal(2.34)
	f(string(fltB))

	strB, _ := json.Marshal("gopher")
	f(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	f(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	f(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	f(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	f(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	f(dat)

	num := dat["num"].(float64)
	f(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	str2 := strs[1].(string)
	f(str1)
	f(str2)

	str := `{"page":1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	f(res)
	f(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
