package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// for visual , we record the document id and document name
	var docNames map[ /*document id*/ int]string

	allData := make([]string, 0)
	docNames, allData = readAllData("./data/benchmark/small")
	analyser := NewAnalyser()
	analyser.AnalyseAndIndex(allData)
	for _, keyword := range []string{
		"just do",
		"just AND do",
		"\"just do\" AND it",
		"+\"just do\" AND tomorrow",
		"-just do",
		"do AND \"you can\" -tomorrow",
		"\"to be or not to be\"",
	} {
		start := time.Now()
		search := analyser.Search(keyword)
		log.Println("[", keyword, "] cost:", time.Since(start))
		for _, res := range search[:min(len(search), 3)] {
			log.Print("[", docNames[int(res.DocId)], "] ", res.Score)
		}
	}
	/**
	  2025/07/12 21:59:46 [ just do ] cost: 44.708µs
	  2025/07/12 21:59:46 [hamlet] 162
	  2025/07/12 21:59:46 [coriolanus] 130
	  2025/07/12 21:59:46 [2henryiv] 127
	  2025/07/12 21:59:46 [ just AND do ] cost: 23.959µs
	  2025/07/12 21:59:46 [hamlet] 162
	  2025/07/12 21:59:46 [2henryiv] 127
	  2025/07/12 21:59:46 [asyoulikeit] 127
	  2025/07/12 21:59:46 [ "just do" AND it ] cost: 22.958µs
	  2025/07/12 21:59:46 [ +"just do" AND tomorrow ] cost: 18.708µs
	  2025/07/12 21:59:46 [ -just do ] cost: 17.208µs
	  2025/07/12 21:59:46 [coriolanus] 130
	  2025/07/12 21:59:46 [lll] 95
	  2025/07/12 21:59:46 [ do AND "you can" -tomorrow ] cost: 732.625µs
	  2025/07/12 21:59:46 [hamlet] 161
	  2025/07/12 21:59:46 [cleopatra] 127
	  2025/07/12 21:59:46 [2henryiv] 121
	  2025/07/12 21:59:46 [ "to be or not to be" ] cost: 4.432041ms
	  2025/07/12 21:59:46 [hamlet] 1
	*/

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func readAllData(documentPath string) (map[int]string, []string) {

	// 检查目录是否存在
	if _, err := os.Stat(documentPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("目录 %s 不存在", documentPath))
	}

	// 读取目录中的文件列表
	files, err := ioutil.ReadDir(documentPath)
	if err != nil {
		panic(fmt.Sprintf("无法读取目录 %s: %v", documentPath, err))
	}

	var documents []string
	names := make(map[int]string)

	for i, file := range files {
		if file.IsDir() {
			continue // 跳过子目录
		}

		filePath := filepath.Join(documentPath, file.Name())
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "无法读取文件 %s: %v\n", filePath, err)
			continue
		}

		documentID := i
		names[documentID] = strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		documents = append(documents, string(content))
	}

	// 打印结果验证
	for id, name := range names {
		fmt.Printf("ID: %d, 文件名: %s, 内容长度: %d\n", id, name, len(documents[id]))
	}
	return names, documents
}
