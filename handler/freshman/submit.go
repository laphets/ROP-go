package freshman

import (
	"github.com/gin-gonic/gin"
	"rop/pkg/errno"
	. "rop/handler"
	"strconv"
	"rop/model"
	form2 "rop/handler/form"
	"encoding/json"
	"github.com/lexkong/log"
	"regexp"
	"errors"
	"strings"
	"rop/service"
	"fmt"
)

//func praseBranch(visited map[int]int, formTemplate []*form2.DataItem, submission map[int]string) error {
//	for _, question := range formTemplate {
//		for _, node := range formTemplate {
//			if visited[node.Tag] == 0 && !node.DefaultJump {
//				// multi branch
//				if node.Required && submission[node.Tag]
//			}
//		}
//	}
//}
//
//func checkValidate(visited map[int]int, formTemplate []*form2.DataItem, submission map[int]string) error {
//	// First find root
//	root :=
//
//	// First check for multi-select
//	for _, question := range formTemplate {
//		if visited[question.Tag] == 0 && question.Type == "SELECT" {
//
//		}
//	}
//
//
//
//	// Check for required
//	for _, question := range formTemplate {
//		if visited[question.Tag] == 0 {
//			// Check itself
//			// Check for required
//			if question.Required {
//				// Then find in submission
//				if _, ok := submission[question.Tag]; ok && submission[question.Tag] != "" {
//
//				} else {
//					// Not exist
//					return errors.New("Not required.")
//				}
//			}
//
//			ans := submission[question.Tag]
//			// TODO: the empty value of slice
//			if question.Choices != nil {
//				// Check for multi choice
//
//			}
//		}
//	}
//
//	// Check for RE
//	for _, question := range formTemplate {
//		if question.Re != "" {
//			if _, err := regexp.MatchString(question.Re, submission[question.Tag]); err != nil {
//				return errors.New("Not RE.")
//			}
//		}
//	}
//}
//
func goThisWay(curTag int, formMap map[int]*form2.DataItem, submission map[int][]string) bool {
	if curTag == -1 {
		return false
	}
	if len(submission[curTag]) != 0 && submission[curTag][0] != "" {
		return true
	}
	return goThisWay(formMap[curTag].Next, formMap, submission)
}

func dfs(curTag int, formMap map[int]*form2.DataItem, submission map[int][]string, inBranch bool) error {
	if curTag == -1 {
		return nil
	}

	curForm := formMap[curTag]

	//submitAns := submission[curTag]
	next := curForm.Next

	if curForm.Type == "TEXT" {
		return dfs(next, formMap, submission, inBranch)
	}

	// Check for required
	//log.Debugf("%t", curForm.Required)

	if curForm.Required && !inBranch {
		if _, ok := submission[curTag]; ok && len(submission[curTag]) != 0 && submission[curTag][0] != "" {

		} else {
			log.Debugf("Required %d", curTag)
			// Not exist
			return errors.New("Not required.")
		}
	}

	// Check for branch
	if curForm.Type == "SELECT" {
		for _, tag := range submission[curTag] {
			taginInt, err := strconv.ParseInt(tag, 10, 32)
			if err != nil {
				return err
			}
			flag := false
			for _, tarChoice := range curForm.Choices {
				if int(taginInt) == tarChoice.Tag {
					flag = true
				}
			}
			if !flag {
				return errors.New("Choice not match.")
			}
		}
		avali := curForm.AvailableCnt
		if avali < len(submission[curTag]) {
			return errors.New("Choice too much.")
		}

		if curForm.DefaultJump {
			// DefaultJump
			// Done
			return dfs(next, formMap, submission, inBranch)
		} else {
			cnt := 0
			// answer contains which choice
			hasThisChoice := make(map[int]bool)
			for _, tag := range submission[curTag] {
				taginInt, err := strconv.ParseInt(tag, 10, 32)
				if err != nil {
					return err
				}
				hasThisChoice[int(taginInt)] = true
			}

			for _, choice := range curForm.Choices {
				// if it is answer branch
				if hasThisChoice[choice.Tag] {
					if goThisWay(choice.Next, formMap, submission) {
						cnt ++
						if err := dfs(choice.Next, formMap, submission, false); err != nil {
							return err
						}
					} else {
						return errors.New("Branches not match.")
					}
				} else {
					if goThisWay(choice.Next, formMap, submission) {
						return errors.New("Submit too much.")
					}
				}

			}
			//if cnt == 0 {
			//	return errors.New("Branches not match.")
			//}
			return dfs(next, formMap, submission, inBranch)
		}

	}

	// Check for RE for common type
	if curForm.Re != "" && len(submission[curTag]) != 0 {
		//log.Debugf("%s 111 %+v %d", curForm.Re, submission[curTag][0], curTag)
		_, _ = regexp.MatchString(curForm.Re, submission[curTag][0])
		//log.Debugf("%s 222")
		if ok, err := regexp.MatchString(curForm.Re, submission[curTag][0]); !ok || err != nil {
			log.Debugf("fail %s %s", curForm.Re, submission[curTag][0])
			return errors.New("Not RE.")
		}
	}

	return dfs(next, formMap, submission, inBranch)
}

