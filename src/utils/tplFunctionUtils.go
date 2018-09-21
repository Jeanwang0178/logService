package utils

func GetMapValue(key string, contentMap map[string]string) string {
	return contentMap[key]
}

func GetSliceMapValue(key string, contentMap map[string]interface{}) string {

	var vdata = contentMap[key]
	return vdata.(string)
}
