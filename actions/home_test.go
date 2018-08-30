package actions

func (as *ActionSuite) Test_RootHandler() {
	res := as.JSON("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "")
}
