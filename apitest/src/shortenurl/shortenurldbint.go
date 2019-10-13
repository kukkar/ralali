package shortenurl

type shortenUrlDBIntf interface {
	saveUrl(url string, sc string) error
	getUrl(sc string) (string, error)
	getStats(sc string) (Stats, error)
}
