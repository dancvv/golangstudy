package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmt"`
	Actors []string `json:"actors:"`
}

func main() {
	movie := Movie{"喜剧之王", 200, 10, []string{"xingye", "张柏芝"}}

	// 编码的过程， 结构体---> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error: ", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程 jsonstr ----> 结构体
	// jsonStr = {"title":"喜剧之王","year":200,"rmt":10,"actors:":["xingye","张柏芝"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshall error", err)
		return
	}
	fmt.Printf("%v \n", myMovie)
}
