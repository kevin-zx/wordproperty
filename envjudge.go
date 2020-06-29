package wordproperty

import "strings"

var judgeWords = []string{
	"怎么",
	"什么",
	"排行",
	"排名",
	"多少",
	"多长",
	"好不好",
	"哪",
	//"哪个",
	//"哪家",
	//"哪些",
	//"哪里",
	"选择",
	"还是",
	"可以",
	"如何",
	"有没有",
	"呢",
}

var lowConfidenceWords = []string{
	"正确",
	//"卖",
	"一般",
	"不要",
	"区别",
	"品牌",
	"牌子",
	"种类",
	"分类",
	"类型",
	"问题",
	"购买",
	"知识",
}

var judgeSuffixes = []string{
	"好",
	"吗",
	"?",
	"？",
}

// EnvWordProperty the function is using to judge word whether environment or not,
// params
// word: the word need to be judge
// return
// isEnv(bool): the word whether environment or not
// confidence(float): the confidence of whether environment word
func EnvWordProperty(word string) (bool, float64) {
	isEnv := false
	var c float64 = 0
	for _, judgeWord := range judgeWords {
		if strings.Contains(word, judgeWord) {
			isEnv = true
			c = 1
			break
		}
	}

	if !isEnv {
		for _, suffix := range judgeSuffixes {
			if strings.HasSuffix(word, suffix) {
				isEnv = true
				c = 1
				break
			}
		}
	}

	if !isEnv {
		for _, confidenceWord := range lowConfidenceWords {
			if strings.Contains(word, confidenceWord) {
				isEnv = true
				c += 0.2
			}
		}
	}

	if c > 1 {
		c = 1
	}
	return isEnv, c
}
