package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    // :=演算子は初期値を宣言する
    doc, err := goquery.NewDocument("https://www.google.com/")
    if err != nil {
        fmt.Printf("Failed")
    }
    title := doc.Find("news_wrp").Text()
    // println : 改行付き出力
    fmt.Println(title)
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        // Attrの戻り値 val string, exists bool
        link, _ := s.Attr("href")
        fmt.Println(link)
    })
}