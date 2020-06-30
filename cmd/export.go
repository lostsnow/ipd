package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "ipd gen",
	Long:  `ipd gen`,
	Run: func(cmd *cobra.Command, args []string) {
		// http://www.mca.gov.cn/article/sj/xzqh/2018/
		// http://www.mca.gov.cn/article/sj/xzqh/2018/201804-12/201804121005.html
		f, err := os.Open("./data/region.source.txt")
		if err != nil {
			fmt.Printf("Open region file error: %s\n", err)
			return
		}
		defer f.Close()

		lastLvl := -1
		sCode := ""
		var states []string
		buf := bufio.NewReader(f)
		for {
			line, err := buf.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Printf("Read error: %s\n", err)
				return
			}

			line = strings.TrimSpace(line)
			var c int
			if line == "" {
				continue
			} else if line[2:6] == "0000" {
				c = 1
			} else if line[4:6] == "00" {
				c = 2
			} else {
				c = 3
			}

			curState := line[:2]
			// Beijing/Tianjin/Shanghai/Chongqing/Taiwan/Hongkong/Macau
			if curState == "11" || curState == "12" || curState == "31" || curState == "50" ||
				curState == "71" || curState == "81" || curState == "82" {
				continue
			}

			if c == 1 {
				sCode = line[:2]
				states = append(states, fmt.Sprintf(`"CN%s": CitiesCN%s,`, sCode, sCode))
				if lastLvl != -1 {
					fmt.Printf("\t}\n\n")
				}
				fmt.Printf(`%sCitiesCN%s = []City{%s`, "\t", sCode, "\n")
			} else {
				if sCode != "46" && c == 3 {
					continue
				}

				city := strings.Fields(line)
				cc := city[1]
				// 吉林省 吉林市
				if cc != "吉林市" {
					cc = strings.TrimRight(cc, "市")
				}

				fmt.Printf(`%s{Code: "CN%s", Title: "%s"},%s`, "\t\t", city[0], cc, "\n")
				// 新疆
				if line[:4] == "6543" {
					fmt.Printf(`%s{Code: "CN659001", Title: "石河子"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN659002", Title: "阿拉尔"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN659003", Title: "图木舒克"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN659004", Title: "五家渠"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN659005", Title: "北屯"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN659006", Title: "铁门关"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN659007", Title: "双河"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN659008", Title: "可克达拉"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN659009", Title: "昆玉"},%s`, "\t\t", "\n")
				}
				// 湖北
				if line[:4] == "4228" {
					fmt.Printf(`%s{Code: "CN429004", Title: "仙桃"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN429005", Title: "潜江"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN429006", Title: "天门"},%s`, "\t\t", "\n")
					fmt.Printf(`%s{Code: "CN429021", Title: "神农架林区"},%s`, "\t\t", "\n")
				}
				// 内蒙古
				if line[:4] == "1522" {
					fmt.Printf(`%s{Code: "CN152201", Title: "乌兰浩特"},%s`, "\t\t", "\n")
				}
				// 河南
				if line[:4] == "4117" {
					fmt.Printf(`%s{Code: "CN419001", Title: "济源"},%s`, "\t\t", "\n")
				}
			}
			lastLvl = c
		}
		fmt.Printf("\t}\n\n")

		for _, v := range states {
			fmt.Printf("%s\n", v)
		}
	},
}

func init() {
	RootCmd.AddCommand(genCmd)
}
