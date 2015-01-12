package markov
import "testing"

func TestChain(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"To generate text using this table we select an initial prefix", ""},
        {"I am not a number! I am a free man!", ""},
    }
    for _, c := range cases {
        got := EasyGenerate(c.in)
        if got != c.want {
            t.Errorf("EasyGenerate(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}

