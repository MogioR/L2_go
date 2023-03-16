package celenderdomain

import "time"

type Event struct {
	Uid   string `json:"id"`
	Date  string `json:"date"`
	Title string `json:"title"`
	User  int    `json:"user"`
}

func (e *Event) Validate() bool {
	_, err := time.Parse("2006-01-02", e.Date)
	return len(e.Uid) == 36 && err == nil
}
