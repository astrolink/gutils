package time

import (
	"github.com/astrolink/gutils/globalization"
	"log"
	"strings"
	"time"
)

//GetTimeNowString return the date now formatted in time zone
func GetTimeNowString(format, timeZone string, onlyDate bool) (string, error) {
	timeNow := time.Now()

	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return "", err
	}
	timeNow = timeNow.In(loc)

	if onlyDate {
		timeNow = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
	}

	timeString := timeNow.Format(format)
	return timeString, nil
}

func GetTimeNow(timeZone string) (time.Time, error) {
	timeNow := time.Now()

	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return time.Time{}, err
	}
	timeNow = timeNow.In(loc)

	return timeNow, nil
}

func GetLocalizedTimeNowByCountryCode(countryCode string) (time.Time, error) {
	now := time.Now()
	countryTimeZone := getCountryTimeZone(countryCode)
	now, err := GetTimeNow(countryTimeZone)

	if err != nil {
		log.Println(err)
	}

	return now, err
}

func GetStringLocalizedTimeNowByCountryCode(countryCode string) (string, error) {
	countryTimeZone := getCountryTimeZone(countryCode)
	localizedTime, err := GetTimeNowString(DateTimeFormat, countryTimeZone, false)

	if err != nil {
		log.Println(err)
	}

	return localizedTime, err
}

func getCountryTimeZone(countryCode string) string {
	var countryTimeZone string

	switch strings.ToLower(countryCode) {
	case globalization.ArgentinaCountryCode:
		countryTimeZone = globalization.AmericaBuenosAiresTimeZone
		break
	case globalization.ChileCountryCode:
		countryTimeZone = globalization.AmericaSantiagoTimeZone
		break
	case globalization.BoliviaCountryCode:
		countryTimeZone = globalization.AmericaLaPazTimeZone
		break
	case globalization.ColombiaCountryCode:
		countryTimeZone = globalization.AmericaBogotaTimeZone
		break
	case globalization.CostaRicaCountryCode:
		countryTimeZone = globalization.AmericaCostaRicaTimeZone
		break
	case globalization.CubaCountryCode:
		countryTimeZone = globalization.AmericaHavanaTimeZone
		break
	case globalization.DominicanRepublicCountryCode:
		countryTimeZone = globalization.AmericaSantoDomingoTimeZone
		break
	case globalization.EcuadorCountryCode:
		countryTimeZone = globalization.AmericaGuayaquilTimeZone
		break
	case globalization.ElSalvadorCountryCode:
		countryTimeZone = globalization.AmericaElSalvadorTimeZone
		break
	case globalization.GuatemalaCountryCode:
		countryTimeZone = globalization.AmericaGuatemalaTimeZone
		break
	case globalization.HondurasCountryCode:
		countryTimeZone = globalization.AmericaTegucigalpaTimeZone
		break
	case globalization.MexicoCountryCode:
		countryTimeZone = globalization.AmericaMexicoCityTimeZone
		break
	case globalization.PanamaCountryCode:
		countryTimeZone = globalization.AmericaPanamaTimeZone
		break
	case globalization.ParaguayCountryCode:
		countryTimeZone = globalization.AmericaAsuncionTimeZone
		break
	case globalization.PeruCountryCode:
		countryTimeZone = globalization.AmericaLimaTimeZone
		break
	case globalization.PortugalCountryCode:
		countryTimeZone = globalization.EuropeLisbon
		break
	case globalization.SpainCountryCode:
		countryTimeZone = globalization.EuropeMadrid
		break
	case globalization.UruguayCountryCode:
		countryTimeZone = globalization.AmericaMontevideoTimeZone
		break
	case globalization.UnitedStatesCountryCode:
		countryTimeZone = globalization.AmericaLosAngeles
		break
	case globalization.VenezuelaCountryCode:
		countryTimeZone = globalization.AmericaCaracasTimeZone
		break
	default:
		countryTimeZone = globalization.AmericaSaoPauloTimeZone
	}

	return countryTimeZone
}

func GetGmtByCountryCode(countryCode string) int {
	var gmt int

	switch strings.ToLower(countryCode) {
	case globalization.ArgentinaCountryCode:
		gmt = globalization.MINUS_THREE_GMT
		break
	case globalization.ChileCountryCode:
		gmt = globalization.MINUS_THREE_GMT
		break
	case globalization.BoliviaCountryCode:
		gmt = globalization.MINUS_FOUR_GMT
		break
	case globalization.ColombiaCountryCode:
		gmt = globalization.MINUS_FIVE_GMT
		break
	case globalization.CostaRicaCountryCode:
		gmt = globalization.MINUS_SIX_GMT
		break
	case globalization.CubaCountryCode:
		gmt = globalization.MINUS_FIVE_GMT
		break
	case globalization.DominicanRepublicCountryCode:
		gmt = globalization.MINUS_FOUR_GMT
		break
	case globalization.EcuadorCountryCode:
		gmt = globalization.MINUS_FIVE_GMT
		break
	case globalization.ElSalvadorCountryCode:
		gmt = globalization.MINUS_SIX_GMT
		break
	case globalization.GuatemalaCountryCode:
		gmt = globalization.MINUS_SIX_GMT
		break
	case globalization.HondurasCountryCode:
		gmt = globalization.MINUS_SIX_GMT
		break
	case globalization.MexicoCountryCode:
		gmt = globalization.MINUS_SIX_GMT
		break
	case globalization.PanamaCountryCode:
		gmt = globalization.MINUS_FIVE_GMT
		break
	case globalization.ParaguayCountryCode:
		gmt = globalization.MINUS_THREE_GMT
		break
	case globalization.PeruCountryCode:
		gmt = globalization.MINUS_FIVE_GMT
		break
	case globalization.PortugalCountryCode:
		gmt = globalization.MINUS_ONE_GMT
		break
	case globalization.SpainCountryCode:
		gmt = globalization.PLUS_ONE_GMT
		break
	case globalization.UruguayCountryCode:
		gmt = globalization.MINUS_THREE_GMT
		break
	case globalization.UnitedStatesCountryCode:
		gmt = globalization.MINUS_EIGHT_GMT
		break
	case globalization.VenezuelaCountryCode:
		gmt = globalization.MINUS_FOUR_GMT
		break
	default:
		gmt = globalization.MINUS_THREE_GMT
	}

	return gmt
}