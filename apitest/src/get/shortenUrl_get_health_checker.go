package get

type ShortenUrlHealthCheck struct {
}

func (n ShortenUrlHealthCheck) GetName() string {
	return "hello world"
}

func (n ShortenUrlHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
