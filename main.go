// Package main provides main ﳑ
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"ig/module"

	"github.com/go-resty/resty/v2"
)

// Variables -
type Variables struct {
	ID    string `json:"id"`
	First int64  `json:"first"`
	After string `json:"after"`
}

//
const (
	BASEURL        = "https://www.instagram.com/"
	QUERYURI       = BASEURL + "graphql/query/"
	QUERYHASH      = "69cba40317214236af40e7efa697781d"
	RELITHEID      = "33374582442"
	PARISSSTREETID = "6716845386"
)


//
var (
    // SLEEPTIME - 每次執行完後睡眠多久(s)
    SLEEPTIME = 3 
    // PAGESIZE - 每次處理幾篇貼文
    PAGESIZE int64 = 12
    // ID - 
    ID = RELITHEID
    // AFTER - 
    AFTER = ""
)


func main() {

	data := make([]module.EdgesNode, 0)

	client := resty.New()


	file, err := os.Create("temp.csv")
	if err != nil {
		panic(err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	total := 0

	for {

		r := client.NewRequest()
		variables := Variables{
			ID:    ID,
			First: PAGESIZE,
			After: AFTER,
		}

		vj, err := json.Marshal(variables)
		if err != nil {
			log.Fatal(err)
		}

		r = r.SetQueryParam("query_hash", QUERYHASH)
		r = r.SetQueryParam("variables", string(vj))

		resp := &module.Response{}
		r.SetResult(resp)

		_, err = r.Get(QUERYURI)


		if err != nil {
			log.Fatal(err)
		}

		for _, v := range resp.Data.User.Media.Edges {
			total++
			text := func() string {
				tmp := ""
				for _, vv := range v.Node.EdgeMediaToCaption.Edges {
					tmp += vv.Node.Text
				}
				return tmp
			}()
			w.Write([]string{fmt.Sprintf("%sp/%s", BASEURL, v.Node.Shortcode), text})
			data = append(data, v.Node)
			fmt.Println(total, text)
		}

		if !resp.Data.User.Media.PageInfo.HasNextPage {
			break
		}

		AFTER = resp.Data.User.Media.PageInfo.EndCursor
		time.Sleep(time.Duration(SLEEPTIME) * time.Second)
	}

    fmt.Println("Close the Crawler " , "Total:" , total)
}
