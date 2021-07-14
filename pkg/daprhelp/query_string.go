package daprhelp

import "strings"

type Values map[string]string

func GetQuery(queryString, key string) string {
	if values, ok := getQueryArray(initQueryCache(queryString), key); ok {
		return values
	}
	return ""
}

//  name=value&name2=value2
func initQueryCache(queryString string) Values {
	if queryString == "" {
		return nil
	}
	val := Values{}
	queryStrArr := strings.Split(queryString, "&")
	for _, s := range queryStrArr {
		keyValArr := strings.Split(s, "=")
		val[keyValArr[0]] = keyValArr[1]
	}

	return val
}
func getQueryArray(values Values, key string) (string, bool) {
	if values, ok := values[key]; ok && len(values) > 0 {
		return values, true
	}

	return "", false
}
