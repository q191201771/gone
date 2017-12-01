package main

import (
	"testing"
)

// 测试replace函数
func TestReplace(t *testing.T) {
	var tmpl string
	mapping := make(map[string]string)
	tmpl = "name={name}"
	mapping = map[string]string{
		"name": "pengrl",
	}
	assert(t, replace(tmpl, mapping), "name=pengrl", "")

	tmpl = "name={name},sex={sex}"
	mapping = map[string]string{
		"name": "pengrl",
		"sex":  "man",
	}
	assert(t, replace(tmpl, mapping), "name=pengrl,sex=man", "")

	tmpl = "name={name},sex={sex}"
	mapping = map[string]string{
		"name": "pengrl",
	}
	assert(t, replace(tmpl, mapping), "name=pengrl,sex={sex}", "")

	tmpl = "name={name},sex={sex}"
	mapping = map[string]string{
		"name":   "pengrl",
		"sex":    "man",
		"number": "7",
	}
	assert(t, replace(tmpl, mapping), "name=pengrl,sex=man", "")

}

func TestSubstr(t *testing.T) {
	assert(t, substr("http://audio.link.inke.cn:7788/publish/trans/inkeaudio/mlinkm/{liveid}", "trans/", "/"), "inkeaudio", "")
}
