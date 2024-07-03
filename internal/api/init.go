package api

import (
	"blaze/internal/api/handlers"
	"blaze/pkg/httpcore"
	"blaze/pkg/storage"
	"blaze/pkg/util"

	"github.com/rs/zerolog/log"

	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitService() http.Handler {
	util.InitLogger()

	router := chi.NewRouter()

	router.Use(cors.New(httpcore.DefaultCorsOptions).Handler)
	router.Use(middleware.Timeout(20 * time.Second))
	router.Use(middleware.Recoverer)
	router.Use(httpcore.LoggerMiddleware)

	env := InitEnv()

	storage, err := storage.NewTodoStorage(env.DatabaseUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not initialize storage")
	}

	controller := handlers.NewApiController(storage)
	applyRoutes(router, controller)

	return router
}
