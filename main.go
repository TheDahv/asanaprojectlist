package main

import (
	"log"
	"os"
	w "github.com/thedahv/asanaprojectlist/web"
	"net/http"
  "github.com/spf13/viper"
	r "regexp"
)

func getAppMode() string {
	env := viper.GetString("appmode")

	if match, _ := r.MatchString("^(production|test|development)$", env); match {
		// Matches a known environment
		return env
	}

	// Default
	return "development"
}

func loadEnvIntoConfig() {
	envSplit := r.MustCompile("(.+)=(.+)")
	for _, env := range os.Environ() {
		split := envSplit.FindStringSubmatch(env)

		if len(split) == 0 {
			continue
		}

		viper.Set(split[1], split[2])
	}
}

func main() {
	// Set up Config for app
	loadEnvIntoConfig()

	var configName string
	mode := getAppMode()
	if mode == "development" {
		configName = "defaults"
	} else {
		configName = mode
	}

	// Load in app configuration
	if _, err := os.Stat("./config/production.json"); mode == "production" && err == nil {
		log.Println("Loading production config")
		viper.SetConfigName(configName)
	} else {
		log.Println("Loading development config")
		viper.SetConfigName(configName)
	}
	viper.AddConfigPath("./config")
	viper.ReadInConfig()

	// Set up Web App routes
	r := w.SetupRoutes()
	http.Handle("/", r)

	port := viper.GetString("port")
	if port == "" {
		port = "5000"
	}

	log.Println("Listening on port " + port)
	log.Println(http.ListenAndServe(":" + port, nil))
}
