package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
	"net/smtp"
	"github.com/robfig/cron/v3"
    "time"
    "log"
)

func main() {
	c := cron.New()
    // c.AddFunc("@every 1s", func() { log.Println("every 1s") })
    // 1時間ごとの定期実行
	c.AddFunc("@hourly", scraping)
	c.Start()
    // 1週間の間は定期実行される
    time.Sleep(168 * time.Hour)
}

func scraping() {
	var message string
    // 対象のURLを設定
	base_url := "https://www.google.com/"
    doc, err := goquery.NewDocument(base_url)
    if err != nil {
        fmt.Printf("スクレイピングに失敗しました")
    }
	selection := doc.Find(".pageList")
	innerSelection := selection.Find(".clearfix")
	innerSelection.Each(func(i int, s *goquery.Selection) {
        // Attrの戻り値 val string, exists bool
		li_item  := s.Find(".title")
		href := li_item.Find("a")
        // リンクを取得
		link, _ := href.Attr("href")
        // タイトルを取得
		title := href.Text()
        // 結果を作成
		message += title+"\n"
		message += base_url+ link+"\n"
    })

	// メール送信
	auth := smtp.PlainAuth(
        "",
        "test@example.com", // 送信に使うアカウント
        "hoge", // アカウントのパスワード or アプリケーションパスワード
        "smtp.gmail.com",
    )
    // 送信処理
	smtp.SendMail(
        "smtp.gmail.com:587",
        auth,
        "test+from@example.com", // 送信元
        []string{"test+to@example.com"}, // 送信先
        []byte(
            "To: <recipient>@gmail.com\r\n" +
            "Subject:スクレイピング結果\r\n" +
            "\r\n" +
            message),
    )
    const layout = "2006-01-02 15:04:05"
    t := time.Now()
    fmt.Println("スクレイピング正常終了" + t.Format(layout))
}
