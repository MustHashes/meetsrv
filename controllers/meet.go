package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"github.com/gorilla/websocket"

	"github.com/MustHashes/meetsrv/models"
)

type MeetController struct {
	BaseController
}

func (this *MeetController) Get() {
	id := this.Ctx.Input.Param(":id")
	event := models.FindEvent(id)
	if event == nil {
		this.Abort("404")
	}

	this.Data["json"] = event
	this.ServeJson()
}

func (this *MeetController) List() {
	this.Data["json"] = models.FindAllEvents()
	this.ServeJson()
}

type MeetJoinRequest struct {
	User string // phone number in format +[countrycode][number]
}

func (this *MeetController) Join() {
	id := this.Ctx.Input.Param(":id")
	event := models.FindEvent(id)
	if event == nil {
		this.Abort("404")
	}

	var req MeetJoinRequest
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Abort("400")
	}

	user := models.FindUser(req.User)

	if user.Number == event.Leader {
		this.Abort("403")
	}
	for _, v := range event.Attendees {
		if v == user.Number {
			this.Abort("403") // don't let them join twice!
		}
	}

	user.Score++
	event.Attendees = append(event.Attendees, user.Number)

	event.Update()
	user.Update()

	this.Ctx.WriteString("OK")
}

type MeetSeenRequest struct {
	User string // phone number in format +[countrycode][number]
}

func (this *MeetController) Seen() {
	id := this.Ctx.Input.Param(":id")
	event := models.FindEvent(id)
	if event == nil {
		this.Abort("404")
	}

	var req MeetSeenRequest
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Abort("400")
	}

	user := models.FindUser(req.User)

	if user.Number != event.Leader {
		got := false
		for _, v := range event.Attendees {
			if v == user.Number {
				got = true
				break
			}
		}
		if !got {
			this.Abort("403") // don't want people who aren't in the event joining!
		}
	}

	user.Score++
	user.Update()

	this.Ctx.WriteString("OK")
}

func (this *MeetController) Socket() {
	id := this.Ctx.Input.Param(":id")
	event := models.FindEvent(id)
	if event == nil {
		this.Abort("404")
	}

	conn, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		this.Abort("400")
		return
	} else if err != nil {
		panic(err)
	}

	for {
		ch := make(chan *models.User)
		needsClosing := true
		defer func() {
			if needsClosing {
				close(ch)
			}
		}()

		models.RegisterSocket(ch)

		user := <-ch

		var w io.WriteCloser
		var bte []byte
		var err error

		if event.Leader != user.Number {
			isAttendee := false
			for _, v := range event.Attendees {
				if v == user.Number {
					isAttendee = true
					break
				}
			}
			if !isAttendee {
				goto CLEAN
			}
		}

		if w, err = conn.NextWriter(websocket.TextMessage); err != nil {
			return
		}

		if bte, err = json.MarshalIndent(user, "", "\t"); err != nil {
			panic(err)
		}
		if _, err := fmt.Fprint(w, bte); err != nil {
			return
		}
		if err := w.Close(); err != nil {
			return
		}

	CLEAN:
		close(ch)
		needsClosing = false
	}

}
