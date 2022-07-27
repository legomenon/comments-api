package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
	// CreateComment(ctx context.Context, cmt Comment) (Comment, error)
	// UpdateComment(ctx context.Context, cmt Comment) error
	// DeleteComment(ctx context.Context, id string) error
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("returning comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return nil
}
