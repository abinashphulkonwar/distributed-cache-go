package handlers

type Body struct {
	Key  string `json:"Key"`
	Data string `json:"Data"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Status  int    `json:"status"`
}
