package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var URL = "https://www.bi.go.id/id/moneter/informasi-kurs/transaksi-bi/Default.aspx"

func main() {
	var r []string
	var rw [][]string
	m := os.Args[1]

	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#ctl00_PlaceHolderMain_biWebKursTransaksiBI_GridView1").Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(j int, tr *goquery.Selection) {
			if j == 0 {
				return
			}
			tr.Find("td").Each(func(l int, td *goquery.Selection) {
				r = append(r, strings.TrimSpace(td.Text()))
			})
			rw = append(rw, r[:len(r)-1])
			r = nil
		})
	})

	for i := range rw {
		if m == rw[i][0] {
			fmt.Println("Kurs Terhadap:", m)
			fmt.Println("Nilai:", rw[i][1])
			fmt.Println("Kurs Jual: Rp.", rw[i][2])
			fmt.Println("Kurs Beli: Rp.", rw[i][3])
		}
	}
}
