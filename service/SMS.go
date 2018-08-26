package service

import (
	"fmt"
	"net/url"
	"github.com/spf13/viper"
	"net/http"
	"io/ioutil"
)

func sendSMS(data url.Values) (string, error) {
	resp, err := http.PostForm(viper.GetString("yunpian.url_send_sms"),data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}


func SendSubmitNotice(name, recruitUrl, ZJUid, mobile, intents, instanceName string) (string, error) {
	text := fmt.Sprintf("【求是潮纳新平台】亲爱的%s您好：<br>我们已经收到了您的报名表，请核对下列信息是否准确。重新提交报名表可以修改信息%s<br>您的学号：%s<br>您的手机号：%s<br>您的志愿：%s<br>  我们将在之后通过短信通知您具体的初试信息，敬请留意。感谢您报名%s，期待您的加入！", name, recruitUrl, ZJUid, mobile, intents, instanceName)
	data := url.Values{"apikey": {viper.GetString("yunpian.apikey")}, "mobile": {mobile},"text":{text}}
	return sendSMS(data)
}

func SendRecruitTime(mobile, name, recruitType, instanceName, URL string) (string, error) {
	text := fmt.Sprintf("【求是潮纳新平台】%s同学，我们已为你生成出了%s的时间与地点，请点击以下链接进行选择与确认。（注意：在链接中我们不会要求你输入任何诸如密码等的敏感信息）感谢参与%s。 %s", name, recruitType, instanceName, URL)
	data := url.Values{"apikey": {viper.GetString("yunpian.apikey")}, "mobile": {mobile},"text":{text}}
	return sendSMS(data)
}


func SendRejectNotice(mobile, name, target, us string) (string, error) {
	text := fmt.Sprintf("【求是潮纳新平台】%s同学，我们很遗憾地通知你，你未能成功加入%s。感谢你的支持，欢迎继续关注%s。", name, target, us)
	data := url.Values{"apikey": {viper.GetString("yunpian.apikey")}, "mobile": {mobile},"text":{text}}
	return sendSMS(data)
}