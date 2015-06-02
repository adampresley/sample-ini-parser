package ini

type IniSection struct {
	Name          string        `json:"name"`
	KeyValuePairs []IniKeyValue `json:"keyValuePairs"`
}
