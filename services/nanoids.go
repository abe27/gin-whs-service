package services

import g "github.com/matoous/go-nanoid/v2"

func GenID() string {
	id, _ := g.New(21)
	return id
}

func TokenID() string {
	id, _ := g.New(65)
	return id
}
