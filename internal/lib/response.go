package lib

type Response struct {
	Success bool
	Status  int
	Message string
	Results any
}

type ResponseUsers struct {
	Success     bool
	Status      int
	Message     string
	Page        string
	ORDER_BY    string
	Order       string
	Data_length string
	Results     any
}

type LoginResponse struct {
	Id    int64
	Token string
}
