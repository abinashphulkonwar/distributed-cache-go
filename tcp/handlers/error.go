package handlers

import (
	"encoding/json"

	"github.com/abinashphulkonwar/dist-cache/tcp/node"
)

func ErrorHandler(data *ErrorRes, c *node.Client) {

	res := Response{
		Message: data.Message,
		Status:  data.Status,
	}
	res.Data = [2]string{"", ""}
	jsonData, err := json.Marshal(res)

	if err != nil {
		println("error: ", err)
		return
	}

	c.Conn.Write(jsonData)
}
