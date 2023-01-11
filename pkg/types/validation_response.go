package types

type ValidationResponse struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
