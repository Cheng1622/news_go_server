package spider

import (
	"go_server/dao/mysql"
	"go_server/models"
	"go_server/utils"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
)

func getlist() {
	url := "https://m.yangzhu360.com/zhujia/chazhujia"
	b := utils.NewBug().SetHeader(headers).DisAutoLoacationClien().Get(url, nil)
	defer b.Close()
	if b.Err != nil {
		zap.L().Debug("err in list get", zap.Error(b.Err))
		return
	}
	dom, err := goquery.NewDocumentFromReader(b.Response.Body)
	if err != nil {
		zap.L().Debug("err in list get", zap.Error(err))
		return
	}

	s1 := dom.Find("div.look-con>table>thead>tr>td:nth-child(3)").Text()
	pattern := regexp.MustCompile(`(\d{4})/(\d{1,2})/(\d{1,2})`).FindStringSubmatch(s1)
	if len(pattern) != 4 {
		return
	}
	s2 := pattern[1] + pattern[2] + pattern[3]
	s3, _ := strconv.ParseInt(s2, 10, 64)
	content1, _ := dom.Find("div.look-tit").Html()
	content2, _ := dom.Find("div.bj-white").Html()
	content := `<div class="look-tit">` + content1 + `</div><div class="bj-white" id="dataDiv">` + content2 + `</div>`
	list := &models.List{
		Id:      s3,
		Content: content,
	}
	err = mysql.InsertList(list)
	if err != nil {
		return
	}
}
