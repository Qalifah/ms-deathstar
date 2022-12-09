package deathstar

import "time"

type Target struct {
	Id        string     `json:"id"`
	Message   string     `json:"message"`
	CreatedOn time.Time  `json:"created_on" bson:"created_on"`
	UpdatedOn *time.Time `json:"updated_on,omitempty" bson:"updated_on,omitempty"`
}

type Event struct {
	Targets []Target `json:"targets" fakesize:"5"`
}
