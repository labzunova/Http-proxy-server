package repository

type RequestsRepo interface {
	SaveRequest(r Request) error
	LoadAllRequests() ([]Request, error)
	LoadOneRequest(id int) (Request, error)
	RepeatRequest(id int) error // todo what else to return
}

type Request struct {
	Id      int
	Host    string
	Path    string
	Method  string
	Headers string
	Body    string
	Params  string
	Cookies string
}
