package twitter

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// Service is the interface that provides tweet methods.
type Service interface {
	FollowUser(context.Context, User) error
	ListUser(context.Context) ([]User, error)
	FollowHome(context.Context, User) error
	ListHome(context.Context) ([]User, error)
}

// NewService returns a new valid twitter service.
func NewService() Service {
	s := &service{}
	http.Handle(
		"/twitter/user/follow",
		httptransport.NewServer(
			followUser(s),
			decode,
			encode,
		),
	)
	return s
}

type service struct {
}

// FollowUser adds a new user to follow and run fetching daemon.
func (s *service) FollowUser(ctx context.Context, user User) error {
	return nil
}

// ListHome list followed users.
func (s *service) ListUser(context.Context) ([]User, error) {
	return nil, nil
}

// FollowHome adds a new home to follow and run fetching daemon.
func (s *service) FollowHome(ctx context.Context, user User) error {
	return nil
}

// ListHome list followed home users.
func (s *service) ListHome(context.Context) ([]User, error) {
	return nil, nil
}
