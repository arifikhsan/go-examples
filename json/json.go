package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type responseOne struct {
	Page   int
	Fruits []string
}

type responseTwo struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	booleanB, _ := json.Marshal(true)
	fmt.Println(string(booleanB))

	integerB, _ := json.Marshal(1)
	fmt.Println(string(integerB))

	floatB, _ := json.Marshal(2.34)
	fmt.Println(string(floatB))

	stringB, _ := json.Marshal("gopher")
	fmt.Println(string(stringB))

	sliceD := []string{"apple", "peach", "pear"}
	sliceB, _ := json.Marshal(sliceD)
	fmt.Println(string(sliceB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	resOneD := &responseOne{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	resOneB, _ := json.Marshal(resOneD)
	fmt.Println(string(resOneB))

	resTwoD := &responseTwo{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	resTwoB, _ := json.Marshal(resTwoD)
	fmt.Println(string(resTwoB))

	byt := []byte(`{"num": 6.13, "strings": ["a", "b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strings := dat["strings"].([]interface{})
	stringOne := strings[0].(string)
	fmt.Println(stringOne)

	fruits := `{"page": 1, "fruits": ["apple", "peach"]}`
	response := responseTwo{}
	json.Unmarshal([]byte(fruits), &response)
	fmt.Println(response)
	fmt.Println(response.Page)
	fmt.Println(response.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
