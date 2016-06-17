package godnsbl

import (
	"fmt"
	"net"
	"testing"
)

// Parallel for loop: https://play.golang.org/p/MVFwbmxbou
// http://www.golangpatterns.info/concurrency/parallel-for-loop

func TestReverseIP(t *testing.T) {
	fmt.Println("TestReverseIP")
	t.Parallel()
	ip := net.IP{192, 168, 1, 1}

	r := Reverse(ip)

	if r != "1.1.168.192" {
		t.Errorf("Expected ip to equal 1.1.168.192, actual %v", r)
	}
}

func TestKnownIP(t *testing.T) {
	fmt.Println("TestKnownIP")
	t.Parallel()

	for i := range Blacklists {
		res := Lookup(Blacklists[i], "127.0.0.2")
		fmt.Println(res.Results)
	}
}

func TestLookupParams(t *testing.T) {
	fmt.Println("TestLookupParams")
	t.Parallel()

	for i := range Blacklists {
		res := Lookup(Blacklists[i], "127.0.0.2")
		//fmt.Println(res.Results)

		if res.List != Blacklists[i] {
			t.Errorf("Expected %v, actual %v", Blacklists[i], res.List)
		}

		if res.Host != "127.0.0.2" {
			t.Errorf("Expected 127.0.0.2, actual %v", res.Host)
		}
	}
}
