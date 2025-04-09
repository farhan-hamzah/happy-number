package main
import "fmt"

func isHappy(n int) bool {
    slow := n
    fast := getNext(n)

    for fast != 1 && slow != fast {
        slow = getNext(slow)
        fast = getNext(getNext(fast))
    }

    return fast == 1
}

func getNext(num int) int {
    sum := 0
    for num > 0 {
        digit := num % 10
        sum += digit * digit
        num = num / 10
    }
    return sum
}

func main() {
    var num int
	fmt.Scan(&num)
	fmt.Print(isHappy(num))
}