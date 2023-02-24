package computations

//
//import (
//	"encoding/json"
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//	"time"
//)
//
//type HolidayResponse struct {
//	Response struct {
//		Holidays []struct {
//			Date struct {
//				ISO string `json:"iso"`
//			} `json:"date"`
//		} `json:"holidays"`
//	} `json:"response"`
//}
//
//func getResponseBodyForPublicHolidays(year time.Time) ([]byte, error) {
//	apiKey := "16d3e283e177e253f5639b0ac91bc315529b60d0"
//	url := fmt.Sprintf("https://calendarific.com/api/v2/holidays?api_key=%s&country=RO&year=%v", apiKey, year.Year())
//	response, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	contentBody, err := io.ReadAll(response.Body)
//
//	err = response.Body.Close()
//	if err != nil {
//		return nil, err
//	}
//	return contentBody, nil
//}
//
//func convertToJSON(year time.Time) ([]time.Time, error) {
//	body, err := getResponseBodyForPublicHolidays(year)
//	if err != nil {
//		return []time.Time{}, err
//	}
//	var publicHoliday HolidayResponse
//	err = json.Unmarshal(body, &publicHoliday)
//	if err != nil {
//		log.Print(err)
//		return []time.Time{}, err
//	}
//
//	var holidays []time.Time
//	for _, holiday := range publicHoliday.Response.Holidays {
//		date, err := time.Parse("2006-01-02T15:04:05-07:00", holiday.Date.ISO)
//		if err != nil {
//			// Try parsing date only
//			date, err = time.Parse("2006-01-02", holiday.Date.ISO[:10])
//			if err != nil {
//				log.Print(err)
//				continue
//			}
//		}
//		holidays = append(holidays, date)
//	}
//	return holidays, nil
//}
//
//func isRomanianPublicHoliday(day int, month time.Month, year time.Time) bool {
//	holidays, err := convertToJSON(year)
//	if err != nil {
//		return false
//	}
//	for _, holiday := range holidays {
//		if holiday.Day() == day && holiday.Month() == month {
//			return true
//		}
//	}
//	return false
//}
