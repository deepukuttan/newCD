package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var cur string
	var new string
	var joined string
	//var prev string
	if len(os.Args) <= 1 {
		fmt.Println("[ARGUMENT ERROR] mycd takes only two input paths.")
		os.Exit(1)
	}
	if len(os.Args) > 3 {
		fmt.Println("[ARGUMENT ERROR] mycd takes only two input paths.")
		os.Exit(1)
	}
	cur = os.Args[1]
	new = os.Args[2]

	curs := strings.Split(cur, "/")
	news := strings.Split(new, "/")
	err := false
	//curl := len(news)

	for _, elm := range news {
		sts := false
		if elm != ".." && elm != "" {
			for _, el := range curs {
				if el == elm && el != ".." {
					joined += el + "/"
					sts = true
					continue
				}
			}
			if sts != true {
				if 
				joined += elm
			}
		} else if elm != "" {
			for _, el := range curs[:(len(curs) - 1)] {
				if strings.Contains(el,".."){
					err = true
					break
				}
				joined += el + "/"
				continue
			}
		}
		if err{
			break
		}
	}

	if(err){
		fmt.Println("[ERROR]")
		os.Exit(1)
	}

	if len(joined) == 0 {
		joined += "/"
	}

	fmt.Println(joined)

}
