package route

import (
	"auto_overnight_api/custom_err"
	"auto_overnight_api/functions"
	"auto_overnight_api/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

// FindStayOutList 외박 신청 내역 조회하여 return
func FindStayOutList(c *gin.Context) {

	// 외박 신청 내역 조회에 필요한 것들 파싱
	var requestsModel model.FindRequest
	value, _ := ioutil.ReadAll(c.Request.Body)

	err := json.Unmarshal(value, &requestsModel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_err.ParsingJsonBodyErr)
	}

	// cookie jar 생성
	jar, err := cookiejar.New(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_err.MakeCookieJarErr)
	}

	// cookie jar에 세션 설정
	functions.MakeCookieJar(requestsModel.Cookies, jar)

	// client에 cookie jar 설정
	client := &http.Client{
		Jar: jar,
	}

	// 파싱 시작
	studentInfo := functions.RequestFindUserNm(client)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	if studentInfo.XML.Parameters.Parameter == "-600" {
		c.JSON(http.StatusBadRequest, custom_err.WrongCookieErr)
	}

	requestInfo := model.RequestInfo{
		YY:       requestsModel.Year,
		TmGbn:    requestsModel.TmGbn,
		SchregNo: studentInfo.XML.Dataset[0].Rows.Row[0].Col[1].Data,
		StdKorNm: studentInfo.XML.Dataset[0].Rows.Row[0].Col[0].Data,
	}

	// 외박 신청 내역 조회
	stayOutList := functions.RequestFindStayOutList(client, requestInfo)

	// 파싱 시작
	sm := functions.ParsingStayoutList(stayOutList)

	// 응답 구조체 만들기
	responseBody := make(map[string]interface{})

	responseBody["outStayFrDt"] = sm.OutStayFrDt
	responseBody["outStayToDt"] = sm.OutStayToDt
	responseBody["outStayStGbn"] = sm.OutStayStGbn

	c.JSON(http.StatusOK, responseBody)
}
