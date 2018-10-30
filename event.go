package model

import (
	"github.com/mvl-at/qbs"
	"time"
)

//Defines an event.
type Event struct {
	Id            int64     `json:"id"`
	Date          time.Time `json:"date" roles:"event"`
	Time          time.Time `json:"time" roles:"event"`
	MusicianTime  time.Time `json:"musicianTime" roles:"event"`
	Name          string    `json:"name" roles:"event"`
	Note          string    `json:"note" roles:"event"`
	Uniform       string    `json:"uniform" roles:"event"`
	Place         string    `json:"place" roles:"event"`
	MusicianPlace string    `json:"musicianPlace" roles:"event"`
	Internal      bool      `json:"internal" roles:"event"`
	Important     bool      `json:"important" roles:"event"`
}

//Defines a declination per member per event
type Declination struct {
	Id       int64     `json:"id"`
	Event    *Event    `json:"event"`
	EventId  int64     `json:"eventId" qbs:"fk:Event"`
	Member   *Member   `json:"member"`
	MemberId int64     `json:"memberId" qbs:"fk:Member"`
	Time     time.Time `json:"time"`
	Declined bool      `json:"declined"`
}

//Validates all association pointers and assign its id fields to the one of LeaderRoleMember.
func (d *Declination) Validate(qbs *qbs.Qbs) error {

	if d.Member != nil {
		d.MemberId = d.Member.Id
	}

	if d.Event != nil {
		d.EventId = d.Event.Id
	}

	return nil
}
