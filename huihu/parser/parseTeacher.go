package huihu

import (
	"huihuCrawler/engine"
	"huihuCrawler/model"
	"regexp"
)

const nameRe = `<div class="data_input"><input type="text" name="name" value="([^"]*)" /></div>`
func ParseTeacher(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(nameRe)
	matches := re.FindAllSubmatch(contents, -1)

	if len(matches) == 0 {
		return engine.ParseResult{}
	}

	item := model.Profile{}
	item.Name = string(matches[0][1])
	item.Gender = string(matches[1][1])
	item.Age = string(matches[2][1])
	item.JobTitle = string(matches[3][1])
	item.Category = string(matches[4][1])
	item.College = string(matches[5][1])
	item.Email = string(matches[6][1])
	item.Phone = string(matches[7][1])
	item.Office = string(matches[8][1])
	result := engine.ParseResult{Items:[]interface{}{item}}
	return result
}
