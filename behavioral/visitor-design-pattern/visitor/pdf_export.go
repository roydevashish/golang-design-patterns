package visitor

import "fmt"

type PDFExportVisitor struct{}

func (p *PDFExportVisitor) VisitLaw(l *LawClient) {
	fmt.Println("export pdf for", l.Email)
}

func (p *PDFExportVisitor) VisitRestaurent(r *RestaurentClient) {
	fmt.Println("export pdf for", r.Email)
}

func (p *PDFExportVisitor) VisitRetail(r *RetailClient) {
	fmt.Println("export pdf for", r.Email)
}
