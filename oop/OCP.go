package oop

//func SendReport(r *Report, method SendType, receiver string) {
//	switch method {
//	case Email:
//	// 이메일 전송
//	case Fax:
//	// fax 전송
//	case PDF:
//	// pdf 파일 생성
//	case Printer:
//		// 프린팅
//	}
//}

type EmailSender struct {
}

func (e *EmailSender) Send(r *Report) {

}

type FaxSender struct {
}

func (f *FaxSender) Send(r *Report) {
	
}
