package main

import (
	"fmt"
	"net/http"
	"strings"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// 从url中提取域名
func extractDomain(url string) (string, error) {
	// 定义正则表达式
	re := regexp.MustCompile(`^(?:https?:)?(?:\/\/)?(?:www\.)?([^\/]+)`)

	// 使用正则表达式匹配URL
	match := re.FindStringSubmatch(url)

	// 如果匹配成功，返回域名
	if len(match) > 1 {
		return match[1], nil
	}

	return "", fmt.Errorf("无法提取域名:" + url)
}

func main(){
	link := "http://www.baidu.com"
	res, err := http.Get(link)
	if err != nil{
		// panic 用于在 HTTP 请求失败或 HTML 解析失败时终止程序
		/*
			panic 是 Go 语言中的一个内置函数，用于引发一个运行时错误并立即终止当前的 goroutine。
			在引发 panic 之后，程序会停止正常执行，并开始从当前函数向上传播，直到找到一个 recover 函数来恢复正常执行，或者终止整个程序。
		*/
		panic(err)
	}
	// defer 用于确保 HTTP 响应主体在函数退出时被关闭，以防止资源泄漏
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil{
		panic(err)
	}

	// 定义一个切片 存储所有a标签下的url 注意要初始化大小为0 如果初始化为3 则会有初始的三个空字符串在里面
	Urls_slice := make([]string, 0)
	doc.Find("a").Each(func(i int, s *goquery.Selection){
		url, exists := s.Attr("href")
		if exists{
			// fmt.Println(url)
			Urls_slice = append(Urls_slice, strings.TrimSpace(string(url)))
		}
	})

	Domains_slice := make([]string, 0)
	for _, value := range Urls_slice{
		// 从字符串中利用正则提取域名
		domain, err := extractDomain(value)
		if err != nil{
			panic(err)
		}

		Domains_slice = append(Domains_slice, domain)
	}

	// 输出切片中的域名
	for _, value := range Domains_slice{
		fmt.Println(value)
	}
}