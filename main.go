package main

import (
	"fmt"
	"graphql/gorest/controller"
	router "graphql/gorest/http"
	"graphql/gorest/repository"
	"graphql/gorest/service"
	"net/http"
)

var (
	postRepository repository.PostRepository = repository.NewFireStoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "up and running. . .")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
