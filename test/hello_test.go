package test
import (
    "localhost/ltian/test/world"
    "testing"
)

/**
使用 go test 命令进行单元测试有两个要求：
第一：测试文件必须以_test 结尾；
第二：测试文件中的测试方法必须以Test开头，如果没有这样的方法，go test则认为测试通过。
*/

func TestHello(t *testing.T) {
    want := "你好，世界。"
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
/*
func TestHello2(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
*/
func TestProverb(t *testing.T) {
    want := "Concurrency is not parallelism."
    if got := Proverb(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}
func TestGetHello(t *testing.T)  {
  want:="hello"
  if got:=world.GetHello(); got!=want{
     t.Errorf("GetHello()=%q,watn %q",got,want)
  }

}


/**
func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got == want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
*/
