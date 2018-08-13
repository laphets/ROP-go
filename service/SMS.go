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