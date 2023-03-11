package node

func WriteConnectionConfig() {
	var data []string

	for _, v := range ConnectionMap {
		data = append(data, v.ID)
	}
	println(data)

}
