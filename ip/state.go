package ip

type State struct {
	Title     string
	Code      string
	Cities    []City
	CityCodes map[string]City
}

func init() {
	for _, v := range allStatesCN {
		stateCodesCN[v.Title] = v
	}
}

func (s *State) GetCites() []City {
	return s.Cities
}

func (s *State) GetCity(c string) *City {
	if v, ok := cityCodes[s.Code][c]; ok {
		return &v
	}
	return nil
}

var (
	stateCodesCN = map[string]State{}

	// <editor-fold desc="all states">
	StatesCN11 = State{Code: "CN11", Title: "北京"}
	StatesCN12 = State{Code: "CN12", Title: "天津"}
	StatesCN13 = State{Code: "CN13", Title: "河北", Cities: CitiesCN13}
	StatesCN14 = State{Code: "CN14", Title: "山西"}
	StatesCN15 = State{Code: "CN15", Title: "内蒙古", Cities: CitiesCN15}
	StatesCN21 = State{Code: "CN21", Title: "辽宁"}
	StatesCN22 = State{Code: "CN22", Title: "吉林"}
	StatesCN23 = State{Code: "CN23", Title: "黑龙江"}
	StatesCN31 = State{Code: "CN31", Title: "上海"}
	StatesCN32 = State{Code: "CN32", Title: "江苏"}
	StatesCN33 = State{Code: "CN33", Title: "浙江"}
	StatesCN34 = State{Code: "CN34", Title: "安徽"}
	StatesCN35 = State{Code: "CN35", Title: "福建"}
	StatesCN36 = State{Code: "CN36", Title: "江西"}
	StatesCN37 = State{Code: "CN37", Title: "山东"}
	StatesCN41 = State{Code: "CN41", Title: "河南"}
	StatesCN42 = State{Code: "CN42", Title: "湖北"}
	StatesCN43 = State{Code: "CN43", Title: "湖南"}
	StatesCN44 = State{Code: "CN44", Title: "广东"}
	StatesCN45 = State{Code: "CN45", Title: "广西"}
	StatesCN46 = State{Code: "CN46", Title: "海南"}
	StatesCN50 = State{Code: "CN50", Title: "重庆"}
	StatesCN51 = State{Code: "CN51", Title: "四川"}
	StatesCN52 = State{Code: "CN52", Title: "贵州"}
	StatesCN53 = State{Code: "CN53", Title: "云南", Cities: CitiesCN53, CityCodes: cityCodes["CN53"]}
	StatesCN54 = State{Code: "CN54", Title: "西藏"}
	StatesCN61 = State{Code: "CN61", Title: "陕西"}
	StatesCN62 = State{Code: "CN62", Title: "甘肃"}
	StatesCN63 = State{Code: "CN63", Title: "青海"}
	StatesCN64 = State{Code: "CN64", Title: "宁夏"}
	StatesCN65 = State{Code: "CN65", Title: "新疆"}
	StatesCN71 = State{Code: "CN71", Title: "台湾"}
	StatesCN81 = State{Code: "CN81", Title: "香港"}
	StatesCN82 = State{Code: "CN82", Title: "澳门"}
	// </editor-fold>

	allStatesCN = []State{
		StatesCN11, StatesCN12, StatesCN13, StatesCN14, StatesCN15,
		StatesCN21, StatesCN22, StatesCN23,
		StatesCN31, StatesCN32, StatesCN33, StatesCN34, StatesCN35, StatesCN36, StatesCN37,
		StatesCN41, StatesCN42, StatesCN43, StatesCN44, StatesCN45, StatesCN46,
		StatesCN50, StatesCN51, StatesCN52, StatesCN53, StatesCN54,
		StatesCN61, StatesCN62, StatesCN63, StatesCN64, StatesCN65,
		StatesCN71,
		StatesCN81,
		StatesCN82,
	}
)
