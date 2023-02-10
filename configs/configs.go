package configs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
)

type Config struct {
	Port           string `json:"port"`
	Domain         string `json:"domain"`
	UBP            string `json:"upload_base_path"`
	AppName        string `json:"app_name"`
	ErrorLog       bool   `json:"error_log"`
	ErrorAccessLog bool   `json:"error_access_log"`
}

var CFG = Config{}

func LoadConfig() {
	file, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Errorf("error wile reading config file, error: %v", err.Error())
	}

	if err == nil {
		if err := json.Unmarshal(file, &CFG); err != nil {
			panic(err)
		}
		fmt.Printf("config loaded.\n")
	}
}
