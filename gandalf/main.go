package main

import (
	"fmt"
	"git.chotot.org/gandalf/wand"
)

func main() {
	gray := "127.0.0.1:13000"
	//gray := "10.60.3.234:32713"
	white := "10.60.3.234:30978"
	w := wand.New(wand.Config{GrayAddrs: gray, WhiteAddrs: white})
	rule, err := w.RetrieveRule("lytestreject23")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", rule.GetStatus() == wand.Rule_ONTRIAL)

	if err = w.WatchRule(rule); err != nil {
		panic(err)
	}
	q := wand.VerdictQuery{
		AdID:     39616036,
		ActionID: 1,
	}
	qv, err := w.GetVerdict(q)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", qv)
}
