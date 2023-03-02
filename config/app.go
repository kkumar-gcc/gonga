package config

import (
	"gonga/utils"
	"log"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name     string
	Env      string
	URL      string
	AssetURL string
	Timezone string
	Debug    bool
}

func LoadAppConfig() *AppConfig {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	debug, err := strconv.ParseBool(strings.ToLower(utils.Env("APP_DEBUG", "false")))
	if err != nil {
		debug = false
	}

	return &AppConfig{

		/*
		   |--------------------------------------------------------------------------
		   | Application Name
		   |--------------------------------------------------------------------------
		   |
		   | This value is the name of your application. This value is used when the
		   | framework needs to place the application's name in a notification or
		   | any other location as required by the application or its packages.
		   |
		*/

		Name: utils.Env("APP_NAME", "Gonga"),

		/*
		   |--------------------------------------------------------------------------
		   | Application Environment
		   |--------------------------------------------------------------------------
		   |
		   | This value determines the "environment" your application is currently
		   | running in. This may determine how you prefer to configure various
		   | services the application utilizes. Set this in your ".env" file.
		   |
		*/

		Env: utils.Env("APP_ENV", "production"),
		/*
		   |--------------------------------------------------------------------------
		   | Application Debug Mode
		   |--------------------------------------------------------------------------
		   |
		   | When your application is in debug mode, detailed error messages with
		   | stack traces will be shown on every error that occurs within your
		   | application. If disabled, a simple generic error page is shown.
		   |
		*/

		Debug: debug,

		/*
		   |--------------------------------------------------------------------------
		   | Application URL
		   |--------------------------------------------------------------------------
		   |
		   | This URL is used by the console to properly generate URLs when using
		   | the Artisan command line tool. You should set this to the root of
		   | your application so that it is used when running Artisan tasks.
		   |
		*/

		URL: utils.Env("APP_URL", "http://localhost"),

		AssetURL: utils.Env("ASSET_URL", "http://localhost"),

		/*
		   |--------------------------------------------------------------------------
		   | Application Timezone
		   |--------------------------------------------------------------------------
		   |
		   | Here you may specify the default timezone for your application, which
		   | will be used by the PHP date and date-time functions. We have gone
		   | ahead and set this to a sensible default for you out of the box.
		   |
		*/

		Timezone: utils.Env("APP_TIMEZONE", "UTC"),
	}

}
