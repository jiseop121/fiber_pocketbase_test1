package main

import (
	"encoding/json"
	"fiber_prac/dtos"
	repo "fiber_prac/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/pluja/pocketbase"
	"log"
)

func main() {

	app := fiber.New()
	defer app.Listen(":3000")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login/email", func(c *fiber.Ctx) error {
		var loginRequest dtos.LoginRequest
		err := c.BodyParser(&loginRequest)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
				"data":  loginRequest,
			})
		}

		response, err := repo.UserAuthInMyAppWithEmail(loginRequest.Username, loginRequest.Password)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
				"data":  loginRequest,
			})
		}

		repo.UserToken = response.Token
		c.Set("Authorization", response.Token)

		log.Println(response.Token)
		log.Println(response.Record.Email)
		log.Println(response.Record.ID)
		log.Println(response.Record.Username)

		return c.SendStatus(fiber.StatusAccepted)
	})

	app.Post("/signup/oauth", func(c *fiber.Ctx) error {

		repo.UserAuthInSMyAppOauth()

		//oauth, err := repo.UserAuthInMyAppWithOauth()
		//if err != nil {
		//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		//		"error": err.Error(),
		//	})
		//}
		//log.Println(oauth.Token)
		//log.Println(oauth)

		return c.SendStatus(fiber.StatusAccepted)
	})

	app.Post("/login/oauth", func(c *fiber.Ctx) error {
		var loginRequest dtos.LoginRequest
		err := c.BodyParser(&loginRequest)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
				"data":  loginRequest,
			})
		}

		response, err := repo.UserAuthInMyAppWithEmail(loginRequest.Username, loginRequest.Password)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
				"data":  loginRequest,
			})
		}

		repo.UserToken = response.Token
		c.Set("Authorization", response.Token)

		log.Println(response.Token)
		log.Println(response.Record.Email)
		log.Println(response.Record.ID)
		log.Println(response.Record.Username)

		return c.SendStatus(fiber.StatusAccepted)
	})

	// post
	app.Get("/posts", func(c *fiber.Ctx) error {

		if c.Get("Authorization") != repo.UserToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization token error",
				"Auth":  c.Get("Authorization"),
				"User":  repo.UserToken,
			})
		}
		newParamsList := pocketbase.ParamsList{
			Page:   1,
			Size:   10,
			Sort:   "created",
			Fields: "id, title, content, user",
		}

		//  "/api/collections/posts/records?page=1& ..."
		// net/http 사용
		body, err := repo.PostsListSearchWithOptionsWithToken(c.Get("Authorization"), newParamsList)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		//log.Println(json.Marshal(postsList.Items))
		return c.Status(fiber.StatusOK).JSON(result)

	})

	// replies
	app.Get("/replies", func(c *fiber.Ctx) error {
		if c.Get("Authorization") != repo.UserToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization token error",
				"Auth":  c.Get("Authorization"),
				"User":  repo.UserToken,
			})
		}

		newParamsList := pocketbase.ParamsList{
			Page: 1,
			Size: 10,
		}

		//  "/api/collections/posts/records?page=1& ..."
		// net/http 사용
		body, err := repo.ReliesListSearchWithOptionsWithToken(c.Get("Authorization"), newParamsList)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		//log.Println(json.Marshal(postsList.Items))
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Post("/logout", func(c *fiber.Ctx) error {

		repo.UserToken = ""

		return c.SendStatus(fiber.StatusOK)
	})

}

func RepliesListSearchWithOptions(c *fiber.Ctx) (pocketbase.ResponseList[repo.Replies], error) {
	repliesList, err := repo.RepliesCollection.List(pocketbase.ParamsList{
		Page: 1,
		Size: 10,
	})
	if err != nil {
		return repliesList, err
	}
	return repliesList, nil
}
