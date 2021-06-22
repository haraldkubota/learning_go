package rotLeft

import ()

// func rotLeft(a []int, d int) []int {
// 	for i := 0; i < d; i += 1 {
// 		t := a[0]
// 		for j:=0; j < len(a)-1; j+=1 {
// 			a[j]=a[j+1]
// 		}
// 		a[len(a)-1]=t
// 	}
// 	return a
// }

func rotLeft(a []int, d int) []int {
	copyFrom := d % len(a)
	res := make([]int, len(a))
	copy(res, a)

	for i := 0; i < len(a); i += 1 {
		res[i] = a[copyFrom]
		copyFrom += 1
		if copyFrom == len(a) {
			copyFrom = 0
		}
	}
	return res
}
