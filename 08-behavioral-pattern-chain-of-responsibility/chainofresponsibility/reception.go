package chainofresponsibility

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.isRegistered {
		fmt.Println("Patient registration has already done")
		r.next.Execute(p)
	}
	fmt.Println("Reception registering patient")
}
