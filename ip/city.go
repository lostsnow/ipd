package ip

var (
	cityCodes = map[string]map[string]map[string]string{
		"CN": cityCodesCN,
	}

	cityCodesCN = map[string]map[string]string{
		"11": cityCodesCN11,
	}

	cityCodesCN11 = map[string]string{}
)

func getCityCode(c, s, cc string) string {
	if v, ok := cityCodes[c][s][cc]; ok {
		return v
	}
	return ""
}
