package model

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"regexp"
	"strings"
	"time"
)

type BookSourceRule struct {
	BookSourceID       int64             `json:"bookSourceID" gorm:"primary_key;comment:'爬虫任务id'"`
	BookSourceName     string            `json:"bookSourceName" gorm:"size:100;not null;comment:'书源名称'"`
	BookSourceURL      string            `json:"bookSourceURL" gorm:"size:255;not null;comment:'书源网址'"`
	Enabled            string            `json:"enabled" gorm:"type:char;default:'1';not null;comment:'已启用 0,不启用 1,启用'"`
	Header             map[string]string `json:"header" gorm:"type:json;serializer:json;comment:'发送请求的header'"`
	Weight             int64             `json:"weight" gorm:"default:0;not null;comment:'权重'"`
	CategoryMapRule    map[string]string `json:"categoryMapRule" gorm:"type:json;serializer:json;comment:'分类'"`
	BookListRule       string            `json:"bookListRule" gorm:"size:255;not null;comment:'书籍列表规则'"`
	BookNameRule       string            `json:"bookNameRule" gorm:"size:255;not null;comment:'书名规则'"`
	AuthorRule         string            `json:"authorRule" gorm:"size:255;not null;comment:'作者规则'"`
	CoverURLRule       string            `json:"coverURLRule" gorm:"size:500;not null;comment:'封面地址'"`
	IntroRule          string            `json:"introRule" gorm:"size:255;not null;comment:'介绍规则'"`
	StatusRule         string            `json:"statusRule" gorm:"size:255;not null;comment:'状态规则'"`
	LastUpdateTimeRule string            `json:"lastUpdateTimeRule" gorm:"size:255;not null;comment:'最后更新时间规则'"`
	ChapterPathRule    string            `json:"chapterPathRule" gorm:"size:255;not null;comment:'章节路径规则'"`
	ChapterListRule    string            `json:"chapterListRule" gorm:"size:255;not null;comment:'章节列表规则'"`
	ContentRule        string            `json:"contentRule" gorm:"size:255;not null;comment:'内容规则'"`
	NextContentRule    string            `json:"nextContentRule" gorm:"size:255;default:'';not null;comment:'下一页内容规则'"`
	RespondTime        int64             `json:"respondTime" gorm:"default:0;not null;comment:'响应时间'"`
	UpdateTime         time.Time         `json:"lastUpdateTime" gorm:"default:current_timestamp();comment:'最后更新时间'"`
	CreateTime         time.Time         `json:"createTime" gorm:"default:current_timestamp();comment:'创建时间'"`
	IsDelete           int               `json:"IsDelete" gorm:"default:0;not null;comment:'是否删除'"`
}

func (BookSourceRule) TableName() string {
	return "book_source_rules"
}

type BookInfo struct {
	BookName       string `json:"bookName"`
	Author         string `json:"author"`
	CoverURL       string `json:"coverURL"`
	Intro          string `json:"intro"`
	Category       string `json:"category"`
	Status         string `json:"status"`
	LastUpdateTime string `json:"lastUpdateTime"`
}

type BookChapter struct {
	ChapterName string `json:"chapterName"`
	ChapterURL  string `json:"chapterURL"`
}

func (r BookSourceRule) getRuleStr(rule string) (string, string, string, error) {
	//#{选择器}
	selectReg, _ := regexp.Compile(`#\{(.*)}@`)
	selector := selectReg.FindStringSubmatch(rule)
	//@{type}
	typeReg, _ := regexp.Compile(`@\{(.*)}\$`)
	typ := typeReg.FindStringSubmatch(rule)
	//${自定义正则表达式}
	customizeReg, _ := regexp.Compile(`\$\{(.*)}`)
	customize := customizeReg.FindStringSubmatch(rule)
	// 防止越界异常
	if len(selector) == 1 || len(typ) == 1 {
		// 返回异常
		return "", "", "", fmt.Errorf("rule error")
	}
	if len(customize) == 1 {
		customize = append(customize, "")
	}
	return selector[1], typ[1], customize[1], nil
}

/**
 * @Description: 解析规则 单返回值
 */
