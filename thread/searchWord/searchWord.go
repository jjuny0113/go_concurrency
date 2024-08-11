package searchWord

//func Init() {
//	if len(os.Args) < 3 {
//		fmt.Println("2개 이상의 실행 인수가 필요합니다.")
//		return
//	}
//	word := os.Args[1]
//	files := os.Args[2:]
//	fmt.Println("찾으려는 단어:", word)
//	printAllFiles(files)
//}

//func getFileList(path string) ([]string, error) {
//	return filepath.Glob(path) // glob는 파일 읽는 것
//}
//
//func printAllFiles(files []string) {
//	for _, path := range files {
//		fileList, err := getFileList(path)
//		if err != nil {
//			fmt.Println("파일 경로가 잘못되었습니다.", err, "path", path)
//			return
//		}
//		fmt.Println("찾으려는 파일 리스트")
//		for _, file := range fileList {
//			fmt.Println(file)
//		}
//	}
//}
