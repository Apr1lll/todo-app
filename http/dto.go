package http

import (
	"encoding/json"
	"errors"
	"time"
)

// DTO: data transfer obj
type TaskDTO struct {
	Title       string
	Description string
}

type CompleteTaskDTO struct {
	Complete bool
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	return string(b)
}
