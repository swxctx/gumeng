package uapp

import (
	"encoding/json"
	"net/url"
)

const (
	// 获取新增数据
	getNewDataApi  = "openapi/param2/1/com.umeng.uapp/umeng.uapp.getNewAccounts"
	getNewDataPath = "param2/1/com.umeng.uapp/umeng.uapp.getNewAccounts"
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

	NewUserDataResult struct {
		NewAccountInfo []*NewUserData `json:"newAccountInfo"`
	}

	// 新增数据
	NewUserData struct {
		// 日期
		Date string `json:"date"`
		// 新增用户
		NewUser int64 `json:"newUser"`
		// 新增账号
		NewAccount int64 `json:"newAccount"`
		// 小时新增用户（按小时查询时）
		HourNewUser []int64 `json:"hourNewUser"`
		// 小时新增账号（按小时查询时）
		HourNewAccount []int64 `json:"hourNewAccount"`
	}
)

// GetNewUserData 获取新增数据
func (uapp *Uapp) GetNewUserData(startDate, endDate string) (*NewUserDataResult, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Add("startDate", startDate)
	params.Add("endDate", endDate)
	params.Add("periodType", "daily")

	// 2. 发起请求
	resp, err := uapp.doGetRequest(getNewDataApi, getNewDataPath, params)
	if err != nil {
		return nil, err
	}

	var (
		newUserData *NewUserDataResult
	)
	if err := json.Unmarshal(resp, &newUserData); err != nil {
		return nil, err
	}
	return newUserData, nil
}

const (
	// 获取活跃数据
	getActiveUserDataApi  = "openapi/param2/1/com.umeng.uapp/umeng.uapp.getActiveAccounts"
	getActiveUserDataPath = "param2/1/com.umeng.uapp/umeng.uapp.getActiveAccounts"
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

	ActiveUserDataResult struct {
		ActiveAccountInfo []*ActiveUserData `json:"activeAccountInfo"`
	}
	// 活跃数据
	ActiveUserData struct {
		// 日期
		Date string `json:"date"`
		// 活跃用户
		ActiveUser int64 `json:"activeUser"`
		// 活跃账号
		ActiveAccount int64 `json:"activeAccount"`
	}
)

// GetNewUserData 获取活跃数据
func (uapp *Uapp) GetActiveUserData(startDate, endDate string) (*ActiveUserDataResult, error) {
	// 1. 处理参数
	params := url.Values{}
	params.Add("startDate", startDate)
	params.Add("endDate", endDate)
	params.Add("periodType", "daily")

	// 2. 发起请求
	resp, err := uapp.doGetRequest(getActiveUserDataApi, getActiveUserDataPath, params)
	if err != nil {
		return nil, err
	}

	var (
		activeUserData *ActiveUserDataResult
	)
	if err := json.Unmarshal(resp, &activeUserData); err != nil {
		return nil, err
	}
	return activeUserData, nil
}
