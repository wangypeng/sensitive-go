package wordFilter

import (
	"fmt"
	"bytes"
	"strings"
)

type wordtree struct {
	tree map[string]*wordtree
	isEnd bool
}

/**
 *	this is store sensitive words in memory
 * 	private variable 
 */
var sensitiveWordMap wordtree

/**
 * 	local sensitive word list to memory
 *	private method
 */
func LoadSensitiveWord (set []string) {
	for _ , value := range set {
		nowMap := &sensitiveWordMap
		wordLenth := len([]rune(value))
		for i := 0 ; i < wordLenth ; i ++ {
			char := string([]rune(value)[i])
			wordMap := nowMap.tree[char]
			if wordMap != nil {  
				nowMap = wordMap
			}else{
				var newtree wordtree
				if nowMap.tree == nil {
					nowMap.tree = make(map[string]*wordtree)
				}
				nowMap.tree[char] = &newtree
				nowMap = &newtree
			}
			if(i == wordLenth-1){
				nowMap.isEnd = true
			}
		}
	}
	fmt.Println(sensitiveWordMap.tree["测"])
	fmt.Println("load sensitive word map end ...")
}
/**
 *	delete sensitive word from memory
 *	private method
 */
func DelSensitiveWord (delword string){

	nowMap := &sensitiveWordMap
	forkNo := 0
LABEL1:
	for i := 0 ; i < len([]rune(delword)) ; i ++ {
		word := string(delword[i])
		if(len(nowMap.tree) > 2){
			forkNo = i;
		}
		nowMap = nowMap.tree[word]
		if nowMap != nil {
			if nowMap.isEnd && i == len(delword) -1 {
				nowMap = &sensitiveWordMap
				for j := 0 ; j < forkNo + 1 ; j ++ {
					if j != forkNo {
						nowMap = nowMap.tree[string(delword[j])]
					}else{
						delete(nowMap.tree,string(delword[j]))
					}
				}
			}
		}else{
			nowMap = &sensitiveWordMap
			nowMap = nowMap.tree[word]
			if nowMap == nil {
				nowMap = &sensitiveWordMap;
			}
		}
		continue LABEL1
	}
}
/**
 *	replace sensitive word in request param text by global params 'SensitiveWordMap'
 */
func ReplaceSensitiveWord (txt string) string {

	var word string
	var failWord bytes.Buffer
	resultText := txt

	nowMap := &sensitiveWordMap

LABEL1:
	for i := 0 ; i < len([]rune(txt)) ; i++ {
		word = string([]rune(txt)[i])
		nowMap = nowMap.tree[word]
		if nowMap != nil {
			failWord.WriteString(word)
			if nowMap.isEnd {
				if len(nowMap.tree) == 0 {
					resultText = replace(len([]rune(failWord.String())),i,resultText)

					nowMap = &sensitiveWordMap
					failWord.Reset()
					continue LABEL1
				} else if len(nowMap.tree) > 1 {
					if i == (len(txt)-1) {
						resultText = replace(len([]rune(failWord.String())),i,resultText)

						nowMap = &sensitiveWordMap
						failWord.Reset()
						continue LABEL1
					}else{
						if nowMap.tree[string(txt[i+1])] == nil {
							resultText = replace(len(failWord.String()),i,resultText)

							nowMap = &sensitiveWordMap
							failWord.Reset()
							continue LABEL1
						}else{
							continue LABEL1
						}
					}
				}
			}
		}else{
			nowMap = &sensitiveWordMap
			failWord.Reset()
			nowMap = nowMap.tree[word]
			if nowMap == nil {
				nowMap = &sensitiveWordMap
				failWord.Reset()
			}else{
				failWord.WriteString(word)
			}
			continue LABEL1
		}
	}
	return resultText
}

/**
 *	replace word's sensitive word by '*'
 */
func replace (failLength int ,local int ,preWord string) string{
	var replaceWord bytes.Buffer
	for j := 0 ; j < failLength ; j++ {
		replaceWord.WriteString("*")
	}
	prefix := string([]rune(preWord)[0:local-(failLength-1)])
	suffix := string([]rune(preWord)[local+1:len(preWord)])
	return strings.Join([]string{prefix , replaceWord.String() , suffix},"")
}


