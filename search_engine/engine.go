package main

import (
	"regexp"
	"slices"
	"strings"
)

type Analyser struct {
	index map[ /*term*/ string][]local
}

type local struct {
	DocId   int64
	Indexes int64
}

type searchModeExp string

const (
	AND     searchModeExp = "AND"
	OR      searchModeExp = "OR"
	INCLUDE searchModeExp = "+"
	EXCLUDE searchModeExp = "-"
)

var AllsearchMode = map[searchModeExp]any{
	AND:     nil,
	OR:      nil,
	INCLUDE: nil,
	EXCLUDE: nil,
}

func isSearchMode(s string) bool {
	_, ok := AllsearchMode[searchModeExp(s)]
	return ok
}

const Quote = "\""

func NewAnalyser() *Analyser {
	return &Analyser{}
}

func (a *Analyser) AnalyseAndIndex(documents []string) {
	for docId, oneDocument := range documents {
		tokens := regexp.MustCompile(`\w+`).FindAllString(oneDocument, -1)
		for index, token := range tokens {
			token = strings.ToLower(token)
			local_ := local{DocId: int64(docId), Indexes: int64(index)} //use inverted index to index terms
			if a.index == nil {
				a.index = make(map[string][]local)
			}
			a.index[token] = append(a.index[token], local_)
		}
	}
}

func (a *Analyser) Search(keyWords string) []SearchRes {
	includeDocScoresMap := make(map[int64]float64)
	excludeDocMap := make(map[int64]any)
	queryExpressions := a.parseQuery(keyWords)
	var searchMode = OR
	for i := 0; i < len(queryExpressions); i++ {
		var locals []local
		if isSearchMode(queryExpressions[i]) {
			searchMode = searchModeExp(queryExpressions[i])
			continue
		} else if queryExpressions[i] == Quote {
			endIndex := stringsIndex(queryExpressions[i+1:], Quote)
			if endIndex == -1 {
				continue // 缺少结束引号，忽略
			}
			if endIndex == 0 {
				continue // 引号中间没有任何exp，忽略
			}
			quoteQueryExpressions := queryExpressions[i+1 : i+1+endIndex]
			i += endIndex + 1

			for i2 := range quoteQueryExpressions {
				quoteQueryExpressions[i2] = strings.ToLower(quoteQueryExpressions[i2])
			}
			firstExp, otherExp := quoteQueryExpressions[0], quoteQueryExpressions[1:]
			var ok bool

			locals, ok = a.index[firstExp]
			if !ok {
				// not found,early continue
				continue
			}
			for _, exp := range otherExp {
				locals = a.findExpInIndex(locals, exp)
			}

		} else {
			//	search
			locals = a.index[queryExpressions[i]]
		}
		switch searchMode {
		case AND:
			for _, l := range locals {
				if _, ok := includeDocScoresMap[l.DocId]; ok {
					includeDocScoresMap[l.DocId] += 1
				}
			}
		case OR:
			for _, l := range locals {
				includeDocScoresMap[l.DocId] += 1
			}
		case INCLUDE:
			locsMap := make(map[int64]float64, len(includeDocScoresMap))
			for _, l := range locals {
				locsMap[l.DocId] = includeDocScoresMap[l.DocId] + 1
			}
			includeDocScoresMap = locsMap
		case EXCLUDE:
			for _, l := range locals {
				excludeDocMap[l.DocId] = nil
			}
		}
		searchMode = OR
	}
	res := make([]SearchRes, 0)
	for k, v := range includeDocScoresMap {
		if _, ok := excludeDocMap[k]; ok {
			continue
		}
		res = append(res, SearchRes{DocId: k, Score: v})
	}
	slices.SortFunc(res, func(a, b SearchRes) int {
		return int(b.Score - a.Score) // order by score desc
	})
	return res
}

type SearchRes struct {
	DocId int64
	Score float64
	//Doc   string
}

// todo add test
func (a *Analyser) findExpInIndex(locals []local, exp string) []local {
	if len(locals) == 0 {
		return make([]local, 0)
	}
	expLocals, ok := a.index[exp]
	if !ok {
		return make([]local, 0)
	}
	res := make([]local, 0)
	for i := 0; i < len(locals); i++ {
		expexctIndex := locals[i].Indexes + 1 //需要在expexctIndex出现exp
		isFind, l := findInDocAndIndex(locals[i].DocId, expexctIndex, expLocals)
		if isFind {
			res = append(res, l)
		}
	}
	return res
}

func findInDocAndIndex(docId int64, index int64, localsB []local) (bool, local) {
	for _, l := range localsB {
		if l.DocId == docId && l.Indexes == index {
			return true, l
		}
	}
	return false, local{}
}

// index strat from 0， not found return -1
func stringsIndex(strings []string, subStr string) int {
	for i := 0; i < len(strings); i++ {
		if strings[i] == subStr {
			return i
		}
	}
	return -1
}

func (a *Analyser) parseQuery(keyWords string) (res []string) {
	reg := regexp.MustCompile(`["\-+]|\w+`)
	keywords := reg.FindAll([]byte(keyWords), -1)
	//for _, keyword := range keywords {
	//	log.Println(string(keyword))
	//}
	for _, keyword := range keywords {
		res = append(res, string(keyword))
	}
	return
}