func (r BookSourceRule) rule(e *colly.HTMLElement, value *string, selector, tye, reg string) {
	var extracted string
	switch tye {
	case "text":
		extracted = e.Text
		break
	case "attr":
		attrReg, _ := regexp.Compile(`\[(.*)]`)
		attrMatches := attrReg.FindStringSubmatch(selector)
		if len(attrMatches) > 1 {
			extracted = e.Attr(attrMatches[1])
		}
		break
	}
	if len(reg) > 1 && extracted != "" {
		customizeReg, _ := regexp.Compile(reg)
		subMatches := customizeReg.FindStringSubmatch(extracted)
		if len(subMatches) > 1 {
			*value = subMatches[1]
			return
		} else {
			return
		}
	}
	*value = extracted
}

/**
* @Description: 获取书籍列表
 */
func (r BookSourceRule) bookList(categoryUrl string) (result []string, err error) {
	c := colly.NewCollector()
	result = []string{}
	selector, tye, reg, err := r.getRuleStr(r.BookListRule)
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		// 如果出现Attr 情况
		bookAddr := ""
		r.rule(e, &bookAddr, selector, tye, reg)
		result = append(result, bookAddr)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit(r.BookSourceURL + categoryUrl)
	return result, err
}

/**
* @Description: 获取书籍信息
 */
func (r BookSourceRule) bookInfo(path string) (BookInfo, error) {
	c := colly.NewCollector()
	bookInfo := BookInfo{
		BookName:       "",
		Author:         "",
		CoverURL:       "",
		Intro:          "",
		Category:       "",
		Status:         "",
		LastUpdateTime: "",
	}
	//书名
	selector, tye, reg, err := r.getRuleStr(r.BookNameRule)
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		r.rule(e, &bookInfo.BookName, selector, tye, reg)
	})
	// 作者
	selectorAct, tyeAct, regAct, err := r.getRuleStr(r.AuthorRule)
	c.OnHTML(selectorAct, func(e *colly.HTMLElement) {
		r.rule(e, &bookInfo.Author, selectorAct, tyeAct, regAct)
	})
	// 封面
	Coverselector, Covertye, Coverreg, err := r.getRuleStr(r.CoverURLRule)
	c.OnHTML(Coverselector, func(e *colly.HTMLElement) {
		r.rule(e, &bookInfo.CoverURL, Coverselector, Covertye, Coverreg)
	})
	// 简介
	Introselector, Introtye, Introreg, err := r.getRuleStr(r.IntroRule)
	c.OnHTML(Introselector, func(e *colly.HTMLElement) {
		r.rule(e, &bookInfo.Intro, Introselector, Introtye, Introreg)
	})
	// 状态
	Statusselector, Statustye, Statusreg, err := r.getRuleStr(r.StatusRule)
	c.OnHTML(Statusselector, func(e *colly.HTMLElement) {
		r.rule(e, &bookInfo.Status, Statusselector, Statustye, Statusreg)
	})
	// 最后更新时间
	LastUpdateTimeselector, LastUpdateTimetye, LastUpdateTimereg, err := r.getRuleStr(r.LastUpdateTimeRule)
	c.OnHTML(LastUpdateTimeselector, func(e *colly.HTMLElement) {
		r.rule(e, &bookInfo.LastUpdateTime, LastUpdateTimeselector, LastUpdateTimetye, LastUpdateTimereg)
	})
	// TODO: 保存info到db
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})
	c.Visit(r.BookSourceURL + path)
	return bookInfo, err
}

/**
* @Description: 获取章节列表
 */
func (r BookSourceRule) chapterList(path string) []BookChapter {
	chapters := make([]BookChapter, 0)
	// 章节列表
	chapterFlag := true
	i := 1
	for chapterFlag {
		c := colly.NewCollector()
		replacer := strings.NewReplacer(
			"${BookSourceURL}", r.BookSourceURL,
			"${BookSourcePath}", path,
			"${page}", fmt.Sprintf("%d", i),
		)
		toUrl := replacer.Replace(r.ChapterPathRule)
		// 章节规则
		// url链接
		selector, _, reg, _ := r.getRuleStr(r.ChapterListRule)
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			chapter := BookChapter{}
			r.rule(e, &chapter.ChapterURL, selector, "attr", reg)
			r.rule(e, &chapter.ChapterName, selector, "text", reg)
			chapters = append(chapters, chapter)
		})
		c.OnRequest(func(r *colly.Request) {
			log.Println("visiting", r.URL.String())
		})
		c.Visit(toUrl)
		chapterFlag = false
		i++
	}
	fmt.Printf("%v\n", chapters)
	return chapters
}
