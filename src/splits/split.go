package splits

type Split struct { data uint32; size uint32 }
type BiSplit struct { a Split; b Split }

// n must be positive
func SplitCount(n uint32) uint32 { return 1 << (n-1) }

func Trivial(s Split) bool {
  return (s.data == 0) || (s.data == SplitCount(s.size) - 1)
}

func Bicolored(s BiSplit) bool {
  // TODO
}
