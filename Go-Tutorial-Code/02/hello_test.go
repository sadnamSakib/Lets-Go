package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Apurbo", "Arpa", "Maruf"},
			result: "Hello, Apurbo, Arpa, Maruf!",
		},
	}
	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("Wanted %s [%v] ,  Got %s", st.result, st.items, s)
		}
	}
}
