package subscription

import "time"

type Subscription struct {
	ID        string
	UserID    string
	Plan      string
	StartDate time.Time
	EndDate   time.Time
	Status    string
}

func NewSubscription(userID, plan string, durationDays int) Subscription {
	startDate := time.Now()
	endDate := startDate.Add(time.Duration(durationDays) * 24 * time.Hour)
	return Subscription{
		ID:        time.Now().Format("20060102150405"),
		UserID:    userID,
		Plan:      plan,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    "Active",
	}
}
