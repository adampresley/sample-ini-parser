package ini

type IniFile struct {
	FileName string       `json:"fileName"`
	Sections []IniSection `json:"sections"`
}
