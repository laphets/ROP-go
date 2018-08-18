package ssr

import "rop/model"

type ScheduleResponse struct {
	*model.FreshmanModel
	IntentGroups []*IntentGroup `json:"intent_groups"`
}

type IntentGroup struct {
	*model.IntentModel
	Interviews []*model.FullInterview `json:"interviews"`
}