package service

import (
	"time"

	"github.com/ValeryChapman/blog-app/pkg/repository"
)

const (
	salt       = "gw@3$vp18tMBT4CvpPVV9v{*qJ@XP46#"
	signingKey = "Ka2gF~Nsx@yo4go$6Cq#Zkig7K9Y}f~k"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// parse token
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	return 1, nil
}
