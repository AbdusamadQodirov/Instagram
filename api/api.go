package api

import (
	"fmt"
	"instagram/api/handler"
	"instagram/config"
	"net/http"
)

func Api(handler handler.Handler) {

	port := config.NewConfig().Server.Port

	// user's paths
	http.HandleFunc("/create-user", handler.CreateUser)
	http.HandleFunc("/get-users", handler.GetUsers)
	http.HandleFunc("/getuser-byid/", handler.GetUserById)
	http.HandleFunc("/update-user/", handler.UpdateUser)
	http.HandleFunc("/delete-user/", handler.DeleteUserById)

	// post paths
	http.HandleFunc("/create-post/", handler.CreatePost)
	http.HandleFunc("/get-posts", handler.GetPosts)
	http.HandleFunc("/getpost-byid/", handler.GetPostsById)
	http.HandleFunc("/update-post/", handler.UpdatePost)
	http.HandleFunc("/delete-post/", handler.DeletePost)

	// comment paths
	http.HandleFunc("/create-comment", handler.CreateComment)
	http.HandleFunc("/getcomment-bypost/", handler.GetCommentByPost)
	http.HandleFunc("/update-comment/", handler.UpdateComment)
	http.HandleFunc("/delete-comment/", handler.DeleteComment)

	fmt.Printf("Server is running on port %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err)
		return
	}

}
