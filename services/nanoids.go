package services

import g "github.com/matoous/go-nanoid/v2"

func GenID() string {
	id, _ := g.New(36)
	return id
}
