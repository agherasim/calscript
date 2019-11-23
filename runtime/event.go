package runtime

import "time"

// Event type
type Event struct {
	Summary     string
	Description string
	Start       *EventDateTime
	End         *EventDateTime
	Recurrence  []string
	Attendees   []*EventAttendees
}

// EventDateTime type
type EventDateTime struct {
	DateTime time.Time
}

// EventAttendees type
type EventAttendees struct {
	Email string
}

// NewEvent factory
func NewEvent() *Event {
	return &Event{
		Recurrence: make([]string, 10),
		Attendees:  make([]*EventAttendees, 64),
	}
}

// SetDateInterval to a recurring event
func (ev *Event) SetDateInterval(start time.Time, end time.Time) {

}
