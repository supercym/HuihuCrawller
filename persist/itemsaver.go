package persist

import (
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"huihuCrawler02/engine"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	//Must turn off sniff in docker
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 1
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)


			id := string(itemCount)
			item.Id = id
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
			itemCount++
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	//client.Index()就相当于存储和修改
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil

}
