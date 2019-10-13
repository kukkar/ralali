package shortenurl

import (
	"fmt"
	"time"
)

type inMemoryDatabase struct {
	data  map[string]string
	stats map[string]Stats
}

func (this inMemoryDatabase) saveUrl(url, sc string) error {

	err := this.validateUrl(sc)
	if err != nil {
		return err
	}
	this.data[sc] = url
	var st = Stats{
		StartDate:     time.Now(),
		LastSeenTime:  time.Now(),
		RedirectCount: 0,
	}
	this.stats[sc] = st
	return nil
}

func (this inMemoryDatabase) getUrl(sc string) (string, error) {

	if val, ok := this.data[sc]; ok {
		return val, nil
	}
	return "", fmt.Errorf("Not Exists")
}

func (this inMemoryDatabase) getStats(sc string) (Stats, error) {
	if val, ok := this.stats[sc]; ok {
		return val, nil
	}
	return Stats{}, fmt.Errorf("Not Exists")
}

func (this inMemoryDatabase) validateUrl(sc string) error {

	if _, ok := this.data[sc]; ok {
		return fmt.Errorf("string exists")
	}
	return nil
}
