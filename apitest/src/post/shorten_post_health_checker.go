package shortenpost

type ShotenPostAPIHealthCheck struct {
}

func (n ShotenPostAPIHealthCheck) GetName() string {
	return "ShotenPostAPI"
}

func (n ShotenPostAPIHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
