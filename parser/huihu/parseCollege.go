package huihu

import (
	"huihuCrawler02/engine"
	"regexp"
)

const collegeRe = `<a href="(/teacher/view.html[?]name=[^"]+)" target="_blank">([^&]+)&nbsp;</a>`

func ParseCollege(contents []byte, url string) engine.ParseResult {
	//contents, err := ioutil.ReadFile("huihuCrawler/parser/huihu/college_test_data.html")
	//if err != nil {
	//	panic(err)
	//}
	re := regexp.MustCompile(collegeRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches{
		url := "http://www.hhkaoyan.com" + string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParseTeacher,
		})
	}
	return result
}
