package persist

import (
	"encoding/json"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"huihuCrawler02/engine"
	"huihuCrawler02/model"
	"testing"
)

func TestSave(t *testing.T) {
	payload := model.Profile{
		Name:		"Peter",
		Gender:		"M",
		Age:		"36",
		JobTitle:	"Professor",
		Category:	"PhD Tutor",
		College:	"Computer",
		Email:		"abc@bupt.edu.cn",
		Phone:		"13511223344",
		Office:		"Main Building 1007",
	}

	testData := engine.Item{
		Url:"https://www.hao123.com",
		Type:"teacher",
		Id:"3568854",
		Payload:payload,
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	//Save expected item
	const index = "testhuihu"
	err = Save(client, index, testData)
	if err != nil {
		panic(err)
	}


	//Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(testData.Type).
		Id(testData.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != testData {
		t.Errorf("Got %v, testData %v", actual, testData)
	}
}
