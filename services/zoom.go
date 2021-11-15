package services

import (
	"encoding/json"
	"fmt"
	tm "github.com/buger/goterm"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Body struct {
	SubjectName   string
	CourseSubject struct {
		Timetables []struct {
			EndDate   int64
			StartDate int64
			WeekIndex int64
			Room      struct {
				Code string
			}
			EndHour struct {
				End int64
			}
			StartHour struct {
				Start int64
			}
		}
	}
}

var base = "/home/duckhue01/coding-data/pro/todos/"

type Data struct {
	SubjectName string
	EndDate     int64
	StartDate   int64
	Room        string
	EndHour     int64
	StartHour   int64
}

func UpdateZoomHandler(user, pass string) {
	token, err := auth(user, pass)
	if err != nil {
		log.Fatal("can not authentication")
	}
	body := getData(token)
	data := tranform(body)
	d, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(base, "zoomData.json"), d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func auth(user, pass string) (string, error) {

	res, err := http.PostForm("http://sinhvien.tlu.edu.vn:8099/education/oauth/token", url.Values{"username": {user}, "password": {pass}, "client_id": {"education_client"},
		"grant_type": {"password"}, "client_secret": {"password"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	raw, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 200 {
		type Body struct {
			Token string `json:"access_token"`
		}
		var body Body
		json.Unmarshal(raw, &body)
		return fmt.Sprintf("Bearer %s", body.Token), nil
	}
	return "", err

}

func getData(token string) []Body {
	var body []Body
	for i := 10; i >= 0; i-- {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", fmt.Sprintf("http://sinhvien.tlu.edu.vn:8099/education/api/StudentCourseSubject/studentLoginUser/%d", i), nil)
		req.Header.Add("Authorization", token)

		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		raw, _ := ioutil.ReadAll(res.Body)

		json.Unmarshal(raw, &body)
		if len(body) > 0 {
			break
		}
	}
	return body
}

func tranform(body []Body) [6][]Data {
	var data [6][]Data

	for i := 0; i < len(body); i++ {

		for a := 0; a < len(body[i].CourseSubject.Timetables); a++ {
			timeTable := body[i].CourseSubject.Timetables[a]
			if timeTable.EndDate > time.Now().UnixMilli() {
				data[timeTable.WeekIndex-1] = append(data[timeTable.WeekIndex-1], Data{
					SubjectName: body[i].SubjectName,
					EndDate:     timeTable.EndDate,
					StartDate:   timeTable.StartDate,
					Room:        timeTable.Room.Code,
					EndHour:     timeTable.EndHour.End,
					StartHour:   timeTable.StartHour.Start,
				})
			}

		}

	}

	return data
}

func StartZoomhandler(username string, password string, auto bool) {

	raw, err := ioutil.ReadFile(filepath.Join(base, "zoomData.json"))
	if err != nil {
		log.Fatal(err)
	}
	var data [][]Data

	err = json.Unmarshal(raw, &data)
	if err != nil {
		log.Fatal(err)
	}

	weekIndex := int(time.Now().Weekday())
	now := time.Now()
	var result Data
	isHave := false
	for i := 0; i < len(data[weekIndex]); i++ {

		// determine outdated subject
		if (data[weekIndex][i].EndDate - now.UnixMilli()) > 0 {
			start := data[weekIndex][i].StartHour % 86400000
			end := data[weekIndex][i].EndHour % 86400000
			i := 1
			tm.Clear()
			if auto {
				for {
					now := time.Now().UnixMilli() % 86400000
					if now > start && now < end {
						result = data[weekIndex][i]
						isHave = true
						break
					}
					tm.MoveCursor(1, 1)
					tm.Print("waiting for next class...")
					tm.Flush() // Call it every time at the end of rendering
					time.Sleep(time.Second)
				}
			} else {
				now := time.Now().UnixMilli() % 86400000
				if now > start && now < end {
					result = data[weekIndex][i]
					isHave = true
					break
				}
			}

		}
	}
	if isHave {

		raw, err := ioutil.ReadFile(filepath.Join(base, "zoomID.json"))
		if err != nil {
			log.Fatal(err)
		}
		var ID map[string]string

		err = json.Unmarshal(raw, &ID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("vao phong: %s(%s) voi ten: %s va password: %s...", result.Room, ID[result.Room], username, password)
		id := strings.Replace(ID[result.Room], " ", "", -1)
		url := fmt.Sprintf("zoommtg://zoom.us/join?action=join&confno=%s&pwd=%s&zc=0&uname=%s", id, username, password)
		open(url)
	} else {
		fmt.Println("you haven't had class yet!!")
	}
}

func open(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func CheckRoomHandler(code string) {
	raw, err := ioutil.ReadFile(filepath.Join(base, "zoomID.json"))
	if err != nil {
		log.Fatal(err)
	}
	var ID map[string]string

	err = json.Unmarshal(raw, &ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ID[code])
}

func CheckTodayHandler() {
	raw, err := ioutil.ReadFile(filepath.Join(base, "zoomData.json"))
	if err != nil {
		log.Fatal(err)
	}
	var data [][]Data

	err = json.Unmarshal(raw, &data)
	if err != nil {
		log.Fatal(err)
	}

	weekIndex := int(time.Now().Weekday())
	now := time.Now()

	for i := 0; i < len(data[weekIndex]); i++ {

		// determine outdated subject
		if (data[weekIndex][i].EndDate - now.UnixMilli()) > 0 {
			tm.Clear()
			fmt.Println(data[weekIndex][i].SubjectName, "-", data[weekIndex][i].Room)
		}

	}

}
