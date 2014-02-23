package models

import (
	"container/list"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
	MONGO_USER = "users"
)

var (
	socketList *list.List
)

func init() {
	socketList = list.New()
}

func RegisterSocket(c chan *User) {
	socketList.PushBack(c)
}

type User struct {
	Number string `bson:"Number"` // +[countrycode][number]
	Name   string `bson:"Name"`
	Score  uint64 `bson:"Score"`
}

func (this *User) Update() {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB(MONGO_DB).C(MONGO_USER)

	err = c.Update(bson.M{"Number": this.Number}, bson.M{"$set": bson.M{"Score": this.Score, "Name": this.Name}})

	if err != nil {
		panic(err)
	}

	go func() {
		for e := socketList.Front(); e != nil; e = e.Next() {
			e.Value.(chan *User) <- this
			socketList.Remove(e)
		}
	}()
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
		result.Score = 0
		result.Name = "Meeter"
		err = c.Insert(&result)
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
