package ch5

import "sort"

func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
