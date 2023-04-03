package spider_test

// import (
// 	"fmt"
// 	"go_server/models"
// 	"go_server/utils"
// 	"strconv"
// 	"testing"

// 	"github.com/PuerkitoBio/goquery"
// 	"go.uber.org/zap"
// )

// var (
// 	headers = map[string]string{

// 		`User-Agent`:                `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36`,
// 		`Accept`:                    `ext/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7`,
// 		`Accept-Language`:           `zh-CN,zh;q=0.9`,
// 		`Referer`:                   `https://www.zhuwang.cc/`,
// 		`Host`:                      `www.zhuwang.cc`,
// 		`Content-Type`:              `text/html; charset=utf-8`,
// 		`Origin`:                    `https://www.zhuwang.cc/`,
// 		`Connection`:                `keep-alive`,
// 		`Upgrade-Insecure-Requests`: `1`,
// 		`Sec-Fetch-Dest`:            `document`,
// 		`Sec-Fetch-Mode`:            `navigate`,
// 		`Sec-Fetch-Site`:            `same-origin`,
// 		`Sec-Fetch-User`:            `?1`,
// 		`Pragma`:                    `no-cache`,
// 		`Cache-Control`:             `no-cache`,
// 		`TE`:                        `trailers`,
// 	}
// 	tabArr = map[string]int{
// 		`xinwen`:   58,
// 		`shengzhu`: 63,
// 		`xinxi`:    90,
// 	}
// )

// func get(k string, v int, page int) {
// 	ts := make([]*models.Post, 0, 50)
// 	url := `https://zhuwang.cc/` + k + `/` + `list-` + strconv.Itoa(v) + `-` + strconv.Itoa(page) + `.html`
// 	b := utils.NewBug().SetHeader(headers).DisAutoLoacationClien().Get(url, nil)
// 	defer b.Close()
// 	if b.Err != nil {
// 		zap.L().Debug("err in spider get", zap.Error(b.Err))
// 		return
// 	}
// 	dom, err := goquery.NewDocumentFromReader(b.Response.Body)
// 	if err != nil {
// 		zap.L().Debug("err in spider get", zap.Error(err))
// 		return
// 	}

// 	dom.Find("div.zxleft>div.zxleft3>ul>li").First().Remove()
// 	dom.Find("div.zxleft>div.zxleft3>ul>li").Last().Remove()
// 	dom.Find("div.zxleft>div.zxleft3>ul>li").Each(func(_ int, s *goquery.Selection) {
// 		Title := s.Find("p.zxleft31>a").Text()
// 		NewsUrl, _ := s.Find("a").Attr("href")
// 		post := &models.Post{
// 			// Isnews:  int32(1),
// 			Title:   Title,
// 			NewsUrl: NewsUrl,
// 		}
// 		ts = append(ts, post)
// 		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")

// 	})
// 	getDeal(ts)
// }

// func getDeal(ts []*models.Post) {

// 	fmt.Println("---")
// 	if len(ts) == 0 {
// 		return
// 	}
// 	for k, v := range ts {
// 		fmt.Println(k, v.NewsUrl)
// 	}
// 	fmt.Println()

// }

// func Test(t *testing.T) {
// 	for k, v := range tabArr {
// 		get(k, v, 1)
// 	}
// }

import (
	"fmt"
	"regexp"
	"testing"
)

