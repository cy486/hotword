package controllers

import (
	"fmt"
	"github.com/Lofanmi/pinyin-golang/pinyin"
	orm2 "github.com/astaxie/beego/orm"
	"hotword/models"
	"io"
	"net/http"
	"regexp"
	"strings"
)

//保存词语的文章列表
func save1(title, concent, urls string, id int) {
	orm := orm2.NewOrm()
	WC := models.WordContent{}
	WC.Title = string(title)
	WC.Concent = orm2.TextField(concent)
	WC.Word = int(id)
	WC.Url = string(urls)
	_, err := orm.Insert(&WC)
	if err != nil {
		fmt.Println(err)
	}
}

//防止数组越界而取最小值
func small(a, b, c int) (result int) {
	result = a
	if a > b {
		result = b
	} else if a > c {
		result = c
	}
	return
}

//查询文章的方法
func findcontent(word string /*page chan int, */, id int) {
	url := "http://app.idcquan.com/?app=search&controller=index&action=search&type=all&wd=" + word
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("读取网页错误", err)
	}
	titlerets := regexp.MustCompile(`<span class="title">(?s:(.*?))</a>`)
	title := titlerets.FindAllStringSubmatch(result, -1)

	concentret := regexp.MustCompile(`<span class="nei_rong">...(?s:(.*?))...</span>`)
	concent := concentret.FindAllStringSubmatch(result, -1)

	urlret := regexp.MustCompile(`<div class="news_nr">(?s:(.*?))" target="_blank" class="d1">`)
	alls := urlret.FindAllStringSubmatch(result, -1)
	a := small(len(title), len(concent), len(alls))
	for i := 0; i < a; i++ {
		title[i][1] = strings.Replace(title[i][1], `<span class="keyword">`, "", -1)
		title[i][1] = strings.Replace(title[i][1], `</span>`, "", -1)
		title[i][1] = strings.TrimSpace(title[i][1])
		alls[i][1] = strings.TrimSpace(alls[i][1])
		alls[i][1] = strings.Replace(alls[i][1], `<a href="`, "", -1)
		save1(title[i][1], concent[i][1], alls[i][1], id)
	}
	/*page <- id*/
}

//读取网页的方法
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

//获取一个词语的首字母
func shouzimu(s string) (is_a string) {
	dict := pinyin.NewDict()
	pin := dict.Sentence(s).Unicode()
	is_a = strings.ToUpper(string(pin[0]))
	return
}

//获取一个词语的百度的词条
func con(word string) string {
	url := "https://baike.baidu.com/item/" + word
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("读取网页错误", err)
	}
	urlret := regexp.MustCompile(`<meta name="description" content="(?s:(.*?))...">`)
	alls := urlret.FindAllStringSubmatch(result, 1)
	if alls != nil {
		return alls[0][1]
	}
	return "a"
}
