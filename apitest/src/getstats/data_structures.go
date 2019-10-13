package getstats

import "time"

type Response struct {
	StatDate      time.Time `json:"startDate"`
	LastSeenDate  time.Time `json:"lastSeenDate"`
	RedirectCount int       `json:"redirectCount"`
}
