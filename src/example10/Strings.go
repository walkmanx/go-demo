package main

import (
	"fmt"
	"strings"
)

func TestContain() {
	// Contains
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}

func TestJoin() {
	// Join
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	//运行结果:foo, bar, baz
}

func TestIndex() {
	// Index
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	//运行结果:
	//    4
	//    -1
}

func TestRepeat() {
	// Repeat
	fmt.Println("ba" + strings.Repeat("na", 2))
	//运行结果:banana
}

func TestReplace() {
	// Replace
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//运行结果:
	//oinky oinky oink
	//moo moo moo
}

func TestSplit() {
	// Split
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	//运行结果:
	//["a" "b" "c"]
	//["" "man " "plan " "canal panama"]
	//[" " "x" "y" "z" " "]
	//[""]
}

func TestTrim() {
	// Trim
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
	//运行结果:["Achtung"]
}

func TestFields() {
	// Fields
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//运行结果:Fields are: ["foo" "bar" "baz"]
}

func TestAppend() {
	// Append
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10) //以10进制方式追加
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')

	fmt.Println(string(str)) //4567false"abcdefg"'单'
}

func TestFormat() {
	a := strconv.FormatBool(false)
	b := strconv.FormatInt(1234, 10)
	c := strconv.FormatUint(12345, 10)
	d := strconv.Itoa(1023)

	fmt.Println(a, b, c, d) //false 1234 12345 1023
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func TestParse() {
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e) //false 123.23 1234 12345 1023
}

func main() {

}
