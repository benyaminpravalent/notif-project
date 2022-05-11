package service

import (
	"log"

	"github.com/project/notif-project/domain/repository"
)

// NotifService manage logical syntax for notif.
type NotifService interface {
}

type notifServiceImpl struct {
	notifRepo repository.NotifRepository
}

// NewNotifService returns new instance of notifServiceImpl.
func NewNotifService() *notifServiceImpl {
	return &notifServiceImpl{}
}

// SetNotifRepo injects notif's repo for notifServiceImpl.
func (s *notifServiceImpl) SetNotifRepo(repo repository.NotifRepository) *notifServiceImpl {
	s.notifRepo = repo
	return s
}

func (s *notifServiceImpl) Validate() *notifServiceImpl {
	if s.notifRepo == nil {
		log.Panic("Notif service need notif repository")
	}
	return s
}
