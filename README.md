# ⚡ Auto_Overnight_API

- 외박 신청 및 로그인 등 여러 요청을 처리하는 API입니다.

- serverless의 nodejs 템플릿을 이용해 AWS lambda에 배포했습니다.

- 자세한 코드는 [여기](https://github.com/AUTO-Overnight/Auto_Overnight_API/blob/main/src/handler.js)를 참조해주세요!

## 📝 API

- 모든 요청은 POST를 사용합니다.


### login

* 로그인을 시도하는 함수입니다.

#### 입력

```
{
    "id" : "201112345678",
    "password" : "1234"
}
```

#### 출력

```
{
    "cookies": "_SSO_Global_Logout_url=get%5 ... amV1c19kb21haW4vanN2XzI=",
    "name": "홍길동",
    "yy" : 2021,
    "tmGbn" : 2,
    "outStayFrDt": [
        "20210830",
        "20210827",
        "20210825"
    ],
    "outStayToDt": [
        "20210830",
        "20210829",
        "20210825"
    ],
    "outStayStGbn": [
        "2",
        "2",
        "2"
    ]
}
```

* yy 년도

* tmGbn  학기 구분 ( 1학기 : 1 / 2학기 : 2 / 여름학기 : 5 / 겨울학기 : 6 )

* outStayStGbn  1 : 미승인 / 2 : 승인

### sendStayOut

* 외박 신청을 하는 함수입니다.

#### 입력

```
{
    "date_list" : [20210831, 20210901, 20210902, 20210903, 20210904, 20210905],
    "is_weekend" : [0, 0, 0, 1, 1, 1],
    "outStayAplyDt" : "20210831",
    "schregNo" : "201112345678",
    "cookies" : "_SSO_Global_Logout_url=get%5 ... amV1c19kb21haW4vanN2XzI= "
}
```

* is_weekend  0 : 평일 / 1 : 주말

* outStayAplyDt  오늘 날짜

* schregNo  학번

#### 출력

```
{
    "outStayFrDt": [
        "20210905",
        "20210904",
        "20210903",
        "20210902",
        "20210901",
        "20210831",
        "20210830",
        "20210827",
        "20210825"
    ],
    "outStayToDt": [
        "20210905",
        "20210904",
        "20210903",
        "20210902",
        "20210901",
        "20210831",
        "20210830",
        "20210829",
        "20210825"
    ],
    "outStayStGbn": [
        "1",
        "1",
        "1",
        "1",
        "1",
        "1",
        "2",
        "2",
        "2"
    ]
}
```

* 외박신청에 성공하면 외박 신청 내역을 return 합니다.

### findStayOutList

* 외박 신청 내역을 조회하는 함수입니다.

#### 입력

```
{
    "yy" : "2021",
    "tmGbn" : "2",
    "schregNo" : "201112345678",
    "userNm" : "홍길동",
    "cookies" : "_SSO_Global_Logout_url=get%5 ... amV1c19kb21haW4vanN2XzI="
}
```

* yy  년도

* tmGbn  학기 구분 / 1학기 : 1 / 2학기 : 2 / 여름학기 : 5 / 겨울학기 : 6

* schregNo  학번

#### 출력

```
{
    "outStayFrDt": [
        "20210830",
        "20210827",
        "20210825"
    ],
    "outStayToDt": [
        "20210830",
        "20210829",
        "20210825"
    ],
    "outStayStGbn": [
        "2",
        "2",
        "2"
    ]
}
```

### findPointList

* 상벌점 내역을 조회하는 함수입니다.

#### 입력

```
{
    "schregNo" : "201112345678",
    "userNm" : "홍길동",
    "cookies" : "_SSO_Global_Logout_url=get%5 ... amV1c19kb21haW4vanN2XzI="
}
```

* schregNo  학번

#### 출력

```
{
    "cmpScr": [
        "-6",
        "-6"
    ],
    "lifSstArdGbn": [
        "2",
        "2"
    ],
    "ardInptDt": [
        "20210808",
        "20210807"
    ],
    "lifSstArdCtnt": [
        "체온미측정",
        "체온미측정"
    ]
}
```

* cmpScr  상벌점 수치

* lifSstArdGbn  상벌점 구분 / 1 : 상점 / 2 : 벌점

* ardInptDt  상벌일자

* lifSstArdCtnt  상벌내용


