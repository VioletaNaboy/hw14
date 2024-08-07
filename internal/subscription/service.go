package subscription

import (
	"log"
	"time"
)

type storage interface {
	Create(s Subscription)
	GetAll() []Subscription
	GetSubscriptionByUserID(userID string) []Subscription
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	return &Service{s: s}
}

func (s *Service) SubscribeUser(userID, plan string, durationDays int) {
	subscription := NewSubscription(userID, plan, durationDays)
	s.s.Create(subscription)
	log.Printf("User %s subscribed to plan %s", userID, plan)
}

func (s *Service) GetSubscriptions() []Subscription {
	subscriptions := s.s.GetAll()
	for i := range subscriptions {
		if time.Now().After(subscriptions[i].EndDate) {
			subscriptions[i].Status = "Expired"
		}
	}
	return subscriptions
}

func (s *Service) GetSubscriptionByUserID(userID string) []Subscription {
	subscriptions := s.s.GetSubscriptionByUserID(userID)
	for i := range subscriptions {
		if time.Now().After(subscriptions[i].EndDate) {
			subscriptions[i].Status = "Expired"
		}
	}
	return subscriptions
}
