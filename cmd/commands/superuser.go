package commands

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/internal/application/services"
	"github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/database"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/database/postgres_repository"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/redis"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/security/hashing"
	"github.com/hamidgh01/Go-Blog-API/pkg/logging"

	"github.com/spf13/cobra"
)

var createSuperuserCmd = &cobra.Command{
	Use:   "create-superuser",
	Short: "start an interactive prompt to create a superuser",
	Run: func(cmd *cobra.Command, args []string) {
		createSuperuser()
	},
}

func createSuperuser() {
	// init configurations
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to init configurations. reason:", err)
	}

	logging.InitLogger(cfg.Logger)

	// establish database connection
	db, err := database.InitDB(&cfg.Postgres)
	if err != nil {
		log.Fatal("failed to establish database connection. reason: ", err)
	}
	defer db.Close()

	// establish redis connection
	redisClient, err := redis.InitRedis(&cfg.Redis)
	if err != nil {
		log.Println("failed to establish redis connection. reason:", err)
		return
	}
	defer redisClient.Close()

	// prepare dependencies
	postgresUserRepo := postgres_repository.NewUserRepository(db)
	pswHasher := hashing.NewPasswordHasher()
	userInfoCache := redis.NewUserInfoCache(redisClient)
	userService := services.NewUserService(postgresUserRepo, pswHasher, userInfoCache)

	// start interactive steps
	time.Sleep(1 * time.Second)
	fmt.Println("\nCreating Superuser")
	time.Sleep(1 * time.Second)
	fmt.Println("to complete the process do all of the steps and don't quit until the end")
	time.Sleep(1 * time.Second)
	fmt.Println("you need to enter a unique `username`, an unregistered `email address`, and a `password`")
	time.Sleep(1 * time.Second)

	username, err := readAndValidateUsername(userService)
	if err != nil {
		log.Println(err) // log.Error
		return
	}
	time.Sleep(700 * time.Millisecond)
	fmt.Printf("'%s' accepted! ✅\n", username)

	email, err := readAndValidateEmail(userService)
	if err != nil {
		log.Println(err) // log.Error
		return
	}
	time.Sleep(700 * time.Millisecond)
	fmt.Printf("'%s' accepted! ✅\n", email)

	password, err := readAndValidatePassword()
	if err != nil {
		log.Println(err) // log.Error
		return
	}
	time.Sleep(700 * time.Millisecond)
	fmt.Println("password accepted! ✅")

	time.Sleep(1 * time.Second)
	fmt.Print("\ncreating superuser")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	hashedPassword, err := pswHasher.Hash(password)
	if err != nil {
		log.Printf("failed to hash password. reason: %s", err.Error()) // log.Error
		return
	}

	createSuperuserQuery := `
		INSERT INTO users (username, email, isSuperuser, password)
		VALUES ($1, $2, true, $3)
		RETURNING id, username, email, enabled, isSuperuser, createdAt
	`

	dbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var su = &entity.User{}
	dbErr := db.QueryRowContext(
		dbCtx, createSuperuserQuery, username, email, hashedPassword,
	).Scan(&su.ID, &su.Username, &su.Email, &su.Enabled, &su.IsSuperuser, &su.CreatedAt)

	if dbErr != nil {
		log.Printf("failed to create superuser. reason: %s", dbErrors.GetDBError(dbErr))
		return
	}

	// save user info in cache
	redisCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	redisErr := userInfoCache.SetAllInfo(redisCtx, su.ID, su.Username, su.IsSuperuser, su.Enabled)
	if redisErr != nil {
		fmt.Println(redisErr) // log.Error
	}

	fmt.Println("superuser created successfully! ✅✅✅") // log.Info
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
	fmt.Println("id:          ", su.ID)
	fmt.Println("username:    ", su.Username)
	fmt.Println("email:       ", su.Email)
	fmt.Println("is superuser:", su.IsSuperuser)
	fmt.Println("created at:  ", su.CreatedAt)
	fmt.Println()
}

