package actions

func (as *ActionSuite) Test_CLIRootHandler() {
	req := as.JSON("/")
	req.Headers["User-Agent"] = "curl"

	res := req.Get()
	as.Equal(200, res.Code)
	as.Equal("", res.Body.String())
}

func (as *ActionSuite) Test_RootHandler() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "<html>\n<head>")
}

func (as *ActionSuite) Test_AllHeadersHandler() {
	req := as.HTML("/all")
	req.Headers["User-Agent"] = "curl"
	res := req.Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "user-agent: curl")
}
