package spider_test

import (
	"fmt"
	"go_server/models"
	"go_server/utils"
	"strconv"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
)

var (
	headers = map[string]string{

		`User-Agent`:                `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36`,
		`Accept`:                    `ext/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7`,
		`Accept-Language`:           `zh-CN,zh;q=0.9`,
		`Referer`:                   `https://www.zhuwang.cc/`,
		`Host`:                      `www.zhuwang.cc`,
		`Content-Type`:              `text/html; charset=utf-8`,
		`Origin`:                    `https://www.zhuwang.cc/`,
		`Connection`:                `keep-alive`,
		`Upgrade-Insecure-Requests`: `1`,
		`Sec-Fetch-Dest`:            `document`,
		`Sec-Fetch-Mode`:            `navigate`,
		`Sec-Fetch-Site`:            `same-origin`,
		`Sec-Fetch-User`:            `?1`,
		`Pragma`:                    `no-cache`,
		`Cache-Control`:             `no-cache`,
		`TE`:                        `trailers`,
	}
	tabArr = map[string]int{
		`xinwen`:   58,
		`shengzhu`: 63,
		`xinxi`:    90,
	}
)

func get(k string, v int, page int) {
	ts := make([]*models.Post, 0, 50)
	url := `https://zhuwang.cc/` + k + `/` + `list-` + strconv.Itoa(v) + `-` + strconv.Itoa(page) + `.html`
	b := utils.NewBug().SetHeader(headers).DisAutoLoacationClien().Get(url, nil)
	defer b.Close()
	if b.Err != nil {
		zap.L().Debug("err in spider get", zap.Error(b.Err))
		return
	}
	dom, err := goquery.NewDocumentFromReader(b.Response.Body)
	if err != nil {
		zap.L().Debug("err in spider get", zap.Error(err))
		return
	}

	dom.Find("div.zxleft>div.zxleft3>ul>li").First().Remove()
	dom.Find("div.zxleft>div.zxleft3>ul>li").Last().Remove()
	dom.Find("div.zxleft>div.zxleft3>ul>li").Each(func(_ int, s *goquery.Selection) {
		Title := s.Find("p.zxleft31>a").Text()
		NewsUrl, _ := s.Find("a").Attr("href")
		post := &models.Post{
			// Isnews:  int32(1),
			Title:   Title,
			NewsUrl: NewsUrl,
		}
		ts = append(ts, post)
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")

	})
	getDeal(ts)
}

func getDeal(ts []*models.Post) {

	fmt.Println("---")
	if len(ts) == 0 {
		return
	}
	for k, v := range ts {
		fmt.Println(k, v.NewsUrl)
	}
	fmt.Println()

}

func Test(t *testing.T) {
	for k, v := range tabArr {
		get(k, v, 1)
	}
}
