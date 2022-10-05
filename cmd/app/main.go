package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"

	"ks-data-prepare/internal/pkg/builder"
	"ks-data-prepare/internal/pkg/config"
	"ks-data-prepare/internal/pkg/db"
	"ks-data-prepare/internal/pkg/migration"
)

func main() {
	flagUpElastic := "up-e"
	flagDownElastic := "down-e"
	flagResetElastic := "reset-e"
	flagUpClassic := "up-c"
	flagDownClassic := "down-c"
	flagResetClassic := "reset-c"

	flagCfg := "config"
	// flagAutoMigration := "auto_migration"
	// flagMigrationDir := "migration_dir"
	aliasesCfg := []string{"c"}
	defaultCfgFile := "internal/pkg/config/file/default.yaml"
	usage := "Configuration file"
	// usageMigrationDir := "Migration directory"
	defaultMigrationDir := "file://./migrations/postgres"
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "api",
				Aliases: []string{},
				Usage:   "Run api server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    flagCfg,
						Aliases: aliasesCfg,
						Value:   defaultCfgFile,
						Usage:   usage,
					},
					// &cli.StringFlag{
					// 	Name:  flagAutoMigration,
					// 	Value: "not run",
					// },
					// &cli.StringFlag{
					// 	Name:  flagMigrationDir,
					// 	Value: defaultMigrationDir,
					// 	Usage: usageMigrationDir,
					// },
				},
				Action: func(c *cli.Context) (err error) {
					err = config.Load(c.String(flagCfg))
					if err != nil {
						return err
					}
					err = db.InitDB()
					if err != nil {
						return err
					}
					apiServer := builder.NewApi()
					return apiServer.Run()

				},
			},
			{
				Name:    "migration",
				Aliases: []string{},
				Usage:   "Run migration server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  flagUpElastic,
						Value: -1,
					},
					&cli.IntFlag{
						Name:  flagDownElastic,
						Value: -1,
					},
					&cli.IntFlag{
						Name:  flagResetElastic,
						Value: -1,
					},
					&cli.IntFlag{
						Name:  flagUpClassic,
						Value: -1,
					},
					&cli.IntFlag{
						Name:  flagDownClassic,
						Value: -1,
					},
					&cli.IntFlag{
						Name:  flagResetClassic,
						Value: -1,
					},
					&cli.StringFlag{
						Name:    flagCfg,
						Aliases: aliasesCfg,
						Value:   defaultCfgFile,
						Usage:   usage,
					},
				},
				Action: func(c *cli.Context) (err error) {
					err = config.Load(c.String(flagCfg))
					if err != nil {
						return err
					}

					upE := c.Int(flagUpElastic)
					downE := c.Int(flagDownElastic)
					resetE := c.Int(flagResetElastic)
					upC := c.Int(flagUpClassic)
					downC := c.Int(flagDownClassic)
					resetC := c.Int(flagResetClassic)

					flagECount, flagCCount := 0, 0
					for _, flag := range []int{upE, downE, resetE} {
						if flag != -1 {
							flagECount += 1
						}
					}

					for _, flag := range []int{upC, downC, resetC} {
						if flag != -1 {
							flagCCount += 1
						}
					}

					// if up == -1 && down == -1 {
					// 	fmt.Println("No up or down migration declared")
					// 	return nil
					// }

					if flagECount > 1 || flagCCount > 1 {
						fmt.Println(flagECount, flagCCount)
						return errors.New("[ERROR] Too many config for elastic/classic at same time. Stop the migration")
					}

					m, err := migration.NewMigration(defaultMigrationDir)
					if err != nil {
						fmt.Println("Can not create migration " + err.Error())
					}

					err = db.InitDB()
					if err != nil {
						return err
					}

					errChan := make(chan error)
					go func(errChan chan error) {
						if downE != -1 {
							m.MigrateElasticDown(downE)
						} else {
							if upE != -1 {
								m.MigrateElasticUp(upE)
							} else {
								m.MigrateElasticDown(0)
								m.MigrateElasticUp(0)
							}
						}

						migration := builder.NewElasticMigration()
						err := migration.Run()
						if err != nil {
							errChan <- err
						}
					}(errChan)
					// go func() {
					// 	if downE != -1 {
					// 		m.MigrateClassicDown(downE)
					// 	} else {
					// 		if upE != -1 {
					// 			m.MigrateClassicUp(upE)
					// 		} else {
					// 			m.MigrateClassicDown(0)
					// 			m.MigrateClassicUp(0)
					// 		}
					// 	}
					// }()
					return <-errChan
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
