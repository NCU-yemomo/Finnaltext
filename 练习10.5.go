package main

type nums1 struct {
	f1 float32
	int
	string
}
type nums2 struct {
	nums1
	d int
}

func main() {
	var b = new(nums2)
	b.nums1 = nums1{3.14, 5, "hello"}
	b.d = 10

}
