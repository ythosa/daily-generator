package models

type DailyMessage struct {
	Yesterday []*Issue
	Today     []*Issue
	Problems  string
}

func NewDailyMessage(yesterday []*Issue, today []*Issue, problems string) *DailyMessage {
	return &DailyMessage{Yesterday: yesterday, Today: today, Problems: problems}
}
