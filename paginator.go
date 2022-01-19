package paginator

import (
	"fmt"
	"strings"
)

func GetPaginatedResponse(data []interface{}, pageSize int, currentPage int) (string, error) {

	total := len(data)
	lastPage := (total / pageSize) + 1
	isDataNull := total < 1
	isPageNumberValid := lastPage >= currentPage
	from := (currentPage-1)*pageSize + 1
	to := currentPage * pageSize

	if to > total {
		to = total
	}
	if from > total {
		from = total
	}

	if isDataNull {
		data = make([]interface{}, 0)
	}

	if isPageNumberValid { //page is not available so we are gonna replace the value with Null
		if !isDataNull {
			data = data[from-1 : to]
		}
	} else {
		data = make([]interface{}, 0)
		from = -1
		to = -1
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

	marshalled = strings.Replace(marshalled, `"-PAGE_IS_NULL-"`, "null", 2) //this is for the links

	if !isPageNumberValid {
		marshalled = strings.Replace(marshalled, `"from": -1`, `"from": null`, 1)
		marshalled = strings.Replace(marshalled, `"to": -1`, `"to": null`, 1)
	}

	return marshalled, nil
}

func getLinks(currentPage int, totalPages int) links {
	var res links
	res.First = "?page=1"
	res.Last = fmt.Sprintf("?page=%v", totalPages)
	res.Prev = "-PAGE_IS_NULL-"
	res.Next = "-PAGE_IS_NULL-"

	if currentPage > 1 {
		res.Prev = fmt.Sprintf("?page=%v", currentPage-1)
	}

	if currentPage < totalPages {
		res.Next = fmt.Sprintf("?page=%v", currentPage+1)
	}

	return res
}
