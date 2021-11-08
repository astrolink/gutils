package general

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// ReadConfigJson Reads Settings file.
func ReadConfigJson(configStruct interface{}, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configStruct)
	if err != nil {
		return err
	}

	return nil
}

// RemoveDuplicate remove duplicate entry in string array
func RemoveDuplicate(slice []string, verbose bool) []string {
	now := time.Now()
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	if verbose == true {
		elapsed := time.Now().Sub(now)
		log.Println("[removeDuplicate] total run time: ", fmt.Sprint(elapsed))
	}
	return list
}

// ConvertIntInterfaceToString tries to convert an interface value to string
func ConvertIntInterfaceToString(value interface{}) (string, error) {
	var err error
	var stringResult string
	var intResult int
	var longResult int64

	switch i := value.(type) {
	case int:
		intResult, _ = value.(int)
		stringResult = strconv.Itoa(intResult)
		return stringResult, nil
	case int64:
		longResult, _ = value.(int64)
		return strconv.FormatInt(longResult, 10), nil
	case float32:
		return strconv.FormatInt(int64(i), 10), nil
	case float64:
		return strconv.FormatInt(int64(i), 10), nil
	}

	longResult, err = strconv.ParseInt(fmt.Sprint(value), 10, 64)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return strconv.FormatInt(longResult, 10), nil
}

// ConvertInterfaceToInt64 converts an interface value to int64
func ConvertInterfaceToInt64(value interface{}) int64 {
	switch i := value.(type) {
	case int:
		return int64(i)
	case int64:
		return value.(int64)
	case string:
		long, _ := strconv.ParseInt(fmt.Sprint(value), 10, 64)
		return long
	case float32:
		return int64(value.(float32))
	case float64:
		return int64(value.(float64))
	}

	return 0
}

// InArray checks if a certain value exists inside the slice/array
func InArray(value interface{}, array interface{}) (bool, int) {
	exists := false
	index := -1

	if reflect.TypeOf(array).Kind() == reflect.Slice {
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				index = i
				exists = true
				return exists, index
			}
		}
	}
	return exists, index
}

// EncodeToLatin1 encode utf8 to ISO8859_9 (latin)
func EncodeToLatin1(utf8Text string) string {
	b := []byte(utf8Text)
	encoded2, _ := charmap.ISO8859_9.NewDecoder().Bytes(b)
	return string(encoded2[:])
}

// EncodeToUtf8 encode ISO8859_9 (latin) to utf8
func EncodeToUtf8(latin1Text string) string {
	r := []byte(latin1Text)
	encoded, _ := charmap.ISO8859_9.NewEncoder().Bytes(r)
	return string(encoded[:])
}

// ReplaceGender replace text infos to custom
func ReplaceGender(gender, username, text string) string {
	newText := strings.ReplaceAll(text, "{user}", username)

	if gender == "M" {
		return strings.ReplaceAll(newText, "{x}", "o")
	}

	return strings.ReplaceAll(newText, "{x}", "a")
}

// FormatDateTimeByLanguage formats dates based on provided language
func FormatDateTimeByLanguage(value, language string) (string, string) {
	var date string
	dateTime := strings.Split(value, " ")
	splitDate := strings.Split(dateTime[0], "-")
	year := splitDate[0]
	month := splitDate[1]
	day := splitDate[2]
	date = fmt.Sprintf("%s/%s/%s", day, month, year)

	if language == "en_us" {
		date = fmt.Sprintf("%s/%s/%s", month, day, year)
	}

	hours := strings.Split(dateTime[1], ":")
	hour := hours[0]
	minutes := hours[1]

	completeHour := fmt.Sprintf("%s:%s", hour, minutes)
	return date, completeHour
}

// UserNeverSigned return true if usu_assinante_datainicio contains a date
func UserNeverSigned(subscribeDate interface{}) bool {
	if subscribeDate == nil {
		return true
	}
	_, err := time.Parse("2006-01-02", subscribeDate.(string))
	if err != nil {
		return  true
	}
	return false
}

// RemoveAccentuation removes any accentuation from a string
func RemoveAccentuation(s string) string {
	if s == "" {
		return s
	}

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ = transform.String(t, s)
	return s
}

// ConvertToCurrency convert price interface to string
func ConvertToCurrency(price interface{}) string {
	intPrice := ConvertInterfaceToInt64(price)
	if intPrice == 0 {
		return "0"
	}

	val := strconv.FormatInt(intPrice, 10)
	charCount := utf8.RuneCountInString(val)

	substr1 := val[0:(charCount - 2)]
	if substr1 == "" {
		substr1 = "0"
	}

	substr2 := val[(charCount - 2):charCount]
	return fmt.Sprintf("%s,%s", substr1, substr2)
}

// ConvertCurrencyToFloat64 convert price interface to float64
func ConvertCurrencyToFloat64(price interface{}) float64 {
	strPrice := ConvertToCurrency(price)
	strPrice = strings.ReplaceAll(strPrice, ",", ".")
	f, _ := strconv.ParseFloat(strPrice, 64)
	return f
}