package service

import (
	"errors"
	"graphql/gorest/entity"
	"graphql/gorest/repository"
	"math/rand"
)

// PostService for
type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct {
	repo repository.PostRepository
}

// NewPostService for
func NewPostService(repo repository.PostRepository) PostService {
	return &service{
		repo: repo,
	}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The Post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}

	return nil

}
func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return s.repo.Save(post)
}
func (s *service) FindAll() ([]entity.Post, error) {
	return s.repo.FindAll()
}