func readAndValidateUsername(userService *services.UserService) (username string, err error) {
	fmt.Print("\nusername: ")
	if _, err = fmt.Scanln(&username); err != nil {
		return "", err
	}

	exists, e := userService.CheckUsernameExists(context.Background(), username)
	if e != nil {
		return "", fmt.Errorf("failed to CheckUsernameExists. reason: %w", err)
	} else if exists {
		fmt.Printf("the username '%s' is already taken. please choose another username.\n", username)
		return readAndValidateUsername(userService)
	} else if !domain.CheckUsernamePattern(username) {
		fmt.Printf("'%s' is not a valid username. %s\n", username, domain.UsernamePatternDescription)
		return readAndValidateUsername(userService)
	}

	return
}

func readAndValidateEmail(userService *services.UserService) (email string, err error) {
	fmt.Print("\nemail address: ")
	if _, err = fmt.Scanln(&email); err != nil {
		return "", err
	}

	exists, e := userService.CheckEmailExists(context.Background(), email)
	if e != nil {
		return "", fmt.Errorf("failed to CheckUsernameExists. reason: %w", err)
	} else if exists {
		fmt.Printf("the email '%s' is already registered. please enter another email.\n", email)
		return readAndValidateEmail(userService)
	} else if !domain.CheckEmailPattern(email) {
		fmt.Printf("'%s' is not a valid email address.\n", email)
		return readAndValidateEmail(userService)
	}

	return
}

func readAndValidatePassword() (password string, err error) {
	fmt.Print("\npassword: ")
	if _, err = fmt.Scanln(&password); err != nil {
		return "", err
	}

	if !domain.CheckPasswordPattern(password) {
		fmt.Printf("'%s' is not accepted as password. %s\n", password, domain.PasswordPatternDescription)
		return readAndValidatePassword()
	}

	var confirmPsw string
	fmt.Print("confirm password: ")
	if _, err = fmt.Scanln(&confirmPsw); err != nil {
		return "", err
	}

	if confirmPsw != password {
		fmt.Println("passwords don't match. try again...")
		return readAndValidatePassword()
	}

	return
}

// ----------------------------------------------------------------------------------

var deleteSuperuserCmd = &cobra.Command{
	Use:   "delete-superuser",
	Short: "start an interactive prompt to delete a superuser",
	Run: func(cmd *cobra.Command, args []string) {
		deleteSuperuser()
	},
}

func deleteSuperuser() {
	// init configurations
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to init configurations. reason:", err)
	}

	logging.InitLogger(cfg.Logger)

	// establish database connection
	db, err := database.InitDB(&cfg.Postgres)
	if err != nil {
		log.Fatal("failed to establish database connection. reason: ", err)
	}
	defer db.Close()

	// establish redis connection
	redisClient, err := redis.InitRedis(&cfg.Redis)
	if err != nil {
		log.Println("failed to establish redis connection. reason:", err)
		return
	}
	defer redisClient.Close()

	// prepare dependencies
	userInfoCache := redis.NewUserInfoCache(redisClient)

	// start interactive steps
	time.Sleep(1 * time.Second)
	fmt.Println("\nDeleting a Superuser")
	time.Sleep(1 * time.Second)
	fmt.Println("you need to enter the unique `username` of a superuser to delete that")
	time.Sleep(1 * time.Second)

	var username string
	fmt.Print("\nusername: ")
	if _, err = fmt.Scanln(&username); err != nil {
		log.Println(err) // log.Error
		return
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("\ndeleting superuser with username='%s'", username)
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	deleteSuperuserQuery := "DELETE FROM users WHERE username = $1 AND isSuperuser = true RETURNING id"

	dbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id uint64
	dbErr := db.QueryRowContext(dbCtx, deleteSuperuserQuery, username).Scan(&id)
	if dbErr != nil {
		if dbErr == sql.ErrNoRows {
			log.Printf("there's no any superuser with username='%s' (not found)", username) // log.Error
			return
		}

		log.Printf(
			"failed to delete superuser with username='%s'. reason: %s",
			username,
			dbErrors.GetDBError(dbErr),
		) // log.Error
		return
	}

	// update cache
	redisCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	redisErr := userInfoCache.DeleteUserInfo(redisCtx, id)
	if redisErr != nil {
		fmt.Println(redisErr) // log.Error
	}

	fmt.Printf("superuser (username='%s') deleted successfully! ✅✅✅\n", username) // log.Info
}
