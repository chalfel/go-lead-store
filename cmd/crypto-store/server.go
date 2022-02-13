package cmd

import (
	"errors"

	env "github.com/Netflix/go-env"
	"github.com/chalfel/go-lead-store/app"
	"github.com/chalfel/go-lead-store/config"
	"github.com/chalfel/go-lead-store/db"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start appname API",
	RunE:  startCmd,
}

func startCmd(cmd *cobra.Command, arg []string) error {
	var cfg config.ApplicationConfig
	var err error

	if err := godotenv.Load(); err != nil {
		return errors.New("error while loading .env file")
	}

	_, err = env.UnmarshalFromEnviron(&cfg)

	if err != nil {
		return err
	}

	validate := validator.New()

	if err := validate.Struct(cfg); err != nil {
		return err
	}

	db, err := db.NewDatabaseConnection(&cfg.Database)

	if err != nil {
		return err
	}

	server := app.NewServer(db, &cfg)
	app.NewAPI(server)

	if err != nil {
		return err
	}

	err = server.Init()

	return err
}
