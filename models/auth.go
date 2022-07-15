package models

type Authentication struct {
	Header string `json:"header" default:"Authorization"`
	Type   string `json:"type" default:"Bearer"`
	Token  string `json:"token"`
}
