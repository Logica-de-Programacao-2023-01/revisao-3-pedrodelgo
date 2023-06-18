package q2

//Escreva uma função para encontrar o prefixo comum mais longo entre um array de strings.
//
//Se não houver um prefixo comum, retorne uma string vazia "".

func LongestCommonPrefix(strs []string) string {
	var longestPrefix string = ""
	var endPrefix = false

	if len(strs) > 0 {
		sort.Strings(strs)
		first := string(strs[0])
		last := string(strs[len(strs)-1])

		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}
