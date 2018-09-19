package actions

func (as *ActionSuite) Test_CLIRootHandler() {
	as.Willie.Headers["User-Agent"] = "curl"
	res := as.JSON("/").Get()
	as.Equal(200, res.Code)
	as.Equal("", res.Body.String())
}

func (as *ActionSuite) Test_RootHandler() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "<html>\n<head>")
}

func (as *ActionSuite) Test_AllHeadersHandler() {
	as.Willie.Headers["User-Agent"] = "curl"
	res := as.HTML("/all").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "user-agent: curl")
}
