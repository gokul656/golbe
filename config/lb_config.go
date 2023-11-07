package config

import (
	"encoding/json"
	"log"
	"os"
)

var LBConfig *LoadBalancer = LoadLBConfig()

type LoadBalancer struct {
	Config map[string]LBConfiguration `json:"config"`
}

type LBConfiguration struct {
	Peers     []string `json:"peers"`
	Algorithm string   `json:"routing_algorithm"`
}

func LoadLBConfig() *LoadBalancer {
	fr, err := os.ReadFile("routes.json")
	if err != nil {
		log.Println("[ERROR]", err)
		return nil
	}

	lb := &LoadBalancer{}
	json.Unmarshal(fr, lb)

	return lb
}
