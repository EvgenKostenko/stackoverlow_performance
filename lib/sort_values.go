package lib

import "sort"

//type sortedMap struct {
//	m map[int]int
//	s []int
//}
//
//func (sm *sortedMap) Len() int {
//	return len(sm.m)
//}
//
//func (sm *sortedMap) Less(i, j int) bool {
//	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
//}
//
//func (sm *sortedMap) Swap(i, j int) {
//	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
//}
//
//func SortedKeysInt(m map[int]int) []int {
//	sm := new(sortedMap)
//	sm.m = m
//	sm.s = make([]int, len(m))
//	i := 0
//	for key, _ := range m {
//		sm.s[i] = key
//		i++
//	}
//	sort.Sort(sm)
//	return sm.s
//}




type Pair struct {
	Key string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SortedKeysStr(m map[string]int) PairList {
	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}


type PairInt struct {
	Key int
	Value int
}

type PairIntList []PairInt

func (p PairIntList) Len() int {
	return len(p)
}

func (p PairIntList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}

func (p PairIntList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SortedKeysInt(m map[int]int) PairIntList {
	pl := make(PairIntList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = PairInt{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

