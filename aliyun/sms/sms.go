// This file is auto-generated, don't edit it. Thanks.
package sms

import (
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
)

var _accessKeyId = ""
var _accessKeySecret = ""

// 使用AK&SK初始化账号Client
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi.Client, _err error) {
	config := &openapi.Config{}
	config.AccessKeyId = accessKeyId
	config.AccessKeySecret = accessKeySecret
	_result = &dysmsapi.Client{}
	_result, _err = dysmsapi.NewClient(config)
	return _result, _err
}

func Init(accessKeyId, accessKeySecret string) {
	_accessKeyId = accessKeyId
	_accessKeySecret = accessKeySecret
}

func Send(phone, sign, templateCode, templateparam string) (err error) {
	config := &openapi.Config{}
	config.AccessKeyId = &_accessKeyId
	config.AccessKeySecret = &_accessKeySecret
	client, err := dysmsapi.NewClient(config)
	if err != nil {
		return
	}

	// 1.发送短信
	sendReq := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  &phone,
		SignName:      &sign,
		TemplateCode:  &templateCode,
		TemplateParam: &templateparam,
	}
	sendResp, err := client.SendSms(sendReq)
	if err != nil {
		return
	}

	code := sendResp.Body.Code
	if code == nil || *code != "OK" {
		err = fmt.Errorf("错误信息: %s", *sendResp.Body.Message)
		return
	}

	return
}
