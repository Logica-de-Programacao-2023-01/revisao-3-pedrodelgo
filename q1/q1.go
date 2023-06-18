package q1

//Dado um array de números inteiros "nums" e um número inteiro "target", retorne os índices dos dois números cuja soma
//seja igual a "target".
//
//Você pode assumir que cada entrada terá exatamente uma solução e não poderá usar o mesmo elemento duas vezes.
//
//Você pode retornar a resposta em qualquer ordem.

func TwoSum(nums []int, target int) []int {
	result := []int{}
	for n := 0; n < len(nums); n++ {
		for x := n + 1; x < len(nums); x++ {
			if nums[n]+nums[x] == target {
				result = append(result, n, x)

			}
		}
	}
	return result
}
