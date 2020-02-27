package render

import (
	"strings"
	"time"

	"github.com/alecthomas/template"
)

func GetFuncMap() template.FuncMap {

	var funcMap = make(template.FuncMap)
	funcMap["Split"] = func(s string, d string) []string {

		return strings.Split(s, d)
	}
	funcMap["Join"] = func(a []string, b string) string {
		return strings.Join(a, b)
	}
	//funcMap["ParseFlashes"] = func(fucks []interface{}) []_struct.Flash {
	//		var flashes []_struct.Flash
	//		for _, k := range fucks {
	//				var flash _struct.Flash
	//				err := json.Unmarshal([]byte(k.(string)), &flash)
	//				if err != nil {
	//						log.Panic(err)
	//				}
	//				flashes = append(flashes, flash)
	//		}
	//		return flashes
	//}

	funcMap["formatPostTime"] = func(t time.Time) string {
		return t.Format(time.UnixDate)
	}

	funcMap["formatTitle"] = func(s string) string {
		title := strings.SplitAfter(s, "/")
		return strings.Title(title[1])
	}
	return funcMap
}
