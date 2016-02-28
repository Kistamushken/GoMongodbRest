package users

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"
)

/*
Create a new *martini.ClassicMartini server.
We'll use a JSON renderer and our MongoDB
database handler. We define three routes:
"GET /users"
"POST /users"
"PUT /users"
*/
func NewServer(session *DatabaseSession) *martini.ClassicMartini {
	// Create the server and set up middleware.
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))
	m.Use(session.Database())

	// Define the "GET /users" route.
	m.Get("/users", func(r render.Render, db *mgo.Database) {
		r.JSON(200, fetchAllUsers(db))
	})

	// Define the "POST /users" route.
	m.Post("/users", binding.Json(User{}),
		func(user User,
		r render.Render,
		db *mgo.Database) {
			if user.valid() {
				// user is valid, insert into database
				err := db.C("users").Insert(user)
				if err == nil {
					// insert successful, 201 Created
					r.JSON(201, user)
				} else {
					// insert failed, 400 Bad Request
					r.JSON(400, map[string]string{
						"error": err.Error(),
					})
				}
			} else {
				// user is invalid, 400 Bad Request
				r.JSON(400, map[string]string{
					"error": "Not a valid user",
				})
			}
		})

	// Define the "PUT /users" route.
	m.Put("/users",binding.Json(User{}),
		func(user User,
		r render.Render,
		db *mgo.Database) {
			if user.valid() {
				change := mgo.Change{
					Update: user,
					Upsert: true,
					ReturnNew: true,
				}
				_, err := db.C("users").Find(user.Id).Apply(change, &user)
				if err == nil {
					// update successful, 201 Created
					r.JSON(201, user)
				} else {
					// insert failed, 400 Bad Request
					r.JSON(400, map[string]string{
						"error": err.Error(),
					})
				}
			} else {
				// user is invalid, 400 Bad Request
				r.JSON(400, map[string]string{
					"error": "Not a valid user",
				})
			}
		})

	// Return the server. Call Run() on the server to
	// begin listening for HTTP requests.
	return m
}