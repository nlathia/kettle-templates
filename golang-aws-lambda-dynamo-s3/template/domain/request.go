package domain

import "errors"

type RequestEvent struct {
	UserName string `json:"user_name"`
}

type ResponseEvent struct {
	Message string `json:"message"`
}

func ValidateRequest(req RequestEvent) error {
	switch {
	case req.UserName == "":
		return errors.New("user_name is missing")
	}
	return nil
}
