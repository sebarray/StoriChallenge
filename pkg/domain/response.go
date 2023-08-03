package domain


type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`

}



type Request struct {
Emails	string `json:"email"`
}