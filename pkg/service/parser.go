package service

import (
	"bufio"
	"fmt"
	webapi "github.com/4from5/TimeHack-webapi"
	"github.com/arran4/golang-ical"
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	_, err = io.Copy(out, resp.Body)
	return err
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Scraper(groupname webapi.Group) []webapi.Event {
	Schedule := make([]webapi.Event, 0)
	var event webapi.Event
	var link string
	loc, _ := time.LoadLocation("Europe/Moscow")

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		gr := []rune(e.Text) //[58:63]
		for i := 0; i < len(gr); i++ {
			if unicode.IsSpace(gr[i]) {
				copy(gr[i:], gr[i+1:])
				gr[len(gr)-1] = ' '
				gr = gr[:len(gr)-1]
				i--
			}
		}

		if string(gr) == groupname.GroupName {
			link = e.Attr("href")
			link = "https://lks.bmstu.ru" + link + ".ics"
			fmt.Println(link)
			if err := DownloadFile("schedule.ics", link); err != nil {
				fmt.Println("файл не грузится")
			}

			lines, err := readLines("schedule.ics")
			if err != nil {
				fmt.Println("не парсится")
			}

			for i := 0; i < len(lines); i++ {
				line2 := []rune(lines[i])
				for _, r := range line2 {
					if r == '\t' {
						l := []rune(lines[i-1])
						line2 = append(line2[1:])
						l = append(l, line2...)

						lines[i-1] = string(l)
						lines = append(lines[:i], lines[i+1:]...)
					}
				}
			}

			for i := 0; i < len(lines); i++ {
				l := ics.ContentLine(lines[i])
				pr, err := ics.ParseProperty(l)
				if err != nil {
					fmt.Println(lines[i])
					line := lines[i-1] + lines[i]
					pr, _ = ics.ParseProperty(ics.ContentLine(line))
					i++

				}

				if pr.IANAToken == "SUMMARY" {
					event.Title = pr.Value
					event.Description = ""
				} else if pr.IANAToken == "DTSTART" {
					layout := "20060102T150405Z"
					event.StartTimestamp, _ = time.Parse(layout, pr.Value)
					event.StartTimestamp = event.StartTimestamp.In(loc)
				} else if pr.IANAToken == "Location" {
					event.EventLocation = pr.Value
				} else if pr.IANAToken == "DESCRIPTION" {
					event.Description = pr.Value

				} else if pr.IANAToken == "DTEND" {
					layout := "20060102T150405Z"
					event.EndTimestamp, _ = time.Parse(layout, pr.Value)
					event.EndTimestamp = event.EndTimestamp.In(loc)

				} else if pr.IANAToken == "RRULE" {
					ss := strings.Split(pr.Value, ";")
					interval := strings.Split(ss[1], "=")
					event.RepeatPeriodDays, _ = strconv.Atoi(interval[1])
					event.RepeatPeriodDays *= 7

					endDay := strings.Split(ss[2], "=")
					layout := "20060102"
					event.EndPeriodTimestamp, _ = time.Parse(layout, endDay[1])
					event.EndPeriodTimestamp = event.EndPeriodTimestamp.In(loc)
					event.CategoryId = groupname.CategoryId
					Schedule = append(Schedule, event)

				}
			}
		}
	})
	err := c.Visit("https://lks.bmstu.ru/schedule/list")
	if err != nil {
		return nil
	}
	return Schedule
}
