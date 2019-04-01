package b

import "go.mod/a"

func Init() {
	a.Direction = make(map[string]string)
}

func Add(k, v string) {
	a.Direction[k] = v
}

func Get(k string) string {
	return a.Direction[k]
}
