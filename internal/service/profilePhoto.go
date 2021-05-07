package service

import (
	"context"
	"github.com/EgorMizerov/kindergarten/internal/storage"
	"io"
)

type ProfilePhotoService struct {
	storage storage.Storage
}

func NewProfilePhotoService(storage storage.Storage) *ProfilePhotoService {
	return &ProfilePhotoService{storage: storage}
}

func (s *ProfilePhotoService) SetProfilePhoto(ctx context.Context, id string, r io.Reader) error {
	return s.storage.SetProfilePhoto(ctx, id, r)
}

func (s *ProfilePhotoService) GetProfilePhoto(ctx context.Context, id string, w io.Writer) error {
	return s.storage.GetProfilePhoto(ctx, id, w)
}
