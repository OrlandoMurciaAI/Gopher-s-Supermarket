package market

//This Module is not done yet.
import (
	"encoding/csv"
	"fmt"
	"os"
	cs "packages/clients"
)

var data []byte
var records [][]string

func init() {
	data, err := os.Open("/files/groseries.csv")
	if err != nil {
		panic(err)
	}
	// Read the file
	r := csv.NewReader(data)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println("Exists : ", len(records))
}

func Buy(c cs.Customer) {

}
