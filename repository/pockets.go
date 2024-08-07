package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pluja/pocketbase"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
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

type Replies struct {
	ID      string `json:"id"`
	Replies string `json:"replies"`
	User    string `json:"user"`
	Post    string `json:"post"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

var SubinId = "kitk1setevfrv"
var JiseopId = "9t2x1sk39d49m27"
var PClient = pocketbase.NewClient("http://localhost:8090")
var UserCollection = pocketbase.CollectionSet[User](PClient, "users")
var PostCollection = pocketbase.CollectionSet[Post](PClient, "Posts")
var RepliesCollection = pocketbase.CollectionSet[Replies](PClient, "Replies")
var UserToken = ""

func PocketbaseNewClientWithAdminEmailPassword(client *pocketbase.Client, postCollection *pocketbase.Collection[Post], userCollection *pocketbase.Collection[User]) *pocketbase.Collection[User] {
	fmt.Println("new client with auth")
	client = pocketbase.NewClient(
		"http://localhost:8090",
		pocketbase.WithAdminEmailPassword("admin@naver.com", "123412341234"),
	)

	postCollection = pocketbase.CollectionSet[Post](client, "posts")
	userCollection = pocketbase.CollectionSet[User](client, "users")
	return userCollection
}

func AuthWithPasswordRoleSilver(password pocketbase.AuthWithPasswordResponse, err error, userCollection *pocketbase.Collection[User]) error {
	password, err = userCollection.AuthWithPassword(
		"user2@naver.com", //subin email
		"123412341234",
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func CreatePost(err error, postCollection *pocketbase.Collection[Post]) {
	posts, err := postCollection.Create(
		Post{
			Title:   "new_title",
			Content: "new content",
			User:    JiseopId,
		},
	)
	if err != nil {
		log.Println("postCollection.Create  ERROR")
		log.Fatal(err)
	}
	log.Println("postCollection.Create ok")
	log.Println(posts.ID)
}

func UserAuthInMyAppWithEmail(username string, password string) (pocketbase.AuthWithPasswordResponse, error) {
	userAuthData, err := UserCollection.AuthWithPassword(
		username,
		password,
	)
	if err != nil {
		return userAuthData, err
	}
	//log.Println("userCollection.AuthWithPassword auth ok")
	//log.Println("user id :", userAuthData.Record.ID)
	return userAuthData, err
}
func UserAuthInSMyAppOauth() {
	// API 엔드포인트
	url := "http://localhost:8090/api/collections/users/auth-with-oauth2"

	// 요청에 사용할 데이터
	payload := map[string]string{
		"provider":     "google",
		"code":         "your_authorization_code",
		"codeVerifier": "your_code_verifier",
		"redirectUrl":  "http://127.0.0.1:8090/api/oauth2-redirect",
	}

	// JSON으로 변환
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// HTTP POST 요청 생성
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// 클라이언트 생성 및 요청 전송
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// 응답 출력
	fmt.Println("Response Status:", resp.Status)
}

func UserAuthInMyAppWithOauth() (pocketbase.AuthWithOauth2Response, error) {
	userAuthData, err := UserCollection.AuthWithOAuth2Code(
		"google",
		"google1",
		"123412341234",
		"http://localhost:8090/api/oauth2-redirect",
	)
	if err != nil {
		return userAuthData, err
	}
	//log.Println("userCollection.AuthWithPassword auth ok")
	//log.Println("user id :", userAuthData.Record.ID)
	return userAuthData, err
}

func AuthWithPasswordRoleGold(userCollection *pocketbase.Collection[User]) (pocketbase.AuthWithPasswordResponse, error) {
	password, err := userCollection.AuthWithPassword(
		"user1123@naver.com", //jiseop email
		"123412341234",
	)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("userCollection.AuthWithPassword auth ok")
	fmt.Println("password:", password.Record.ID)
	return password, err
}

func ReliesListSearchWithOptionsWithToken(token string, paramsList pocketbase.ParamsList) ([]byte, error) {
	baseURL := "http://localhost:8090/api/collections/replies/records"

	urlWithParams, err := BaseUrlAddParams(baseURL, paramsList)

	log.Println("POST :", urlWithParams)

	//c.Client.List(c.Name, params)

	// 새로운 요청 생성 (GET 요청)
	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// 헤더 설정

	reqSetBasic(req)

	req.Header.Set("Authorization", token) // 실제 토큰 값으로 대체

	// 클라이언트 생성 및 요청 전송
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}

	// 응답 데이터 읽기
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	if err != nil {
		return body, err
	}
	return body, err
}

func PostsListSearchWithOptionsWithToken(token string, paramsList pocketbase.ParamsList) ([]byte, error) {

	baseURL := "http://localhost:8090/api/collections/posts/records"

	urlWithParams, err := BaseUrlAddParams(baseURL, paramsList)

	log.Println("POST :", urlWithParams)

	//c.Client.List(c.Name, params)

	// 새로운 요청 생성 (GET 요청)
	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// 헤더 설정

	reqSetBasic(req)

	req.Header.Set("Authorization", token) // 실제 토큰 값으로 대체

	// 클라이언트 생성 및 요청 전송
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}

	// 응답 데이터 읽기
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	if err != nil {
		return body, err
	}
	return body, err
}

func BaseUrlAddParams(baseURL string, paramsList pocketbase.ParamsList) (string, error) {
	parsedURL, err := url.Parse(baseURL)

	query := parsedURL.Query()

	if paramsList.Page > 0 {
		query.Set("page", strconv.Itoa(paramsList.Page))
	}
	if paramsList.Size > 0 {
		query.Set("perPage", strconv.Itoa(paramsList.Size))
	}
	if paramsList.Filters != "" {
		query.Set("filter", paramsList.Filters)
	}
	if paramsList.Sort != "" {
		query.Set("sort", paramsList.Sort)
	}
	if paramsList.Expand != "" {
		query.Set("expand", paramsList.Expand)
	}
	if paramsList.Fields != "" {
		query.Set("fields", paramsList.Fields)
	}
	return baseURL + "?" + query.Encode(), err
}

func reqSetBasic(req *http.Request) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
}

func PocketbaseNewClient() {
	PClient = pocketbase.NewClient("http://localhost:8090")
	PostCollection = pocketbase.CollectionSet[Post](PClient, "posts")
	UserCollection = pocketbase.CollectionSet[User](PClient, "users")
}
