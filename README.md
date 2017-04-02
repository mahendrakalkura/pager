# Pager

[![Build Status](https://travis-ci.org/mahendrakalkura/pager.png?branch=master)](https://travis-ci.org/mahendrakalkura/pager)

[![GoDoc](https://godoc.org/github.com/mahendrakalkura/pager?status.svg)](https://godoc.org/github.com/mahendrakalkura/pager)

Pager is a library which you can use for the purpose of generating and rendering pagination links.

## How to install?

```
$ go get github.com/mahendrakalkura/pager
```

## How to use?

**Step 1:** Initialize the pager.

```
import (
    "github.com/mahendrakalkura/pager"
)

// ...total number of items under consideration
count := 100

// ...number of items per page
limit := 10

// ...current page number
page := 1

// ...the size of the sliding window of page numbers
//    (on either side of the current page)
numbers := 3

// ...the URL syntax
//    (%d will be replaced with the page number in context)
url := "/%d/"

p = pager.NewPager(count, settings.Container.Pagers.Limit, page, settings.Container.Pagers.Numbers, url)

```

**Step 2:** Use it in your view.

```
query := "SELECT * FROM ... WHERE ... ORDER BY ... LIMIT %d OFFSET %d"
query = fmt.Sprintf(query, p.Records.Limit, p.Records.Offset)
```

**Step 3:** Use it in your template.

```
<div>
    <p class="pagination pull-right">
        Showing {{ $.p.Records.From }} to {{ $.p.Records.To }} of {{ $.p.Records.Total }}
        items
    </p>
    <ul class="pagination">
        <li {{ if eq $.p.Pages.Number 1 }}class="disabled"{{ end }}>
            <a href="{{ $.p.GetURL $.p.Pages.First }}">
                <span>&laquo;</span>
            </a>
        </li>
        <li {{ if eq $.p.Pages.Number $.p.Pages.Previous }}class="disabled"{{ end }}>
            <a href="{{ $.p.GetURL $.p.Pages.Previous }}">
                <span>&lt;</span>
            </a>
        </li>
        {{ range $.p.Pages.Numbers }}
            <li {{ if eq $.p.Pages.Number . }}class="disabled"{{ end }}>
                <a href="{{ $.p.GetURL . }}">
                    <span>{{ . }}</span>
                </a>
            </li>
        {{ end }}
        <li {{ if eq $.p.Pages.Number $.p.Pages.Next }}class="disabled"{{ end }}>
            <a href="{{ $.p.GetURL $.p.Pages.Next }}">
                <span>&gt;</span>
            </a>
        </li>
        <li {{ if eq $.p.Pages.Number $.p.Pages.Last }}class="disabled"{{ end }}>
            <a href="{{ $.p.GetURL $.p.Pages.Last }}">
                <span>&raquo;</span>
            </a>
        </li>
    </ul>
</div>
```
