package string_test

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) //string会被初始化为""
	s = "hello"
	t.Log(s[1])
	// s[1] = "3" //string是不可变的byte slice
	t.Log(s, len(s))
	s = "\xE4\xB8\xA5" //16进制串
	t.Log(s)
	s = "中"
	t.Log(len(s)) //byte数
	c := []rune(s)
	t.Logf("杜 Unicode %x", c[0])
	t.Logf("杜 UTF-8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d %[1]x", c) //%[1]c代表将第一个参数以%c格式化
	}
}

func TestStringSplit(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for id, part := range parts {
		t.Logf("No %d is : %v", id, part)
	}
	s2 := strings.Join(parts, "-")
	t.Log(s2)
}
