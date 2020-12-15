package uapp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/swxctx/gumeng"

	"github.com/swxctx/ghttp"
)

// doGetRequest 获取数据request
func (uapp *Uapp) doGetRequest(path string, params url.Values) ([]byte, error) {
	// 基础参数
	params.Add("appkey", uapp.AppKey)

	// 1.2 处理签名参数
	sig := gumeng.Sign(uapp.ApiSecurity, path, params)
	params.Add("_aop_signature", sig)
	params.Add("_aop_timestamp", fmt.Sprintf("%d", time.Now().Unix()))

	// 2. 处理请求参数
	req := ghttp.Request{
		Url:       fmt.Sprintf("%s/%s/%s", uapp.GateWay, path, uapp.ApiKey),
		Method:    "GET",
		Query:     params,
		ShowDebug: uapp.Debug,
	}

	// 3. 发起请求
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("doGetRequest: status code not normal, code-> %d, path-> %s", resp.StatusCode, path)
	}
	if resp == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	// 4. 读取返回
	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBs, nil
}
