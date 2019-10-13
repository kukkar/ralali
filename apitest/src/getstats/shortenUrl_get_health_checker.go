package getstats

type ShortenUrlStatsHealthCheck struct {
}

func (n ShortenUrlStatsHealthCheck) GetName() string {
	return "hello world"
}

func (n ShortenUrlStatsHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
