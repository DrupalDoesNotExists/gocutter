package main

import (
	"github.com/EduUnID/jwtmarshal"
	"github.com/fvbock/endless"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"os"
	"time"
)

var (
	HttpAddr    = os.Getenv("SERVICE_HTTP_ADDR")
	DatabaseDSN = os.Getenv("SERVICE_DATABASE_DSN")
	JwtKeyFile  = os.Getenv("SERVICE_JWT_KEY")
)

func newRouter(jwtCtx *jwtmarshal.JWTContext, db *sqlx.DB) *gin.Engine {
	// create a new gin router
	r := gin.New()

	// middlewares
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))
	// TODO: add custom middlewares

	// controllers
	// TODO: add custom routes (use closures for DI)

	return r
}

func main() {
	// configure zap logging
	l := zap.Must(zap.NewProduction())
	zap.ReplaceGlobals(l)

	// create a jwt context
	jwtCtx := &jwtmarshal.JWTContext{
		SigningMethod: jwt.SigningMethodRS512,
		ExpireAfter:   time.Minute * 30,
	}

	// read jwt keys
	l.Debug("reading jwt key pair", zap.String("path", JwtKeyFile))
	if err := jwtCtx.ReadKeyFiles(JwtKeyFile, JwtKeyFile+".pub"); err != nil {
		l.Fatal("error reading jwt keys", zap.Error(err))
	}

	// connect to a database
	l.Debug("connecting to a database", zap.String("dsn", DatabaseDSN))
	db, err := sqlx.Open("pgx", DatabaseDSN)
	if err != nil {
		l.Fatal("error connecting to a database", zap.Error(err))
	}

	// http server
	r := newRouter(jwtCtx, db)
	l.Debug("listening http", zap.String("addr", HttpAddr))
	if err := endless.ListenAndServe(HttpAddr, r); err != nil {
		l.Fatal("http server error", zap.Error(err))
	}

	l.Debug("graceful shutdown")
}
