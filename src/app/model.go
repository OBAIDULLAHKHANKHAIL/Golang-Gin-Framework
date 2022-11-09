package main

import "time"

var employees = map[string]Employee{
	"962134": {
		ID:        962134,
		FirstName: "Obaid Ullah",
		LastName:  "Khan",
		Position:  "CEO",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
	},
	"176158": {
		ID:        176158,
		FirstName: "Zia Ullah",
		LastName:  "Khan",
		Position:  "COO",
		StartDate: time.Now().Add(-4 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  20,
	},
	"160898": {
		ID:        160898,
		FirstName: "Saad Ullah",
		LastName:  "Khan",
		Position:  "CTO",
		StartDate: time.Now().Add(-6 * time.Hour * 24 * 365),
		TotalPTO:  20,
	},
	"297365": {
		ID:        297365,
		FirstName: "Mr",
		LastName:  "Abid",
		Position:  "Worker Bee",
		StartDate: time.Now().Add(-12 * time.Hour * 24 * 365),
		TotalPTO:  30,
	},
}

var TimesOff = map[string][]TimeOff{
	"962134": {
		{
			Type:      "Holiday",
			Amount:    8.,
			StartDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Taken",
		}, {
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 8, 16, 0, 0, 0, 0, time.UTC),
			Status:    "Scheduled",
		}, {
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 12, 8, 0, 0, 0, 0, time.UTC),
			Status:    "Requested",
		},
	},
}

type Employee struct {
	ID        uint
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
	StartDate time.Time
	Position  string  `form:"position"`
	TotalPTO  float32 `form:"pto"`
	Status    string
	TimesOff  []TimeOff
}

type TimeOff struct {
	Type      string    `json:"reason" binding:"required"`
	Amount    float32   `json:"hours" binding:"required,gt=0"`
	StartDate time.Time `json:"startDate" binding:"required"`
	Status    string    `json:"status" binding:"required"`
}
