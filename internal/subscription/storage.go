package subscription

import (
	"sync"
)

type InMemStorage struct {
	subscriptionsM sync.Mutex
	subscriptions  []Subscription
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{}
}

func (s *InMemStorage) Create(subscription Subscription) {
	s.subscriptionsM.Lock()
	defer s.subscriptionsM.Unlock()
	s.subscriptions = append(s.subscriptions, subscription)
}

func (s *InMemStorage) GetAll() []Subscription {
	s.subscriptionsM.Lock()
	defer s.subscriptionsM.Unlock()
	return s.subscriptions
}

func (s *InMemStorage) GetSubscriptionByUserID(userID string) []Subscription {
	s.subscriptionsM.Lock()
	defer s.subscriptionsM.Unlock()

	var userSubscriptions []Subscription
	for _, sub := range s.subscriptions {
		if sub.UserID == userID {
			userSubscriptions = append(userSubscriptions, sub)
		}
	}
	return userSubscriptions
}
