package domain

type Rule struct {
	Rule  string `json:"rule"`
	Value int    `json:"value"`
}

type UserInput struct {
	Password string `json:"password"`
	Rules    []Rule `json:"rules"`
}

type ValidationResponse struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
