package singleton

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleton(t *testing.T) {

	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("test failed, instance is not equal")
	}

}

func BenchmarkSingleton(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ins := GetInstance()
		fmt.Println("ins type: ", reflect.TypeOf(ins))
	}
}
