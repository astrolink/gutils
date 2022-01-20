package language

import (
	"fmt"
	"github.com/astrolink/gutils/general"
	"strconv"
	"strings"
)

const (
	PtBr = "pt_br"
	Es   = "es"
	EnUs = "en_us"
)

func getAllowedLangs() []string {
	return []string{EnUs, Es, PtBr}
}

var translationKeys map[string]string
var contexts []string

// LoadLang carrega em memória o arquivo de tradução para determinado contexto e idioma
func LoadLang(lang map[string]map[string]string, context string, idiom string) {
	contextIdiom := context + "_" + idiom

	testInArray, _ := general.InArray(contextIdiom, contexts)

	if testInArray == true {
		return
	}

	if len(translationKeys) == 0 {
		translationKeys = make(map[string]string)
	}

	if val, ok := lang[idiom]; ok {
		for key, value := range val {
			translationKeys[key+"_"+idiom] = value
		}
	}

	contexts = append(contexts, contextIdiom)
}

// Translate substitui a chave de idiomas pelo seu valor correspondente definido
// no arquivo de tradução
func Translate(line string, idiom string, replacements []string) string {
	value := ""

	if val, ok := translationKeys[line+"_"+idiom]; ok {
		value = val
	}

	if value == "" || replacements == nil {
		return value
	}

	for index, replace := range replacements {
		value = strings.Replace(value, "{"+strconv.Itoa(index)+"}", replace, 1)
	}

	return value
}

// ReplaceI18nQueries funcao responsavel por substituir #i18n nas queries
// exemplo: SELECT fee_tag_adjetivo_m#i18n AS feeling_male FROM feelings_tags
func ReplaceI18nQueries(i18nLang, query string) string {
	switch i18nLang {
	case "en_us":
		query = strings.ReplaceAll(query, "#i18n", "_en_us")
		break
	case "es":
		query = strings.ReplaceAll(query, "#i18n", "_es")
		break
	default:
		query = strings.ReplaceAll(query, "#i18n", "")
		break
	}

	return query
}

// SetLang atribui valor à propriedade de idiomas do repositório
func SetLang(lang *string, langArg string) {
	if langArg == PtBr {
		return
	}

	allowedLangs := getAllowedLangs()
	if ok, _ := general.InArray(langArg, allowedLangs); ok {
		*lang = fmt.Sprintf("_%s", langArg)
	}
}