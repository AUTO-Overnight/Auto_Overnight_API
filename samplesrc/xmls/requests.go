package xmls

import (
	"auto_overnight_api/error_response"
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func RequestFindYYtmgbn(client *http.Client, findYYtmgbnChan chan Root, cookies map[string]string) {
	req, err := http.NewRequest(
		"POST",
		"https://dream.tukorea.ac.kr/aff/dorm/DormCtr/findYyTmGbnList.do?menuId=MPB0022&pgmId=PPB0021",
		bytes.NewBuffer(FindYYtmgbnXML))
	if err != nil {
		panic(err)
	}

	if cookies != nil {
		req.AddCookie(&http.Cookie{Name: "_SSO_Global_Logout_url", Value: cookies["_SSO_Global_Logout_url"]})
		req.AddCookie(&http.Cookie{Name: "kalogin", Value: cookies["kalogin"]})
		req.AddCookie(&http.Cookie{Name: "JSVSESSIONID", Value: cookies["JSVSESSIONID"]})
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	var yytmGbnInfo Root
	xml.Unmarshal(body, &yytmGbnInfo)

	findYYtmgbnChan <- yytmGbnInfo
}

func RequestFindUserNm(client *http.Client, findUserNmChan chan Root, cookies map[string]string) {
	req, err := http.NewRequest(
		"POST",
		"https://dream.tukorea.ac.kr/com/SsoCtr/findMyGLIOList.do?menuId=MPB0022&pgmId=PPB0021",
		bytes.NewBuffer(FindUserNmXML))
	if err != nil {
		panic(err)
	}

	if cookies != nil {
		req.AddCookie(&http.Cookie{Name: "_SSO_Global_Logout_url", Value: cookies["_SSO_Global_Logout_url"]})
		req.AddCookie(&http.Cookie{Name: "kalogin", Value: cookies["kalogin"]})
		req.AddCookie(&http.Cookie{Name: "JSVSESSIONID", Value: cookies["JSVSESSIONID"]})
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	var studentInfo Root
	xml.Unmarshal(body, &studentInfo)

	findUserNmChan <- studentInfo
}

// RequestFindStayOutList 외박 신청 내역을 요청하여 가지고 오는 함수
func RequestFindStayOutList(client *http.Client,
	yy, tmGbn, schregNo, stdKorNm string,
	cookies map[string]string) (Root, *http.Request, error) {

	// 요청 위한 XML 만들기
	findLiveStuNoXML := MakefindLiveStuNoXML(yy, tmGbn, schregNo, stdKorNm)

	// 응답 XML 저장 위한 구조체
	var temp Root

	// 외박 신청 내역 조회를 위한 http request 만들기
	req, err := http.NewRequest(
		"POST",
		"https://dream.tukorea.ac.kr/aff/dorm/DormCtr/findStayAplyList.do?menuId=MPB0022&pgmId=PPB0021",
		bytes.NewBuffer(findLiveStuNoXML))
	if err != nil {
		return temp, nil, error_response.MakeHttpRequestError
	}

	// 입력 받은 쿠키가 존재한다면 설정하기
	if cookies != nil {
		req.AddCookie(&http.Cookie{Name: "_SSO_Global_Logout_url", Value: cookies["_SSO_Global_Logout_url"]})
		req.AddCookie(&http.Cookie{Name: "kalogin", Value: cookies["kalogin"]})
		req.AddCookie(&http.Cookie{Name: "JSVSESSIONID", Value: cookies["JSVSESSIONID"]})
	}

	// http request 보내기
	res, err := client.Do(req)
	if err != nil {
		return temp, nil, error_response.SendHttpRequestError
	}

	// body 읽어서 구조체에 저장
	body, _ := ioutil.ReadAll(res.Body)
	err = xml.Unmarshal(body, &temp)
	if err != nil {
		return temp, nil, error_response.ParsingXMLBodyError
	}

	return temp, req, nil
}

func RequestFindPointList(client *http.Client,
	yy, tmGbn, schregNo, stdKorNm string, cookies map[string]string) Root {
	findPointListXML := MakefindLiveStuNoXML(yy, tmGbn, schregNo, stdKorNm)

	req, err := http.NewRequest(
		"POST",
		"https://dream.tukorea.ac.kr/aff/dorm/DormCtr/findFindArdListList.do?menuId=MPB0024&pgmId=PPB0023",
		bytes.NewBuffer(findPointListXML))
	if err != nil {
		panic(err)
	}

	if cookies != nil {
		req.AddCookie(&http.Cookie{Name: "_SSO_Global_Logout_url", Value: cookies["_SSO_Global_Logout_url"]})
		req.AddCookie(&http.Cookie{Name: "kalogin", Value: cookies["kalogin"]})
		req.AddCookie(&http.Cookie{Name: "JSVSESSIONID", Value: cookies["JSVSESSIONID"]})
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	var temp Root
	err = xml.Unmarshal(body, &temp)
	if err != nil {
		panic(err)
	}

	return temp
}

func RequestSendStayOut(client *http.Client, studentInfo, yytmGbnInfo Root,
	DateList []string, IsWeekend []int, OutStayAplyDt string, cookies map[string]string) error {

	findLiveStuNoXML := MakefindLiveStuNoXML(
		yytmGbnInfo.Dataset[0].Rows.Row[0].Col[0].Data,
		yytmGbnInfo.Dataset[0].Rows.Row[0].Col[1].Data,
		studentInfo.Dataset[0].Rows.Row[0].Col[1].Data,
		studentInfo.Dataset[0].Rows.Row[0].Col[0].Data,
	)

	req, err := http.NewRequest(
		"POST",
		"https://dream.tukorea.ac.kr/aff/dorm/DormCtr/findMdstrmLeaveAplyList.do?menuId=MPB0022&pgmId=PPB0021",
		bytes.NewBuffer(findLiveStuNoXML))
	if err != nil {
		panic(err)
	}

	if cookies != nil {
		req.AddCookie(&http.Cookie{Name: "_SSO_Global_Logout_url", Value: cookies["_SSO_Global_Logout_url"]})
		req.AddCookie(&http.Cookie{Name: "kalogin", Value: cookies["kalogin"]})
		req.AddCookie(&http.Cookie{Name: "JSVSESSIONID", Value: cookies["JSVSESSIONID"]})
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	var liveStuNo Root
	err = xml.Unmarshal(body, &liveStuNo)
	if err != nil {
		panic(err)
	}

	var outStayGbn string
	for i := 0; i < len(DateList); i++ {
		if IsWeekend[i] == 0 {
			outStayGbn = "07"
		} else {
			outStayGbn = "04"
		}

		send(
			MakeSendStayOutXML(
				yytmGbnInfo.Dataset[0].Rows.Row[0].Col[0].Data,
				yytmGbnInfo.Dataset[0].Rows.Row[0].Col[1].Data,
				liveStuNo.Dataset[0].Rows.Row[0].Col[12].Data,
				outStayGbn,
				DateList[i],
				DateList[i],
				OutStayAplyDt,
			),
			client,
			cookies,
		)
	}

	return nil
}

func send(sendStayOutXML []byte, client *http.Client, cookies map[string]string) {
	req, err := http.NewRequest(
		"POST",
		"https://dream.tukorea.ac.kr/aff/dorm/DormCtr/saveOutAplyList.do?menuId=MPB0022&pgmId=PPB0021",
		bytes.NewBuffer(sendStayOutXML))

	if err != nil {
		panic(err)
	}

	if cookies != nil {
		req.AddCookie(&http.Cookie{Name: "_SSO_Global_Logout_url", Value: cookies["_SSO_Global_Logout_url"]})
		req.AddCookie(&http.Cookie{Name: "kalogin", Value: cookies["kalogin"]})
		req.AddCookie(&http.Cookie{Name: "JSVSESSIONID", Value: cookies["JSVSESSIONID"]})
	}

	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
