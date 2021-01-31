package repository

import "graphql/gorest/entity"

// PostRepository for
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
