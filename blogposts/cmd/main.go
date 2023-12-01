package main

import (
	"fmt"
	"log"
	"os"

	blogposts "github.com/aviralvardiya/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	for _,post:=range posts{

		fmt.Println("##############################")
		fmt.Println(post.Title)
		fmt.Println(post.Description)
		fmt.Println(post.Body)
	}
}