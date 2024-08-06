package main

import (
	"fmt"
	"github.com/pluja/pocketbase"
	"log"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	User    string `json:"user"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

var subinId = "kitk1setevfrv"
var jiseopId = "9t2x1sk39d49m27"

func main() {
	client := pocketbase.NewClient("http://localhost:8090")
	postCollection := pocketbase.CollectionSet[Post](client, "posts")
	userCollection := pocketbase.CollectionSet[User](client, "users")

	// [RECORD CRUD]
	// List with pagination
	response, err := postCollection.List(pocketbase.ParamsList{
		Page: 1, Size: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Items)

	// FullList also available for collections:
	response, err = postCollection.FullList(pocketbase.ParamsList{
		Sort: "-created",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Items)

	log.Printf("%+v", response.Items)

	// [AUTH]
	//err = postCollection.RequestVerification("auth@naver.com")
	//if err != nil {
	//	log.Println("postCollection.RequestVerification()")
	//	log.Fatal(err)
	//}

	// authwithpassword 후에 posts 다루기
	// posts view rule : @request.auth.email = user.email 적용
	password, err := userCollection.AuthWithPassword(
		"user1@naver.com", //jiseop email
		"123412341234",
	)
	if err != nil {
		return
	}

	fmt.Println("userCollection.AuthWithPassword auth ok")
	fmt.Println("password:", password.Record.ID)

	// [Auth]
	// users crud rules : id = @request.auth.id
	userList, err := userCollection.FullList(pocketbase.ParamsList{
		Page: 1, Size: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userList.Items)

	err = postCollection.Update(
		"wp8qnf4gagtbfyh",
		Post{
			Title:   "new title",
			Content: "new content",
			User:    subinId,
		},
	)
	if err != nil {
		log.Println("postCollection.Update  ERROR")
		log.Fatal(err)
	}
	log.Println("postCollection.Update ok")

	//&& @request.data.user.role = "gold" _> error
	password, err = userCollection.AuthWithPassword(
		"user2@naver.com", //subin email
		"123412341234",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = postCollection.Update(
		"wp8qnf4gagtbfyh",
		Post{
			Title:   "new title",
			Content: "new content",
			User:    subinId,
		},
	)
	if err != nil {
		log.Println("postCollection.Update  ERROR")
		log.Fatal(err)
	}

	// auth 로그인 하면서 newClient
	fmt.Println("new client with auth")
	client = pocketbase.NewClient(
		"http://localhost:8090",
		pocketbase.WithAdminEmailPassword("admin@naver.com", "123412341234"),
	)
	postCollection = pocketbase.CollectionSet[Post](client, "posts")
	userCollection = pocketbase.CollectionSet[User](client, "users")

	// [Auth]
	// users crud rules : id = @request.auth.id
	userList, err = userCollection.FullList(pocketbase.ParamsList{
		Page: 1, Size: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userList.Items)
}
