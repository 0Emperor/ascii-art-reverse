package ascii

import (
	"fmt"
)

func Reverse(fileName string) {
	a, err := ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	b := [][]string{}
	n := []string{}
	count := 0
	for _, v := range a {

		if count == 8 {
			b = append(b, n)
			n = []string{}
			count = 1
		} else if v == "\n" {
			b = append(b, n)
		} else {
			count++
		}
		n = append(n, v)
	}

	shaadow, _ := ReadFile("banners/shadow.txt")
	standard, _ := ReadFile("banners/standard.txt")
	thinkertoy, _ := ReadFile("banners/thinkertoy.txt")
	c := [][]string{}
	for _, v := range b {
		o := split(v)
		c = append(c, o)
	}
	j := ""
	space := []string{" ", " ", " ", " ", " ", " ", " ", " "}
	cc := 0
	for r, v := range c {
		for i := 0; i < len(v); i += 8 {
			t := v[i : i+8]
			if compare(t, space) {
				cc++
				if cc == 6 {
					j += " "
					cc = 0
				}
			} else {
				io, oi := contains(standard, t)
				if io {
					j += string(rune((oi / 9) + 32))
				}
				io, oi = contains(shaadow, t)
				if io {
					j += string(rune((oi / 9) + 32))
				}
				io, oi = contains(thinkertoy, t)
				if io {
					j += string(rune((oi / 9) + 32))
				}
			}
		}
		if r != len(c)-1 {
			j += "\\n"
		}

	}
	fmt.Println(j)
}

func compare(v, x []string) bool {
	if len(v) != len(x) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if v[i] != x[i] {
			return false
		}
	}
	return true
}

func split(v []string) []string {
	p := []string{}
	for _, i := range v {
		if i != "" {
			for o := 0; o < len(i); o++ {
				y := i[o]

				if y == ' ' {
					if check(v, o) {
						for e := 0; e < 8; e++ {
							f := ""
							f += v[e][:o+1]
							p = append(p, f)
						}

						for e := 0; e < 8; e++ {
							if len(v) == 0 {
								break
							}
							v[e] = v[e][o+1:]
						}
						i = i[o+1:]
						o = -1
					}
				}
			}
			break
		} else {
			p = append(p, "")
		}
	}

	return p
}

func contains(s, u []string) (bool, int) {
	if len(u) == 0 {
		return true, 0 // An empty slice is always a subsequence.
	}
	if len(s) < len(u) {
		return false, 0 // `u` cannot be contained in `s` if it's longer.
	}

	for i := 0; i <= len(s)-len(u); i++ {
		match := true
		for j := 0; j < len(u); j++ {
			if s[i+j] != u[j] {
				match = false
				break
			}
		}
		if match {
			return true, i // Return true and the starting index.
		}
	}
	return false, 0 // Return false if no match found.
}

func check(i []string, o int) bool {
	for _, v := range i {
		if o < len(v) {
			if v[o] != ' ' {
				return false
			}
		}
	}
	return true
}
