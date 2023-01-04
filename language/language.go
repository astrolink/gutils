package language

import (
	"fmt"
	"github.com/astrolink/gutils/general"
	"strconv"
	"strings"
)

const (
	PtBr     = "pt_br"
	Es       = "es"
	EnUs     = "en_us"
	Fr       = "fr"
	It       = "it"
	De       = "de"
	I18nHash = "#i18n"
)

func getAllowedLangs() []string {
	return []string{PtBr, Es, EnUs, Fr, It, De}
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

// ReplaceI18nQueries função responsável por substituir #i18n nas queries
// exemplo: SELECT fee_tag_adjetivo_m#i18n AS feeling_male FROM feelings_tags
func ReplaceI18nQueries(i18nLang, query string) string {
	switch i18nLang {
	case EnUs:
		query = strings.ReplaceAll(query, I18nHash, "_en_us")
		break
	case Es:
		query = strings.ReplaceAll(query, I18nHash, "_es")
		break
	case Fr:
		query = strings.ReplaceAll(query, I18nHash, "_fr")
		break
	case It:
		query = strings.ReplaceAll(query, I18nHash, "_it")
		break
	case De:
		query = strings.ReplaceAll(query, I18nHash, "_de")
		break
	default:
		query = strings.ReplaceAll(query, I18nHash, "")
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

// SetI18nQueryFields substitui o hash #i18n pelo idioma correspondente configurado
// no atributo lang do repositório
func SetI18nQueryFields(lang, query string) string {
	return strings.ReplaceAll(query, I18nHash, lang)
}
