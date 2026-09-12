package req

import (
	"net/http"
	"todo/pkg/res"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.JSONWrite(w, err, http.StatusBadRequest)
		return nil, err
	}

	if err := IsValid(body); err != nil {
		res.JSONWrite(w, err, http.StatusBadRequest)
		return nil, err
	}

	return body, nil
}
