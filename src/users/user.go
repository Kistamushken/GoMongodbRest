package users

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

/*
Each signature is composed of a first name, last name,
email, age, and short message. When represented in
JSON, ditch TitleCase for snake_case.
*/
type User struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func (user *User) valid() bool {
	return len(user.FirstName) > 0 &&
	len(user.LastName) > 0 &&
	len(user.Email) > 0 &&
	user.Age >= 18 && user.Age <= 180
}

/*
I'll use this method when displaying all signatures for
"GET /signatures". Consult the mgo docs for more info:
http://godoc.org/labix.org/v2/mgo
*/
func fetchAllUsers(db *mgo.Database) []User {
	users := []User{}
	err := db.C("users").Find(nil).All(&users)
	if err != nil {
		panic(err)
	}

	return users
}