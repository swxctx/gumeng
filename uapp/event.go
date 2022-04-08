package uapp

import (
	"encoding/json"
	"net/url"
)

const (
	// 获取事件独立用户数
	eventGetUniqueUsersApi  = "openapi/param2/1/com.umeng.uapp/umeng.uapp.event.getUniqueUsers"
	eventGetUniqueUsersPath = "param2/1/com.umeng.uapp/umeng.uapp.event.getUniqueUsers"

	// 获取事件PV
	eventGetPVApi  = "openapi/param2/1/com.umeng.uapp/umeng.uapp.event.getData"
	eventGetPVPath = "param2/1/com.umeng.uapp/umeng.uapp.event.getData"
)

type (
	/*
		{
			"uniqueUsers":[
				{
					"data":"1234,5678",
					"dates":""
				}
			]
		}
	*/
	// 事件独立用户数
	EventKeyUV struct {
		EventUniqueUsers []*EventUniqueUser `json:"uniqueUsers"`
	}
	EventUniqueUser struct {
		// 日期数组
		Dates []string `json:"dates"`
		// 数据数组
		Data []int64 `json:"data"`
	}

	// 事件PV
	EventKeyPV struct {
		EventData []*EventData `json:"eventData"`
	}
	EventData struct {
		// 日期数组
		Dates []string `json:"dates"`
		// 数据数组
		Data []int64 `json:"data"`
	}
)

// GetEventKeyUV 获取自定义事件UV
func (uapp *Uapp) GetEventKeyUV(eventName, startDate, endDate string) (*EventKeyUV, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Add("startDate", startDate)
	params.Add("endDate", endDate)
	params.Add("eventName", eventName)

	// 2. 发起请求
	resp, err := uapp.doGetRequest(eventGetUniqueUsersApi, eventGetUniqueUsersPath, params)
	if err != nil {
		return nil, err
	}

	var (
		eventKvUV *EventKeyUV
	)
	if err := json.Unmarshal(resp, &eventKvUV); err != nil {
		return nil, err
	}
	return eventKvUV, nil
}

// GetEventKeyPV 获取自定义事件PV
func (uapp *Uapp) GetEventKeyPV(eventName, startDate, endDate string) (*EventKeyPV, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Add("startDate", startDate)
	params.Add("endDate", endDate)
	params.Add("eventName", eventName)

	// 2. 发起请求
	resp, err := uapp.doGetRequest(eventGetPVApi, eventGetPVPath, params)
	if err != nil {
		return nil, err
	}

	var (
		eventKvPV *EventKeyPV
	)
	if err := json.Unmarshal(resp, &eventKvPV); err != nil {
		return nil, err
	}
	return eventKvPV, nil
}
