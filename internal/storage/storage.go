package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"io"
	"log"
)

type Storage interface {
	SetProfilePhoto(ctx context.Context, id string, r io.Reader) error
	GetProfilePhoto(ctx context.Context, id string, w io.Writer) error
}

type GoogleStorage struct {
	client            *storage.Client
	bucket            string
	profilePhotosPath string
}

func NewGoogleStorage(keyPath, bucket, profilePhotosPath string) *GoogleStorage {
	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(keyPath))
	if err != nil {
		log.Fatalf("error connection to google cloud storage")
	}

	return &GoogleStorage{
		client:            storageClient,
		bucket:            bucket,
		profilePhotosPath: profilePhotosPath,
	}
}

func (s *GoogleStorage) SetProfilePhoto(ctx context.Context, id string, r io.Reader) error {
	w := s.client.Bucket(s.bucket).Object(s.profilePhotosPath + id).NewWriter(ctx)
	_, err := io.Copy(w, r)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *GoogleStorage) GetProfilePhoto(ctx context.Context, id string, w io.Writer) error {
	r, err := s.client.Bucket(s.bucket).Object(s.profilePhotosPath + id).NewReader(ctx)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	err = r.Close()
	if err != nil {
		return err
	}

	return nil
}
