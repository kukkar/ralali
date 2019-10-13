package shortenurl

import "time"

type Error struct {
	Code    int
	Message string
}

func (this Error) String() {

}

type Stats struct {
	StartDate     time.Time
	LastSeenTime  time.Time
	RedirectCount int
}
