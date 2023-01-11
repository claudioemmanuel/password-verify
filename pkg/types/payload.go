package types

type Payload struct {
	Password string `json:"password" valid:"required"`
	Rules    []Rule `json:"rules" valid:"required"`
}
