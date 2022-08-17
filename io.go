package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type param struct {
	Mi1 float64    `json:"mi1"`
	Mi2 float64    `json:"mi2"`
	Mi3 float64    `json:"mi3"`
	Days float64    `json:"days"`
	R1 float64    `json:"r1"`
	R2 float64    `json:"r2"`
	R3 float64    `json:"r3"`
	SP float64    `json:"sp"`
}

func main(){
	c := param{
		Mi1: 0,
		Mi2: 0,
		Mi3: 0,
		Days: 0,
		R1: 0,
		R2: 0,
		R3: 0,
		SP: 0,
	}

//savJson(c)
c = loaJson()
disp(c)
fmt.Println("Hello")
}
func savJson(c param){

	dat, err := json.Marshal(c)
	if err != nil {
		//return err
	}
	err = ioutil.WriteFile("file.json", dat, 0644)
	if err != nil {
		//return err
	}

}
func loaJson() param {

	dat, err := ioutil.ReadFile("file.json")
	if err != nil {
		//return err
	}
	c := param{}
	err = json.Unmarshal(dat, &c)
	if err != nil {
		//return err
	}
return c
}

func disp(c param) {
	fmt.Println(c.Mi1)
	fmt.Println(c.Mi2)
	fmt.Println(c.Mi3)
	fmt.Println(c.Days)
	fmt.Println(c.R1)
	fmt.Println(c.R2)
	fmt.Println(c.R3)
	fmt.Println(c.SP)

}
