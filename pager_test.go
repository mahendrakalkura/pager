package pager

import (
	"reflect"
	"testing"
)

func TestPager1(t *testing.T) {
	count := 100
	limit := 10
	number := 1
	numbers := 3
	url := "/%d/"
	pager := NewPager(count, limit, number, numbers, url)
	if pager.Records.Total != 100 {
		t.Error("Invalid pager.Records.Total")
	}
	if pager.Records.Limit != 10 {
		t.Error("Invalid pager.Records.Limit")
	}
	if pager.Records.Offset != 0 {
		t.Error("Invalid pager.Records.Offset")
	}
	if pager.Records.From != 1 {
		t.Error("Invalid pager.Records.From")
	}
	if pager.Records.To != 10 {
		t.Error("Invalid pager.Records.To")
	}
	if pager.Pages.Total != 10 {
		t.Error("Invalid pager.Pages.Total")
	}
	if pager.Pages.Number != 1 {
		t.Error("Invalid pager.Pages.Number")
	}
	if pager.Pages.First != 1 {
		t.Error("Invalid pager.Pages.First")
	}
	if pager.Pages.Previous != 1 {
		t.Error("Invalid pager.Pages.Previous")
	}
	if !reflect.DeepEqual(pager.Pages.Numbers, []int{1, 2, 3, 4}) {
		t.Error("Invalid pager.Pages.Numbers")
	}
	if pager.Pages.Next != 2 {
		t.Error("Invalid pager.Pages.Next")
	}
	if pager.Pages.Last != 10 {
		t.Error("Invalid pager.Pages.Last")
	}
	if pager.GetURL(1) != "/1/" {
		t.Error("Invalid pager.GetURL()")
	}
}

func TestPager2(t *testing.T) {
	count := 100
	limit := 10
	number := 10
	numbers := 5
	url := "/%d/"
	pager := NewPager(count, limit, number, numbers, url)
	if pager.Records.Total != 100 {
		t.Error("Invalid pager.Records.Total")
	}
	if pager.Records.Limit != 10 {
		t.Error("Invalid pager.Records.Limit")
	}
	if pager.Records.Offset != 90 {
		t.Error("Invalid pager.Records.Offset")
	}
	if pager.Records.From != 91 {
		t.Error("Invalid pager.Records.From")
	}
	if pager.Records.To != 100 {
		t.Error("Invalid pager.Records.To")
	}
	if pager.Pages.Total != 10 {
		t.Error("Invalid pager.Pages.Total")
	}
	if pager.Pages.Number != 10 {
		t.Error("Invalid pager.Pages.Number")
	}
	if pager.Pages.First != 1 {
		t.Error("Invalid pager.Pages.First")
	}
	if pager.Pages.Previous != 9 {
		t.Error("Invalid pager.Pages.Previous")
	}
	if !reflect.DeepEqual(pager.Pages.Numbers, []int{5, 6, 7, 8, 9, 10}) {
		t.Error("Invalid pager.Pages.Numbers")
	}
	if pager.Pages.Next != 10 {
		t.Error("Invalid pager.Pages.Next")
	}
	if pager.Pages.Last != 10 {
		t.Error("Invalid pager.Pages.Last")
	}
	if pager.GetURL(1) != "/1/" {
		t.Error("Invalid pager.GetURL()")
	}
}

func TestPager3(t *testing.T) {
	count := 0
	limit := 10
	number := 1
	numbers := 5
	url := "/%d/"
	pager := NewPager(count, limit, number, numbers, url)
	if pager.Records.Total != 0 {
		t.Error("Invalid pager.Records.Total")
	}
	if pager.Records.Limit != 10 {
		t.Error("Invalid pager.Records.Limit")
	}
	if pager.Records.Offset != 0 {
		t.Error("Invalid pager.Records.Offset")
	}
	if pager.Records.From != 0 {
		t.Error("Invalid pager.Records.From")
	}
	if pager.Records.To != 0 {
		t.Error("Invalid pager.Records.To")
	}
	if pager.Pages.Total != 0 {
		t.Error("Invalid pager.Pages.Total")
	}
	if pager.Pages.Number != 1 {
		t.Error("Invalid pager.Pages.Number")
	}
	if pager.Pages.First != 0 {
		t.Error("Invalid pager.Pages.First")
	}
	if pager.Pages.Previous != 0 {
		t.Error("Invalid pager.Pages.Previous")
	}
	if len(pager.Pages.Numbers) != 0 {
		t.Error("Invalid pager.Pages.Numbers")
	}
	if pager.Pages.Next != 0 {
		t.Error("Invalid pager.Pages.Next")
	}
	if pager.Pages.Last != 0 {
		t.Error("Invalid pager.Pages.Last")
	}
	if pager.GetURL(1) != "/1/" {
		t.Error("Invalid pager.GetURL()")
	}
}
