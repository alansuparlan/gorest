package repository

import (
	"context"
	"graphql/gorest/entity"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct{}

//NewFireStoreRepository for
func NewFireStoreRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "alsu-2d5a1"
	collectionName string = "posts"
)

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a FireStore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a FireStore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	doneSymbol := iterator.Done
	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()

		if err == doneSymbol {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)

	}

	return posts, nil
}

// //Delete: TODO
// func (r *repo) Delete(post *entity.Post) error {
// 	return nil
// }
