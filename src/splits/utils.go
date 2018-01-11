package splits

// Considers the bits of `in` flagged by `filter
// If they're not all the same, returns 0
// If they're all 1, returns 1
// If they're all 0, returns -1
func all_on (inp uint32, filter uint32) int8 {
  inp &= filter;

  if (inp == filter) {
    return 1;
  }

  if (inp == 0) {
    return -1;
  }

  return 0;
}

func invert (s *BiSplit) {
  s.data = inverted(*s);
}

func inverted (s BiSplit) uint32 {
  return s.data ^ ((1 << BiSplitSize(s)) - 1);
}

// This will only work for standardized splits
func compatible(a uint32, b uint32) int8 {
  // first grab only the bits that differ
  var diff uint32 = a ^ b;
  // then check that they all differ in the same way
  return all_on(a, diff);
}
