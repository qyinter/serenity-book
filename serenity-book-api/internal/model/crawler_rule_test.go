package model

import (
	"fmt"
	"testing"
)

func TestCrawlerRule(t *testing.T) {
	source := BookSourceRule{
		BookSourceID:   1,
		BookSourceName: "书趣阁",
		BookSourceURL:  "https://www.22shuquge.com",
		Enabled:        "1",
		Header:         map[string]string{"test": "test"},
		Weight:         0,
		CategoryMapRule: map[string]string{
			"玄幻": "/class/1_1.html",
			"修真": "/class/2_1.html",
			"都市": "/class/3_1.html",
			"历史": "/class/4_1.html",
			"网游": "/class/6_1.html",
			"科幻": "/class/7_1.html",
			"耽美": "/class/8_1.html",
			"恐怖": "",
			"军事": "",
			"其他": "",
		},
		BookListRule: "#{.txt-list.txt-list-row5 li .s2 a[href]}@{attr}${}",

		BookNameRule:       "#{.top h1}@{text}${}",
		AuthorRule:         "#{.fix p}@{text}${作.*者：(.*)}",
		CoverURLRule:       "#{.imgbox img[src]}@{attr}${}",
		IntroRule:          "#{.desc.xs-hidden}@{text}${}",
		StatusRule:         "#{.xs-show}@{text}${状.*态：(.*)}",
		LastUpdateTimeRule: "#{p}@{text}${更新时间：(.*)}",

		ChapterPathRule: "${BookSourceURL}${BookSourcePath}${page}", // 注意斜杠
		ChapterListRule: "#{.section-list.fix li a[href]}@{}${}",

		ContentRule:     ".content>#p",
		NextContentRule: "~next_url",
		RespondTime:     3000,
	}
	// #{选择器}@{type}${自定义正则表达式}
	// type: text、attr、html

	// 查找小说列表 取第一分类测试
	categoryUrl := source.CategoryMapRule["玄幻"]
	bookList, _ := source.bookList(categoryUrl)
	fmt.Printf("%v\n", bookList)
	//TODO 开启多协程
	info, _ := source.bookInfo(bookList[0])
	info.Category = "玄幻"
	fmt.Printf("%v\n", info)
	// 查找章节列表
	source.chapterList(bookList[0])
	//TODO 翻页分类翻页

}
