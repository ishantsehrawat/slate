package config

import (
	"github.com/ishant/slate-backend/internal/utils"
)

var (
	GoogleRedirectURL = utils.GetEnv("GOOGLE_REDIRECT_URL", "")
	GoogleClientID    = utils.GetEnv("GOOGLE_CLIENT_ID", "")
	GoogleClientSecret = utils.GetEnv("GOOGLE_CLIENT_SECRET", "")
	JWTSecret         = []byte(utils.GetEnv("JWT_SECRET", "default-secret"))
)
