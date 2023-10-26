package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"ozon-test/internal/entities"
	api "ozon-test/pkg/api/proto"
	"ozon-test/pkg/helpers"
	"sync"
)

type MemoryStorage struct {
	originalAsKey map[string]string
	shortAsKey    map[string]string

	logger *log.Entry

	sync.RWMutex
}

func NewMemoryStorage(l *log.Logger) *MemoryStorage {
	logger := l.WithField("storage", "memory")

	return &MemoryStorage{
		originalAsKey: make(map[string]string),
		shortAsKey:    make(map[string]string),
		logger:        logger,
	}
}

func (m *MemoryStorage) GetStorageType() string {
	return "memory"
}

func (m *MemoryStorage) AddURL(ctx context.Context, request *api.AddURLRequest) (*api.AddURLResponse, error) {
	var (
		e        error
		response *api.AddURLResponse
	)

	if savedLink, ok := m.originalAsKey[request.GetUrl()]; ok {
		response = &api.AddURLResponse{Url: &api.ShortenedURL{OriginalURL: request.GetUrl(), ShortenedURL: savedLink}}
	} else if shortLink, err := m.getSecureToken(10); err == nil {
		m.RWMutex.Lock()
		m.originalAsKey[request.GetUrl()] = shortLink
		m.shortAsKey[shortLink] = request.GetUrl()
		response = &api.AddURLResponse{Url: &api.ShortenedURL{OriginalURL: request.GetUrl(), ShortenedURL: shortLink}}
		m.RWMutex.Unlock()
	} else {
		e = entities.ServerError
	}

	if response != nil {
		m.logger.WithFields(log.Fields{
			"request":  request,
			"response": response,
			"code":     0,
		}).Info("addUrl success")
	}
	return response, e
}

func (m *MemoryStorage) GetURL(ctx context.Context, request *api.GetURLRequest) (*api.GetURLResponse, error) {
	var (
		response *api.GetURLResponse
		e        error
	)

	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()

	if originalLink, ok := m.shortAsKey[request.GetUrl()]; ok {
		response = &api.GetURLResponse{Url: &api.ShortenedURL{OriginalURL: originalLink, ShortenedURL: request.Url}}
	} else {
		e = entities.NotFound
	}

	if response != nil {
		m.logger.WithFields(log.Fields{
			"request":  request,
			"response": response,
			"code":     0,
		}).Info("getUrl success")
	}
	return response, e
}

func (m *MemoryStorage) getSecureToken(length int) (string, error) {
	for i := 0; i < 5; i++ {
		token, err := helpers.GenToken(length)
		if err != nil {
			return "", err
		}

		m.RWMutex.RLock()
		_, exists := m.shortAsKey[token]
		m.RWMutex.RUnlock()

		if !exists {
			return token, nil
		}
	}

	return "", entities.ServerError
}
