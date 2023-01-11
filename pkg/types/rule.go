package types

type Rule struct {
	Rule  string `json:"rule" valid:"required"`
	Value int    `json:"value" valid:"required"`
}
