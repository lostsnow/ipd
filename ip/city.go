package ip

type City struct {
	Title string
	Code  string
}

func init() {
	for k, c := range allCitiesCN {
		if cityCodes[k] == nil {
			cityCodes[k] = map[string]City{}
		}
		for _, v := range c {
			cityCodes[k][v.Title] = v
		}
	}
}

var (
	cityCodes = map[string]map[string]City{}

	// <editor-fold desc="all cities">
	CitiesCN13 = []City{
		{Code: "CN130100", Title: "石家庄"},
		{Code: "CN130200", Title: "唐山"},
		{Code: "CN130300", Title: "秦皇岛"},
		{Code: "CN130400", Title: "邯郸"},
		{Code: "CN130500", Title: "邢台"},
		{Code: "CN130600", Title: "保定"},
		{Code: "CN130700", Title: "张家口"},
		{Code: "CN130800", Title: "承德"},
		{Code: "CN130900", Title: "沧州"},
		{Code: "CN131000", Title: "廊坊"},
		{Code: "CN131100", Title: "衡水"},
	}

	CitiesCN14 = []City{
		{Code: "CN140100", Title: "太原"},
		{Code: "CN140200", Title: "大同"},
		{Code: "CN140300", Title: "阳泉"},
		{Code: "CN140400", Title: "长治"},
		{Code: "CN140500", Title: "晋城"},
		{Code: "CN140600", Title: "朔州"},
		{Code: "CN140700", Title: "晋中"},
		{Code: "CN140800", Title: "运城"},
		{Code: "CN140900", Title: "忻州"},
		{Code: "CN141000", Title: "临汾"},
		{Code: "CN141100", Title: "吕梁"},
	}

	CitiesCN15 = []City{
		{Code: "CN150100", Title: "呼和浩特"},
		{Code: "CN150200", Title: "包头"},
		{Code: "CN150300", Title: "乌海"},
		{Code: "CN150400", Title: "赤峰"},
		{Code: "CN150500", Title: "通辽"},
		{Code: "CN150600", Title: "鄂尔多斯"},
		{Code: "CN150700", Title: "呼伦贝尔"},
		{Code: "CN150800", Title: "巴彦淖尔"},
		{Code: "CN150900", Title: "乌兰察布"},
		{Code: "CN152200", Title: "兴安盟"},
		{Code: "CN152201", Title: "乌兰浩特"},
		{Code: "CN152500", Title: "锡林郭勒盟"},
		{Code: "CN152900", Title: "阿拉善盟"},
	}

	CitiesCN21 = []City{
		{Code: "CN210100", Title: "沈阳"},
		{Code: "CN210200", Title: "大连"},
		{Code: "CN210300", Title: "鞍山"},
		{Code: "CN210400", Title: "抚顺"},
		{Code: "CN210500", Title: "本溪"},
		{Code: "CN210600", Title: "丹东"},
		{Code: "CN210700", Title: "锦州"},
		{Code: "CN210800", Title: "营口"},
		{Code: "CN210900", Title: "阜新"},
		{Code: "CN211000", Title: "辽阳"},
		{Code: "CN211100", Title: "盘锦"},
		{Code: "CN211200", Title: "铁岭"},
		{Code: "CN211300", Title: "朝阳"},
		{Code: "CN211400", Title: "葫芦岛"},
	}

	CitiesCN22 = []City{
		{Code: "CN220100", Title: "长春"},
		{Code: "CN220200", Title: "吉林市"},
		{Code: "CN220300", Title: "四平"},
		{Code: "CN220400", Title: "辽源"},
		{Code: "CN220500", Title: "通化"},
		{Code: "CN220600", Title: "白山"},
		{Code: "CN220700", Title: "松原"},
		{Code: "CN220800", Title: "白城"},
		{Code: "CN222400", Title: "延边朝鲜族自治州"},
	}

	CitiesCN23 = []City{
		{Code: "CN230100", Title: "哈尔滨"},
		{Code: "CN230200", Title: "齐齐哈尔"},
		{Code: "CN230300", Title: "鸡西"},
		{Code: "CN230400", Title: "鹤岗"},
		{Code: "CN230500", Title: "双鸭山"},
		{Code: "CN230600", Title: "大庆"},
		{Code: "CN230700", Title: "伊春"},
		{Code: "CN230800", Title: "佳木斯"},
		{Code: "CN230900", Title: "七台河"},
		{Code: "CN231000", Title: "牡丹江"},
		{Code: "CN231100", Title: "黑河"},
		{Code: "CN231200", Title: "绥化"},
		{Code: "CN232700", Title: "大兴安岭地区"},
	}

	CitiesCN32 = []City{
		{Code: "CN320100", Title: "南京"},
		{Code: "CN320200", Title: "无锡"},
		{Code: "CN320300", Title: "徐州"},
		{Code: "CN320400", Title: "常州"},
		{Code: "CN320500", Title: "苏州"},
		{Code: "CN320600", Title: "南通"},
		{Code: "CN320700", Title: "连云港"},
		{Code: "CN320800", Title: "淮安"},
		{Code: "CN320900", Title: "盐城"},
		{Code: "CN321000", Title: "扬州"},
		{Code: "CN321100", Title: "镇江"},
		{Code: "CN321200", Title: "泰州"},
		{Code: "CN321300", Title: "宿迁"},
	}

	CitiesCN33 = []City{
		{Code: "CN330100", Title: "杭州"},
		{Code: "CN330200", Title: "宁波"},
		{Code: "CN330300", Title: "温州"},
		{Code: "CN330400", Title: "嘉兴"},
		{Code: "CN330500", Title: "湖州"},
		{Code: "CN330600", Title: "绍兴"},
		{Code: "CN330700", Title: "金华"},
		{Code: "CN330800", Title: "衢州"},
		{Code: "CN330900", Title: "舟山"},
		{Code: "CN331000", Title: "台州"},
		{Code: "CN331100", Title: "丽水"},
	}

	CitiesCN34 = []City{
		{Code: "CN340100", Title: "合肥"},
		{Code: "CN340200", Title: "芜湖"},
		{Code: "CN340300", Title: "蚌埠"},
		{Code: "CN340400", Title: "淮南"},
		{Code: "CN340500", Title: "马鞍山"},
		{Code: "CN340600", Title: "淮北"},
		{Code: "CN340700", Title: "铜陵"},
		{Code: "CN340800", Title: "安庆"},
		{Code: "CN341000", Title: "黄山"},
		{Code: "CN341100", Title: "滁州"},
		{Code: "CN341200", Title: "阜阳"},
		{Code: "CN341300", Title: "宿州"},
		{Code: "CN341500", Title: "六安"},
		{Code: "CN341600", Title: "亳州"},
		{Code: "CN341700", Title: "池州"},
		{Code: "CN341800", Title: "宣城"},
	}

	CitiesCN35 = []City{
		{Code: "CN350100", Title: "福州"},
		{Code: "CN350200", Title: "厦门"},
		{Code: "CN350300", Title: "莆田"},
		{Code: "CN350400", Title: "三明"},
		{Code: "CN350500", Title: "泉州"},
		{Code: "CN350600", Title: "漳州"},
		{Code: "CN350700", Title: "南平"},
		{Code: "CN350800", Title: "龙岩"},
		{Code: "CN350900", Title: "宁德"},
	}

	CitiesCN36 = []City{
		{Code: "CN360100", Title: "南昌"},
		{Code: "CN360200", Title: "景德镇"},
		{Code: "CN360300", Title: "萍乡"},
		{Code: "CN360400", Title: "九江"},
		{Code: "CN360500", Title: "新余"},
		{Code: "CN360600", Title: "鹰潭"},
		{Code: "CN360700", Title: "赣州"},
		{Code: "CN360800", Title: "吉安"},
		{Code: "CN360900", Title: "宜春"},
		{Code: "CN361000", Title: "抚州"},
		{Code: "CN361100", Title: "上饶"},
	}

	CitiesCN37 = []City{
		{Code: "CN370100", Title: "济南"},
		{Code: "CN370200", Title: "青岛"},
		{Code: "CN370300", Title: "淄博"},
		{Code: "CN370400", Title: "枣庄"},
		{Code: "CN370500", Title: "东营"},
		{Code: "CN370600", Title: "烟台"},
		{Code: "CN370700", Title: "潍坊"},
		{Code: "CN370800", Title: "济宁"},
		{Code: "CN370900", Title: "泰安"},
		{Code: "CN371000", Title: "威海"},
		{Code: "CN371100", Title: "日照"},
		{Code: "CN371200", Title: "莱芜"},
		{Code: "CN371300", Title: "临沂"},
		{Code: "CN371400", Title: "德州"},
		{Code: "CN371500", Title: "聊城"},
		{Code: "CN371600", Title: "滨州"},
		{Code: "CN371700", Title: "菏泽"},
	}

	CitiesCN41 = []City{
		{Code: "CN410100", Title: "郑州"},
		{Code: "CN410200", Title: "开封"},
		{Code: "CN410300", Title: "洛阳"},
		{Code: "CN410400", Title: "平顶山"},
		{Code: "CN410500", Title: "安阳"},
		{Code: "CN410600", Title: "鹤壁"},
		{Code: "CN410700", Title: "新乡"},
		{Code: "CN410800", Title: "焦作"},
		{Code: "CN410900", Title: "濮阳"},
		{Code: "CN411000", Title: "许昌"},
		{Code: "CN411100", Title: "漯河"},
		{Code: "CN411200", Title: "三门峡"},
		{Code: "CN411300", Title: "南阳"},
		{Code: "CN411400", Title: "商丘"},
		{Code: "CN411500", Title: "信阳"},
		{Code: "CN411600", Title: "周口"},
		{Code: "CN411700", Title: "驻马店"},
		{Code: "CN419001", Title: "济源"},
	}

	CitiesCN42 = []City{
		{Code: "CN420100", Title: "武汉"},
		{Code: "CN420200", Title: "黄石"},
		{Code: "CN420300", Title: "十堰"},
		{Code: "CN420500", Title: "宜昌"},
		{Code: "CN420600", Title: "襄阳"},
		{Code: "CN420700", Title: "鄂州"},
		{Code: "CN420800", Title: "荆门"},
		{Code: "CN420900", Title: "孝感"},
		{Code: "CN421000", Title: "荆州"},
		{Code: "CN421100", Title: "黄冈"},
		{Code: "CN421200", Title: "咸宁"},
		{Code: "CN421300", Title: "随州"},
		{Code: "CN422800", Title: "恩施土家族苗族自治州"},
		{Code: "CN429004", Title: "仙桃"},
		{Code: "CN429005", Title: "潜江"},
		{Code: "CN429006", Title: "天门"},
		{Code: "CN429021", Title: "神农架林区"},
	}

	CitiesCN43 = []City{
		{Code: "CN430100", Title: "长沙"},
		{Code: "CN430200", Title: "株洲"},
		{Code: "CN430300", Title: "湘潭"},
		{Code: "CN430400", Title: "衡阳"},
		{Code: "CN430500", Title: "邵阳"},
		{Code: "CN430600", Title: "岳阳"},
		{Code: "CN430700", Title: "常德"},
		{Code: "CN430800", Title: "张家界"},
		{Code: "CN430900", Title: "益阳"},
		{Code: "CN431000", Title: "郴州"},
		{Code: "CN431100", Title: "永州"},
		{Code: "CN431200", Title: "怀化"},
		{Code: "CN431300", Title: "娄底"},
		{Code: "CN433100", Title: "湘西土家族苗族自治州"},
	}

	CitiesCN44 = []City{
		{Code: "CN440100", Title: "广州"},
		{Code: "CN440200", Title: "韶关"},
		{Code: "CN440300", Title: "深圳"},
		{Code: "CN440400", Title: "珠海"},
		{Code: "CN440500", Title: "汕头"},
		{Code: "CN440600", Title: "佛山"},
		{Code: "CN440700", Title: "江门"},
		{Code: "CN440800", Title: "湛江"},
		{Code: "CN440900", Title: "茂名"},
		{Code: "CN441200", Title: "肇庆"},
		{Code: "CN441300", Title: "惠州"},
		{Code: "CN441400", Title: "梅州"},
		{Code: "CN441500", Title: "汕尾"},
		{Code: "CN441600", Title: "河源"},
		{Code: "CN441700", Title: "阳江"},
		{Code: "CN441800", Title: "清远"},
		{Code: "CN441900", Title: "东莞"},
		{Code: "CN442000", Title: "中山"},
		{Code: "CN445100", Title: "潮州"},
		{Code: "CN445200", Title: "揭阳"},
		{Code: "CN445300", Title: "云浮"},
	}

	CitiesCN45 = []City{
		{Code: "CN450100", Title: "南宁"},
		{Code: "CN450200", Title: "柳州"},
		{Code: "CN450300", Title: "桂林"},
		{Code: "CN450400", Title: "梧州"},
		{Code: "CN450500", Title: "北海"},
		{Code: "CN450600", Title: "防城港"},
		{Code: "CN450700", Title: "钦州"},
		{Code: "CN450800", Title: "贵港"},
		{Code: "CN450900", Title: "玉林"},
		{Code: "CN451000", Title: "百色"},
		{Code: "CN451100", Title: "贺州"},
		{Code: "CN451200", Title: "河池"},
		{Code: "CN451300", Title: "来宾"},
		{Code: "CN451400", Title: "崇左"},
	}

	CitiesCN46 = []City{
		{Code: "CN460100", Title: "海口"},
		{Code: "CN460105", Title: "秀英区"},
		{Code: "CN460106", Title: "龙华区"},
		{Code: "CN460107", Title: "琼山区"},
		{Code: "CN460108", Title: "美兰区"},
		{Code: "CN460200", Title: "三亚"},
		{Code: "CN460202", Title: "海棠区"},
		{Code: "CN460203", Title: "吉阳区"},
		{Code: "CN460204", Title: "天涯区"},
		{Code: "CN460205", Title: "崖州区"},
		{Code: "CN460300", Title: "三沙"},
		{Code: "CN460400", Title: "儋州"},
		{Code: "CN469001", Title: "五指山"},
		{Code: "CN469002", Title: "琼海"},
		{Code: "CN469005", Title: "文昌"},
		{Code: "CN469006", Title: "万宁"},
		{Code: "CN469007", Title: "东方"},
		{Code: "CN469021", Title: "定安县"},
		{Code: "CN469022", Title: "屯昌县"},
		{Code: "CN469023", Title: "澄迈县"},
		{Code: "CN469024", Title: "临高县"},
		{Code: "CN469025", Title: "白沙黎族自治县"},
		{Code: "CN469026", Title: "昌江黎族自治县"},
		{Code: "CN469027", Title: "乐东黎族自治县"},
		{Code: "CN469028", Title: "陵水黎族自治县"},
		{Code: "CN469029", Title: "保亭黎族苗族自治县"},
		{Code: "CN469030", Title: "琼中黎族苗族自治县"},
	}

	CitiesCN51 = []City{
		{Code: "CN510100", Title: "成都"},
		{Code: "CN510300", Title: "自贡"},
		{Code: "CN510400", Title: "攀枝花"},
		{Code: "CN510500", Title: "泸州"},
		{Code: "CN510600", Title: "德阳"},
		{Code: "CN510700", Title: "绵阳"},
		{Code: "CN510800", Title: "广元"},
		{Code: "CN510900", Title: "遂宁"},
		{Code: "CN511000", Title: "内江"},
		{Code: "CN511100", Title: "乐山"},
		{Code: "CN511300", Title: "南充"},
		{Code: "CN511400", Title: "眉山"},
		{Code: "CN511500", Title: "宜宾"},
		{Code: "CN511600", Title: "广安"},
		{Code: "CN511700", Title: "达州"},
		{Code: "CN511800", Title: "雅安"},
		{Code: "CN511900", Title: "巴中"},
		{Code: "CN512000", Title: "资阳"},
		{Code: "CN513200", Title: "阿坝藏族羌族自治州"},
		{Code: "CN513300", Title: "甘孜藏族自治州"},
		{Code: "CN513400", Title: "凉山彝族自治州"},
	}

	CitiesCN52 = []City{
		{Code: "CN520100", Title: "贵阳"},
		{Code: "CN520200", Title: "六盘水"},
		{Code: "CN520300", Title: "遵义"},
		{Code: "CN520400", Title: "安顺"},
		{Code: "CN520500", Title: "毕节"},
		{Code: "CN520600", Title: "铜仁"},
		{Code: "CN522300", Title: "黔西南布依族苗族自治州"},
		{Code: "CN522600", Title: "黔东南苗族侗族自治州"},
		{Code: "CN522700", Title: "黔南布依族苗族自治州"},
	}

	CitiesCN53 = []City{
		{Code: "CN530100", Title: "昆明"},
		{Code: "CN530300", Title: "曲靖"},
		{Code: "CN530400", Title: "玉溪"},
		{Code: "CN530500", Title: "保山"},
		{Code: "CN530600", Title: "昭通"},
		{Code: "CN530700", Title: "丽江"},
		{Code: "CN530800", Title: "普洱"},
		{Code: "CN530900", Title: "临沧"},
		{Code: "CN532300", Title: "楚雄彝族自治州"},
		{Code: "CN532500", Title: "红河哈尼族彝族自治州"},
		{Code: "CN532600", Title: "文山壮族苗族自治州"},
		{Code: "CN532800", Title: "西双版纳傣族自治州"},
		{Code: "CN532900", Title: "大理白族自治州"},
		{Code: "CN533100", Title: "德宏傣族景颇族自治州"},
		{Code: "CN533300", Title: "怒江傈僳族自治州"},
		{Code: "CN533400", Title: "迪庆藏族自治州"},
	}

	CitiesCN54 = []City{
		{Code: "CN540100", Title: "拉萨"},
		{Code: "CN540200", Title: "日喀则"},
		{Code: "CN540300", Title: "昌都"},
		{Code: "CN540400", Title: "林芝"},
		{Code: "CN540500", Title: "山南"},
		{Code: "CN540600", Title: "那曲"},
		{Code: "CN542500", Title: "阿里地区"},
	}

	CitiesCN61 = []City{
		{Code: "CN610100", Title: "西安"},
		{Code: "CN610200", Title: "铜川"},
		{Code: "CN610300", Title: "宝鸡"},
		{Code: "CN610400", Title: "咸阳"},
		{Code: "CN610500", Title: "渭南"},
		{Code: "CN610600", Title: "延安"},
		{Code: "CN610700", Title: "汉中"},
		{Code: "CN610800", Title: "榆林"},
		{Code: "CN610900", Title: "安康"},
		{Code: "CN611000", Title: "商洛"},
	}

	CitiesCN62 = []City{
		{Code: "CN620100", Title: "兰州"},
		{Code: "CN620200", Title: "嘉峪关"},
		{Code: "CN620300", Title: "金昌"},
		{Code: "CN620400", Title: "白银"},
		{Code: "CN620500", Title: "天水"},
		{Code: "CN620600", Title: "武威"},
		{Code: "CN620700", Title: "张掖"},
		{Code: "CN620800", Title: "平凉"},
		{Code: "CN620900", Title: "酒泉"},
		{Code: "CN621000", Title: "庆阳"},
		{Code: "CN621100", Title: "定西"},
		{Code: "CN621200", Title: "陇南"},
		{Code: "CN622900", Title: "临夏回族自治州"},
		{Code: "CN623000", Title: "甘南藏族自治州"},
	}

	CitiesCN63 = []City{
		{Code: "CN630100", Title: "西宁"},
		{Code: "CN630200", Title: "海东"},
		{Code: "CN632200", Title: "海北藏族自治州"},
		{Code: "CN632300", Title: "黄南藏族自治州"},
		{Code: "CN632500", Title: "海南藏族自治州"},
		{Code: "CN632600", Title: "果洛藏族自治州"},
		{Code: "CN632700", Title: "玉树藏族自治州"},
		{Code: "CN632800", Title: "海西蒙古族藏族自治州"},
	}

	CitiesCN64 = []City{
		{Code: "CN640100", Title: "银川"},
		{Code: "CN640200", Title: "石嘴山"},
		{Code: "CN640300", Title: "吴忠"},
		{Code: "CN640400", Title: "固原"},
		{Code: "CN640500", Title: "中卫"},
	}

	CitiesCN65 = []City{
		{Code: "CN650100", Title: "乌鲁木齐"},
		{Code: "CN650200", Title: "克拉玛依"},
		{Code: "CN650400", Title: "吐鲁番"},
		{Code: "CN650500", Title: "哈密"},
		{Code: "CN652300", Title: "昌吉回族自治州"},
		{Code: "CN652700", Title: "博尔塔拉蒙古自治州"},
		{Code: "CN652800", Title: "巴音郭楞蒙古自治州"},
		{Code: "CN652900", Title: "阿克苏地区"},
		{Code: "CN653000", Title: "克孜勒苏柯尔克孜自治州"},
		{Code: "CN653100", Title: "喀什地区"},
		{Code: "CN653200", Title: "和田地区"},
		{Code: "CN654000", Title: "伊犁哈萨克自治州"},
		{Code: "CN654200", Title: "塔城地区"},
		{Code: "CN654300", Title: "阿勒泰地区"},
		{Code: "CN659001", Title: "石河子"},
		{Code: "CN659002", Title: "阿拉尔"},
		{Code: "CN659003", Title: "图木舒克"},
		{Code: "CN659004", Title: "五家渠"},
		{Code: "CN659005", Title: "北屯"},
		{Code: "CN659006", Title: "铁门关"},
		{Code: "CN659007", Title: "双河"},
		{Code: "CN659008", Title: "可克达拉"},
		{Code: "CN659009", Title: "昆玉"},
	}
	// </editor-fold>

	allCitiesCN = map[string][]City{
		"CN13": CitiesCN13,
		"CN14": CitiesCN14,
		"CN15": CitiesCN15,
		"CN21": CitiesCN21,
		"CN22": CitiesCN22,
		"CN23": CitiesCN23,
		"CN32": CitiesCN32,
		"CN33": CitiesCN33,
		"CN34": CitiesCN34,
		"CN35": CitiesCN35,
		"CN36": CitiesCN36,
		"CN37": CitiesCN37,
		"CN41": CitiesCN41,
		"CN42": CitiesCN42,
		"CN43": CitiesCN43,
		"CN44": CitiesCN44,
		"CN45": CitiesCN45,
		"CN46": CitiesCN46,
		"CN51": CitiesCN51,
		"CN52": CitiesCN52,
		"CN53": CitiesCN53,
		"CN54": CitiesCN54,
		"CN61": CitiesCN61,
		"CN62": CitiesCN62,
		"CN63": CitiesCN63,
		"CN64": CitiesCN64,
		"CN65": CitiesCN65,
	}
)
