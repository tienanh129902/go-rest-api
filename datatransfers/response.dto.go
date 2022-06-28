package datatransfers

type Response struct {
	Status       interface{} `json:"status,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	AccessToken  interface{} `json:"access_token,omitempty"`
	RefreshToken interface{} `json:"refresh_token,omitempty"`
	Error        interface{} `json:"error,omitempty"`
}