func Submit(c *gin.Context) {
	instanceId, err := strconv.ParseUint(c.Query("instanceId"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	instance, err := model.GetInstanceById(uint(instanceId))
	if err != nil {
		SendResponse(c, errno.ErrInstanceNotFound, nil)
		return
	}

	req := &SubmitRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	form, err := model.GetFormByID(instance.FormId)
	if err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}

	//log.Debugf("%s", form.Data)

	submitArray := req.Data

	// Trans submission into map style
	submission := make(map[int][]string)
	for _, item := range submitArray {
		submission[item.Key] = item.Value
	}

	// Get form template
	formArray := make([]*form2.DataItem, 0)
	json.Unmarshal([]byte(form.Data), &formArray)
	root := form.RootTag

	formMap := make(map[int]*form2.DataItem)
	for _, item := range formArray {
		formMap[item.Tag] = item
	}

	//log.Debugf("%d", root)
	//log.Debugf("%+v", req.Data[0].Value)

	// Tags for visted(prased)

	if err := dfs(root, formMap, submission, false); err != nil {
		SendResponse(c, errno.ErrOperation, err.Error())
		return
	}

	otherInfoArray := make([]*OtherInfo, 0)
	//go func() {
	for _, submit := range submitArray {
		question := formMap[submit.Key]
		otherInfo := &OtherInfo{}
		if question.Type != "SELECT" {
			otherInfo.Key = question.Text
			otherInfo.Value = submit.Value[0]
		} else {
			// SELECT
			otherInfo.Key = question.Text
			ansList := make([]string, 0)
			for _, choice := range submit.Value {
				curTag, err := strconv.ParseInt(choice, 10, 32)
				if err != nil {
					SendResponse(c, errno.ErrTypeNotMatch, nil)
					return
				}
				for _, item := range question.Choices {
					if int(curTag) == item.Tag {
						ansList = append(ansList, item.Text)
						break
					}
				}
			}
			otherInfo.Value = strings.Join(ansList, "%")
		}
		//log.Debugf("%+v", otherInfo)
		otherInfoArray = append(otherInfoArray, otherInfo)
	}
	//}()
	otherInfoToJson, err := json.Marshal(otherInfoArray)
	if err != nil {
		SendResponse(c, errno.ErrPrase, err)
		return
	}

	//log.Debugf("%s", string(otherInfoToJson))

	freshman := &model.FreshmanModel{
		OtherInfo: string(otherInfoToJson),
	}

	intentList := make([]string, 0)

	for _, ans := range submitArray {
		if form, ok := formMap[ans.Key]; ok {
			switch form.Spec {
			case "NAME":
				freshman.Name = ans.Value[0]
			case "ZJUID":
				freshman.ZJUid = ans.Value[0]
			case "MOBILE":
				freshman.Mobile = ans.Value[0]
			case "EMAIL":
				freshman.Email = ans.Value[0]
			case "GENDER":
				freshman.Gender = ans.Value[0]
			case "PHOTO":
				freshman.Photo = ans.Value[0]
			case "DEPART":
				// Intent
				for _, choice := range ans.Value {
					curTag, err := strconv.ParseInt(choice, 10, 32)
					if err != nil {
						SendResponse(c, errno.ErrTypeNotMatch, nil)
						return
					}

					for _, tar := range form.Choices {
						if int(curTag) == tar.Tag {
							department := tar.Text
							intentList = append(intentList, department)
						}
					}
				}
			default:
			}
		}
	}

	if len(intentList) > instance.MaxIntent {
		SendResponse(c, errno.TooMuchIntent, nil)
		return
	}

	if err := freshman.Create(); err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}

	intents := make([]*model.IntentModel, 0)

	for _, intentString := range intentList {
		intent := &model.IntentModel{
			FreshmanId: freshman.ID,
			Department: intentString,
			MainStage: "Public Sea",
			SubStage: 1,
		}
		intents = append(intents, intent)
	}


	if err := model.CreateIntents(intents); err != nil {
		SendResponse(c, errno.DBError, err)
	}

	encryptedInstanceId, err := service.Encrypt(strconv.FormatUint(uint64(freshman.InstanceId), 10))

	_, err = service.SendSubmitNotice(freshman.Mobile, freshman.Name, fmt.Sprintf("https://101.132.66.238:8081?uid=%s", encryptedInstanceId), instance.Name, freshman.ZJUid, freshman.Mobile, intents[0].Department, intents[1].Department, instance.Name )
	if err != nil {
		SendResponse(c, errno.ErrSMS, err.Error())
		return
	}
	//freshman := &model.FreshmanModel{
	//	InstanceId: uint(instanceId),
	//	ZJUid: "3170111705",
	//	Mobile: "18888922004",
	//	Name: "罗文卿",
	//	//MainStage: "Public Sea",
	//	//SubStage: "None",
	//	OtherInfo: "{a json here}",
	//}
	//
	//if err := freshman.Create(); err != nil {
	//	SendResponse(c, errno.DBError, err)
	//	return
	//}
	//
	//intent1 := &model.IntentModel{
	//	FreshmanId: freshman.ID,
	//	Department: "技术研发中心",
	//	//GroupId: 0,
	//	MainStage: "Public Sea",
	//	SubStage: 1,
	//}
	//
	//intent2 := &model.IntentModel{
	//	FreshmanId: freshman.ID,
	//	Department: "人力资源部门",
	//	//GroupId: 0,
	//	MainStage: "Public Sea",
	//	SubStage: 1,
	//}
	//
	//intents := []*model.IntentModel{intent1, intent2}
	//
	//if err := model.CreateIntents(intents); err != nil {
	//	SendResponse(c, errno.DBError, err.Error())
	//	return
	//}

	SendResponse(c, nil, nil)
}


//
//res1, err1 := service.SendRejectNotice("18867136212", "博亚", "求是潮2018秋纳", "求是潮")
//if err1 != nil {
//	SendResponse(c, errno.ErrSMS, err.Error())
//	return
//}
//SendResponse(c, nil, res1)