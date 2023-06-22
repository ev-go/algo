package plat2_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

//type MyData struct {
//	One int	'json:"one"'
//	Two string 'json:"-"' 'json:","' ///
//}

type MyData struct {
	One int    `json:"one"`
	Two string `json:"two"`
}

func TestPlat2(t *testing.T) {
	in := MyData{
		One: 1,
		Two: "two",
	}
	fmt.Printf("%#v\n", in)
	encoded, _ := json.Marshal(in)

	fmt.Println(string(encoded))

	var out MyData
	err := json.Unmarshal(encoded, &out)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", out)
}
