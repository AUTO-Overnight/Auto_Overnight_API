package custom_error

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
)

var (
	ParsingJsonBodyError   = errors.New("json body 파싱 에러")
	MakeJsonBodyError      = errors.New("json body 생성 에러")
	ParsingXMLBodyError    = errors.New("XML body 파싱 에러")
	EmptyIdOrPasswordError = errors.New("id와 password를 입력해주세요")
	WrongIdOrPasswordError = errors.New("id 또는 비밀번호를 잘못 입력했습니다")
	WrongCookieError       = errors.New("쿠키/세션이 만료되거나 잘못됐습니다.")
	MakeCookieJarError     = errors.New("cookiejar 생성 에러")
	MakeHttpRequestError   = errors.New("http request 생성 에러")
	SendHttpRequestError   = errors.New("http request 요청 에러")
)

// MakeErrorResponse 에러 응답을 만드는 함수
func MakeErrorResponse(err error, statusCode int) (events.APIGatewayProxyResponse, error) {
	body := make(map[string]string)
	body["error"] = err.Error()

	responseJson, _ := json.Marshal(body)
	response := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(responseJson),
	}
	return response, nil
}
