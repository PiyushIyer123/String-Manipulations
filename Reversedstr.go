package main

func Reversedstr(r string) string {
	reversed := ""
	for _, ch := range r {
		reversed = string(ch) + reversed
	}
	return reversed
}

// func Reversedstring(r string) string {
// 	reversed := ""
// 		for i := len(r) - 1; i >= 0; i-- {
// 			reversed += string(str[i])
// 		}
// 	return reversed
// }
