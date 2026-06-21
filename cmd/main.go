package main

import (
	"github.com/hamidgh01/Go-Blog-API/cmd/commands"

	_ "github.com/hamidgh01/Go-Blog-API/docs"
)

// @title			Go Blog API
// @version         1.0
// @description     Go powered RESTful Blog API.

// @contact.name	Hamid Ghahremani
// @contact.url		https://github.com/hamidgh01
// @contact.email	hamidghahremani2001@gmail.com

// @license.name  MIT
// @license.url   https://mit-license.org/

// @tag.name auth
// @tag.name users
// @tag.name links
// @tag.name posts
// @tag.name tags
// @tag.name comments
// @tag.name lists

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token: "Bearer eyJhbGciOiJIUzI1Ni..."
func main() {
	commands.Execute()
}
