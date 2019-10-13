package shortenurl

import (
	"math/rand"
	"regexp"
	"time"
)

var dbInst shortenUrlDBIntf

func SaveUrl(url, sc string) (string, *Error) {

	//bussienss logic

	db := getDbInstance()
	var validID = regexp.MustCompile(`^[0-9a-zA-Z_]{6}$`)

	if sc != "" {
		if validID.MatchString(sc) {

		} else {
			//error
			return "", &Error{
				Code:    1505,
				Message: "wrong regex",
			}
		}
	} else {
		sc = getRandom(6)
		if validID.MatchString(sc) {

		} else {
			return "", &Error{
				Code:    1505,
				Message: "wrong regex",
			}
		}
	}
	err := db.saveUrl(url, sc)
	if err != nil {
		if err.Error() == "string exists" {
			return "", &Error{
				Code:    1501,
				Message: "already in use",
			}
		}
	}
	return sc, nil
}

func GetUrl(sc string) (string, *Error) {
	i := getDbInstance()
	url, err := i.getUrl(sc)
	if err != nil && err.Error() == "Not Exists" {
		return "", &Error{
			Code:    1601,
			Message: "Not found",
		}
	}
	return url, nil
}

func GetStats(sc string) (Stats, *Error) {
	i := getDbInstance()
	st, err := i.getStats(sc)
	if err != nil && err.Error() == "Not Exists" {
		return Stats{}, &Error{
			Code:    1601,
			Message: "Not found",
		}
	}
	return st, nil
}

func getDbInstance() shortenUrlDBIntf {

	if dbInst == nil {
		data := make(map[string]string)
		stats := make(map[string]Stats)
		dbInst = inMemoryDatabase{
			data:  data,
			stats: stats,
		}
	}
	return dbInst
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func getRandom(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
