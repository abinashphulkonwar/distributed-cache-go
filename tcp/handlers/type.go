package handlers

type Body struct {
	Type string `json:"Type"`

	Data any `json:"Data"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Status  int    `json:"status"`
}
