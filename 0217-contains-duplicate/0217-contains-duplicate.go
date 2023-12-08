// Time: O(n)
// Space: O(n)
func containsDuplicate(nums []int) bool {
    // Space: O(n)
    set := NewSet()
    
    // Time: O(n)
    for _, num := range nums {
        // Time: O(1)
        if set.Has(num) {
            return true
        }
        // Time: O(1)
        set.Add(num)
    }

    return false
}

type Set struct {
    items map[int]struct{}
}

func NewSet() *Set{
    return &Set{
        items: make(map[int]struct{}),
    }
}

// Time: O(1)
// Space: O(1)
func (s *Set) Add(val int) {
    // Time: O(1)
    s.items[val] = struct{}{}
}

// Time: O(1)
// Space: O(1)
func (s *Set) Has(val int) bool {
    // Time: O(1)
    _, ok := s.items[val]
    return ok
}