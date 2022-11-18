package main

import (
	"fmt"
	"os"
	"strings"
)

func sub(a []string, c string) bool {
	var sts bool = false
	for _, k := range a {
		if c == k {
			fmt.Println(k, c)
			sts = true
			break
		}
	}
	return sts
}

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
	temp := curs
	//err := false
	//curl := len(news)

	for _, m := range news {
		if !sub(curs, m) && m != ".." && m != "" {
			temp = append(temp, m)
		} else if !sub(curs, m) && m == ".." && m != "" {
			temp = temp[:len(temp)-1]
		}
	}

	for _, k := range temp {
		joined += k + "/"
	}

	if len(joined) == 0 {
		joined = "/"
	}

	fmt.Println(joined)
}
