package shortenpost

type Request struct {
	Url       string `json:"url"`
	ShortCode string `json:"shortcode"`
}

type Response struct {
	ShortCode string `json:"shortcode"`
}
