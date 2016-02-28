package users

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

/*
Each user is composed of a first name, last name,
email, age.
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
I'll use this method when displaying all users for
"GET /users".
*/
func fetchAllUsers(db *mgo.Database) []User {
	users := []User{}
	err := db.C("users").Find(nil).All(&users)
	if err != nil {
		panic(err)
	}

	return users
}