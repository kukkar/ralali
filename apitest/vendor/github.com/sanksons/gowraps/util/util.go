package util

import (
	"fmt"
	"strings"

	"github.com/go-errors/errors"
)

func GetMultiInsertQuery(table string, fields []string, values [][]interface{}) string {

	sqlqp := make([]string, 0, len(values))
	var fvalues []interface{} = make([]interface{}, 0, len(fields)*len(values))
	for _, vSet := range values {
		s := make([]string, 0, len(fields))

		for _, v1 := range vSet {
			s = append(s, "?")
			fvalues = append(fvalues, v1)
		}
		qstring := fmt.Sprintf("(%s)", strings.Join(s, ", "))
		sqlqp = append(sqlqp, qstring)
	}
	vstring := strings.Join(sqlqp, ",")
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES %s`,
		table, strings.Join(fields, ","), vstring)

	return query
}

func GeneratePlaceholderString(len int, placeholder string) string {
	h := GeneratePlaceholderSlice(len, placeholder)
	return strings.Join(h, ",")
}

func GeneratePlaceholderSlice(len int, placeholder string) []string {
	h := make([]string, 0, len)
	for i := 0; i < len; i++ {
		h = append(h, "?")
	}
	return h
}

//set related functions
func OuterJoinInt(set1 []int, set2 []int) (s1 []int, s2 []int) {
	set1Size := len(set1)
	set2Size := len(set2)
	if set1Size == 0 || set2Size == 0 {
		return set1, set2
	}
	s1 = make([]int, 0)
	for _, one := range set1 {
		var isofuse bool = true
		for _, two := range set2 {
			if one == two {
				isofuse = false
			}
		}
		if isofuse {
			s1 = append(s1, one)
		}
	}
	s2 = make([]int, 0)
	for _, two := range set2 {
		var isofuse bool = true
		for _, one := range set1 {
			if one == two {
				isofuse = false
			}
		}
		if isofuse {
			s2 = append(s2, two)
		}
	}
	return
}

//panic
func PanicHandler(msg string) {
	if r := recover(); r != nil {
		fmt.Println("Attention!!!   Panic Occurred !!!")
		fmt.Println(msg)
		fmt.Println(errors.Wrap(r, 2).ErrorStack())
	}
}
