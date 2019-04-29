package huihu

import (
	"huihuCrawler02/engine"
	"huihuCrawler02/model"
	"regexp"
)

const nameRe = `<div class="data_input"><input type="text" name="name" value="([^"]*)" /></div>`
func ParseTeacher(contents []byte, url string) engine.ParseResult {
	re := regexp.MustCompile(nameRe)
	matches := re.FindAllSubmatch(contents, -1)

	if len(matches) == 0 {
		return engine.ParseResult{}
	}

	profile := model.Profile{}
	profile.Name = string(matches[0][1])
	profile.Gender = string(matches[1][1])
	profile.Age = string(matches[2][1])
	profile.JobTitle = string(matches[3][1])
	profile.Category = string(matches[4][1])
	profile.College = string(matches[5][1])
	profile.Email = string(matches[6][1])
	profile.Phone = string(matches[7][1])
	profile.Office = string(matches[8][1])
	result := engine.ParseResult{
		Items:[]engine.Item{
			{
				Url:     url,
				Type:    "teacher",
				Id:      "",
				Payload: profile,
			},
		},
	}
	return result
}
