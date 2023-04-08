package device

import (
	"context"
	"github.com/kaiserbh/tachiyomi-sync-server/internal/domain"
	"github.com/kaiserbh/tachiyomi-sync-server/internal/logger"
	"github.com/rs/zerolog"
)

type Service interface {
	Store(ctx context.Context, device *domain.Device) error
	Delete(ctx context.Context, id int) error
	ListDevices(ctx context.Context, apikey string) ([]domain.Device, error)
}

func NewService(log logger.Logger, repo domain.DeviceRepo) Service {
	return &service{
		log:  log.With().Str("module", "device").Logger(),
		repo: repo,
	}
}

type service struct {
	log  zerolog.Logger
	repo domain.DeviceRepo
}

func (s service) Store(ctx context.Context, device *domain.Device) error {
	err := s.repo.Store(ctx, device)
	if err != nil {
		s.log.Error().Err(err).Msgf("could not store device: %+v", device)
		return err
	}

	return nil
}

func (s service) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.log.Error().Err(err).Msgf("could not delete device with id: %v", id)
		return err
	}

	return nil
}

func (s service) ListDevices(ctx context.Context, apikey string) ([]domain.Device, error) {
	devices, err := s.repo.ListDevices(ctx, apikey)
	if err != nil {
		s.log.Error().Err(err).Msg("could not list devices")
		return nil, err
	}

	return devices, nil
}
