package handlers

type Data struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Commend string `json:"commend"`
}

type Body struct {
	Type string `json:"Type"`
	Data Data   `json:"Data"`
}

const (
	SET    = "SET"
	GET    = "GET"
	DELETE = "DELETE"
)

func IsValidCommend(commend Data) bool {
	switch commend.Commend {
	case SET:
		return true
	case GET:
		return true
	case DELETE:
		return true
	default:
		return false
	}
}

type Response struct {
	Message string    `json:"message"`
	Data    [2]string `json:"data"`
	Status  int       `json:"status"`
}

type ErrorRes struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

const (
	SUCCESS = 200
	FAIL    = 500
)
