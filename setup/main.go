package setup

import (
	"github.com/joho/godotenv"
	"os"
)

var setupDir = ""
var envFile = setupDir + "/.env"
var configFile = setupDir + "/my-config.json"
var logFile = setupDir + "/my.log"

func GetSetupDir() string {
	return setupDir
}

func GetEnvFile() string {
	return envFile
}

func GetConfigFile() string {
	return configFile
}

func GetLogFile() string {
	return logFile
}

func EnsureEnvironmentIsReady() {
	createSetupDirIfNotExists()
	ensureEnvFileIsReady()
}

func ensureEnvFileIsReady() {
	// get if env file exists. If not, create it with default values.
	_, err := os.Stat(envFile)
	if os.IsNotExist(err) {
		defaultEnv := getDefaultEnv()

		fErr := os.WriteFile(envFile, defaultEnv, 0644)
		if fErr != nil {
			panic(fErr)
		}
	}

	err = godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}
}

func createSetupDirIfNotExists() {
	_, err := os.Stat(setupDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(setupDir, 0755)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
}

func getDefaultEnv() []byte {
	// Edit this function to set the default environment variables
	return []byte("" +
		"SAMPLE_KEY=default value\n" +
		"")
}
