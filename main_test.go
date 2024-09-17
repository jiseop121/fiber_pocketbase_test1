package main

import (
	"fmt"
	"github.com/pluja/pocketbase"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 * 해당 테스트는 pluja의 기능 설명서에 가까움
 */

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
	Role     string `json:"role"`
	AnyCol   string `json:"any_col"`
}

func getAdminClient(url, email, password string) *pocketbase.Client {
	return pocketbase.NewClient(
		url,
		pocketbase.WithAdminEmailPassword(
			email, password),
	)
}

func getUserClient(url, email, password string) *pocketbase.Client {
	return pocketbase.NewClient(
		url,
		pocketbase.WithUserEmailPassword(
			email, password),
	)
}

func Test_CollectionSet_대소문자구분없이정상생성(t *testing.T) {
	client := getAdminClient(
		"http://localhost:8090",
		"admin@naver.com",
		"123412341234",
	)

	collectionName := "USERS"

	usersCollection := pocketbase.CollectionSet[User](client, collectionName)

	//생성된 collectionSet의 이름은 작성자가 넣은 이름에 의존함
	assert.Equal(t, usersCollection.Name, collectionName)

	list, err := usersCollection.List(
		pocketbase.ParamsList{
			Sort: "created",
		})
	if err != nil {
		assert.FailNow(t, err.Error())
		return
	}
	//올바르게 값들이 반환됨
	assert.NotEmpty(t, list)

	//포켓베이스에 실제로 없는 값들은 빈 값이 들어가고, struct에 없는 컬럼은 들어가지 않음
	//결론적으로 오류 없이 반환됨
	assert.Empty(t, list.Items[0].AnyCol)
	fmt.Println("collection set list:", list)
	fmt.Println(list.Items[0].AnyCol)
}

// authorize를 수행하지 않으면 초기에 토큰값이 들어있지 않음
func Test_Admin_Token(t *testing.T) {
	client := getAdminClient(
		"http://localhost:8090",
		"admin@naver.com",
		"123412341234",
	)

	token := client.AuthStore().Token()
	assert.Empty(t, token)

	client.Authorize()

	token = client.AuthStore().Token()
	fmt.Println(token)
	assert.NotEmpty(t, token)
}

func Test_User_Token(t *testing.T) {
	adminClient := getAdminClient(
		"http://localhost:8090",
		"admin@naver.com",
		"123412341234",
	)

	usersCollectionSet := pocketbase.CollectionSet[User](adminClient, "Users")

	jsAuth, err := usersCollectionSet.AuthWithPassword(
		"js@naver.com",
		"123412341234")
	if err != nil {
		assert.FailNow(t, err.Error())
		return
	}

	fmt.Println(jsAuth.Record)
	assert.NotEmpty(t, jsAuth.Record)

	fmt.Println(jsAuth.Token)
	assert.NotEmpty(t, jsAuth.Token)

}

// test_collection_access
// Post List/search api rules : @request.auth.email = "js@naver.com"
func Test_Collection_OnlyJsAccess(t *testing.T) {
	jsClient := getUserClient(
		"http://localhost:8090",
		"js@naver.com",
		"123412341234")

	postCollectionSet := pocketbase.CollectionSet[Post](jsClient, "Post")

	list, err := postCollectionSet.List(pocketbase.ParamsList{})
	if err != nil {
		assert.FailNow(t, err.Error())
		return
	}

	fmt.Println(list)
}

// status : 403, Only admins can perform this action 에러 메세지 반환
func Test_Collection_Denied_Only_Admins(t *testing.T) {
	jsClient := getUserClient(
		"http://localhost:8090",
		"sb@naver.com",
		"123412341234")

	postCollectionSet := pocketbase.CollectionSet[Post](jsClient, "Post")

	_, err := postCollectionSet.Create(Post{
		Title:   "new",
		Content: "new content",
	})
	if err != nil {
		assert.Error(t, err)
		fmt.Println(err.Error())
		return
	}
	assert.FailNow(t, "Test failed")
}

// api과 맞지 않은 경우 빈 값, 혹은 api 룰과 맞는 값들만 필터링되어서 반환됨
func Test_Collection_Denied_From_Api_Rules(t *testing.T) {
	jsClient := getUserClient(
		"http://localhost:8090",
		"sb@naver.com",
		"123412341234")

	postCollectionSet := pocketbase.CollectionSet[Post](jsClient, "Post")

	list, err := postCollectionSet.List(pocketbase.ParamsList{})
	if err != nil {
		assert.FailNow(t, err.Error())
		return
	}
	fmt.Println(list)
}

// collection 생성에 대한 api 호출를 다루는 함수는 존재하지 않는다.
// 다만 이미 생성된 collection에 대해서 record를 crud할 수 있는 함수가 존재한다.
func Test_Record_CRUD(t *testing.T) {
	adminClient := getAdminClient(
		"http://localhost:8090",
		"admin@naver.com",
		"123412341234",
	)

	// List/Search
	postCollectionSet := pocketbase.CollectionSet[Post](adminClient, "Post")

	list, err := postCollectionSet.List(pocketbase.ParamsList{
		Sort: "created",
	})
	if err != nil {
		assert.FailNow(t, err.Error())
		return
	}

	assert.NotEmpty(t, list)

	//Create
	create, err := postCollectionSet.Create(Post{
		Title:   "new",
		Content: "new content",
	})
	if err != nil {
		assert.FailNow(t, err.Error())
		return
	}
	assert.NotEmpty(t, create)

	//Update
	//이때 컬럼 값을 비우면, update가 되지 않는 것이 아니라, 빈 값이 체로 update가 된다.
	err = postCollectionSet.Update("aaaaaaaaaaaaaaa", Post{
		Title: "new",
	})
	if err != nil {
		assert.FailNow(t, err.Error())
		return
	}

	//Delete
	err = postCollectionSet.Delete("aaaaaaaaaaaaaaa")
	if err != nil {
		assert.FailNow(t, err.Error())
		return
	}
}
