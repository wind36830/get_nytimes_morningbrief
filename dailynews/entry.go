/**
 * Create by GoLand
 * User: SCAR
 * Date: 2018/4/16
 * Time: 20:05
 */
package main

import (
	"net/http"
	"io/ioutil"
	"time"
	"log"
	"regexp"
	"fmt"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"

)
func main()  {
	for {
		if x := time.Now().Hour(); x == 8 {
			url := "https://cn.nytimes.com/morning-brief"
			//url := "https://cn.nytimes.com/china/"
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			today := time.Now()
			year := strconv.Itoa(today.Year())
			var month string
			if x := int(today.Month()); x < 10 {
				month = "0" + strconv.Itoa(x)
			} else {
				month = strconv.Itoa(x)
			}
			var day string
			if x := int(today.Day()); x < 10 {
				day = "0" + strconv.Itoa(x)
			} else {
				day = strconv.Itoa(x)
			}
			searchday := year + month + day
			fmt.Println(searchday)
			r := regexp.MustCompile("\"/morning-brief/" + searchday + "/.*?\"")
			results := strings.Trim(r.FindString(string(body)), "\"")
			host := "https://cn.nytimes.com"
			news := getInfo(host, results, searchday)
			ioutil.WriteFile(searchday+"早报.txt", []byte(news), 0666)
			fmt.Println(time.Now())
			t := time.NewTimer(time.Hour*1)
			<-t.C
			fmt.Println("日报被写入")

		}else {
			fmt.Println("时间未到开头")
			t := time.NewTimer(time.Hour*1)

			<-t.C
			fmt.Println("时间未到")
		}
	}

}
func getInfo(host, url string,date string) string  {
	uri := host+url
	res, err:= http.Get(uri)
	if err!= nil{
		log.Fatal(err)
	}
	defer  res.Body.Close()
	doc,err:= goquery.NewDocumentFromReader(res.Body)
	 news := date+"早报\r\n"
	 doc.Find(".article-body > p").Each(func(i int, selection *goquery.Selection) {

		x:= selection.Text()
		if strings.Index(x,"注册")<0{
			news+= x+"\r\n"
		}
	})
	return news

}

