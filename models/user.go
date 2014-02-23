package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
	MONGO_USER = "users"
)

type User struct {
	Number string   `bson:"Number"` // +[countrycode][number]
	Events []uint64 `bson:"Events"`
}

func (this *User) Update() {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB(MONGO_DB).C(MONGO_USER)

	err = c.Update(bson.M{"Number": this.Number}, bson.M{"$set": bson.M{"Events": this.Events}})

	if err != nil {
		panic(err)
	}
}

// Find a user. If the user does not exist,
// create the user.
func FindUser(number string) *User {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB(MONGO_DB).C(MONGO_USER)

	result := User{}
	err = c.Find(bson.M{"Number": number}).One(&result)
	if err == mgo.ErrNotFound {
		result.Number = number
		result.Events = make([]uint64, 0)
		err = c.Insert(bson.M{"Number": number, "Events": make([]uint64, 0)})
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		result.Number = number
	}

	return &result
}
