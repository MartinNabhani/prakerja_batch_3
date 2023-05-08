package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		defer resp.Body.Close()

		var posts []Post
		err = json.NewDecoder(resp.Body).Decode(&posts)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		return c.JSON(http.StatusOK, posts)
	})

	e.GET("/posts/:id", func(c echo.Context) error {
		id := c.Param("id")
		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id)

		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get post with id %s", id)
		}

		post := &struct {
			UserID int    `json:"userId"`
			ID     int    `json:"id"`
			Title  string `json:"title"`
			Body   string `json:"body"`
		}{}

		if err := c.Bind(post); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, post)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
