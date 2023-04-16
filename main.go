package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/sebastian009w/go_run/message"
	post "github.com/sebastian009w/go_run/posts"
)

func main() {

	file, err := os.OpenFile("posts.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var posts []post.Post

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)

		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &posts)
		if err != nil {
			panic(err)
		}
	} else {
		posts = []post.Post{}
	}
	if len(os.Args) < 2 {
		message.Message_init()
	}

	if os.Args[1] == "--help" {
		message.Help_me()
	}
}
