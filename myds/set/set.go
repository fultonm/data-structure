package set

type Set[T comparable] struct {
	col []T
}

func (s *Set[T]) Add(ele T) bool {
	for _, e := range s.col {
		if e == ele {
			return false
		}
	}
	s.col = append(s.col, ele)
	return true
}

func (s *Set[T]) Remove(ele T) bool {
	for i, e := range s.col {
		if e == ele {
			s.col[i] = s.col[len(s.col)-1]
			s.col = s.col[:len(s.col)-1]
			return true
		}
	}
	return false
}

// We do not get elements from a Set
// func (s *Set[T]) Get(index int) T {
// 	if index < s.Len() {
// 		return s.col[index]
// 	} else {
// 		panic(fmt.Sprintf("Index %v out of range. Set length is: %v", index, s.Len()))
// 	}
// }

func (s *Set[T]) Contains(ele T) bool {
	for _, e := range s.col {
		if e == ele {
			return true
		}
	}
	return false
}

func (s *Set[T]) Len() int {
	return len(s.col)
}
