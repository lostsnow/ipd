package ip

type Country struct {
	Title      string
	Code       string
	States     []State
	StateCodes map[string]State
}

func init() {
	for _, v := range allCountries {
		countryCodes[v.Title] = v
	}
}

func GetCountry(c string) *Country {
	if v, ok := countryCodes[c]; ok {
		return &v
	}
	return nil
}

func (c *Country) GetStates() []State {
	return c.States
}

func (c *Country) GetState(s string) *State {
	if v, ok := c.StateCodes[s]; ok {
		return &v
	}
	return nil
}

var (
	// http://zh.wikipedia.org/wiki/ISO_3166-1

	countryCodes = map[string]Country{}

	// <editor-fold desc="country codes">
	CountryAD = Country{Code: "AD", Title: "安道尔"}
	CountryAE = Country{Code: "AE", Title: "阿联酋"}
	CountryAF = Country{Code: "AF", Title: "阿富汗"}
	CountryAG = Country{Code: "AG", Title: "安提瓜和巴布达"}
	CountryAI = Country{Code: "AI", Title: "安圭拉"}
	CountryAL = Country{Code: "AL", Title: "阿尔巴尼亚"}
	CountryAM = Country{Code: "AM", Title: "亚美尼亚"}
	CountryAO = Country{Code: "AO", Title: "安哥拉"}
	CountryAQ = Country{Code: "AQ", Title: "南极洲"}
	CountryAR = Country{Code: "AR", Title: "阿根廷"}
	CountryAS = Country{Code: "AS", Title: "美属萨摩亚"}
	CountryAT = Country{Code: "AT", Title: "奥地利"}
	CountryAU = Country{Code: "AU", Title: "澳大利亚"}
	CountryAW = Country{Code: "AW", Title: "阿鲁巴"}
	CountryAX = Country{Code: "AX", Title: "奥兰群岛"}
	CountryAZ = Country{Code: "AZ", Title: "阿塞拜疆"}
	CountryBA = Country{Code: "BA", Title: "波斯尼亚和黑塞哥维那"}
	CountryBB = Country{Code: "BB", Title: "巴巴多斯"}
	CountryBD = Country{Code: "BD", Title: "孟加拉"}
	CountryBE = Country{Code: "BE", Title: "比利时"}
	CountryBF = Country{Code: "BF", Title: "布基纳法索"}
	CountryBG = Country{Code: "BG", Title: "保加利亚"}
	CountryBH = Country{Code: "BH", Title: "巴林"}
	CountryBI = Country{Code: "BI", Title: "布隆迪"}
	CountryBJ = Country{Code: "BJ", Title: "贝宁"}
	CountryBL = Country{Code: "BL", Title: "圣巴泰勒米岛"}
	CountryBM = Country{Code: "BM", Title: "百慕大"}
	CountryBN = Country{Code: "BN", Title: "文莱"}
	CountryBO = Country{Code: "BO", Title: "玻利维亚"}
	CountryBQ = Country{Code: "BQ", Title: "荷兰加勒比"}
	CountryBR = Country{Code: "BR", Title: "巴西"}
	CountryBS = Country{Code: "BS", Title: "巴哈马"}
	CountryBT = Country{Code: "BT", Title: "不丹"}
	CountryBV = Country{Code: "BV", Title: "布韦岛"}
	CountryBW = Country{Code: "BW", Title: "博茨瓦纳"}
	CountryBY = Country{Code: "BY", Title: "白俄罗斯"}
	CountryBZ = Country{Code: "BZ", Title: "伯利兹"}
	CountryCA = Country{Code: "CA", Title: "加拿大"}
	CountryCC = Country{Code: "CC", Title: "科科斯(基林)群岛"}
	CountryCD = Country{Code: "CD", Title: "刚果(金)"}
	CountryCF = Country{Code: "CF", Title: "中非"}
	CountryCG = Country{Code: "CG", Title: "刚果(布)"}
	CountryCH = Country{Code: "CH", Title: "瑞士"}
	CountryCI = Country{Code: "CI", Title: "科特迪瓦"}
	CountryCK = Country{Code: "CK", Title: "库克群岛"}
	CountryCL = Country{Code: "CL", Title: "智利"}
	CountryCM = Country{Code: "CM", Title: "喀麦隆"}
	CountryCN = Country{Code: "CN", Title: "中国", States: allStatesCN, StateCodes: stateCodesCN}
	CountryCO = Country{Code: "CO", Title: "哥伦比亚"}
	CountryCR = Country{Code: "CR", Title: "哥斯达黎加"}
	CountryCU = Country{Code: "CU", Title: "古巴"}
	CountryCV = Country{Code: "CV", Title: "佛得角"}
	CountryCW = Country{Code: "CW", Title: "库拉索"}
	CountryCX = Country{Code: "CX", Title: "圣诞岛"}
	CountryCY = Country{Code: "CY", Title: "塞浦路斯"}
	CountryCZ = Country{Code: "CZ", Title: "捷克"}
	CountryDE = Country{Code: "DE", Title: "德国"}
	CountryDJ = Country{Code: "DJ", Title: "吉布提"}
	CountryDK = Country{Code: "DK", Title: "丹麦"}
	CountryDM = Country{Code: "DM", Title: "多米尼克"}
	CountryDO = Country{Code: "DO", Title: "多米尼加"}
	CountryDZ = Country{Code: "DZ", Title: "阿尔及利亚"}
	CountryEC = Country{Code: "EC", Title: "厄瓜多尔"}
	CountryEE = Country{Code: "EE", Title: "爱沙尼亚"}
	CountryEG = Country{Code: "EG", Title: "埃及"}
	CountryEH = Country{Code: "EH", Title: "西撒哈拉"}
	CountryER = Country{Code: "ER", Title: "厄立特里亚"}
	CountryES = Country{Code: "ES", Title: "西班牙"}
	CountryET = Country{Code: "ET", Title: "埃塞俄比亚"}
	CountryFI = Country{Code: "FI", Title: "芬兰"}
	CountryFJ = Country{Code: "FJ", Title: "斐济"}
	CountryFK = Country{Code: "FK", Title: "福克兰群岛"}
	CountryFM = Country{Code: "FM", Title: "密克罗尼西亚联邦"}
	CountryFO = Country{Code: "FO", Title: "法罗群岛"}
	CountryFR = Country{Code: "FR", Title: "法国"}
	CountryGA = Country{Code: "GA", Title: "加蓬"}
	CountryGB = Country{Code: "GB", Title: "英国"}
	CountryGD = Country{Code: "GD", Title: "格林纳达"}
	CountryGE = Country{Code: "GE", Title: "格鲁吉亚"}
	CountryGF = Country{Code: "GF", Title: "法属圭亚那"}
	CountryGG = Country{Code: "GG", Title: "根西岛"}
	CountryGH = Country{Code: "GH", Title: "加纳"}
	CountryGI = Country{Code: "GI", Title: "直布罗陀"}
	CountryGL = Country{Code: "GL", Title: "格陵兰"}
	CountryGM = Country{Code: "GM", Title: "冈比亚"}
	CountryGN = Country{Code: "GN", Title: "几内亚"}
	CountryGP = Country{Code: "GP", Title: "瓜德罗普"}
	CountryGQ = Country{Code: "GQ", Title: "赤道几内亚"}
	CountryGR = Country{Code: "GR", Title: "希腊"}
	CountryGS = Country{Code: "GS", Title: "南乔治亚岛和南桑威奇群岛"}
	CountryGT = Country{Code: "GT", Title: "危地马拉"}
	CountryGU = Country{Code: "GU", Title: "关岛"}
	CountryGW = Country{Code: "GW", Title: "几内亚比绍"}
	CountryGY = Country{Code: "GY", Title: "圭亚那"}
	CountryHK = Country{Code: "HK", Title: "香港"}
	CountryHM = Country{Code: "HM", Title: "赫德岛和麦克唐纳群岛"}
	CountryHN = Country{Code: "HN", Title: "洪都拉斯"}
	CountryHR = Country{Code: "HR", Title: "克罗地亚"}
	CountryHT = Country{Code: "HT", Title: "海地"}
	CountryHU = Country{Code: "HU", Title: "匈牙利"}
	CountryID = Country{Code: "ID", Title: "印度尼西亚"}
	CountryIE = Country{Code: "IE", Title: "爱尔兰"}
	CountryIL = Country{Code: "IL", Title: "以色列"}
	CountryIM = Country{Code: "IM", Title: "马恩岛"}
	CountryIN = Country{Code: "IN", Title: "印度"}
	CountryIO = Country{Code: "IO", Title: "英属印度洋领地"}
	CountryIQ = Country{Code: "IQ", Title: "伊拉克"}
	CountryIR = Country{Code: "IR", Title: "伊朗"}
	CountryIS = Country{Code: "IS", Title: "冰岛"}
	CountryIT = Country{Code: "IT", Title: "意大利"}
	CountryJE = Country{Code: "JE", Title: "泽西岛"}
	CountryJM = Country{Code: "JM", Title: "牙买加"}
	CountryJO = Country{Code: "JO", Title: "约旦"}
	CountryJP = Country{Code: "JP", Title: "日本"}
	CountryKE = Country{Code: "KE", Title: "肯尼亚"}
	CountryKG = Country{Code: "KG", Title: "吉尔吉斯斯坦"}
	CountryKH = Country{Code: "KH", Title: "柬埔寨"}
	CountryKI = Country{Code: "KI", Title: "基里巴斯"}
	CountryKM = Country{Code: "KM", Title: "科摩罗"}
	CountryKN = Country{Code: "KN", Title: "圣基茨和尼维斯"}
	CountryKP = Country{Code: "KP", Title: "朝鲜"}
	CountryKR = Country{Code: "KR", Title: "韩国"}
	CountryKW = Country{Code: "KW", Title: "科威特"}
	CountryKY = Country{Code: "KY", Title: "开曼群岛"}
	CountryKZ = Country{Code: "KZ", Title: "哈萨克斯坦"}
	CountryLA = Country{Code: "LA", Title: "老挝"}
	CountryLB = Country{Code: "LB", Title: "黎巴嫩"}
	CountryLC = Country{Code: "LC", Title: "圣卢西亚"}
	CountryLI = Country{Code: "LI", Title: "列支敦士登"}
	CountryLK = Country{Code: "LK", Title: "斯里兰卡"}
	CountryLR = Country{Code: "LR", Title: "利比里亚"}
	CountryLS = Country{Code: "LS", Title: "莱索托"}
	CountryLT = Country{Code: "LT", Title: "立陶宛"}
	CountryLU = Country{Code: "LU", Title: "卢森堡"}
	CountryLV = Country{Code: "LV", Title: "拉脱维亚"}
	CountryLY = Country{Code: "LY", Title: "利比亚"}
	CountryMA = Country{Code: "MA", Title: "摩洛哥"}
	CountryMC = Country{Code: "MC", Title: "摩纳哥"}
	CountryMD = Country{Code: "MD", Title: "摩尔多瓦"}
	CountryME = Country{Code: "ME", Title: "黑山"}
	CountryMF = Country{Code: "MF", Title: "法属圣马丁"}
	CountryMG = Country{Code: "MG", Title: "马达加斯加"}
	CountryMH = Country{Code: "MH", Title: "马绍尔群岛"}
	CountryMK = Country{Code: "MK", Title: "马其顿"}
	CountryML = Country{Code: "ML", Title: "马里"}
	CountryMM = Country{Code: "MM", Title: "缅甸"}
	CountryMN = Country{Code: "MN", Title: "蒙古"}
	CountryMO = Country{Code: "MO", Title: "澳门"}
	CountryMP = Country{Code: "MP", Title: "北马里亚纳群岛"}
	CountryMQ = Country{Code: "MQ", Title: "马提尼克岛"}
	CountryMR = Country{Code: "MR", Title: "毛里塔尼亚"}
	CountryMS = Country{Code: "MS", Title: "蒙特塞拉特岛"}
	CountryMT = Country{Code: "MT", Title: "马耳他"}
	CountryMU = Country{Code: "MU", Title: "毛里求斯"}
	CountryMV = Country{Code: "MV", Title: "马尔代夫"}
	CountryMW = Country{Code: "MW", Title: "马拉维"}
	CountryMX = Country{Code: "MX", Title: "墨西哥"}
	CountryMY = Country{Code: "MY", Title: "马来西亚"}
	CountryMZ = Country{Code: "MZ", Title: "莫桑比克"}
	CountryNA = Country{Code: "NA", Title: "纳米比亚"}
	CountryNC = Country{Code: "NC", Title: "新喀里多尼亚"}
	CountryNE = Country{Code: "NE", Title: "尼日尔"}
	CountryNF = Country{Code: "NF", Title: "诺福克岛"}
	CountryNG = Country{Code: "NG", Title: "尼日利亚"}
	CountryNI = Country{Code: "NI", Title: "尼加拉瓜"}
	CountryNK = Country{Code: "NK", Title: "纳戈尔诺-卡拉巴赫"}
	CountryNL = Country{Code: "NL", Title: "荷兰"}
	CountryNO = Country{Code: "NO", Title: "挪威"}
	CountryNP = Country{Code: "NP", Title: "尼泊尔"}
	CountryNR = Country{Code: "NR", Title: "瑙鲁"}
	CountryNU = Country{Code: "NU", Title: "纽埃岛"}
	CountryNZ = Country{Code: "NZ", Title: "新西兰"}
	CountryOM = Country{Code: "OM", Title: "阿曼"}
	CountryPA = Country{Code: "PA", Title: "巴拿马"}
	CountryPE = Country{Code: "PE", Title: "秘鲁"}
	CountryPF = Country{Code: "PF", Title: "法属波利尼西亚"}
	CountryPG = Country{Code: "PG", Title: "巴布亚新几内亚"}
	CountryPH = Country{Code: "PH", Title: "菲律宾"}
	CountryPK = Country{Code: "PK", Title: "巴基斯坦"}
	CountryPL = Country{Code: "PL", Title: "波兰"}
	CountryPM = Country{Code: "PM", Title: "圣皮埃尔和密克隆群岛"}
	CountryPN = Country{Code: "PN", Title: "皮特凯恩群岛"}
	CountryPR = Country{Code: "PR", Title: "波多黎各"}
	CountryPS = Country{Code: "PS", Title: "巴勒斯坦"}
	CountryPT = Country{Code: "PT", Title: "葡萄牙"}
	CountryPW = Country{Code: "PW", Title: "帕劳"}
	CountryPY = Country{Code: "PY", Title: "巴拉圭"}
	CountryQA = Country{Code: "QA", Title: "卡塔尔"}
	CountryRE = Country{Code: "RE", Title: "留尼汪岛"}
	CountryRO = Country{Code: "RO", Title: "罗马尼亚"}
	CountryRS = Country{Code: "RS", Title: "塞尔维亚"}
	CountryRU = Country{Code: "RU", Title: "俄罗斯"}
	CountryRW = Country{Code: "RW", Title: "卢旺达"}
	CountrySA = Country{Code: "SA", Title: "沙特阿拉伯"}
	CountrySB = Country{Code: "SB", Title: "所罗门群岛"}
	CountrySC = Country{Code: "SC", Title: "塞舌尔"}
	CountrySD = Country{Code: "SD", Title: "苏丹"}
	CountrySE = Country{Code: "SE", Title: "瑞典"}
	CountrySG = Country{Code: "SG", Title: "新加坡"}
	CountrySH = Country{Code: "SH", Title: "圣赫勒拿"}
	CountrySI = Country{Code: "SI", Title: "斯洛文尼亚"}
	CountrySJ = Country{Code: "SJ", Title: "斯瓦尔巴群岛和扬马延岛"}
	CountrySK = Country{Code: "SK", Title: "斯洛伐克"}
	CountrySL = Country{Code: "SL", Title: "塞拉利昂"}
	CountrySM = Country{Code: "SM", Title: "圣马力诺"}
	CountrySN = Country{Code: "SN", Title: "塞内加尔"}
	CountrySO = Country{Code: "SO", Title: "索马里"}
	CountrySR = Country{Code: "SR", Title: "苏里南"}
	CountrySS = Country{Code: "SS", Title: "南苏丹"}
	CountryST = Country{Code: "ST", Title: "圣多美和普林西比"}
	CountrySV = Country{Code: "SV", Title: "萨尔瓦多"}
	CountrySX = Country{Code: "SX", Title: "荷属圣马丁"}
	CountrySY = Country{Code: "SY", Title: "叙利亚"}
	CountrySZ = Country{Code: "SZ", Title: "斯威士兰"}
	CountryTC = Country{Code: "TC", Title: "特克斯和凯科斯群岛"}
	CountryTD = Country{Code: "TD", Title: "乍得"}
	CountryTF = Country{Code: "TF", Title: "法属南部领地"}
	CountryTG = Country{Code: "TG", Title: "多哥"}
	CountryTH = Country{Code: "TH", Title: "泰国"}
	CountryTJ = Country{Code: "TJ", Title: "塔吉克斯坦"}
	CountryTK = Country{Code: "TK", Title: "托克劳群岛"}
	CountryTL = Country{Code: "TL", Title: "东帝汶"}
	CountryTM = Country{Code: "TM", Title: "土库曼斯坦"}
	CountryTN = Country{Code: "TN", Title: "突尼斯"}
	CountryTO = Country{Code: "TO", Title: "汤加"}
	CountryTR = Country{Code: "TR", Title: "土耳其"}
	CountryTT = Country{Code: "TT", Title: "特立尼达和多巴哥"}
	CountryTV = Country{Code: "TV", Title: "图瓦卢"}
	CountryTW = Country{Code: "TW", Title: "台湾"}
	CountryTZ = Country{Code: "TZ", Title: "坦桑尼亚"}
	CountryUA = Country{Code: "UA", Title: "乌克兰"}
	CountryUG = Country{Code: "UG", Title: "乌干达"}
	CountryUM = Country{Code: "UM", Title: "美国本土外小岛屿"}
	CountryUS = Country{Code: "US", Title: "美国"}
	CountryUY = Country{Code: "UY", Title: "乌拉圭"}
	CountryUZ = Country{Code: "UZ", Title: "乌兹别克斯坦"}
	CountryVA = Country{Code: "VA", Title: "梵蒂冈"}
	CountryVC = Country{Code: "VC", Title: "圣文森特和格林纳丁斯"}
	CountryVE = Country{Code: "VE", Title: "委内瑞拉"}
	CountryVG = Country{Code: "VG", Title: "英属维尔京群岛"}
	CountryVI = Country{Code: "VI", Title: "美属维尔京群岛"}
	CountryVN = Country{Code: "VN", Title: "越南"}
	CountryVU = Country{Code: "VU", Title: "瓦努阿图"}
	CountryWF = Country{Code: "WF", Title: "瓦利斯和富图纳群岛"}
	CountryWS = Country{Code: "WS", Title: "萨摩亚"}
	CountryYE = Country{Code: "YE", Title: "也门"}
	CountryYT = Country{Code: "YT", Title: "马约特"}
	CountryZA = Country{Code: "ZA", Title: "南非"}
	CountryZM = Country{Code: "ZM", Title: "赞比亚"}
	CountryZW = Country{Code: "ZW", Title: "津巴布韦"}
	// </editor-fold>

	// <editor-fold desc="all countries">
	allCountries = []Country{
		CountryAD, CountryAE, CountryAF, CountryAG, CountryAI, CountryAL, CountryAM, CountryAO,
		CountryAQ, CountryAR, CountryAS, CountryAT, CountryAU, CountryAW, CountryAX, CountryAZ,
		CountryBA, CountryBB, CountryBD, CountryBE, CountryBF, CountryBG, CountryBH, CountryBI,
		CountryBJ, CountryBL, CountryBM, CountryBN, CountryBO, CountryBQ, CountryBR, CountryBS,
		CountryBT, CountryBV, CountryBW, CountryBY, CountryBZ,
		CountryCA, CountryCC, CountryCD, CountryCF, CountryCG, CountryCH, CountryCI, CountryCK,
		CountryCL, CountryCM, CountryCN, CountryCO, CountryCR, CountryCU, CountryCV, CountryCW,
		CountryCX, CountryCY, CountryCZ,
		CountryDE, CountryDJ, CountryDK, CountryDM, CountryDO, CountryDZ,
		CountryEC, CountryEE, CountryEG, CountryEH, CountryER, CountryES, CountryET, CountryFI,
		CountryFJ, CountryFK, CountryFM, CountryFO, CountryFR, CountryGA, CountryGB, CountryGD,
		CountryGE, CountryGF, CountryGG, CountryGH, CountryGI, CountryGL, CountryGM, CountryGN,
		CountryGP, CountryGQ, CountryGR, CountryGS, CountryGT, CountryGU, CountryGW, CountryGY,
		CountryHK, CountryHM, CountryHN, CountryHR, CountryHT, CountryHU,
		CountryID, CountryIE, CountryIL, CountryIM, CountryIN, CountryIO, CountryIQ, CountryIR,
		CountryIS, CountryIT,
		CountryJE, CountryJM, CountryJO, CountryJP,
		CountryKE, CountryKG, CountryKH, CountryKI, CountryKM, CountryKN, CountryKP, CountryKR,
		CountryKW, CountryKY, CountryKZ,
		CountryLA, CountryLB, CountryLC, CountryLI, CountryLK, CountryLR, CountryLS, CountryLT,
		CountryLU, CountryLV, CountryLY,
		CountryMA, CountryMC, CountryMD, CountryME, CountryMF, CountryMG, CountryMH, CountryMK,
		CountryML, CountryMM, CountryMN, CountryMO, CountryMP, CountryMQ, CountryMR, CountryMS,
		CountryMT, CountryMU, CountryMV, CountryMW, CountryMX, CountryMY, CountryMZ,
		CountryNA, CountryNC, CountryNE, CountryNF, CountryNG, CountryNI, CountryNK, CountryNL,
		CountryNO, CountryNP, CountryNR, CountryNU, CountryNZ,
		CountryOM,
		CountryPA, CountryPE, CountryPF, CountryPG, CountryPH, CountryPK, CountryPL, CountryPM,
		CountryPN, CountryPR, CountryPS, CountryPT, CountryPW, CountryPY,
		CountryQA,
		CountryRE, CountryRO, CountryRS, CountryRU, CountryRW,
		CountrySA, CountrySB, CountrySC, CountrySD, CountrySE, CountrySG, CountrySH, CountrySI,
		CountrySJ, CountrySK, CountrySL, CountrySM, CountrySN, CountrySO, CountrySR, CountrySS,
		CountryST, CountrySV, CountrySX, CountrySY, CountrySZ,
		CountryTC, CountryTD, CountryTF, CountryTG, CountryTH, CountryTJ, CountryTK, CountryTL,
		CountryTM, CountryTN, CountryTO, CountryTR, CountryTT, CountryTV, CountryTW, CountryTZ,
		CountryUA, CountryUG, CountryUM, CountryUS, CountryUY, CountryUZ,
		CountryVA, CountryVC, CountryVE, CountryVG, CountryVI, CountryVN, CountryVU,
		CountryWF, CountryWS,
		CountryYE, CountryYT,
		CountryZA, CountryZM, CountryZW,
	}
	// </editor-fold>

	irregularCountries = map[string]bool{
		"012.NET 骨干网":                  true,
		"0LOSS.COM":                    true,
		"114DNS.COM":                   true,
		"1AND1.COM":                    true,
		"1AND1.COM 骨干网":                true,
		"3M.COM":                       true,
		"3QSDN.COM":                    true,
		"51EIDC.COM":                   true,
		"6CONNECT.COM":                 true,
		"A2WEBHOSTING.COM":             true,
		"AAGICO.COM":                   true,
		"AARNET.EDU.AU 骨干网":            true,
		"ABSATELLITE.COM 骨干网":          true,
		"ACEDC 骨干网":                    true,
		"AC.ME":                        true,
		"ADOBE.COM":                    true,
		"ADVANCEDHOSTERS.COM":          true,
		"AEDNS.AE":                     true,
		"AEROTEK.COM.TR":               true,
		"AETHOSTING.COM":               true,
		"AFILIAS.INFO":                 true,
		"AGARTO.COM":                   true,
		"A.GTLD.BIZ":                   true,
		"AIRTEL.COM 骨干网":               true,
		"AKAMAI.COM":                   true,
		"ALCOM.AX 骨干网":                 true,
		"ALENTUS":                      true,
		"ALIBABA.COM 骨干网":              true,
		"ALIDNS.COM":                   true,
		"ALIYUN.COM":                   true,
		"ALIYUN":                       true,
		"ALTER.NET 骨干网":                true,
		"AMAZON.COM":                   true,
		"AMAZON.COM 骨干网":               true,
		"AMCBB 骨干网":                    true,
		"AMNIC.NET":                    true,
		"ANEXIA-IT.COM":                true,
		"ANYCAST.IO":                   true,
		"AORTA 骨干网":                    true,
		"APNIC.NET":                    true,
		"APPLE.COM":                    true,
		"APPLE.COM 骨干网":                true,
		"APPNEXUS.COM":                 true,
		"APPRIVER.COM":                 true,
		"ARKADIN.COM":                  true,
		"A-ROOT-SERVERS":               true,
		"ARTFILES.COM":                 true,
		"ARVAN.CLOUD":                  true,
		"ARYAKA.COM":                   true,
		"ASC.EDU 骨干网":                  true,
		"ASK.COM":                      true,
		"ASTUTEINTERNET.COM":           true,
		"ASTUTEINTERNET.COM 骨干网":       true,
		"ATECHMEDIA":                   true,
		"ATHEA.AR":                     true,
		"ATLANTECH.NET 骨干网":            true,
		"ATRATO.COM 骨干网":               true,
		"ATT.COM":                      true,
		"ATT.COM 骨干网":                  true,
		"AUDA.ORG.AU":                  true,
		"AUSREGISTRY.COM.AU":           true,
		"AUTOMATTIC":                   true,
		"AVANTIPLC.COM 骨干网":            true,
		"AVAST.COM":                    true,
		"AZION.COM.BR":                 true,
		"Azure-DNS":                    true,
		"BACKBONE.IS 骨干网":              true,
		"BAHNHOF.NET 骨干网":              true,
		"BAIDU.COM":                    true,
		"BAREFRUIT.CO.UK":              true,
		"BELGACOM 骨干网":                 true,
		"BELL.CA 骨干网":                  true,
		"BELLCA 骨干网":                   true,
		"BETAIDC.COM":                  true,
		"BGP.NET 骨干网":                  true,
		"B.GTLD.BIZ":                   true,
		"BITGRAVITY.COM":               true,
		"BLACKBERRY.COM":               true,
		"BLIX.COM":                     true,
		"BLIZZARD.COM 骨干网":             true,
		"BOOKING.COM":                  true,
		"B-ROOT-SERVERS":               true,
		"BSONETWORK.COM 骨干网":           true,
		"BT.COM 骨干网":                   true,
		"BUNNYCDN.COM":                 true,
		"BUYVM.NET":                    true,
		"C4L.CO.UK 骨干网":                true,
		"CACHEFLY.COM":                 true,
		"CARRIER2CARRIER 骨干网":          true,
		"CA-SERVERS.CA":                true,
		"CDLAN.IT 骨干网":                 true,
		"CDN77.COM":                    true,
		"CDNETWORKS.COM":               true,
		"CEDEXIS.COM":                  true,
		"CENTRALNIC.COM":               true,
		"CERT.AT":                      true,
		"C.GTLD.BIZ":                   true,
		"CHELLO 骨干网":                   true,
		"CHINANETCENTER.COM":           true,
		"CHOOPADNS.COM":                true,
		"CHTGLOBAL.COM 骨干网":            true,
		"CIRA.CA":                      true,
		"CITI.COM":                     true,
		"CITRIX.COM 骨干网":               true,
		"CITYTELECOM.RU 骨干网":           true,
		"CLARKSYS.COM":                 true,
		"CLEARHOST.CO.UK":              true,
		"CLOUDACCESS.NET":              true,
		"CLOUDBROKERS.AT":              true,
		"CLOUDDNS.NET":                 true,
		"CLOUDFLARE.COM":               true,
		"CLOUDFLOORDNS.COM":            true,
		"CLOUD.GOOGLE.COM":             true,
		"CLOUVIDER.NET":                true,
		"CMC.IQ":                       true,
		"CNET.NET":                     true,
		"CNNIC.CN":                     true,
		"COGECOPEER1.COM":              true,
		"COGENTCO.COM":                 true,
		"COGENTCO.COM 骨干网":             true,
		"COLT.NET 骨干网":                 true,
		"COLUMBUS-NETWORKS 骨干网":        true,
		"COMNET 骨干网":                   true,
		"COMODO.COM":                   true,
		"CONEXIMGLOBAL.NET":            true,
		"CONPLEX.IO":                   true,
		"CONSOLECONNECT.COM 骨干网":       true,
		"CONTINENT8.COM 骨干网":           true,
		"CORE-BACKBONE.COM 骨干网":        true,
		"COREINTERNET.CO.UK":           true,
		"CORELINK.IO":                  true,
		"CPRM.NET 骨干网":                 true,
		"CRASHLYTICS.COM":              true,
		"C-ROOT-SERVERS":               true,
		"CVNT.NET":                     true,
		"CWASIA.NET 骨干网":               true,
		"CYANLINK.NET":                 true,
		"CYBERA.NET":                   true,
		"DAILYMOTION.COM":              true,
		"DAL.NET":                      true,
		"DATAHOP 骨干网":                  true,
		"DATAPATH.IO":                  true,
		"DATAPIPE.COM 骨干网":             true,
		"DATPIXEL.NET":                 true,
		"DDITSERVICES 骨干网":             true,
		"DENIC.DE":                     true,
		"DEVCAPSULE.COM":               true,
		"DEVELOPMENTENTITY.WORK":       true,
		"DIALPAD.COM 骨干网":              true,
		"DIGIDESERT.NET":               true,
		"DIGITAL.HT":                   true,
		"DIGIWEB.IE 骨干网":               true,
		"DIPEX.RU":                     true,
		"DISNEY.COM":                   true,
		"DISTILNETWORKS.COM":           true,
		"DNS.AR":                       true,
		"DNS.BE":                       true,
		"DNSBIRD.NET":                  true,
		"DNS.BR":                       true,
		"DNSFILTER.COM":                true,
		"DNSIMPLE.COM":                 true,
		"DNS.LU":                       true,
		"DNSNOVA.COM":                  true,
		"DNS-OARC.NET":                 true,
		"DNSPOD.COM":                   true,
		"DNSREC.MEO.WS":                true,
		"DOMENESHOP.NO":                true,
		"DOSARREST.COM":                true,
		"DOTQUAD.COM":                  true,
		"DOTQUAD.COM 骨干网":              true,
		"DQECOM.COM 骨干网":               true,
		"DREAMSCAPENETWORKS.COM":       true,
		"D-ROOT-SERVERS":               true,
		"DROPBOX.COM 骨干网":              true,
		"DROPBOX 骨干网":                  true,
		"DUDICOM.DK":                   true,
		"DUODECADITS.COM":              true,
		"DYN.COM":                      true,
		"E24DNS.COM":                   true,
		"EASYDNS.COM":                  true,
		"EASYNET.NET 骨干网":              true,
		"EBAY.COM 骨干网":                 true,
		"ECOMLTD 骨干网":                  true,
		"EDGECAST.COM":                 true,
		"EDPNET.NET 骨干网":               true,
		"EDSI-TECH.COM":                true,
		"EGIHOSTING":                   true,
		"E.GTLD.BIZ":                   true,
		"ELAUWIT.COM":                  true,
		"ELISA.FI 骨干网":                 true,
		"ELLIEMAE.COM 骨干网":             true,
		"EQUINIX.COM 骨干网":              true,
		"ERICSSONUDN.COM":              true,
		"E-ROOT-SERVERS":               true,
		"ESNET 骨干网":                    true,
		"ESOSOFT.NET":                  true,
		"euNetworks 骨干网":               true,
		"EXA.NET.UK 骨干网":               true,
		"FACEBOOK.COM":                 true,
		"FACEBOOK.COM 骨干网":             true,
		"FALCO-NETWORKS.COM":           true,
		"FARSNEWS.COM":                 true,
		"FASTLY.COM":                   true,
		"FFASTFILL.COM":                true,
		"F.GTLD.BIZ":                   true,
		"FIBERAX.COM 骨干网":              true,
		"FIBERDROP.NET":                true,
		"FIBER.GOOGLE.COM 骨干网":         true,
		"fibernetdirect.com 骨干网":       true,
		"FIBTEL.NET":                   true,
		"FIDO.NET":                     true,
		"FIORD.RU 骨干网":                 true,
		"FLOPS.RU":                     true,
		"FLYCAT.IO 骨干网":                true,
		"FONDATIONRESTENA":             true,
		"FORCEPOINT.COM":               true,
		"FORTINET.COM":                 true,
		"FREENOM.COM":                  true,
		"F-ROOT-SERVERS":               true,
		"GADITEK.COM":                  true,
		"GANDI.NET":                    true,
		"GANDI.NET 骨干网":                true,
		"GCORE.LU":                     true,
		"GDNS.SO":                      true,
		"GEANT 骨干网":                    true,
		"GEOSCALING.COM":               true,
		"GHX.COM":                      true,
		"GIGAS.COM":                    true,
		"GILATSATCOM 骨干网":              true,
		"GINERNET.COM":                 true,
		"GlobalCloudXchange 骨干网":       true,
		"GLOBALCOM 骨干网":                true,
		"GLOBALCONNECT.NET 骨干网":        true,
		"GLOBALTRANSIT.NET 骨干网":        true,
		"GLOBENET.NET 骨干网":             true,
		"GODADDY.COM":                  true,
		"GODADDY.COM 骨干网":              true,
		"GOOGLE.COM":                   true,
		"GOOGLE.COM 骨干网":               true,
		"GOOGLEVPN":                    true,
		"GOSCOMB 骨干网":                  true,
		"GRADWELL.COM":                 true,
		"GRANSY.COM":                   true,
		"GRATISDNS.DK":                 true,
		"G-ROOT-SERVERS":               true,
		"GTT.NET 骨干网":                  true,
		"HAWKHOST.COM":                 true,
		"HEG.COM 骨干网":                  true,
		"HE.NET":                       true,
		"HE.NET 骨干网":                   true,
		"HERE.COM 骨干网":                 true,
		"HGC.COM.HK 骨干网":               true,
		"HIBERNIACDN.COM":              true,
		"HIGHWINDS.COM":                true,
		"HIGHWINDS.COM 骨干网":            true,
		"HILLSIDENEWMEDIA.COM":         true,
		"HOPONE.NET 骨干网":               true,
		"HOST1PLUS.COM":                true,
		"HOSTEMO 骨干网":                  true,
		"HOSTING4REAL.NET":             true,
		"HOSTISERVER.COM":              true,
		"HOST.NET 骨干网":                 true,
		"HOSTNOC 骨干网":                  true,
		"HOSTUS.US":                    true,
		"HOSTVIRTUAL.COM":              true,
		"HOSTWAY.COM":                  true,
		"HOTWIRECOMMUNICATION.COM 骨干网": true,
		"HP 骨干网":                       true,
		"I3D.NET":                      true,
		"IANA-SERVERS.NET":             true,
		"ICANN.ORG":                    true,
		"ICS.FORTH.GR":                 true,
		"IFXNETWORKS.COM 骨干网":          true,
		"IGLOO.TO":                     true,
		"IGNUM.CZ":                     true,
		"IIJ.AD.JP":                    true,
		"IINET.NET.AU 骨干网":             true,
		"IMPLETEC.COM":                 true,
		"INCAPSULA.COM":                true,
		"INIT7.NET 骨干网":                true,
		"INNOVA":                       true,
		"INSTARTLOGIC.COM":             true,
		"INSTART":                      true,
		"INTELSATONE 骨干网":              true,
		"INTERNAP.COM":                 true,
		"INTERNET2 骨干网":                true,
		"INTERNETNAMESFORBUSINESS.COM": true,
		"INTEROUTE.COM 骨干网":            true,
		"INTERROUTE 骨干网":               true,
		"INVITEL.NET":                  true,
		"IPIP.NET":                     true,
		"IP-MAX 骨干网":                   true,
		"IP-ONLY.SE 骨干网":               true,
		"IPTP 骨干网":                     true,
		"IPTRANSIT 骨干网":                true,
		"I-ROOT-SERVERS":               true,
		"ISC.ORG":                      true,
		"ISC 骨干网":                      true,
		"ISPRIME.COM":                  true,
		"J-ROOT-SERVERS":               true,
		"KAIAGLOBAL 骨干网":               true,
		"KAMATERA.COM":                 true,
		"KARTINAWORLD.COM":             true,
		"KDDI.COM 骨干网":                 true,
		"K.GTLD.BIZ":                   true,
		"KLAYER.COM":                   true,
		"KNET.CN":                      true,
		"KNIPP.DE":                     true,
		"KOBALT.DK":                    true,
		"KOBO.COM 骨干网":                 true,
		"KPN.COM 骨干网":                  true,
		"K-ROOT-SERVERS":               true,
		"KT.COM 骨干网":                   true,
		"LAYER42 骨干网":                  true,
		"LAZADA.COM":                   true,
		"LEASEWEBCDN.COM":              true,
		"LEASEWEB 骨干网":                 true,
		"LEVEL3.COM":                   true,
		"LEVEL3.COM 骨干网":               true,
		"LG DACOM 骨干网":                 true,
		"LIFELINK.RU":                  true,
		"LIGHTEDGE.COM 骨干网":            true,
		"LIGHTOWER.COM 骨干网":            true,
		"LIMELIGHT.COM":                true,
		"LIMELIGHT.COM 骨干网":            true,
		"LINK11.DE 骨干网":                true,
		"LINKEDIN 骨干网":                 true,
		"LINKEY.RU 骨干网":                true,
		"LINKWORK.DK":                  true,
		"LINODE.COM":                   true,
		"LINXTELECOM 骨干网":              true,
		"LIQUIDTELECOM.COM 骨干网":        true,
		"LITTLEMOE.NET":                true,
		"LIVEAIR.NET 骨干网":              true,
		"LOGIC.BM 骨干网":                 true,
		"L-ROOT-SERVERS":               true,
		"LUADNS.COM":                   true,
		"M247.COM":                     true,
		"M247.COM 骨干网":                 true,
		"MAINLOOP.SE":                  true,
		"MANAGEDWAY.COM":               true,
		"MASERGY.COM":                  true,
		"MASERGY 骨干网":                  true,
		"MASTERBASE.COM":               true,
		"MAVIDENIZ.COM.TR":             true,
		"MAXCDN.COM":                   true,
		"MAYERBROWN.COM":               true,
		"MCPROHOSTING.COM":             true,
		"M-D.NET":                      true,
		"MEDIANETWORKSERVICES.COM":     true,
		"MEDIANETWORKSERVICES.COM 骨干网": true,
		"MELBICOM.NET":                 true,
		"MERAKI.COM":                   true,
		"MESH.EU 骨干网":                  true,
		"MICHAEL-KEHOE":                true,
		"MICROSOFT.COM":                true,
		"MICROSOFT 骨干网":                true,
		"MIMECAST.COM":                 true,
		"MIRROR-IMAGE.COM":             true,
		"MISAKA.IO":                    true,
		"MISAKA.IO 骨干网":                true,
		"MOJOHOST.COM":                 true,
		"MONOSOLUTIONS.COM":            true,
		"M-ROOT-SERVERS":               true,
		"MTC 骨干网":                      true,
		"MTN 骨干网":                      true,
		"MTS.RU 骨干网":                   true,
		"MYNIC.MY":                     true,
		"NEOCITIES.ORG":                true,
		"NEOLOGY.CO.ZA":                true,
		"Neotelecoms 骨干网":              true,
		"NETAPPLICATIONS.COM":          true,
		"NETASSIST.UA 骨干网":             true,
		"NETFLIX.COM":                  true,
		"NETNAMES.COM":                 true,
		"NETNOD.SE":                    true,
		"NETREGISTRY.COM.AU":           true,
		"NETRIPLEX.COM":                true,
		"NEWWORLDTEL 骨干网":              true,
		"NEXCESS.NET":                  true,
		"NEXTHOP.NO":                   true,
		"NEXUSGUARD.COM":               true,
		"NEXUSGUARD.COM 骨干网":           true,
		"NIC.AT":                       true,
		"NIC.CH":                       true,
		"NIC.CL":                       true,
		"NIC.CZ":                       true,
		"NIC.DK":                       true,
		"NIC.FR":                       true,
		"NIC.MX":                       true,
		"NIKE.COM":                     true,
		"NJAL.LA":                      true,
		"NORDU.NET 骨干网":                true,
		"NPIX.NET.NP":                  true,
		"NS1.COM":                      true,
		"NSOF.IO":                      true,
		"NTT.COM":                      true,
		"NTT.COM 骨干网":                  true,
		"NUBELA.CO":                    true,
		"NUMERICABLE.FR 骨干网":           true,
		"O3B-NETWORKS.COM":             true,
		"OMNIS.COM":                    true,
		"ONAPP.COM":                    true,
		"OPENBSD.ORG.IL":               true,
		"OPENDNS.COM":                  true,
		"OPTICFUSION.NET 骨干网":          true,
		"ORACLE.COM 骨干网":               true,
		"ORANGE.COM 骨干网":               true,
		"OVH.COM":                      true,
		"OVH.COM 骨干网":                  true,
		"PACKET.NET":                   true,
		"PACNET.COM":                   true,
		"PACNET.COM 骨干网":               true,
		"PAGESAT":                      true,
		"PCCW 骨干网":                     true,
		"PCH.NET":                      true,
		"PCH.NET 骨干网":                  true,
		"PEAK10.COM":                   true,
		"PEARSON.COM":                  true,
		"PEER1 骨干网":                    true,
		"PHOTONVPS.COM":                true,
		"PLANETHOSTER.NET":             true,
		"PLANISYS.COM":                 true,
		"PORTLANE 骨干网":                 true,
		"PRAGER-IT.COM":                true,
		"PRETECS.COM":                  true,
		"PRIMUSTEL.CA 骨干网":             true,
		"PRODUBAN.COM":                 true,
		"PROFIHOST.COM":                true,
		"PROLEXIC.COM":                 true,
		"PROLEXIC.COM 骨干网":             true,
		"PROLEXIC 骨干网":                 true,
		"PSYCHZ.NET":                   true,
		"PUBLICDOMAINREGISTRY.COM":     true,
		"PULSEPOINT.COM":               true,
		"QRATOR.NET":                   true,
		"QSC.DE 骨干网":                   true,
		"QUAD9.NET":                    true,
		"QUADRANET.COM":                true,
		"QUANTCAST.COM":                true,
		"QWEST.NET 骨干网":                true,
		"QWEST.NET骨干网":                 true,
		"RAGE4.COM":                    true,
		"RASCOM.RU 骨干网":                true,
		"RCN 骨干网":                      true,
		"REDCLARA 骨干网":                 true,
		"REFLECTED.NET":                true,
		"RIM 骨干网":                      true,
		"RINGCENTRAL.COM":              true,
		"RINGCENTRAL.COM 骨干网":          true,
		"RIOTGAMES 骨干网":                true,
		"RIPE.NET":                     true,
		"RIU.EDU.AR":                   true,
		"ROBTEX.COM":                   true,
		"ROLLERNET.US":                 true,
		"RT.RU 骨干网":                    true,
		"SAFEDNS.COM":                  true,
		"SAFEVALUE.PRO":                true,
		"SAGLAYICI.COM":                true,
		"SALESFORCE 骨干网":               true,
		"SAMIR.CA":                     true,
		"SAVVIS.NET 骨干网":               true,
		"SCIP.ES":                      true,
		"SDNS.CN":                      true,
		"SEACOM.MU 骨干网":                true,
		"SECTION.IO":                   true,
		"SECUREDSERVERS.COM":           true,
		"SEEWEB":                       true,
		"SELECTEL.RU":                  true,
		"SERV-AC.COM":                  true,
		"SERVERS.COM":                  true,
		"SFERA.NET":                    true,
		"SG.GS":                        true,
		"SG.GS 骨干网":                    true,
		"SHOPIFY.COM":                  true,
		"SIEL.SI":                      true,
		"SIFY.COM 骨干网":                 true,
		"SINGTEL.COM 骨干网":              true,
		"SITA.AERO 骨干网":                true,
		"SKHOSTING.EU":                 true,
		"SKYEBYNOMINUM.COM":            true,
		"SKYPARKCDN.RU":                true,
		"SKYVISION.NET 骨干网":            true,
		"SMC.LI":                       true,
		"SMULE.COM":                    true,
		"SOFTLAYER.COM":                true,
		"SOFTLAYER.COM 骨干网":            true,
		"SPACEDUMP.NET":                true,
		"SpectrumNet 骨干网":              true,
		"SPECTRUMNET 骨干网":              true,
		"SPIRITCOM.COM 骨干网":            true,
		"SPRINT.COM":                   true,
		"SQUIZ.NET":                    true,
		"STACKPATH.COM":                true,
		"STAMINUS.NET":                 true,
		"STAMINUS.NET 骨干网":             true,
		"STARRYDNS.COM":                true,
		"STREAMING.CO.NZ":              true,
		"SUCURI.NET":                   true,
		"SUNGARDRS.COM 骨干网":            true,
		"SUNRISE 骨干网":                  true,
		"SUPREMEBYTES.COM":             true,
		"SUPREMEBYTES.COM 骨干网":         true,
		"SURF.NL":                      true,
		"SYMANTEC.COM":                 true,
		"SYNAPSEGLOBAL.COM":            true,
		"SYNOPSYS.COM":                 true,
		"TAGHOS.COM.BR":                true,
		"TALAMUS.COM":                  true,
		"TATACOMMUNICATIONS.COM 骨干网":   true,
		"TDC.DK 骨干网":                   true,
		"TELE2.COM 骨干网":                true,
		"TELECOMITALIA.SM":             true,
		"TELECOM.KZ":                   true,
		"TELECOM.KZ 骨干网":               true,
		"TELEFONICA.COM 骨干网":           true,
		"TELEKOM.DE 骨干网":               true,
		"TELENOR.COM 骨干网":              true,
		"TELENOR 骨干网":                  true,
		"TELETOK.TK":                   true,
		"TELIA.COM 骨干网":                true,
		"TELIN 骨干网":                    true,
		"TELKOM.CO.ZA 骨干网":             true,
		"TELSTRA.COM 骨干网":              true,
		"TENCENT.COM":                  true,
		"TERAGO 骨干网":                   true,
		"TERRA.COM.BR":                 true,
		"TERREMARK.COM 骨干网":            true,
		"THREATSTOP.COM":               true,
		"TIBUS 骨干网":                    true,
		"TIERPOINT.COM":                true,
		"TIGGEE.COM":                   true,
		"TISPARKLE.COM 骨干网":            true,
		"TMNET 骨干网":                    true,
		"T-MOBILE.CZ 骨干网":              true,
		"TMXATRIUM.COM 骨干网":            true,
		"TOPNET.UA 骨干网":                true,
		"TOTALUPTIME.COM":              true,
		"TRANSTELCO.NET 骨干网":           true,
		"TRANSTELECOM 骨干网":             true,
		"TRENDMICRO.COM":               true,
		"TTK.RU 骨干网":                   true,
		"TURKTELEKOM.COM.TR 骨干网":       true,
		"TWCC.COM 骨干网":                 true,
		"TWGATE.NET 骨干网":               true,
		"TWITCH.TV":                    true,
		"TWITTER":                      true,
		"TWITTER 骨干网":                  true,
		"TYPCN.COM 骨干网":                true,
		"UCDN.COM":                     true,
		"UGO.COMPANY":                  true,
		"ULTRADNS.COM":                 true,
		"UNDERNET.ORG":                 true,
		"UNIVIE.AC.AT":                 true,
		"USSIGNAL.COM 骨干网":             true,
		"UUNET 骨干网":                    true,
		"UWORLD.SE":                    true,
		"VALVESOFTWARE":                true,
		"VELDER.LI":                    true,
		"VELLANCE.COM":                 true,
		"VERISIGN.COM":                 true,
		"VERISIGN.COM 骨干网":             true,
		"VERIZON.COM":                  true,
		"VERIZON.COM 骨干网":              true,
		"VIAWEST.COM":                  true,
		"VIAWEST.COM 骨干网":              true,
		"VIDILLION.COM":                true,
		"VIKINGHOST.COM":               true,
		"VIRTELA 骨干网":                  true,
		"VIRTUASERVER.COM.BR":          true,
		"VIRTUTEL.COM.AU":              true,
		"VIRTUTEL.COM.AU 骨干网":          true,
		"VITALWERKS.COM":               true,
		"VOCUS.COM.AU 骨干网":             true,
		"VODAFONE.COM 骨干网":             true,
		"VOXEL.NET":                    true,
		"VOXEL.NET 骨干网":                true,
		"VOXILITY.COM":                 true,
		"VOXILITY.COM 骨干网":             true,
		"VPSQUAN.COM":                  true,
		"VULTR.COM":                    true,
		"WANWIRE.COM":                  true,
		"WEB2OBJECTS.COM":              true,
		"WEBAIR.COM":                   true,
		"WEBEX.COM 骨干网":                true,
		"WINKSTREAMING.COM":            true,
		"WIPRO.COM":                    true,
		"WIPRO 骨干网":                    true,
		"WOLFE 骨干网":                    true,
		"WORLDNS.NET":                  true,
		"X4B.NET":                      true,
		"XCONNECT24 骨干网":               true,
		"XTCN.COM":                     true,
		"XTOM.COM":                     true,
		"YAHOO.COM":                    true,
		"YAHOO.COM 骨干网":                true,
		"YOUTUBE.COM":                  true,
		"YUN-IDC.COM":                  true,
		"YUN-IDC.COM 骨干网":              true,
		"ZAYO.COM 骨干网":                 true,
		"ZILORE":                       true,
		"ZORINS.CO.UK":                 true,
		"ZSCALER.COM":                  true,
		"ZVONKOVA":                     true,
		"ZXITY.UK":                     true,
		"亚太地区":                         true,
		"保留地址":                         true,
		"共享地址":                         true,
		"北美地区":                         true,
		"局域网":                          true,
		"拉美地区":                         true,
		"本地链路":                         true,
		"本机地址":                         true,
		"欧盟":                           true,
		"科索沃":                          true, // XK
		"非洲地区":                         true,
	}
)
