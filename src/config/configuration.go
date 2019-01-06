package config

// Configuration represents database connection settings
type Configuration struct {
	Books struct {
		Host      string `json:"host"`
		Port      string `json:"port"`
		User      string `json:"user"`
		Password  string `json:"password"`
		Databasae string `json:"database"`
	} `json:"books"`
	Service struct {
		Port string `json:"port"`
	} `json:"service"`
}

var CONFIG Configuration
