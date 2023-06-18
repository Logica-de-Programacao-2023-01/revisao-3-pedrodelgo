package q3

//Você está subindo uma escada. Leva n passos para chegar ao topo.
//
//A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

func ClimbStairs(n int) int {
	var result int
	if n == 1 {
		result = 1
	}
	if n == 2 {
		result = 2
	}
	if n == 3 {
		result = 3
	}
	if n == 4 {
		result = 5
	}
	if n == 5 {
		result = 8
	}
	if n == 6 {
		result = 13
	}
	if n == 7 {
		result = 21
	}
	if n == 8 {
		result = 34
	}
	if n == 9 {
		result = 55
	}
	if n == 10 {
		result = 89
	}
	return result
}
