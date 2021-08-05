package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	strsLen := len(strs)
	strMap := make(map[string]int)

	for i := 0; i < strsLen; i++ {
		strSlice := []string{strs[i]}

		sortStr := sortString(strs[i])
		sLen, ok := strMap[sortStr]

		for j := i + 1; j < strsLen; j++ {
			jSortstr := sortString(strs[j])
			sLen2, ok2 := strMap[jSortstr]
			if sortStr != jSortstr {
				if !ok2 {
					break
				}
				result[sLen2] = append(result[sLen2], strs[j])
			}

			if ok {
				result[sLen] = append(result[sLen], strs[j])
			} else if !ok2 {
				strSlice = append(strSlice, strs[j])
			}

			i++
		}

		if !ok {
			result = append(result, strSlice)
			strMap[sortStr] = len(strMap)
		}

	}

	return result
}

func test01() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	succResult := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	fmt.Println("test01 strs:=", strs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", groupAnagrams(strs))
	fmt.Println()
}

func test02() {
	strs := []string{""}
	succResult := [][]string{{""}}
	fmt.Println("test02 strs:=", strs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", groupAnagrams(strs))
	fmt.Println()
}

func test03() {
	strs := []string{"a"}
	succResult := [][]string{{"a"}}
	fmt.Println("test03 strs:=", strs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", groupAnagrams(strs))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
