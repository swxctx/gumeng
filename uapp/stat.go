package uapp

import (
	"encoding/json"
	"net/url"
)

const (
	// 获取新增数据
	getNewDataApi = "openapi/param2/1/com.umeng.uapp/umeng.uapp.getNewAccounts"
)

type (
	/*
		{
			"newAccountInfo":[
				{
					"date":"2018-01-01",
					"hourNewUser":"[11,65,49,4,4,8,25,29,31,29,32,29,38,63,39,33,34,41,40,53,12,77,86,50]",
					"newUser":0,
					"newAccount":0,
					"hourNewAccount":"[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]"
				}
			]
		}
	*/
	// 新增数据
	NewUserData struct {
		// 日期
		date string `json:"data"`
		// 新增用户
		newUser int64 `json:"new_user"`
		// 新增账号
		newAccount int64 `json:"new_account"`
		// 小时新增用户（按小时查询时）
		hourNewUser []int64 `json:"hour_new_user"`
		// 小时新增账号（按小时查询时）
		hourNewAccount []int64 `json:"hour_new_account"`
	}
)

// GetNewUserData 获取新增数据
func (uapp *Uapp) GetNewUserData(eventName, startDate, endDate string) (*NewUserData, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Add("startDate", startDate)
	params.Add("endDate", endDate)
	params.Add("periodType", "daily")

	// 2. 发起请求
	resp, err := uapp.doGetRequest(getNewDataApi, params)
	if err != nil {
		return nil, err
	}

	var (
		newUserData *NewUserData
	)
	if err := json.Unmarshal(resp, &newUserData); err != nil {
		return nil, err
	}
	return newUserData, nil
}

const (
	// 获取活跃数据
	getActiveUserDataApi = "openapi/param2/1/com.umeng.uapp/umeng.uapp.getActiveAccounts"
)

type (
	/*
		{
			"activeAccountInfo":[
				{
					"date":"2018-01-01",
					"activeAccount":0,
					"activeUser":0
				}
			]
		}
	*/
	// 新增数据
	ActiveUserData struct {
		// 日期
		date string `json:"data"`
		// 新增用户
		activeUser int64 `json:"activeUser"`
		// 新增账号
		activeAccount int64 `json:"activeAccount"`
	}
)

// GetNewUserData 获取新增数据
func (uapp *Uapp) GetActiveUserData(eventName, startDate, endDate string) (*ActiveUserData, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Add("startDate", startDate)
	params.Add("endDate", endDate)
	params.Add("periodType", "daily")

	// 2. 发起请求
	resp, err := uapp.doGetRequest(eventGetUniqueUsersApi, params)
	if err != nil {
		return nil, err
	}

	var (
		activeUserData *ActiveUserData
	)
	if err := json.Unmarshal(resp, &activeUserData); err != nil {
		return nil, err
	}
	return activeUserData, nil
}
