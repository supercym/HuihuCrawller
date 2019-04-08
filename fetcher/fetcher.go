package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100*time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	log.Printf("Fetching url %s", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("NewRequest is err ", err)
		return nil, fmt.Errorf("NewRequest is err %v\n", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	//其实readAll()里面参数直接写resp.body就行，因为现在珍爱网也是utf8了，不需要再转码
	return ioutil.ReadAll(utf8Reader)
}

//确定爬取的页面内容的编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//如果直接读1024个byte用来判别编码，那么这1024个byte之后就被跳过了不能再读了
	//所以这里我们先用buffer存下来body里面前1024个byte
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	//这个函数读取resp.body前1024个字节，判断是什么编码并返回
	e, _, _ := charset.DetermineEncoding(bytes,"")
	return e
}
