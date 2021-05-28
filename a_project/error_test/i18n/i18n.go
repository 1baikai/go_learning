package i18n

type Language string

const (
	// LangZh 中文
	LangZh Language = "zh"
	// LangEn 英文
	LangEn Language = "en"
	// LangJp 日文
	LangJp Language = "jp"
)

var LanguageMap = map[Language]string{
	LangZh: "中文",
	LangEn: "英文",
	LangJp: "日语",
}

func IsValid(language string) (Language, bool) {
	l := Language(language)
	if _, ok := LanguageMap[l]; ok {
		return l, true
	}
	return "", false
}
