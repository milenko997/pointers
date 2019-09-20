package main

func main() {
	count := 10
	println("Before: ", count, " -------- ", &count)

	increment(&count)

	println("After: ", count, " -------- ", &count)
}
func increment(inc *int) {
	*inc++
	println("Inc: ", *inc, " -------- ", &inc, " -------- ", inc)

}
