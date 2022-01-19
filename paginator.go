package paginator

import (
	"fmt"
)

func GetPaginatedResponse(data []interface{}, pageSize int, currentPage int) (string, error) {

	total := len(data)
	lastPage := (total / pageSize) + 1
	isDataNull := total < 1
	isPageNumberValid := lastPage >= currentPage
	fromValue := (currentPage-1)*pageSize + 1
	toValue := currentPage * pageSize
	from := &fromValue
	to := &toValue

	if *to > total {
		to = &total
	}
	if *from > total {
		from = &total
	}

	if isDataNull {
		data = make([]interface{}, 0)
	}

	if isPageNumberValid { //page is not available so we are gonna replace the value with Null
		if !isDataNull {
			data = data[*from-1 : *to]
		}
	} else {
		data = make([]interface{}, 0)
		from = nil
		to = nil
	}
	marshalled, err := MarshalToString(map[string]interface{}{
		"data": data,
		"meta": Meta{
			PerPage:     pageSize,
			CurrentPage: currentPage,
			LastPage:    lastPage,
			From:        from,
			To:          to,
			Path:        fmt.Sprintf("?page=%v", currentPage),
			Total:       total,
		},
		"links": getLinks(currentPage, lastPage),
	})
	if err != nil {
		return "", fmt.Errorf("cannot convert data to json")
	}

	return marshalled, nil
}

func getLinks(currentPage int, totalPages int) links {
	var res links
	res.First = "?page=1"
	res.Last = fmt.Sprintf("?page=%v", totalPages)

	if currentPage > 1 {
		prev := fmt.Sprintf("?page=%v", currentPage-1)
		res.Prev = &prev
	}

	if currentPage < totalPages {
		next := fmt.Sprintf("?page=%v", currentPage+1)
		res.Next = &next
	}

	return res
}
