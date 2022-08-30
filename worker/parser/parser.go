package parser

import (
	"QLProject/engine"
	"fmt"
	"regexp"
)

const QuanLiangName = `<div class="_3_JaaUmGUCjKZIdiLhqtfr">([^<]+)</div>`
const QuanLiangPublishData = `<div class="_3TzAhzBA-XQQruZs-bwWjE"><div class="_3NQyvscVRjQNS9gEgQoDDo"></div>(["\d{4}-\d{2}-\d{2}]")</div>`
const QuanLiangPublishLearner = `<div class="__2gvAnxa4Xc7IT14d5w8MI1">([1-9][0-9]*])</div>`
const QuanLiangPublishNumber = `<a rel="nofollow">([^<]+)</a>`

//传入contents，输出是一个request的一个item的集合
func ParseQLList(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	//分页数
	re := regexp.MustCompile(QuanLiangPublishNumber)
	matches := re.FindAllSubmatch(contents, -1)
	number := matches[len(matches)-1][1]

	fmt.Println(number)

	reName := regexp.MustCompile(QuanLiangName)
	matchesQLName := reName.FindAllSubmatch(contents, -1)

	reQLPublishData := regexp.MustCompile(QuanLiangPublishData)
	matchesQLPublishData := reQLPublishData.FindAllSubmatch(contents, -1)

	reQLPublishLearner := regexp.MustCompile(QuanLiangPublishLearner)
	matchesQLPublishLearner := reQLPublishLearner.FindAllSubmatch(contents, -1)

	/*
		1.将分页的number字节转成int
		2.将爬出的数值放入result中
		3.写一个写excel方法依次在每一个out channel写出去
	*/
	for i := 0; i < 59; i++ {
		//分页处理
		qlItem := engine.QLItem{
			QuanLiangName:           string(matchesQLName[i][1]),
			QuanLiangPublishLearner: string(matchesQLPublishLearner[i][1]),
			QuanLiangPublishData:    string(matchesQLPublishData[i][1]),
		}
		result.Items = append(result.Items, qlItem)
	}
	return result
}
