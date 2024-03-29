package config

var (
	LoginUrl = "https://ksc.tukorea.ac.kr/sso/login_proc.jsp?returnurl=null"

	SchoolUrl = "https://dream.tukorea.ac.kr/"

	SessionUrl = SchoolUrl + "com/SsoCtr/initPageWork.do?loginGbn=sso&loginPersNo="

	YearSemesterUrl = SchoolUrl + "aff/dorm/DormCtr/findYyTmGbnList.do?menuId=MPB0022&pgmId=PPB0021"
	NameIdUrl       = SchoolUrl + "com/SsoCtr/findMyGLIOList.do?menuId=MPB0022&pgmId=PPB0021"
	ApplyListUrl    = SchoolUrl + "aff/dorm/DormCtr/findStayAplyList.do?menuId=MPB0022&pgmId=PPB0021"
	RewardListUrl   = SchoolUrl + "aff/dorm/DormCtr/findFindArdListList.do?menuId=MPB0024&pgmId=PPB0023"
	DormStuIdUrl    = SchoolUrl + "aff/dorm/DormCtr/findMdstrmLeaveAplyList.do?menuId=MPB0022&pgmId=PPB0021"
	SendApplyUrl    = SchoolUrl + "aff/dorm/DormCtr/saveOutAplyList.do?menuId=MPB0022&pgmId=PPB0021"
)
