package main

import (
	"fmt"
	"regexp"
	"testing"
)

func Test(t *testing.T) {
	s1 := `
	<div class="zxxwleft">
		<div class="zxxw">
			<p class="zxxw1">湖南省生猪稳产保供划定年度目标 能繁母猪存栏356万头 生猪出栏6000万头</p>
			<b></b>来源：湖南日报 2023-04-02 11:10:55|  查看：<span style="float: none;" id="hits">2070</span>次 <p></p>
			<script type="text/javascript" src="//www.zhuwang.cc/api.php?op=count&amp;id=534888&amp;modelid=12"></script>
				<div style="position:relative; margin:10px 0; width:650px; height:80px;" class="gg-mark">
					<script src="//www.zhuwang.cc/api.php?op=ad_data&amp;id=77"></script>
				</div>
				<div style="font-size: 16px;font-family:'宋体'" class="show_content">
          		<div style="text-align: justify;"><span style="font-family:宋体;"><span style="font-size:16px;"><span style="line-height:2em;">　　湖南日报3月30日讯(全媒体记者张航)记者从近日召开的全省畜牧兽医工作会议上获悉,聚焦稳产保供,今年全省生猪产业划定目标任务:能繁母猪存栏356万头、规模<a href="http://hangqing.zhuwang.cc/jishu/zzjs/index.html"><u>养猪场</u></a>保有量1万个以上,生猪出栏6000万头左右。<br>
<br>
　　过去一年,湖南生猪产业提质升级,特色畜禽发展提速。全省出栏生猪6248万头,比上年增长2.1%,居全国第二、中部第一;出栏肉牛183万头、肉羊1101万头、家禽5.5亿羽,有效确保市场供给。<br>
<br>
　　<a href="http://www.zhuego.com/"><u>畜牧业</u></a>是我省农业农村经济的第一大产业。会议强调,要切实保障生猪等重要畜禽产品稳定安全供给,推动<a href="http://www.zhuego.com/"><u>畜牧业</u></a>持续健康发展,助力农业强省建设。依托大型龙头企业,推进家禽标准化养殖,提升肉禽、蛋禽现代化、智能化养殖水平;扩大<a href="http://www.zhuego.com/"><u>畜牧业</u></a>农机购置<a href="https://news.zhuwang.cc/butie/index.html"><u>补贴</u></a>,鼓励发展设施<a href="http://www.zhuego.com/"><u>畜牧业</u></a>,推进家禽牛羊集中屠宰,支持冷链配送体系建设,推进城乡肉品一体化供应。<br>
<br>
　　同时,对国家级、省级核心育种场、地方保种场和标准化示范场,全面开展监测评估;加大无疫小区和疫病净化场创建;用好全省动物卫生监督信息平台,全面实行动物和动物产品无纸化出证。加快绿色发展、强化行业监管,确保全省畜禽粪污综合利用率稳定在83%以上。健全病死动物无害化收集体系,升级改造收集运输设施,防止病原交叉感染。</span></span></span></div>
				</div>
          	<p></p>
			<link type="text/css" rel="stylesheet" href="https://www.zhuwang.cc/statics/css/content/iconfont.css">
          	<div class="pages" id="pages">
          					</div>
			
			<div style="border: 1px solid #e0e0e0;margin: 20px 0px;padding: 10px 20px;"><strong>【版权声明】</strong>养猪网旗下所有平台转载的文章均已注明来源、养猪网原创文章其他平台转载需注明来源且保持图文完整性、养猪网特别说明的文章未经允许不可转载，感谢您的支持与配合；我们所有刊登的文章仅供养猪人参考学习，不构成投资意见。若有不妥，请及时联系我们，可添加官方微信号“zgyangzhuwang”！</div>
			<div style="position:relative;margin:10px 0; " class="zxright1tp other-gg-mark">
				<script src="//www.zhuwang.cc/api.php?op=ad_data&amp;id=86"></script>
			</div>
			<div style="position:relative;margin:10px 0; padding-bottom:10px;"></div>
			<!--重要的qrcode-->
			<div class="clear"></div>
			<div class="yangzhub">
				<div class="yzapp">
					<a href="https://yzb.zhuwang.com.cn/" target="_blank">---养猪网手机APP端</a>
				</div>
			</div>
			<div class="share">
				<div class="bdsharebuttonbox bdshare-button-style0-16" data-bd-bind="1680442764920"><a href="#" class="bds_more" data-cmd="more"></a><a href="#" class="bds_qzone" data-cmd="qzone" title="分享到QQ空间"></a><a href="#" class="bds_tsina" data-cmd="tsina" title="分享到新浪微博"></a><a href="#" class="bds_tqq" data-cmd="tqq" title="分享到腾讯微博"></a><a href="#" class="bds_renren" data-cmd="renren" title="分享到人人网"></a><a href="#" class="bds_weixin" data-cmd="weixin" title="分享到微信"></a></div>
				<!--<div class="box">
					<div class="btn">分享</div>
					<div class="weixin commen css"><a href="javascript:;"><i class="iconfont">&#x3488;</i></a></div>
					<img id="weixinImg" src="" alt="页面二维码" />
					<div id="share_close">
						<img src="https://www.zhuwang.cc/statics/images/close.png" width="20" height="20" alt="关闭"/>
					</div>
				</div>-->
			</div>
			<iframe id="mood_frame" width="640" height="200" src="https://zhujia.zhuwang.com.cn/api1/mood/id/534888/catid/58" marginwidth="0" marginheight="0" frameborder="0" scrolling="no"></iframe>
		</div>
		<div class="zxread"><b>相关阅读</b>
<!-- 		相关标签 -->
               		<a href="//www.zhuwang.cc/index.php?m=content&amp;c=tag&amp;a=lists&amp;tag=%E6%B9%96%E5%8D%97%E7%9C%81" target="_blank">湖南省</a>
		       		<a href="//www.zhuwang.cc/index.php?m=content&amp;c=tag&amp;a=lists&amp;tag=%E7%94%9F%E7%8C%AA" target="_blank">生猪</a>
		       		<a href="//www.zhuwang.cc/index.php?m=content&amp;c=tag&amp;a=lists&amp;tag=%E7%A8%B3%E4%BA%A7%E4%BF%9D%E4%BE%9B" target="_blank">稳产保供</a>
		       		<a href="//www.zhuwang.cc/index.php?m=content&amp;c=tag&amp;a=lists&amp;tag=" target="_blank"></a>
		       		<a href="//www.zhuwang.cc/index.php?m=content&amp;c=tag&amp;a=lists&amp;tag=%E8%83%BD%E7%B9%81%E6%AF%8D%E7%8C%AA" target="_blank">能繁母猪</a>
		       		<a href="//www.zhuwang.cc/index.php?m=content&amp;c=tag&amp;a=lists&amp;tag=%E5%AD%98%E6%A0%8F" target="_blank">存栏</a>
				</div>
	<div class="zxread1" id="pinglun">
		<ul>
<!-- 		相关文章 -->
							<li><a href="https://news.zhuwang.cc/xinwen/20230402/534888.html" title="湖南省生猪稳产保供划定年度目标 能繁母猪存栏356万头 生猪出栏6000万头" target="_blank">湖南省生猪稳产保供划定年度目标 能繁母猪存栏356万头 生猪出栏6000万头</a><b>2023-04-02</b></li>
					<li><a href="https://news.zhuwang.cc/xinwen/20230401/534866.html" title="新希望大幅下修业绩预告，亏损或超15亿！此次大幅下修究竟是什么原因？" target="_blank">新希望大幅下修业绩预告，亏损或超15亿！此次大幅下修究竟是什么原因？</a><b>2023-04-01</b></li>
					<li><a href="https://news.zhuwang.cc/xinwen/20230331/534795.html" title="受销量及猪价下降影响，双汇发展2022年营收下滑6.16%，净利润增长15.51%" target="_blank">受销量及猪价下降影响，双汇发展2022年营收下滑6.16%，净利润增长15.51%</a><b>2023-03-31</b></li>
					<li><a href="https://news.zhuwang.cc/xinwen/20230331/534794.html" title="五部门联合发布:2月末能繁母猪存栏4343万头，下降0.6%" target="_blank">五部门联合发布:2月末能繁母猪存栏4343万头，下降0.6%</a><b>2023-03-31</b></li>
					<li><a href="https://news.zhuwang.cc/xinwen/20230331/534784.html" title="国内外猪价走弱，万洲国际业绩稳了？" target="_blank">国内外猪价走弱，万洲国际业绩稳了？</a><b>2023-03-31</b></li>
				 
		</ul>
	</div>
	<div class="zxread3">
	</div>
    </div>`
	pattern := regexp.MustCompile(`来源：(.+) (.+)\|`).FindStringSubmatch(s1)
	fmt.Println(pattern)
}
