package spider

import (
	"go_server/dao/mysql"
	"go_server/dao/redis"
	"go_server/models"
	"go_server/utils"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

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
		`xinwen`:               58,  //国内新闻
		`qihuo`:                263, //生猪期货
		`guojixinwen`:          147, //国际新闻
		`xingyedianping`:       148, //行业点评
		`dujiafenxi`:           149, //原创分析
		`zhuping`:              70,  //每日猪评
		`baodao`:               118, //展会报道
		`zhongzhuzixun`:        170, //种猪资讯
		`zhongzhuxingyexinwen`: 166, //种猪行业新闻
		`zhongzhuqiye`:         143, //种猪企业
		`zhongzhujishu`:        173, //种猪企业访谈
		`mingqituijianzz`:      221, //名企推荐
		`zhuchangjs`:           31,  //猪场建设
		`shoujing`:             32,  //繁育管理
		`siliaoyy`:             91,  //饲养管理
		`kxyangzhu`:            35,  //猪场管理
		`pxhshenchan`:          233, //批次化生产
		`liman`:                261, //养猪大会
		`hangqingfenxi`:        81,  //行情分析
		`yumi`:                 68,  //玉米价格
		`doupo`:                67,  //豆粕价格
		`zhuliangbi`:           257, //猪粮比
		`siliaogongxu`:         256, //饲料供需
		`siliaofenxi`:          267, //饲料分析
		`shengzhu`:             63,  //生猪价格
		`zizhu`:                64,  //仔猪价格
		`zhurou`:               65,  //猪肉价格
		`shengshi`:             115, //各省市猪价
		`xinxi`:                90,  //养猪新闻
		`jishushipin`:          88,  //养猪技术
		`meirizhujia`:          260, //每日猪价
		`jiangzuo`:             92,  //专家讲座
		`fangtanshipin`:        93,  //人物访谈
		`qiye`:                 94,  //企业展示
		`qitashipin`:           111, //养猪致富
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
		Isnews := int32(1)
		CommunityId := int64(v)
		Title := s.Find("p.zxleft31>a").Text()
		NewsUrl, _ := s.Find("a").Attr("href")
		pattern := regexp.MustCompile(`https://[\s\S]*/(.+)/(.+).html`).FindStringSubmatch(NewsUrl)
		k := ""
		if len(pattern) == 3 {
			k = strconv.Itoa(v) + pattern[1] + pattern[2]
		}
		re := regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
		result := re.FindAllStringSubmatch(NewsUrl, -1)
		if result == nil {
			NewsUrl = "https:" + NewsUrl
			pattern := regexp.MustCompile(`https://[\s\S]*-(.+)-(.+)-(.+).html`).FindStringSubmatch(NewsUrl)
			if len(pattern) == 4 {
				k = pattern[1] + pattern[2] + pattern[3]
			}
		}
		if NewsUrl == "https:" {
			return
		}
		Id, _ := strconv.ParseInt(k, 10, 64)
		post := &models.Post{
			CommunityId: CommunityId,
			Isnews:      Isnews,
			Title:       Title,
			NewsUrl:     NewsUrl,
			Id:          Id,
		}
		ts = append(ts, post)
	})
	// 遍历指针切片
	for _, v := range ts {
		getDeal(v)
	}
}

func getDeal(s *models.Post) {
	c := utils.NewBug().SetHeader(headers).DisAutoLoacationClien().Get(s.NewsUrl, nil)
	defer c.Close()
	if c.Err != nil {
		zap.L().Debug("err in spider getDeal", zap.Error(c.Err))
		return
	}
	dom, err := goquery.NewDocumentFromReader(c.Response.Body)
	if err != nil {
		zap.L().Debug("err in spider getDeal", zap.Error(err))
		return
	}
	s1 := dom.Find("div.zxxwleft>div.zxxw").Text()
	pattern := regexp.MustCompile(`来源：(.+) (\d{4}-\d{1,2}-\d{1,2} \d{1,2}:\d{1,2}:\d{1,2})`).FindStringSubmatch(s1)
	if len(pattern) == 3 {
		s.NewsSource = pattern[1]
		s.NewsTime = pattern[2]
	}
	if s.NewsSource == "" && s.NewsTime == "" {
		s.NewsSource = "腾讯视频"
		s.NewsTime = dom.Find("div.articleft> div.contant >p.cishu> span:nth-child(2)").Text()
	}
	dom.Find("div.show_content>style").Remove().Html()
	dom.Find("div.show_content>table>tbody>tr:nth-child(1)").Remove().Html()
	dom.Find("div.show_content>table>tbody").RemoveAttr("style")
	s.Content, _ = dom.Find("div.show_content").Html()
	// 图片处理
	r1 := regexp.MustCompile(`https://(.*?).(jpg|png|jpeg)`)
	image := r1.FindAllString(s.Content, -1)
	if len(image) > 0 && len(image) < 3 {
		s.Image1 = image[0]
		s.Isimage = 1
	}
	if len(image) >= 3 {
		s.Image1 = image[0]
		s.Image2 = image[1]
		s.Image3 = image[2]
		s.Isimage3 = 1
	}
	// 视频处理
	videourltx, ok := dom.Find("div.articleft>div.contant>iframe").Attr("src")
	if ok {
		s.Isvideo = 1
		r2 := regexp.MustCompile(`vid=(.*)`)
		videourlnum := r2.FindString(videourltx)
		videotihuan := strings.Replace(videourlnum, "vid=", "", -1)
		s.Videoimage = "https://puui.qpic.cn/vpic_cover/" + videotihuan + "/" + videotihuan + "_hz.jpg"
		videourltxsp := "http://vv.video.qq.com/getinfo?vids=" + videotihuan + "&platform=101001&charge=0&otype=json"
		txsp := utils.NewBug().SetHeader(headers).DisAutoLoacationClien().Get(videourltxsp, nil).Body()
		txspt := regexp.MustCompile(`"fvkey":"(.*?)".*?"fn":"(.*?.mp4)".*?"url":"(.*?)"`)
		videourl := txspt.FindStringSubmatch(string(txsp))

		if videourl != nil {
			videozl := videourl[3] + videourl[2] + "?vkey=" + videourl[1]
			s.Video, _ = downLoad(videozl, s.Id)
		}
	}
	// 数据库增加记录
	err = mysql.InsertPost(s)
	if err != nil {
		return
	}
	// 去点赞数量的 zset 新增一条记录
	err = redis.AddPost(s.Id)
	if err != nil {
		return
	}
}

func downLoad(url string, s int64) (string, error) {
	b, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer b.Body.Close()
	k := strconv.FormatInt(s, 10) + ".mp4"
	path := "./video/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// mkdir 创建目录，mkdirAll 可创建多层级目录
		os.MkdirAll(path, os.ModePerm)
	}
	out, err := os.Create(path + k)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, b.Body)
	return k, err

}
