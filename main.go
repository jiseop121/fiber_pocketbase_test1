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
	Role     string `json:"role"`
}

var subinId = "kitk1setevfrv"
var jiseopId = "9t2x1sk39d49m27"

func main() {
	client, postCollection, userCollection := pocketbaseNewClient()

	// [RECORD CRUD]
	// List with pagination
	err := postsViewWithOptions(postCollection)

	// [AUTH]
	//err = postCollection.RequestVerification("auth@naver.com")
	//if err != nil {
	//	log.Println("postCollection.RequestVerification()")
	//	log.Fatal(err)
	//}

	// authwithpassword 후에 posts 다루기
	// posts view rule : @request.auth.email = user.email 적용
	// rules : id = @request.auth.id
	password := authWithPasswordRoleGold(err, userCollection)

	// [Auth]
	// users crud rules : id = @request.auth.id
	createPost(err, postCollection)

	//&& @request.data.user.role = "gold" _> error
	err = authWithPasswordRoleSilver(password, err, userCollection)

	// auth 로그인 하면서 newClient
	userCollection = pocketbaseNewClientWithAdminEmailPassword(client, postCollection, userCollection)

	// [Auth]
	// users crud rules : id = @request.auth.id

}

func pocketbaseNewClientWithAdminEmailPassword(client *pocketbase.Client, postCollection *pocketbase.Collection[Post], userCollection *pocketbase.Collection[User]) *pocketbase.Collection[User] {
	fmt.Println("new client with auth")
	client = pocketbase.NewClient(
		"http://localhost:8090",
		pocketbase.WithAdminEmailPassword("admin@naver.com", "123412341234"),
	)
	postCollection = pocketbase.CollectionSet[Post](client, "posts")
	userCollection = pocketbase.CollectionSet[User](client, "users")
	return userCollection
}

func authWithPasswordRoleSilver(password pocketbase.AuthWithPasswordResponse, err error, userCollection *pocketbase.Collection[User]) error {
	password, err = userCollection.AuthWithPassword(
		"user2@naver.com", //subin email
		"123412341234",
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func createPost(err error, postCollection *pocketbase.Collection[Post]) {
	posts, err := postCollection.Create(
		Post{
			Title:   "new_title",
			Content: "new content",
			User:    jiseopId,
		},
	)
	if err != nil {
		log.Println("postCollection.Create  ERROR")
		log.Fatal(err)
	}
	log.Println("postCollection.Create ok")
	log.Println(posts.ID)
}

func authWithPasswordRoleGold(err error, userCollection *pocketbase.Collection[User]) pocketbase.AuthWithPasswordResponse {
	password, err := userCollection.AuthWithPassword(
		"user1@naver.com", //jiseop email
		"123412341234",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("userCollection.AuthWithPassword auth ok")
	fmt.Println("password:", password.Record.ID)
	return password
}

func postsViewWithOptions(postCollection *pocketbase.Collection[Post]) error {
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
	return err
}

func pocketbaseNewClient() (*pocketbase.Client, *pocketbase.Collection[Post], *pocketbase.Collection[User]) {
	client := pocketbase.NewClient("http://localhost:8090")
	postCollection := pocketbase.CollectionSet[Post](client, "posts")
	userCollection := pocketbase.CollectionSet[User](client, "users")
	return client, postCollection, userCollection
}
