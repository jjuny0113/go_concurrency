package oop

type Report interface {
	Report() string
}

type FinanceReport struct {
	report string
}

func (r *FinanceReport) Report() string {
	return r.report
}

type ReportSender struct {
}

func (s *ReportSender) SendReport(report Report) {
	
}
