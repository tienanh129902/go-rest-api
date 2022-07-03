package datatransfers

type Error struct {
	Error interface{} `json:"error,omitempty"`
}

type Status struct {
	Status interface{} `json:"status,omitempty"`
}

type Data struct {
	Data interface{} `json:"data,omitempty"`
}

type Token struct {
	AccessToken  interface{} `json:"access_token,omitempty"`
	RefreshToken interface{} `json:"refresh_token,omitempty"`
}
