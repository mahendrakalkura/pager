package pager

import (
	"fmt"
	"math"
)

type pagerRecords struct {
	Total  int
	Limit  int
	Offset int
	From   int
	To     int
}

type pagerPages struct {
	Total    int
	Number   int
	First    int
	Previous int
	Numbers  []int
	Next     int
	Last     int
}

// Pager ...
type Pager struct {
	Records pagerRecords
	Pages   pagerPages
	URL     string
}

// NewPager ...
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

// GetURL ...
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
