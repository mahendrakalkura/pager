// Package pager is a library which you can use for the purpose of generating
// and rendering pagination links.
package pager

import (
	"fmt"
	"math"
)

// Pager is a container struct. It encapsulates Records and Pages structs.
type Pager struct {
	// (see below)
	Records Records
	// (see below)
	Pages Pages
	// ...URL syntax
	URL string
}

// Records struct contains variables pertaining to the records under
// consideration.
type Records struct {
	// ...total number of records under consideration
	Total int
	// ...LIMIT (number of records per page)
	Limit int
	// ...OFFSET
	Offset int
	// ...serial number of the first object in the current page
	From int
	// ...serial number of the last object in the current page
	To int
}

// Pages struct contains variables pertaining to the pages that were generated.
type Pages struct {
	// ...total number of pages
	Total int
	// ...current page number
	Number int
	// ...first page number
	First int
	// ...previous page number
	Previous int
	// ...sliding window of page numbers (on either side of the current page)
	Numbers []int
	// ...next page number
	Next int
	// ...last page number
	Last int
}

// NewPager creates a new instance of a Pager struct. It accepts the following
// variables and computes the remaining variables on-the-fly. The computed
// variables are persisted throughout the lifecycle of the Pager struct (no
// costly recomputation is performed).
func NewPager(recordsTotal int, recordsLimit int, pagesNumber int, pagesNumbers int, url string) *Pager {
	pager := &Pager{}
	pager.Records.Total = recordsTotal
	pager.Records.Limit = recordsLimit
	pager.Pages.Number = pagesNumber
	pager.URL = url
	if pager.Records.Total > 0 {
		pager.initPagesTotal()
		pager.initRecordsOffset()
		pager.initRecordsFrom()
		pager.initRecordsTo()
		pager.initPagesFirst()
		pager.initPagesPrevious()
		pager.initPagesNumbers(pagesNumbers)
		pager.initPagesNext()
		pager.initPagesLast()
	}
	return pager
}

// GetURL accepts a page number and returns a neatly formatted URL using the
// given URL syntax.
func (pager *Pager) GetURL(number int) string {
	url := fmt.Sprintf(pager.URL, number)
	return url
}

func (pager *Pager) initPagesTotal() {
	totalFloat := math.Ceil(float64(pager.Records.Total) / float64(pager.Records.Limit))
	totalInt := int(totalFloat)
	pager.Pages.Total = totalInt
}

func (pager *Pager) initRecordsOffset() {
	offset := (pager.Records.Limit * (pager.Pages.Number - 1))
	pager.Records.Offset = offset
}

func (pager *Pager) initRecordsFrom() {
	from := pager.Records.Offset + 1
	pager.Records.From = from
}

func (pager *Pager) initRecordsTo() {
	toFloat64 := math.Min(float64(pager.Records.Limit+pager.Records.From-1), float64(pager.Records.Total))
	toInt := int(toFloat64)
	pager.Records.To = toInt
}

func (pager *Pager) initPagesFirst() {
	first := 1
	pager.Pages.First = first
}

func (pager *Pager) initPagesPrevious() {
	if pager.Pages.Number == 1 {
		previous := pager.Pages.First
		pager.Pages.Previous = previous
		return
	}
	previous := pager.Pages.Number - 1
	pager.Pages.Previous = previous
}

func (pager *Pager) initPagesNumbers(numbers int) {
	for index := numbers; index > 0; index-- {
		number := pager.Pages.Number - index
		if number > 0 {
			pager.Pages.Numbers = append(pager.Pages.Numbers, number)
		}
	}
	number := pager.Pages.Number
	pager.Pages.Numbers = append(pager.Pages.Numbers, number)
	for index := 1; index <= numbers; index++ {
		number := pager.Pages.Number + index
		if number <= pager.Pages.Total {
			pager.Pages.Numbers = append(pager.Pages.Numbers, number)
		}
	}
}

func (pager *Pager) initPagesNext() {
	if pager.Pages.Number >= pager.Pages.Total {
		next := pager.Pages.Total
		pager.Pages.Next = next
		return
	}
	next := pager.Pages.Number + 1
	pager.Pages.Next = next
}

func (pager *Pager) initPagesLast() {
	last := pager.Pages.Total
	pager.Pages.Last = last
}
