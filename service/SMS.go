package service

import (
	"encoding/json"
	"fmt"
	"git.zjuqsc.com/rop/ROP-go/model"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type SMSInfo struct {
	Balance float32 `json:"balance"`
	Mobile string `json:"mobile"`
}

func GetAccountInfo() (*SMSInfo, error) {
	u := "https://sms.yunpian.com/v2/user/get.json"

	payload := strings.NewReader(fmt.Sprintf("apikey=%s", viper.GetString("yunpian.tpl_send_sms")))

	req, _ := http.NewRequest("POST", u, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(res)

	info := SMSInfo{}
	json.Unmarshal(body, &info)

	return &info, nil
}

func sendSMS(data url.Values, tpl bool) (string, error) {
	var resp *http.Response
	var err error
	if tpl {
		resp, err = http.PostForm(viper.GetString("yunpian.tpl_send_sms"),data)
	} else {
		resp, err = http.PostForm(viper.GetString("yunpian.url_send_sms"),data)
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	sms4Save := model.SmsModel{
		Mobile: data["mobile"][0],
		Text: data["text"][0],
		Result: string(body),
		//Error: err.Error(),
	}
	go sms4Save.Create()

	if err != nil {
		return "", err
	}

	//log.Debugf("%s", string(body))
	return string(body), nil
}


func SendSubmitNotice(name, ZJUid, mobile, intents, instanceName string) (string, error) {
	//text := fmt.Sprintf("【求是潮纳新平台】亲爱的%s您好：<br>我们已经收到了您的报名表，请核对下列信息是否准确。重新提交报名表可以修改信息%s<br>您的学号：%s<br>您的手机号：%s<br>您的志愿：%s<br>  我们将在之后通过短信通知您具体的初试信息，敬请留意。感谢您报名%s，期待您的加入！", name, recruitUrl, ZJUid, mobile, intents, instanceName)
	text := fmt.Sprintf("【求是潮纳新平台】亲爱的%s，我们已经收到了您的报名表，请核对下列信息，重复提交可覆盖原表。学号：%s 。部门志愿： %s 。我们将在之后通过短信通知您具体的初试信息，敬请留意。感谢您报名%s，期待您的加入！", name, ZJUid, intents, instanceName)
	data := url.Values{"apikey": {viper.GetString("yunpian.apikey")}, "mobile": {mobile},"text":{text}}
	return sendSMS(data, false)
}

func SendRecruitTime(mobile, name, recruitType, instanceName, URL string) (string, error) {
	text := fmt.Sprintf("【求是潮纳新平台】%s同学，我们已为你生成出了%s的时间与地点，请点击以下链接进行选择与确认。（注意：在链接中我们不会要求你输入任何诸如密码等的敏感信息）感谢参与%s。 %s", name, recruitType, instanceName, URL)
	data := url.Values{"apikey": {viper.GetString("yunpian.apikey")}, "mobile": {mobile},"text":{text}}
	return sendSMS(data, false)
}


func SendRejectNotice(mobile, name, target, us string) (string, error) {
	text := fmt.Sprintf("【求是潮纳新平台】%s同学，我们很遗憾地通知你，你未能成功加入%s。感谢你的支持，欢迎继续关注%s。", name, target, us)
	data := url.Values{"apikey": {viper.GetString("yunpian.apikey")}, "mobile": {mobile},"text":{text}}
	return sendSMS(data, false)
}

func SendInterviewerNotice(mobile, name, association, department, instanceName string) (string, error) {
	text := fmt.Sprintf("【求是潮纳新平台】%s%s%s同学您好，求是潮纳新开放平台温馨提醒您注意最近为您安排的%s日程。详情请登录ROP管理后台查看: https://rop.zjuqsc.com/console", association, department, name, instanceName)
	data := url.Values{"apikey": {viper.GetString("yunpian.apikey")}, "mobile": {mobile},"text":{text}}
	return sendSMS(data, false)
}
