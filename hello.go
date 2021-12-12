package hello

import "rsc.io/quote"
import "encoding/csv"
import "fmt"
import "strings"

func Hello() string {
	in := `first_name,last_name,username
Rob,Pike,rob
Ken,Thompson,ken
Robert,Griesemer,gri`

	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		fmt.Print("fatal error =", err)
	}

	fmt.Print(records)
	fmt.Print("\n")

    return quote.Hello()
}


