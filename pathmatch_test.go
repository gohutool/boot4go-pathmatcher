package path

import (
	"fmt"
	util4go "github.com/gohutool/boot4go-util"
	"path"
	"testing"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : pathmatch_test.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/5/3 10:31
* 修改历史 : 1. [2022/5/3 10:31] 创建文件 by LongYong
*/
/*
* 字符测试实例

//* 匹配0或多个非/的字符
path.Match("*", "a")            // true
path.Match("*", "sefesfe/")     // false
? 字符测试实例

//？匹配一个非/的字符
path.Match("a?b", "aab")    // true
path.Match("a?b", "a/b")    // false
[] 格式测试实例

path.Match("[abc][123]", "b2")        // true

path.Match("[abc][1-3]", "b2")        // true

path.Match("[abc][^123]", "b2")        // false

path.Match("[abc][^123]", "b4")        // true
字符或者特殊用途字符(  \\   ?  *   [ )测试实例

path.Match("a\\\\", "a\\")	// true
path.Match("a\\[", "a[")		// true
path.Match("a\\?", "a?")		// true
path.Match("a\\*", "a*")		// true
path.Match("abc", "abc")		// true

*/
func TestPathMatch(t *testing.T) {
	printMatch("*", "a")
	printMatch("*", "sefesfe/")
	printMatch("*/b", "b")
	printMatch("/*/b", "b")
	printMatch("/*", "/b")
	printMatch("/*/b", "/b")
	printMatch("**/b", "/b")

	printMatch("/api/*", "/api/a/a/b")
	printMatch("/api/*/?b", "/api/a/ab")

	printMatch("/*", "/a")
	printMatch("/*", "/")
	printMatch("/*", "/abc/abd")

	printMatch("a?b", "a0b")
	printMatch("a?b", "a/b")
	printMatch("a?b", "abb")
	printMatch("a?b", "axb")

	printMatch("a\\\\", "a\\")
	printMatch("a\\[", "a[")
	printMatch("a\\?", "a?")
	printMatch("a\\*", "a*")
	printMatch("abc", "abc")
}

func printMatch(a, b string) {
	if r, err := path.Match(a, b); err != nil {
		fmt.Printf("Error %v\n", err)
	} else {
		fmt.Printf("path.Match(%q, %q)=%v\n", a, b, r)
	}
}

type TestCase struct {
	match, sample string
	isMatch       bool
	isStd         bool
	error         bool
}

var CaseSample []TestCase

func init() {
	CaseSample = make([]TestCase, 0, 300)

	CaseSample = append(CaseSample, TestCase{
		match: "*", sample: "", isMatch: true, isStd: true, error: false,
	})
	for idx := 'a'; idx < 'f'; idx++ {
		CaseSample = append(CaseSample, TestCase{
			match: "*", sample: string(idx), isMatch: true, isStd: true, error: false,
		})
	}
	for idx := 0; idx < 10; idx++ {
		v := ""
		for i := 0; i < 3; i++ {
			v = v + string(util4go.RandRune())
		}
		CaseSample = append(CaseSample, TestCase{
			match: "*", sample: v, isMatch: true, isStd: true, error: false,
		})
	}

	CaseSample = append(CaseSample, TestCase{
		match: "*", sample: "/", isMatch: false, isStd: true, error: false,
	})
	for idx := 2; idx < 4; idx++ {
		CaseSample = append(CaseSample, TestCase{
			match: "*", sample: "/" + util4go.LeftPad("a", idx, util4go.RandRune()), isMatch: false, isStd: true, error: false,
		})
	}
	CaseSample = append(CaseSample, TestCase{
		match: "/*", sample: "/", isMatch: true, isStd: true, error: false,
	})
	for idx := 2; idx < 4; idx++ {
		CaseSample = append(CaseSample, TestCase{
			match: "/*", sample: "/" + util4go.LeftPad("a", idx, util4go.RandRune()), isMatch: true, isStd: true, error: false,
		})
	}

	for idx := 0; idx < 10; idx++ {
		v := "/"
		for i := 0; i < 3; i++ {
			v = v + string(util4go.RandRune())
		}
		CaseSample = append(CaseSample, TestCase{
			match: "/*", sample: v, isMatch: true, isStd: true, error: false,
		})
	}

	CaseSample = append(CaseSample, TestCase{
		match: "*c", sample: "abc", isMatch: true, isStd: true, error: false,
	})

	for idx := 1; idx < 4; idx++ {
		CaseSample = append(CaseSample, TestCase{
			match: "*c", sample: util4go.LeftPad("c", idx, util4go.RandRune()), isMatch: true, isStd: true, error: false,
		})
	}

	CaseSample = append(CaseSample, TestCase{
		match: "*c", sample: util4go.RightPad("c", 2, util4go.RandRune()) + "/", isMatch: false, isStd: true, error: false,
	})
	CaseSample = append(CaseSample, TestCase{
		match: "*c", sample: "/" + util4go.RightPad("c", 2, util4go.RandRune()), isMatch: false, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "a*/b", sample: "abc/b", isMatch: true, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "a*/b", sample: "a/c/b", isMatch: false, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "a*b*c*d*e*", sample: "axbxcxdxe", isMatch: true, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "a*b*/a", sample: "axbxxxx/a", isMatch: true, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "ab[c]", sample: "abc", isMatch: true, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "ab[b-d]", sample: "abc", isMatch: true, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "ab[b-d]", sample: "abf", isMatch: false, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "ab[^c]", sample: "abc", isMatch: false, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "ab[^b-d]", sample: "abf", isMatch: true, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "ab[^b-d]", sample: "abf", isMatch: true, isStd: true, error: false,
	})

	CaseSample = append(CaseSample, TestCase{
		match: "*.log", sample: "a/a/a.log", isMatch: false, isStd: false, error: false,
	})
	CaseSample = append(CaseSample, TestCase{
		match: "**/*.log", sample: "a/a/a.log", isMatch: true, isStd: false, error: false,
	})
	CaseSample = append(CaseSample, TestCase{
		match: "[a-b-d]", sample: "a", isMatch: true, isStd: false, error: false,
	})
	CaseSample = append(CaseSample, TestCase{
		match: "[a-bb-d]*", sample: "d", isMatch: true, isStd: false, error: false,
	})
	CaseSample = append(CaseSample, TestCase{
		match: "[a-be-g]*", sample: "f", isMatch: true, isStd: false, error: false,
	})
	CaseSample = append(CaseSample, TestCase{
		match: "[a-be-g]*", sample: "c", isMatch: false, isStd: false, error: false,
	})

	for idx := 0; idx < 10; idx++ {
		CaseSample = append(CaseSample, TestCase{
			match: "[a-fo-z]*", sample: string(util4go.RandRune2('a', 'f')) + string(util4go.RandRune()) + string(util4go.RandRune()),
			isMatch: true, isStd: false, error: false,
		})
	}

	for idx := 0; idx < 10; idx++ {
		CaseSample = append(CaseSample, TestCase{
			match: "[a-fo-z]*", sample: string(util4go.RandRune2('o', 'z')) + string(util4go.RandRune()) + string(util4go.RandRune()),
			isMatch: true, isStd: false, error: false,
		})
	}

	for idx := 0; idx < 10; idx++ {
		CaseSample = append(CaseSample, TestCase{
			match: "[a-fo-z]*", sample: string(util4go.RandRune2('g', 'm')) + string(util4go.RandRune()) + string(util4go.RandRune()),
			isMatch: false, isStd: false, error: false,
		})
	}

	CaseSample = append(CaseSample, TestCase{
		match: "ab[!c]", sample: "abc", isMatch: false, isStd: false, error: false,
	})

	for idx := 0; idx < 10; idx++ {
		CaseSample = append(CaseSample, TestCase{
			match: "ab[!c]", sample: "ab" + string(util4go.RandRune2('d', 'z')), isMatch: true, isStd: false, error: false,
		})
	}
}

func TestCaseMatch(t *testing.T) {

	t.Logf("TestCase : %v", len(CaseSample))

	for idx, tt := range CaseSample {
		// Since Match() always uses "/" as the separator, we
		// don't need to worry about the tt.testOnDisk flag
		matchRun(t, idx, tt)
	}
}

func runAll() {
	for idx, tt := range CaseSample {
		// Since Match() always uses "/" as the separator, we
		// don't need to worry about the tt.testOnDisk flag
		run(idx, tt)
	}
}

func run(idx int, tt TestCase) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("#%v. Match(%#q, %#q) panicked: %#v", idx, tt.match, tt.sample, r)
		}
	}()

	// Match() always uses "/" as the separator
	matched, err := Match(tt.match, tt.sample)
	if matched != tt.isMatch || (err != nil && !tt.error) {
		fmt.Errorf("#%v. Match(%#q, %#q) = %v, %v. but expect %v, %v", idx, tt.match, tt.sample, matched, err, tt.isMatch, err)
	}

	if tt.isStd {
		stdOk, err := path.Match(tt.match, tt.sample)
		if matched != stdOk || (err != nil && !tt.error) {
			fmt.Errorf("#%v. path.Match(%#q, %#q) = %v, %v. but expect %v, %v", idx, tt.match, tt.sample, matched, err, stdOk, err)
		}
	}
}

func matchRun(t *testing.T, idx int, tt TestCase) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("#%v. Match(%#q, %#q) panicked: %#v", idx, tt.match, tt.sample, r)
		}
	}()

	// Match() always uses "/" as the separator
	matched, err := Match(tt.match, tt.sample)
	if matched != tt.isMatch || (err != nil && !tt.error) {
		t.Errorf("#%v. Match(%#q, %#q) = %v, %v. but expect %v, %v", idx, tt.match, tt.sample, matched, err, tt.isMatch, err)
	}

	if tt.isStd {
		stdOk, err := path.Match(tt.match, tt.sample)
		if matched != stdOk || (err != nil && !tt.error) {
			t.Errorf("#%v. path.Match(%#q, %#q) = %v, %v. but expect %v, %v", idx, tt.match, tt.sample, matched, err, stdOk, err)
		}
	}
}

func BenchmarkPathMatcher(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, c := range CaseSample {
			Match(c.match, c.sample)
		}
	}
}

func BenchmarkGoStdMatch(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, c := range CaseSample {
			path.Match(c.match, c.sample)
		}
	}
}

func BenchmarkGoStdMatchParallel(b *testing.B) {
	b.SetParallelism(8)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			runAll()
		}
	})
}
