package comment

import (
	"context"
	"fmt"
)

// Comment - a representation of comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all the methods
// that service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct on which system
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new
// service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}
