package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type alert struct {
	Code  string
	Level string
}

var alerts map[string] /*code*/ alert

func getAlertInHtml() map[string]alert {
	file, err := os.Open("/tmp/code.html")
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		panic(err)
	}
	alertList := make(map[string]alert, 0)

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		alert := alert{}
		s.Find("td").Each(func(i int, innerS *goquery.Selection) {
			if i == 1 { //code is in second <td>
				code := strings.TrimSpace(innerS.Text())
				alert.Code = code
			}
			if i == 2 { //code is in third <td>
				level := strings.TrimSpace(innerS.Text())
				alert.Level = level
			}
		})
		if alert.Code != "" {
			//alertList = append(alertList, alert)
			alertList[alert.Code] = alert
		}
	})
	return alertList
}

func getAlertDesc(code string) *alert {
	if a, ok := alerts[code]; ok {
		return &alert{
			Code:  a.Code,
			Level: a.Level,
		}
	}
	return nil
}

func parseAlertJson(bs []byte) {
	if err := json.Unmarshal(bs, &alerts); err != nil {
		panic(err)
	}
}

func main() {
	alertList := getAlertInHtml()
	bs, err := json.MarshalIndent(alertList, "", "    ")
	if err != nil {
		panic(err)
	}
	parseAlertJson(bs)
	for _, alert := range alertList {
		fmt.Printf(`%s Alert= Alert{
			code:  "%s",
			level: level_%s,
		   reference: referencePrefix+codeToAnchor("%s"),
		}`, strings.ToUpper(alert.Code), strings.ToUpper(alert.Code), alert.Level, strings.ToUpper(alert.Code))

		//fmt.Printf(`"%s": %s,`, alert.Code, alert.Code)
		////fmt.Printf(`code_%v code="%v"`, alert.Code, alert.Code)
		fmt.Println()
	}
	//if a := getAlertDesc("BIND_SIP_FAILED"); a != nil {
	//	fmt.Println(a.Level)
	//	a.Level = "test"
	//}
	//if a := getAlertDesc("BIND_SIP_FAILED"); a != nil {
	//	fmt.Println(a.Level)
	//}
}
