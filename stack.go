package blueprints

type stack[T any] []T

func (st *stack[T]) empty() bool { return len(*st) == 0 }
func (st *stack[T]) len() int    { return len(*st) }
func (st *stack[T]) peek() T     { return (*st)[len(*st)-1] }
func (st *stack[T]) push(e T)    { *st = append(*st, e) }
func (st *stack[T]) pop() T      { x := (*st)[len(*st)-1]; *st = (*st)[:len(*st)-1]; return x }

// MonotonousStack shows an example how to use monotonous increase stack.
func MonotonousStack(nums []int) []int {
	st := stack[int]{}

	for i := 0; i < len(nums); i++ {
		for !st.empty() && st.peek() > nums[i] {
			st.pop()
		}
		st.push(nums[i])
	}
	return st
}
