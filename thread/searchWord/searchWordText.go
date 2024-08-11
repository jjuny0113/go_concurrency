package searchWord

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

//
//func InitV3() {
//	if len(os.Args) < 3 {
//		fmt.Println("2개 이상의 실행 인수가 필요합니다.")
//		return
//	}
//
//	word := os.Args[1] // 찾으려는 단어
//	files := os.Args[2:]
//	findInfos := []FindInfo{}
//
//	for _, path := range files {
//		//파일 찾기
//		findInfos = append(findInfos, findWordInAllFiles(word, path)...)
//	}
//	for _, findInfo := range findInfos {
//		fmt.Println(findInfo.filename)
//		fmt.Println("_________________")
//		for _, lineInfo := range findInfo.lines {
//			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
//		}
//		fmt.Println("_________________")
//		fmt.Println()
//	}
//}
//
//func getFileList(path string) ([]string, error) {
//	return filepath.Glob(path)
//}
//
//func findWordInAllFiles(word string, path string) []FindInfo {
//	findInfos := []FindInfo{}
//	fileList, err := getFileList(path)
//	if err != nil {
//		fmt.Println("파일 경로가 잘못되었습니다.", err, "path", path)
//		return findInfos
//	}
//	for _, filename := range fileList {
//		findInfos = append(findInfos, findWordInFile(word, filename))
//	}
//	return findInfos
//}
//
//func findWordInFile(word, filename string) FindInfo {
//	findInfo := FindInfo{filename, []LineInfo{}}
//	file, err := os.Open(filename)
//	if err != nil {
//		fmt.Println("파일을 찾을 수 없습니다.", filename)
//		return findInfo
//	}
//	defer file.Close()
//
//	lineNo := 1
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		line := scanner.Text()
//		if strings.Contains(line, word) {
//			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
//		}
//		lineNo++
//	}
//
//	return findInfo
//}
