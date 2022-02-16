package main

//パッケージインポート
import (
    "encoding/json"
    "time"
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "strconv"
)

//お天気構造体
type Weather struct {
	PublishingOffice string    `json:"publishingOffice"`
	ReportDatetime   time.Time `json:"reportDatetime"`
	TargetArea       string    `json:"targetArea"`
	// HeadlineText     string    `json:"headlineText"`
	Text             string    `json:"text"`
}

func main() {
	//日付を取得
    t := time.Now()
	//曜日を日本語に変換
	weekday := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	//接尾辞を付けて変数に格納
	var weekday_format = weekday[t.Weekday()]+"曜日"
    const (
        //区切り線
        hr = "▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲▽▲"
		//日付フォーマット
		date = "2006年01月02日"
		//時刻フォーマット
		time = "15時04分"
	)
	var (
		nowyear int = t.Year()
		// nowmonth = t.Month()
		// nowday int = t.Day()
		nowyearday int = t.YearDay()
	)

	//残り日数を計算
	var daynum int
	if isLeapYear(nowyear){
        //うるう年の日数
		daynum = 366
	}else{
        //うるう年でない年の日数
		daynum = 365
	}
    //残りの日数
	var lastyearday int = daynum - nowyearday

    //区切り線
    fmt.Println(hr)
	//1行目
    fmt.Println("今日は"+t.Format(date)+weekday_format+"です。"+"今年は残り"+strconv.Itoa(lastyearday)+"日です")
	//2行目
    fmt.Println("現在の時刻は"+t.Format(time)+"です。")
    //区切り線
    fmt.Println(hr)

    //気象庁のjsonから取得
    //140000は神奈川県の気象情報
    url := "https://www.jma.go.jp/bosai/forecast/data/overview_forecast/140000.json"
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    byteArray, _ := ioutil.ReadAll(resp.Body)

    jsonBytes := ([]byte)(byteArray)
    data := new(Weather)

    if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }

    //二重改行を消去
    rep := regexp.MustCompile(`\n\n`)
    data.Text = rep.ReplaceAllString(data.Text, "\n")

    rep2 := regexp.MustCompile(`　`)
    data.Text = rep2.ReplaceAllString(data.Text, "")

    fmt.Println(data.TargetArea+"の天気")
    fmt.Println(data.Text)
    fmt.Println("("+data.ReportDatetime.Format(date)+data.ReportDatetime.Format(time)+data.PublishingOffice+"発表）")
	
    //区切り線
    fmt.Println(hr)
}

// うるう年かどうか判定する
func isLeapYear(year int) bool {
    if year%400 == 0 { // 400で割り切れたらうるう年
        return true
    } else if year%100 == 0 { // 100で割り切れたらうるう年じゃない
        return false
    } else if year%4 == 0 { // 4で割り切れたらうるう年
        return true
    } else {
        return false
    }
}