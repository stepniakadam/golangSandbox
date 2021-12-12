package exercise1

import (
	"io"
	"os" 
	"log"
	"encoding/csv"
)

type Person struct {
	Name string
	Age int
}

func ReadPersons() []Person {
	var f, e = os.OpenFile("filesHandling/5m_Sales_Records.csv", os.O_RDONLY, 0755) 
	if e != nil{
		log.Fatal(e)
		return nil
	}
	
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(record) == 0 {
			log.Fatal("Empty record found")
		}
	}

	return nil
}