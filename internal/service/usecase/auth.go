package usecase

import (
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)


type authUseCase struct{
	repo repo.AuthRepository
	log log.Logger
}
func NewAuthUseCase(
	repo repo.AuthRepository,
	log log.Logger,
) AuthUseCase {
	return &authUseCase{
		repo: repo,
		log: log,
	}
}

