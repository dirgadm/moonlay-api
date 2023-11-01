package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"project-version3/moonlay-api/pkg/ehttp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_listsHandler "project-version3/moonlay-api/service/lists/delivery"
	_listRepo "project-version3/moonlay-api/service/lists/repository"
	_listsUsecase "project-version3/moonlay-api/service/lists/usecase"
	_subListsHandler "project-version3/moonlay-api/service/sublists/delivery"
	_subListRepo "project-version3/moonlay-api/service/sublists/repository"
	_subListsUsecase "project-version3/moonlay-api/service/sublists/usecase"
	_uploadHandler "project-version3/moonlay-api/service/upload/delivery"
	_uploadUsecase "project-version3/moonlay-api/service/upload/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbUser, dbName, dbPass, dbHost, dbPort)

	// val := url.Values{}
	// val.Add("parseTime", "1")
	// val.Add("loc", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := gormDB.DB()
	defer func() {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// setup machine and middleware
	e := echo.New()
	// setup cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header"},
	}))

	// setup echo for request id
	e.Use(middleware.RequestID())

	// setup echo for secure
	e.Use(middleware.Secure())

	// setup echo for gzip compres
	e.Use(middleware.Gzip())

	// setup custom context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ehttp.Context{
				Context:        c,
				ResponseFormat: ehttp.NewResponse(),
				ResponseData:   nil,
			}
			return next(cc)
		}
	})

	// setup timeout
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// setup repo
	listsRepo := _listRepo.NewlListsRepository(gormDB)
	sublistsRepo := _subListRepo.NewlSubListsRepository(gormDB)

	// setup usecase
	listsUsecase := _listsUsecase.NewListsUsecase(listsRepo, timeoutContext)
	subListsUsecase := _subListsUsecase.NewSubListsUsecase(sublistsRepo, listsRepo, timeoutContext)
	uploadUsecase := _uploadUsecase.NewUploadUsecase(timeoutContext)

	// setup handler
	_listsHandler.NewListsHandler(e, listsUsecase)
	_subListsHandler.NewSubListsHandler(e, subListsUsecase)
	_uploadHandler.NewUploadsHandler(e, uploadUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
