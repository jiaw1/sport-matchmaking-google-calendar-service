package model

import "time"

type CalendarDetails struct {
	UserIDs      []string     `json:"userIds"`
	MatchDetails MatchDetails `json:"matchDetails"`
}

type MatchDetails struct {
	ID                string    `json:"id"`
	Sport             string    `json:"sport"`
	MinParticipants   *int32    `json:"minParticipants"`
	MaxParticipants   *int32    `json:"maxParticipants"`
	StartsAt          time.Time `json:"startsAt"`
	EndsAt            time.Time `json:"endsAt"`
	Location          string    `json:"location"`
	Description       string    `json:"description"`
	ParticipationFee  int64     `json:"participationFee"`
	RequiredEquipment []string  `json:"requiredEquipment"`
	Level             string    `json:"level"`
	ChatLink          string    `json:"chatLink"`
	HostUserID        string    `json:"hostUserId"`
}
