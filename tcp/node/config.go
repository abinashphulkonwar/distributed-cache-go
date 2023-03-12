package node

import (
	"encoding/json"
	"os"
)

type josnStructure struct {
	Title       string   `json:"title"`
	Connections []string `json:"connections"`
}

func WriteConnectionConfig() {




	var data []string




	for _, v := range ConnectionMap {
		data = append(data, v.ID)
	}




	if os.Getenv("mode") == "test" {
		data = append(data, "node1", "node2")
	}

	

	var jsonStructure = josnStructure{
		Title:       "Connections",
		Connections: data,
	}

	println(jsonStructure.Connections)

	jsonData, err := json.Marshal(jsonStructure)
	if err != nil {
		return
	}
	err = os.WriteFile("config.json", jsonData, 0644)
	if err != nil {
		println(err)
		return
	}

}
