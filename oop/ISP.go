package oop

import "time"

//type Report interface {
//	Report() string
//	Pages() int
//	Author() string
//	WrittenDate() time.Time
//}
//
//func SendReport(report Report) {}

type ISPReport interface {
	Report()
}

type WrittenInfo interface {
	Page() int
	Author() string
	Written() time.Time
}

//func ISPSendReport(r Report) {
//	send(r.Report())
//}
