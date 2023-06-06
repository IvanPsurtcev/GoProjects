package program

import (
	fhttpRouter "github.com/fasthttp/router"
	"github.com/jmoiron/sqlx"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"metamask-auth/config"
	authHttp "metamask-auth/internal/auth/delivery/http"
	authRepository "metamask-auth/internal/auth/repository"
	authUseCase "metamask-auth/internal/auth/usecase"
	contractUseCase "metamask-auth/internal/contract/usecase"
	gameHttp "metamask-auth/internal/game/delivery/http"
	gameRepository "metamask-auth/internal/game/repository"
	gameUseCase "metamask-auth/internal/game/usecase"
	"metamask-auth/internal/middleware"
	requestHttp "metamask-auth/internal/request/delivery/http"
	requestRepository "metamask-auth/internal/request/repository"
	requestUseCase "metamask-auth/internal/request/usecase"
	appCloser "metamask-auth/pkg/app_closer"
	"metamask-auth/pkg/jwt"
	"metamask-auth/pkg/reservoir"
	"time"
)

type coreProgram struct {
	cfg       *config.Config
	logger    *zap.Logger
	closer    appCloser.Closer
	appRouter *fhttpRouter.Router
}

func NewCoreProgram(cfg *config.Config, logger *zap.Logger, closer appCloser.Closer, db *sqlx.DB) Program {
	authRepo := authRepository.NewRepository(logger, db)
	requestRepo := requestRepository.NewRepository(logger, db)
	gameRepo := gameRepository.NewRepository(logger, db)

	jwtHmacProvider := jwt.NewHmacProvider(
		cfg.Jwt.HmacSecret,
		cfg.Jwt.Issuer,
		time.Duration(cfg.Jwt.Duration)*time.Second)

	reservoirCli := reservoir.NewClient(logger, cfg.ReservoirApiUrl)

	authUC := authUseCase.NewUseCase(logger, jwtHmacProvider, authRepo)
	requestUC := requestUseCase.NewUseCase(logger, requestRepo, authRepo)
	contractUC := contractUseCase.NewUseCase(logger, cfg)
	gameUC := gameUseCase.NewUseCase(logger, gameRepo, reservoirCli, contractUC, authRepo)

	authHandlers := authHttp.NewHandlers(logger, authUC)
	requestHandlers := requestHttp.NewHandlers(logger, requestUC)
	gameHandlers := gameHttp.NewHandlers(logger, gameUC)

	mw := middleware.NewMw(logger, jwtHmacProvider)

	appRouter := fhttpRouter.New()

	authGroup := appRouter.Group("/auth")
	requestGroup := appRouter.Group("/invite")
	gameGroup := appRouter.Group("/game")

	authHttp.InitRouter(authGroup, authHandlers, mw)
	requestHttp.InitRouter(requestGroup, requestHandlers, mw)
	gameHttp.InitRouter(gameGroup, gameHandlers, mw)

	return &coreProgram{
		cfg:       cfg,
		logger:    logger,
		closer:    closer,
		appRouter: appRouter,
	}
}

func (core *coreProgram) Run() {
	core.logger.Info("starting coreProgram")
	server := fasthttp.Server{
		Handler: middleware.CORS(core.appRouter.Handler),
	}

	core.closer.AddCloser(func() {
		if err := server.Shutdown(); err != nil {
			core.logger.Panic("cannot stop coreProgram", zap.Error(err))
		}
	}, "coreProgram")

	listenAddr := core.cfg.System.Host + ":" + core.cfg.System.Port
	if err := server.ListenAndServe(listenAddr); err != nil {
		core.logger.Panic("Error in running server", zap.Error(err))
	}
}
