package splits

type Split struct { data uint32; size uint32 }
type BiSplit struct { data uint32; n uint32; d uint32 }

func SplitA(s BiSplit) Split { return Split{0,0} }
func SplitB(s BiSplit) Split { return Split{0,0} }
func BiSplitSize(s BiSplit) uint32 { return s.n + s.d }

// n must be positive
func SplitCount(n uint32) uint32 { return 1 << (n-1) }

// checks whether this split is trivial, i.e. all
// on one side (all 0s or all 1s)
func Trivial(s Split) bool {
  return (s.data == 0) || (s.data == SplitCount(s.size) - 1)
}

func Bicolored(s BiSplit) bool {
  return true;
}

//////////////////////////////////////////////////////////
// If we exclude trivial Splits/monocolored BiSplits,
// then there are always two representations of each
// split. To handle permutations correctly, we need
// to be able to revert to a standard representation.
//
// Our stanard will simply be that the last of the n + d
// bits of BiSplit#data must be zero
//////////////////////////////////////////////////////////
func Standard(s BiSplit) bool {
  topbit := s.data & (1 << (BiSplitSize(s) - 1))
  return 0 != topbit
}

func Standardize(s *BiSplit) {
  if !Standard(*s) { invert(s) }
}

// Assumes that a.n == b.n and a.d == b.d
func Compatible(a BiSplit, b BiSplit) int8 {
  tmp := compatible(a.data, b.data)

  if tmp != 0 {
    return tmp
  } else {
    return compatible(inverted(a), b.data);
  }
}
