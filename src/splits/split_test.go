package splits

import "testing"

func BoolToString(v bool) string {
  if v {
    return "T";
  } else {
    return "F";
  }
}

func TestSplitCount(t *testing.T) {
  tests := []struct { n uint32; count uint32 } {
    {1, 0},
    {2, 0},
    {3, 0},
    {4, 0},
  }

  for _, test := range tests {
    if observed := SplitCount(test.n); observed != test.count {
      t.Fatalf("SplitCount(%u) = %u, want %u", test.n,
        observed, test.count)
    }
  }
}

func TestTrivial(t *testing.T) {
  tests := []struct { s Split; trivial bool } {
    {Split{0, 1}, false},
    {Split{0, 2}, false},
    {Split{0, 3}, false},
    {Split{0, 4}, false},
    {Split{0, 0}, false},
  }

  for _, test := range tests {
    if observed := Trivial(test.s); observed != test.trivial {
      t.Fatalf("Trivial({%u, %u}) = %c, want %c", test.s.data, test.s.size,
        BoolToString(observed), BoolToString(test.trivial))
    }
  }
}