func Test(t *testing.T) {
	s1 := `<div class="bj-white" id="dataDiv">
    <!--生猪(外三元)-->
<div class="look-con">
    <table border="0" cellpadding="0" cellspacing="0">
        <thead>
        <tr>
            <td style="width:20%;">排名</td>
            <td>地区</td>
            <td>2023/04/03价格</td>
            <td>今日跌涨</td>
        </tr>
        </thead>
        <tbody>
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                                
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                <tr data-ud="0.30">
            <td><em>1</em></td>
            <td>北京市</td>
            <td>14.82元/公斤</td>
            <td class="red"><span class="jt-up"></span>0.30</td>
        </tr><tr data-ud="0.07">
            <td><em>2</em></td>
            <td>宁夏</td>
            <td>14.69元/公斤</td>
            <td class="red"><span class="jt-up"></span>0.07</td>
        </tr><tr data-ud="0.07">
            <td><em>3</em></td>
            <td>海南省</td>
            <td>14.00元/公斤</td>
            <td class="red"><span class="jt-up"></span>0.07</td>
        </tr><tr data-ud="0.05">
            <td><em>4</em></td>
            <td>浙江省</td>
            <td>15.52元/公斤</td>
            <td class="red"><span class="jt-up"></span>0.05</td>
        </tr><tr data-ud="0.04">
            <td><em>5</em></td>
            <td>陕西省</td>
            <td>13.96元/公斤</td>
            <td class="red"><span class="jt-up"></span>0.04</td>
        </tr><tr data-ud="0.04">
            <td><em>6</em></td>
            <td>新疆</td>
            <td>13.49元/公斤</td>
            <td class="red"><span class="jt-up"></span>0.04</td>
        </tr><tr data-ud="0.02">
            <td><em>7</em></td>
            <td>黑龙江省</td>
            <td>13.96元/公斤</td>
            <td class="red"><span class="jt-up"></span>0.02</td>
        </tr><tr data-ud="0.01">
            <td><em>8</em></td>
            <td>云南省</td>
            <td>14.29元/公斤</td>
            <td class="red"><span class="jt-up"></span>0.01</td>
        </tr><tr data-ud="0.01">
            <td><em>9</em></td>
            <td>甘肃省</td>
            <td>14.20元/公斤</td>
            <td class="red"><span class="jt-up"></span>0.01</td>
        </tr><tr data-ud="0.00">
            <td><em>10</em></td>
            <td>福建省</td>
            <td>15.54元/公斤</td>
            <td class="green"><span class="jt-down"></span>0.00</td>
        </tr><tr data-ud="0.00">
            <td><em>11</em></td>
            <td>江西省</td>
            <td>14.74元/公斤</td>
            <td class="green"><span class="jt-down"></span>0.00</td>
        </tr><tr data-ud="0.00">
            <td><em>12</em></td>
            <td>湖北省</td>
            <td>14.63元/公斤</td>
            <td class="green"><span class="jt-down"></span>0.00</td>
        </tr><tr data-ud="0.00">
            <td><em>13</em></td>
            <td>山东省</td>
            <td>14.57元/公斤</td>
            <td class="green"><span class="jt-down"></span>0.00</td>
        </tr><tr data-ud="0.00">
            <td><em>14</em></td>
            <td>贵州省</td>
            <td>14.48元/公斤</td>
            <td class="green"><span class="jt-down"></span>0.00</td>
        </tr><tr data-ud="0.00">
            <td><em>15</em></td>
            <td>山西省</td>
            <td>13.94元/公斤</td>
            <td class="green"><span class="jt-down"></span>0.00</td>
        </tr><tr data-ud="-0.01">
            <td><em>16</em></td>
            <td>安徽省</td>
            <td>14.73元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.01</td>
        </tr><tr data-ud="-0.01">
            <td><em>17</em></td>
            <td>四川省</td>
            <td>14.19元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.01</td>
        </tr><tr data-ud="-0.02">
            <td><em>18</em></td>
            <td>青海省</td>
            <td>14.95元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.02</td>
        </tr><tr data-ud="-0.02">
            <td><em>19</em></td>
            <td>西藏</td>
            <td>14.95元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.02</td>
        </tr><tr data-ud="-0.02">
            <td><em>20</em></td>
            <td>河南省</td>
            <td>14.37元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.02</td>
        </tr><tr data-ud="-0.02">
            <td><em>21</em></td>
            <td>河北省</td>
            <td>14.35元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.02</td>
        </tr><tr data-ud="-0.03">
            <td><em>22</em></td>
            <td>重庆市</td>
            <td>14.72元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.03</td>
        </tr><tr data-ud="-0.04">
            <td><em>23</em></td>
            <td>上海市</td>
            <td>15.62元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.04</td>
        </tr><tr data-ud="-0.04">
            <td><em>24</em></td>
            <td>湖南省</td>
            <td>14.77元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.04</td>
        </tr><tr data-ud="-0.05">
            <td><em>25</em></td>
            <td>辽宁省</td>
            <td>13.91元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.05</td>
        </tr><tr data-ud="-0.06">
            <td><em>26</em></td>
            <td>天津市</td>
            <td>14.61元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.06</td>
        </tr><tr data-ud="-0.06">
            <td><em>27</em></td>
            <td>吉林省</td>
            <td>13.99元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.06</td>
        </tr><tr data-ud="-0.11">
            <td><em>28</em></td>
            <td>广东省</td>
            <td>15.43元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.11</td>
        </tr><tr data-ud="-0.16">
            <td><em>29</em></td>
            <td>江苏省</td>
            <td>15.16元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.16</td>
        </tr><tr data-ud="-0.18">
            <td><em>30</em></td>
            <td>广西</td>
            <td>14.61元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.18</td>
        </tr><tr data-ud="-0.39">
            <td><em>31</em></td>
            <td>内蒙古</td>
            <td>13.80元/公斤</td>
            <td class="green"><span class="jt-down"></span>-0.39</td>
        </tr></tbody>
    </table>
</div>
<!--玉米(14%水分)-->
<div style="display: none" class="look-con">
    <table border="0" cellpadding="0" cellspacing="0">
        <thead>
        <tr>
            <td style="width:20%;">排名</td>
            <td>地区</td>
            <td>2023/04/03价格</td>
            <td>今日跌涨</td>
        </tr>
        </thead>
        <tbody>
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
                        
         
                <tr data-ud="375">
            <td><em>1</em></td>
            <td>新疆</td>
            <td>5200元/吨</td>
            <td class="red"><span class="jt-up"></span>375</td>
        </tr><tr data-ud="105">
            <td><em>2</em></td>
            <td>黑龙江省</td>
            <td>4518元/吨</td>
            <td class="red"><span class="jt-up"></span>105</td>
        </tr><tr data-ud="80">
            <td><em>3</em></td>
            <td>吉林省</td>
            <td>4587元/吨</td>
            <td class="red"><span class="jt-up"></span>80</td>
        </tr><tr data-ud="52">
            <td><em>4</em></td>
            <td>河南省</td>
            <td>4169元/吨</td>
            <td class="red"><span class="jt-up"></span>52</td>
        </tr><tr data-ud="47">
            <td><em>5</em></td>
            <td>山西省</td>
            <td>4087元/吨</td>
            <td class="red"><span class="jt-up"></span>47</td>
        </tr><tr data-ud="25">
            <td><em>6</em></td>
            <td>河北省</td>
            <td>4089元/吨</td>
            <td class="red"><span class="jt-up"></span>25</td>
        </tr><tr data-ud="21">
            <td><em>7</em></td>
            <td>安徽省</td>
            <td>3998元/吨</td>
            <td class="red"><span class="jt-up"></span>21</td>
        </tr><tr data-ud="16">
            <td><em>8</em></td>
            <td>四川省</td>
            <td>4059元/吨</td>
            <td class="red"><span class="jt-up"></span>16</td>
        </tr><tr data-ud="15">
            <td><em>9</em></td>
            <td>陕西省</td>
            <td>4215元/吨</td>
            <td class="red"><span class="jt-up"></span>15</td>
        </tr><tr data-ud="10">
            <td><em>10</em></td>
            <td>辽宁省</td>
            <td>4010元/吨</td>
            <td class="red"><span class="jt-up"></span>10</td>
        </tr><tr data-ud="4">
            <td><em>11</em></td>
            <td>贵州省</td>
            <td>4643元/吨</td>
            <td class="red"><span class="jt-up"></span>4</td>
        </tr><tr data-ud="2">
            <td><em>12</em></td>
            <td>广西</td>
            <td>4000元/吨</td>
            <td class="red"><span class="jt-up"></span>2</td>
        </tr><tr data-ud="0">
            <td><em>13</em></td>
            <td>北京市</td>
            <td>5000元/吨</td>
            <td class="green"><span class="jt-down"></span>0</td>
        </tr><tr data-ud="-8">
            <td><em>14</em></td>
            <td>浙江省</td>
            <td>4360元/吨</td>
            <td class="green"><span class="jt-down"></span>-8</td>
        </tr><tr data-ud="-11">
            <td><em>15</em></td>
            <td>海南省</td>
            <td>4627元/吨</td>
            <td class="green"><span class="jt-down"></span>-11</td>
        </tr><tr data-ud="-14">
            <td><em>16</em></td>
            <td>甘肃省</td>
            <td>4580元/吨</td>
            <td class="green"><span class="jt-down"></span>-14</td>
        </tr><tr data-ud="-15">
            <td><em>17</em></td>
            <td>江西省</td>
            <td>4809元/吨</td>
            <td class="green"><span class="jt-down"></span>-15</td>
        </tr><tr data-ud="-23">
            <td><em>18</em></td>
            <td>青海省</td>
            <td>4028元/吨</td>
            <td class="green"><span class="jt-down"></span>-23</td>
        </tr><tr data-ud="-23">
            <td><em>19</em></td>
            <td>西藏</td>
            <td>4027元/吨</td>
            <td class="green"><span class="jt-down"></span>-23</td>
        </tr><tr data-ud="-33">
            <td><em>20</em></td>
            <td>宁夏</td>
            <td>4058元/吨</td>
            <td class="green"><span class="jt-down"></span>-33</td>
        </tr><tr data-ud="-55">
            <td><em>21</em></td>
            <td>云南省</td>
            <td>4585元/吨</td>
            <td class="green"><span class="jt-down"></span>-55</td>
        </tr><tr data-ud="-79">
            <td><em>22</em></td>
            <td>山东省</td>
            <td>3985元/吨</td>
            <td class="green"><span class="jt-down"></span>-79</td>
        </tr><tr data-ud="-86">
            <td><em>23</em></td>
            <td>广东省</td>
            <td>3757元/吨</td>
            <td class="green"><span class="jt-down"></span>-86</td>
        </tr><tr data-ud="-89">
            <td><em>24</em></td>
            <td>湖南省</td>
            <td>4336元/吨</td>
            <td class="green"><span class="jt-down"></span>-89</td>
        </tr><tr data-ud="-110">
            <td><em>25</em></td>
            <td>江苏省</td>
            <td>4295元/吨</td>
            <td class="green"><span class="jt-down"></span>-110</td>
        </tr><tr data-ud="-187">
            <td><em>26</em></td>
            <td>上海市</td>
            <td>4520元/吨</td>
            <td class="green"><span class="jt-down"></span>-187</td>
        </tr><tr data-ud="-237">
            <td><em>27</em></td>
            <td>福建省</td>
            <td>4400元/吨</td>
            <td class="green"><span class="jt-down"></span>-237</td>
        </tr><tr data-ud="-241">
            <td><em>28</em></td>
            <td>湖北省</td>
            <td>4379元/吨</td>
            <td class="green"><span class="jt-down"></span>-241</td>
        </tr><tr data-ud="-303">
            <td><em>29</em></td>
            <td>天津市</td>
            <td>4767元/吨</td>
            <td class="green"><span class="jt-down"></span>-303</td>
        </tr><tr data-ud="-429">
            <td><em>30</em></td>
            <td>内蒙古</td>
            <td>3975元/吨</td>
            <td class="green"><span class="jt-down"></span>-429</td>
        </tr><tr data-ud="-532">
            <td><em>31</em></td>
            <td>重庆市</td>
            <td>4100元/吨</td>
            <td class="green"><span class="jt-down"></span>-532</td>
        </tr></tbody>
    </table>
</div></div>`
	pattern := regexp.MustCompile(`(\d{4})/(\d{1,2})/(\d{1,2})`).FindStringSubmatch(s1)
	// s3, _ := strconv.ParseInt(s2, 10, 64)
	fmt.Println(pattern, len(pattern)) //2019-07-31
}
