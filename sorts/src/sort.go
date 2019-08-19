// sorting algoritm
package main

import "fmt"

func main() {
	numbers := []int{5, 6, 1, 4, 9, 8, 8, 2, 7}
  fmt.Printf("strArray1: %v\n", numbers)

  sortednumbers := mergeIntArraySort(numbers)
  fmt.Printf("mergesorted: %v\n", sortednumbers)
}


func mergeIntArraySort(numbers []int) []int{
	length := len(numbers)
	// if the length is greater than 2, split and sort array 
	// then merge the 2 sorted in order
	if length > 2{
	  mid := length/2
	  numbers1 := mergeIntArraySort(numbers[:mid])
	  numbers2 := mergeIntArraySort(numbers[mid:])
	  i := 0
	  j := 0
	  maxi := len(numbers1)
	  maxj := len(numbers2)
	  result := make([]int, length)
	  k := 0
	  for k < length{
		if i < maxi && j < maxj{
		  if numbers1[i] < numbers2[j]{
			result[k] = numbers1[i]
			k++
			i++
		  } else {
			result[k] = numbers2[j]
			k++
			j++
		  }
		} else if i < maxi {
		  result[k] = numbers1[i]
		  k++
		  i++
		} else if j < maxj {
		  result[k] = numbers2[j]
		  k++
		  j++
		}
		
	  }
	  return result
	// if less than 2 items in array, return the array. it is already sorted
	} else if length < 2 {
	  return numbers
	// else the length is 2, sort from lowest to highest
	} else {
	  if numbers[0] < numbers[1]{
		return numbers
	  } else {
		result := []int{numbers[1], numbers[0]}
		return result
	  }
	}
  }