package main

import (
	"os"

	"github.com/moov-io/base/log"
	"github.com/sethlivingston/reponotifications"
	"github.com/sethlivingston/reponotifications/pkg/server"
)

func main() {
	logger := log.NewDefaultLogger().
		Set("app", log.String("sethlivingston/reponotifications")).
		Set("version", log.String(reponotifications.Version))

	env := &server.Environment{
		Logger: logger,
	}

	_, err := server.NewEnvironment(env)
	if err != nil {
		logger.Fatal().LogErrorf("creating environment: %v", err)
		os.Exit(1)
	}
}
