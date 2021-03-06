package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
	MONGO_EVENT = "events"
)

type Event struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"Name"`

	Leader    string   `bson:"Leader"`    // telephone#
	Attendees []string `bson:"Attendees"` // telephone#s
}

func (this *Event) Update() {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB(MONGO_DB).C(MONGO_EVENT)

	err = c.Update(bson.M{"_id": this.Id}, bson.M{"$set": bson.M{"Name": this.Name, "Leader": this.Leader, "Attendees": this.Attendees}})

	if err != nil {
		panic(err)
	}
}

func FindEventsByAttendee(number string) []string {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB(MONGO_DB).C(MONGO_EVENT)

	var result []Event
	err = c.Find(bson.M{"Attendees": number}).All(&result)
	if err != nil {
		panic(err)
	}

	ids := make([]string, 0)
	for _, v := range result {
		ids = append(ids, v.Id.Hex())
	}

	return ids
}

func FindEventsByLeader(number string) []string {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB(MONGO_DB).C(MONGO_EVENT)

	var result []Event
	err = c.Find(bson.M{"Leader": number}).All(&result)
	if err != nil {
		panic(err)
	}

	ids := make([]string, 0)
	for _, v := range result {
		ids = append(ids, v.Id.Hex())
	}

	return ids
}

func FindEvent(id string) *Event {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB(MONGO_DB).C(MONGO_EVENT)

	result := Event{}
	err = c.FindId(bson.ObjectIdHex(id)).One(&result)
	if err == mgo.ErrNotFound {
		return nil
	} else if err != nil {
		panic(err)
	}

	return &result
}

type AllEventsReturn struct {
	Id   string
	Name string
}

func FindAllEvents() []AllEventsReturn {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB(MONGO_DB).C(MONGO_EVENT)

	var results []*Event
	err = c.Find(nil).Select(bson.M{"Leader": 0, "Attendees": 0}).All(&results)
	if err != nil && err != mgo.ErrNotFound {
		panic(err)
	}

	ids := make([]AllEventsReturn, 0)
	for _, v := range results {
		ids = append(ids, AllEventsReturn{v.Id.Hex(), v.Name})
	}

	return ids
}

func CreateEvent(name, leader string) *Event {
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Strong, true)

	c := session.DB(MONGO_DB).C(MONGO_EVENT)

	result := &Event{Id: bson.NewObjectId(), Name: name, Leader: leader}
	err = c.Insert(result)
	if err != nil {
		panic(err)
	}

	return result
}
