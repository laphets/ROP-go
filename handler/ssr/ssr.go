package ssr

import "rop/model"

type ScheduleResponse struct {
	*model.FreshmanModel
	IntentGroups []*IntentGroup `json:"intent_groups"`
}

type IntentGroup struct {
	*model.IntentModel
	ChineseStage string `json:"chinese_stage"`
	Interviews []*model.FullInterview `json:"interviews"`
}