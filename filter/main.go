package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
	"unicode"

	"encoding/json"

	spineTypes "git.chotot.org/ad-service/spine/apps/types"
)

var (
	ruleFile = flag.String("file", "./json_rules.txt", "All the rules data in json format")
	dataFile = flag.String("data", "./data.txt", "data input")

	COMBINE_RULE_AND = 1
	COMBINE_RULE_OR  = 2
)

type (
	Rule struct {
		Id          int
		Target      *string
		ListItems   map[int]*ListItems
		Action      *int
		CombineRule int
	}

	ListItems struct {
		Id        int
		Ignore    bool
		Fields    []string
		WholeWord bool
		Items     *RadixTree
	}

	// ad struct {
	// 	Body     string `json:"body"`
	// 	Category string `json:"category"`
	// }
)

func parseBlockRules(blockRules []spineTypes.BlockRule) map[string](map[int]*Rule) {
	tmpRules := make(map[string](map[int]*Rule))

	for _, res := range blockRules {
		if res.RuleID == nil {
			continue
		}

		var ok bool
		var applicationRules map[int]*Rule

		if applicationRules, ok = tmpRules[*res.Application]; !ok {
			applicationRules = make(map[int]*Rule)
		}

		rule := (*Rule)(nil)
		if rule, ok = applicationRules[*res.RuleID]; !ok {
			rule = &Rule{
				Id:          *res.RuleID,
				Action:      res.Action,
				Target:      res.Target,
				ListItems:   make(map[int]*ListItems),
				CombineRule: res.CombineRule,
			}
		}

		l := ListItems{
			Id:        *res.ListID,
			WholeWord: *res.WholeWord,
			Fields:    strings.Split(*res.Fields, ","),
		}

		if res.Ignore != nil {
			l.Ignore = *res.Ignore
		}

		tmpRd := RadixTree{}

		if res.Items != nil {
			for _, item := range strings.Split(*res.Items, "badboyd") {
				// if !strings.Contains(item, "cọc") && item != "2020" && !strings.Contains(item, "không") {
				// 	continue
				// }
				// fmt.Printf("ListID %d Item %s\n", *res.ListID, item)
				lowerItem := strings.ToLower(item)
				if !*res.WholeWord {
					for _, vv := range strings.Split(lowerItem, " ") {
						tmpRd.Add(vv)
					}
				} else {
					tmpRd.Add(lowerItem)
				}
			}
		}
		l.Items = &tmpRd

		rule.ListItems[*res.ListID] = &l
		applicationRules[*res.RuleID] = rule
		tmpRules[*res.Application] = applicationRules
	}

	return tmpRules
}

func main() {
	data, err := ioutil.ReadFile(*dataFile)
	if err != nil {
		panic(err)
	}
	ad := map[string]string{}
	if err := json.Unmarshal(data, &ad); err != nil {
		panic(err)
	}

	rulesData, err := ioutil.ReadFile(*ruleFile)
	if err != nil {
		panic(err)
	}

	blockRules := []spineTypes.BlockRule{}
	if err = json.Unmarshal(rulesData, &blockRules); err != nil {
		panic(err)
	}

	rules := parseBlockRules(blockRules)["newad"]

	start := time.Now()
	defer func(t time.Time) {
		fmt.Print("Run time: %d", time.Since(t))
	}(start)

	for _, r := range rules {
		target, action := r.check(ad, nil)
		if target != nil {
			fmt.Printf("target %s", *target)
		}

		if action != nil {
			fmt.Printf(" action %d\n", *action)
			break
		}
	}
}

func (r *Rule) check(input map[string]string, targetPrio map[string]int) (*string, *int) {
	exist := false
	done := false

	// fmt.Printf("Start checking rule %d for input %+v\n", r.Id, input)
	for _, li := range r.ListItems {
		if li.Ignore {
			continue
		}
		for _, fi := range li.Fields {
			if val, ok := input[fi]; ok {
				// fmt.Printf("rule %d list %d check field %s\n", r.Id, li.Id, val)
				if exist = li.Items.Lookup(val); exist {
					// fmt.Printf("match rule %d items %d val-rule: %s\n", r.Id, li.Id, val)
					if r.CombineRule == COMBINE_RULE_OR {
						fmt.Printf("rule is or: ", r)
						done = true
					}
					break
				}
			}
		}

		if exist == false || done {
			break
		}
	}
	if exist {
		return r.Target, r.Action
	}

	return nil, nil
}

type (
	// RadixTree ...
	RadixTree struct {
		Root *Node
	}
	// Node of tree
	Node struct {
		IsLeaf   bool
		NextNode map[string]*Node
		Attrs    map[string]interface{}
	}
)

// Init tree
func (rt *RadixTree) Init() {
	rt.Root = &Node{
		IsLeaf:   false,
		NextNode: make(map[string]*Node),
		Attrs:    make(map[string]interface{}),
	}
}

// Add node to tree
func (rt *RadixTree) Add(input string) bool {
	if rt.Root == nil {
		rt.Init()
	}

	tmp := rt.Root
	nextNode := rt.Root
	hasNode := false
	for _, v := range strings.Split(input, " ") {
		// fmt.Printf("Add %s of %s\n", v, input)
		if nextNode, hasNode = tmp.NextNode[v]; !hasNode {
			nextNode = &Node{
				NextNode: make(map[string]*Node),
				Attrs:    make(map[string]interface{}),
			}

			tmp.NextNode[v] = nextNode
		}
		tmp = nextNode
	}

	tmp.IsLeaf = true
	tmp.Attrs["is_word"] = true

	return false
}

// Lookup input in tree
func (rt *RadixTree) Lookup(input string) bool {
	if rt.Root == nil {
		rt.Init()
	}

	wg := sync.WaitGroup{}
	fc := make(chan bool)

	tmp := input

	for {
		if len(tmp) == 0 {
			break
		}

		wg.Add(1)
		go func(in string, c chan bool) {
			defer wg.Done()
			c <- rt.lookup(in)
		}(tmp, fc)

		// tmp[1:] is for remove any space in the subject or name
		tmp = strings.TrimLeftFunc(tmp[1:], func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r)
		})
	}

	go func() {
		wg.Wait()
		close(fc)
	}()

	match := false
	for v := range fc {
		if v {
			match = v
		}
	}

	return match
}

// Lookup input in tree
func (rt *RadixTree) lookup(input string) bool {
	if rt.Root == nil {
		rt.Init()
	}

	tmp := rt.Root
	for _, v := range strings.Split(input, " ") {
		// fmt.Printf("Look %s in %s\n", v, input)
		if nextNode, ok := tmp.NextNode[v]; ok {
			// fmt.Printf("Found %s in %s\n", v, input)
			tmp = nextNode
			if isWord, ok := nextNode.Attrs["is_word"]; ok && isWord.(bool) {
				return true
			}
			continue
		}
		tmp = rt.Root
	}
	return false
}

// LookupWithWord input in tree return word
func (rt *RadixTree) LookupWithWord(input string) (string, bool) {
	if rt.Root == nil {
		rt.Init()
	}

	tmp := rt.Root
	phrase := ""
	for _, v := range strings.Split(input, " ") {
		if nextNode, ok := tmp.NextNode[v]; ok {
			fmt.Println(v)
			phrase += v + " "
			tmp = nextNode
			if isWord, ok := nextNode.Attrs["is_word"]; ok && isWord.(bool) {
				return strings.Trim(phrase, " "), true
			}
			continue
		}
		phrase = ""
		tmp = rt.Root
	}
	return "", false
}

// Del node from tree
func (rt *RadixTree) Del(input string) bool {
	return false
}
