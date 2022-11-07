package server

import (
	"encoding/json"
	//"fmt"
)

type Relation struct {
	Index []struct {
		//Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func GetRelations() {
	Relations := Relation{}
	data, _ := Harvest("https://groupietrackers.herokuapp.com/api/relation")
	json.Unmarshal(data, &Relations)
	//fmt.Println(Relations)
	//fmt.Println("jee")
	//fmt.Println(string(data))
	//return Relations
}
