package app

import (
	"crypto/sha1"
	"net/http"

	"github.com/alamin-mahamud/arya/auth/pkg/config"
	"github.com/alamin-mahamud/arya/auth/pkg/database"
	"github.com/alamin-mahamud/arya/auth/pkg/router"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ribice/gorsk/pkg/utl/rbac"
	"github.com/ribice/gorsk/pkg/utl/secure"
	"github.com/ribice/gorsk/pkg/utl/zlog"
)

func Start(cfg *config.Configuration) error {
	db, err := database.New(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	secureService := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.New()
	jwt := jwt.New(cfg.JWT.Secret,warfaze cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
	log := zlog.New()

	r := router.New()
	http.Handle("/", r)
	return http.ListenAndServe(cfg.Server.Port, nil)
}
