package authentication_model

type AuthLoginPassword struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}
