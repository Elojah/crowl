package twitter

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func decode(ctx context.Context, r *http.Request) (interface{}, error) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func encode(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func followUser(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		user := r.(User)
		_ = user
		return nil, nil
	}
}
