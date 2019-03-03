package ssr

import "git.zjuqsc.com/rop/ROP-go/model"

type ScheduleResponse struct {
	*model.FreshmanModel
	IntentGroups []*IntentGroup `json:"intent_groups"`
}

type IntentGroup struct {
	*model.IntentModel
	ChineseStage string `json:"chinese_stage"`
	Interviews []*model.FullInterview `json:"interviews"`
}

type RegisterResponse struct {
	Name string `json:"name"`
	Association *model.AssociationModel `json:"association"`
}
