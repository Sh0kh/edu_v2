package utils

import (
	"edu_v2/graph/model"
)

func AbsResponseChecking(err error, msg string) (*model.Response, error) {
	if err != nil {
		return &model.Response{
			StatusCode: 409,
			Message:    err.Error(),
		}, nil
	}
	return &model.Response{
		StatusCode: 200,
		Message:    msg,
	}, nil
}

type Response struct {
	UserID int `json:"user_id"`
	Code   int `json:"code"`
}

func OffSetGenerator(page, size *int) int {
	if page == nil || *page < 1 {
		p := 1
		page = &p
	}
	if size == nil || *size < 1 {
		s := 10
		size = &s
	}

	return *size * (*page - 1)
}
