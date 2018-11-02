package model

import (
	"github.com/mvl-at/qbs"
	"strconv"
	"strings"
	"time"
)

//Defines a member.
type Member struct {
	Id           int64     `json:"id"`
	FirstName    string    `json:"firstName" roles:"member"`
	LastName     string    `json:"lastName" roles:"member"`
	Joined       int       `json:"joined" roles:"member"`
	Picture      string    `json:"picture" roles:"member"`
	Active       bool      `json:"active" roles:"member"`
	Birthday     time.Time `json:"birthday" roles:"member"`
	LoginAllowed bool      `json:"loginAllowed"`
	Username     string    `json:"username"`
	Password     string    `json:"-"`
	Gender       rune      `json:"gender"`

	Instrument   *Instrument `json:"instrument" roles:"member"`
	InstrumentId int64       `qbs:"fk:Instrument" json:"instrumentId" roles:"member"`
}

//Validates all association pointers and assign its id fields to the one of Member.
func (m *Member) Validate(qbs *qbs.Qbs) error {

	if m.Instrument != nil {
		m.InstrumentId = m.Instrument.Id
	}

	if len(m.Username) == 0 {
		m.Username = identifier(m)
	}

	if len(m.Picture) == 0 {
		m.Picture = identifier(m)
	}
	return nil
}

//generates human readable identifier
func identifier(m *Member) string {
	replacements := map[string]string{
		"ä": "ae",
		"ö": "oe",
		"ü": "ue",
		"ß": "s"}
	identifier := strings.ToLower(strings.Join([]string{m.FirstName, m.LastName, strconv.Itoa(m.Joined)}, "-"))
	for k, v := range replacements {
		identifier = strings.Replace(identifier, k, v, -1)
	}
	return identifier
}
