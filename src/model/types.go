package model

import uuid "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID
	Name string
}

type Item struct {
	ID   uuid.UUID
	Name string
}
