package huihu

import (
	"huihuCrawler02/engine"
	"regexp"
)

const collegeListRe = `<a href="(/teacher/major.html[?]uid=[0-9a-z]+)">([^<]+)</a>`
func ParseCollegeList(contents []byte, url string) engine.ParseResult {
	re := regexp.MustCompile(collegeListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := "http://www.hhkaoyan.com" + string(m[1])
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        url,
				ParserFunc: ParseCollege,
			})
	}
	return result
}
