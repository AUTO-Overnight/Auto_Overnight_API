package route

import (
	"auto_overnight_api/config"
	"auto_overnight_api/custom_err"
	"auto_overnight_api/functions"
	"auto_overnight_api/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// Login Id와 Password를 Json으로 입력 받아 로그인하고 이름, 년도, 학기, 쿠키, 외박 신청 내역을 return
func Login(c *gin.Context) {

	// Id, Password 파싱
	var requestsModel model.LoginRequest
	value, _ := ioutil.ReadAll(c.Request.Body)

	err := json.Unmarshal(value, &requestsModel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_err.ParsingJsonBodyErr)
	}
	if requestsModel.Id == "" || requestsModel.PassWord == "" {
		c.JSON(http.StatusBadRequest, custom_err.EmptyIdOrPasswordErr)
	}

	// PassWord URIDecode
	decodeValue, err := url.QueryUnescape(requestsModel.PassWord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_err.ParsingJsonBodyErr)
	}

	// x-www-form-urlencoded 방식으로 로그인 하기 위해 form 생성
	loginInfo := url.Values{
		"internalId": {requestsModel.Id},
		"internalPw": {decodeValue},
		"gubun":      {"inter"},
	}

	// cookie jar 생성
	jar, err := cookiejar.New(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_err.MakeCookieJarErr)
	}

	// 로그인 http request 생성
	req, err := http.NewRequest("POST", config.LoginUrl, bytes.NewBufferString(loginInfo.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_err.MakeHttpRequestErr)
	}

	// Content-Type 헤더 설정, client에 cookie jar 설정
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		Jar: jar,
	}
	// 로그인 시도
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_err.SendHttpRequestErr)
	}

	defer res.Body.Close()

	// 통합 정보 시스템 세션 얻기 시도
	req, err = http.NewRequest("GET", config.SessionUrl, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_err.MakeHttpRequestErr)
	}

	res, err = client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_err.SendHttpRequestErr)
	}

	// 파싱 시작
	studentInfo := functions.RequestFindUserNm(client)
	yytmGbnInfo := functions.RequestFindYYtmgbn(client)

	if studentInfo.Error != nil || yytmGbnInfo.Error != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	if studentInfo.XML.Parameters.Parameter == "-600" {
		c.JSON(http.StatusBadRequest, custom_err.WrongIdOrPasswordErr)
	}

	requestInfo := model.RequestInfo{
		YY:       yytmGbnInfo.XML.Dataset[0].Rows.Row[0].Col[0].Data,
		TmGbn:    yytmGbnInfo.XML.Dataset[0].Rows.Row[0].Col[1].Data,
		SchregNo: studentInfo.XML.Dataset[0].Rows.Row[0].Col[1].Data,
		StdKorNm: studentInfo.XML.Dataset[0].Rows.Row[0].Col[0].Data,
	}

	// 외박 신청 내역 조회
	stayOutList := functions.RequestFindStayOutList(client, requestInfo)

	if stayOutList.Error != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// 파싱 시작
	sm := functions.ParsingStayoutList(stayOutList)
	cookie := functions.ParsingCookies(client)

	// 응답 구조체 만들기
	responseBody := make(map[string]interface{})

	// 이름, 년도, 학기 저장
	responseBody["name"] = studentInfo.XML.Dataset[0].Rows.Row[0].Col[0].Data
	responseBody["yy"] = yytmGbnInfo.XML.Dataset[0].Rows.Row[0].Col[0].Data
	responseBody["tmGbn"] = yytmGbnInfo.XML.Dataset[0].Rows.Row[0].Col[1].Data

	responseBody["cookies"] = cookie
	responseBody["outStayFrDt"] = sm.OutStayFrDt
	responseBody["outStayToDt"] = sm.OutStayToDt
	responseBody["outStayStGbn"] = sm.OutStayStGbn

	c.JSON(http.StatusOK, responseBody)
}
