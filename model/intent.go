package model

import (
	"time"
	"github.com/lexkong/log"
)

type IntentModel struct {
	ID        uint `gorm:"primary_key" json:"intent_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"unique_index:idx_freshman_department"`
	FreshmanId uint `gorm:"not null;unique_index:idx_freshman_department" json:"freshman_id"`
	Department string `gorm:"not null;unique_index:idx_freshman_department" json:"department"`
	InterviewId uint `gorm:"index" json:"interview_id"`
	MainStage string `json:"main_stage"`
	SubStage int `json:"sub_stage"`
	TargetInterviewId uint `json:"target_interview_id"`
}

func (x *IntentModel) Create() error {
	return DB.Local.Create(&x).Error
}

// You need set ID in update method
func (x *IntentModel) Update() error {
	if _, err := GetIntentByID(x.ID); err != nil {
		return err
	}
	return DB.Local.Model(&x).Updates(&x).Error
}

// This method need to be checked
// Do delete in FreshmanModel
func CreateIntents(intentsData []*IntentModel) error {
	 //curIntents := make([]*IntentModel, 0)
	 //if !DB.Local.Where("freshman_id = ?", lastfreshmanId).Find(&curIntents).RecordNotFound() {
		//// if record exist, then all replace
		//// Delete and insert
		//log.Debugf("%d %d", len(curIntents), lastfreshmanId)
		// for _, item := range curIntents {
		// 	log.Debugf("%d", item.ID)
		//	 if err := DeleteIntent(item.ID); err != nil {
		//		 return err
		//	 }
		// }
	 //}
	 // Then insert
	for _, item := range intentsData {
		if err := item.Create(); err != nil {
			return err
		}
	}
	return nil
}

func ListIntentByFreshman(freshmanId uint) ([]*IntentModel, error) {
	intents := make([]*IntentModel, 0)
	d := DB.Local.Where("freshman_id = ?", freshmanId).Find(&intents)
	return intents, d.Error
}

func ListIntentByInterview(interviewId uint) ([]*FullIntent, error) {
	intents := make([]*IntentModel, 0)
	if err := DB.Local.Where("interview_id = ?", interviewId).Find(&intents).Error; err != nil {
		return nil, err
	}
	fulIntents := make([]*FullIntent, 0)
	for _, item := range intents {
		tmp, err := GetFullIntentByID(item.ID)
		if err != nil {
			return nil, err
		}
		fulIntents = append(fulIntents, tmp)
	}
	return fulIntents, nil
}


func DeleteIntent(intentId uint) error {
	log.Debugf("%d", intentId)
	intent := &IntentModel{}
	intent.ID = intentId
	return DB.Local.Delete(&intent).Error
}

func GetIntentByID(intentId uint) (*IntentModel, error) {
	intent := &IntentModel{}
	d := DB.Local.Where("ID = ?", intentId).First(&intent)
	//log.Debugf("%d", intentId)
	return intent, d.Error
}

type FullIntent struct {
	*IntentModel
	*FreshmanModel
	Interview *InterviewModel `json:"interview"`
}

// for conditions, [0] is stage, [1] is department
func ListFullIntentByInstance(instanceId uint, conditions ...string) ([]*FullIntent, error) {
	mainStage, department := conditions[0], conditions[1]
	freshmans, err := ListFreshman(instanceId)
	if err != nil {
		return nil, err
	}
	fulIntents := make([]*FullIntent, 0)
	for _, freshman := range freshmans {
		intents, err := ListIntentByFreshman(freshman.ID)
		if err != nil {
			return nil, err
		}
		for _, intent := range intents {
			// Check condition here
			if mainStage != "" && intent.MainStage != mainStage {
				continue
			}
			if department != "" && intent.Department != department {
				continue
			}
			fulIntent, err := GetFullIntentByID(intent.ID)
			if err != nil {
				return nil, err
			}
			fulIntents = append(fulIntents, fulIntent)
		}
	}
	return fulIntents, nil
}

func GetFullIntentByID(intentId uint) (*FullIntent, error) {
	intent, err := GetIntentByID(intentId)
	if err != nil {
		return nil, err
	}
	freshman, err := GetFreshmanById(intent.FreshmanId);
	if err != nil {
		return nil, err
	}
	fulIntent := &FullIntent{
		IntentModel: intent,
		FreshmanModel: freshman,
	}
	if intent.InterviewId != 0 {
		interview, err := GetInterviewByID(intent.InterviewId)
		if err != nil {
			return nil, err
		}
		fulIntent.Interview = interview
	}
	return fulIntent, nil
}