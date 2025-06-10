package api

import (
	"log"
	"net/http"

	"github.com/Chenorlive/brainy/middleware"
	"github.com/Chenorlive/brainy/pkg/auth"
	"github.com/Chenorlive/brainy/pkg/core/role"
	"github.com/Chenorlive/brainy/pkg/core/userrole"
	"github.com/Chenorlive/brainy/pkg/ping"

	"gorm.io/gorm"
)

type AIPServer struct {
	addr string
	db   *gorm.DB
}

func NewServer(addr string, db *gorm.DB) *AIPServer {
	return &AIPServer{
		addr: addr,
		db:   db,
	}
}

func (s *AIPServer) Run() error {

	// Normal router
	router := http.NewServeMux()

	// ping
	pingHander := ping.NewHander()
	pingHander.RegisterRoutes(router)

	// auth
	authHander := auth.NewHander()
	authHander.RegisterRoutes(router)

	//Authentication router
	authRouter := http.NewServeMux()

	// core.role
	roleStore := role.NewStore(s.db)
	roleHandler := role.NewHander(roleStore)
	roleHandler.RegisterRoutes(authRouter)

	//core.userRole
	userRoleStore := userrole.NewStore(s.db)
	userRoleHandler := userrole.NewHandler(userRoleStore)
	userRoleHandler.RegisterRoutes(authRouter)

	//userRoleStore.

	// Authentication middleware
	authRouterWithMiddleware := middleware.AuthMiddleware(authRouter)
	router.Handle("/", authRouterWithMiddleware)

	// nomarl middleware
	stack := middleware.ChainMiddleware(
		middleware.LoggingMiddleware,
	)
	chain := stack(router)

	// Apply the middleware to the /api/v1 router
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", chain))

	server := &http.Server{
		Addr:    s.addr,
		Handler: v1,
	}

	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}
