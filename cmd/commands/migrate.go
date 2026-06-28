package commands

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hamidgh01/Go-Blog-API/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:               "migrate",
	Short:             "apply database migrations",
	PersistentPreRun:  func(cmd *cobra.Command, args []string) { initMigrateInstance() },
	PersistentPostRun: func(cmd *cobra.Command, args []string) { giveReportsAndCloseSources() },
}

var migrateUpCmd = &cobra.Command{
	Use:     "up",
	Short:   "apply all or N up-migrations",
	Example: "  go run ./cmd migrate up  # apply all up migrations \n  go run ./cmd migrate up --steps 1  # apply 1 up migration",
	Run:     migrateUp,
}

var migrateDownCmd = &cobra.Command{
	Use:     "down",
	Short:   "apply N down-migrations",
	Example: "  go run ./cmd migrate down --steps 1  # apply 1 down migration \n  go run ./cmd migrate down --steps 2  # apply 2 down migrations",
	Run:     migrateDown,
}

var migrateForceCmd = &cobra.Command{
	Use:   "force [V]",
	Short: "set database migration version to V (V is a positive integer)",
	Example: "  go run ./cmd force 7  # set migration version to 7",
	Args:  cobra.ExactArgs(1),
	Run:   forceMigrationVersion,
}

var migrateInstance *migrate.Migrate

func initMigrateInstance() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to init configurations. reason:", err)
	}

	migrationsPath := "file://migrations"
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	)

	migrateInstance, err = migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("failed to initialize migrate: %v", err)
	}

	_ = migrateInstance.Log
}

// this function gives reports current migration version and
// dirtiness, then closes migration source and database.
func giveReportsAndCloseSources() {
	version, dirty, err := migrateInstance.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Printf("failed to get migration version and dirtiness: %v \n", err)
	} else {
		log.Printf("current migration version: %d (dirty: %v)\n", version, dirty)
	}

	se, de := migrateInstance.Close()
	if se != nil || de != nil {
		log.Printf(
			"there's error while closing migrate source and/or database.\nsource close error: %s \ndatabase close error: %s \n",
			se.Error(), de.Error(),
		)
	} else {
		log.Printf("migrate source and database closed")
	}
}

func migrateUp(cmd *cobra.Command, args []string) {
	steps, _ := cmd.Flags().GetInt("steps")

	if cmd.Flags().Changed("steps") {
		if steps <= 0 {
			log.Println("step must be a positive integer (step > 0). skip applying migrations (nothing changed)")
			return
		}
	} else { // step is not set
		log.Println("direction: UP")
		if err := migrateInstance.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migration up failed: %v", err)
		}
		log.Println("all UP migrations applied successfully")
		return
	}

	// step is certainly a positive integer
	log.Println("direction: UP")
	if err := migrateInstance.Steps(steps); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration up failed: %v", err)
	}
	log.Printf("Up migrations applied successfully (steps: %d)\n", steps)
}

func migrateDown(cmd *cobra.Command, args []string) {
	steps, _ := cmd.Flags().GetInt("steps")
	if steps <= 0 {
		log.Println("step must be a positive integer (step > 0). skip applying migrations (nothing changed)")
		return
	}

	log.Println("direction: DOWN")
	if err := migrateInstance.Steps(-steps); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration down failed: %v", err)
	}
	log.Printf("migrations rolled '%d steps' back successfully\n", steps)
}

func forceMigrationVersion(cmd *cobra.Command, args []string) {
	version, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("entered arg is not a valid integer. origin: %v \n", err)
	}
	if version < 0 {
		log.Fatal("invalid version. it should be a positive integer (V >= 0)")
	}

	if err := migrateInstance.Force(version); err != nil {
		log.Fatalf("force migration failed: %v", err)
	}
	log.Printf("migration version forced to %d \n", version)
}
