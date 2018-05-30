package types

type User struct {
	DN       string `form:"dn" json:"dn" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Generator struct {
	UserName string `json:"username,omitempty"`
	CA       string `json:"ca,omitempty"`
	Token    string `json:"token,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	Message  string `json:"message,omitempty"`
	Status   string `json:"status"`
}

const (
	Authorized   = "authorized"
	Unauthorized = "unauthorized"
	Error        = "error"
)
