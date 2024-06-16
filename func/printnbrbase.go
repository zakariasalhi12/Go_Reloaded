package reroaded

func PrintNbrBase(n int, base string) string {
	res := ""
	negative := false

	if IsValidBase(base) {
		return "NV"
	}

	if n == 0 {
		return "0"
	}

	if n < 0 {
		negative = true
		n = n * -1
	}

	if n == -9223372036854775808 {
		return "-" + PrintNbrBase(922337203685477580, base) + string(base[8])
	}

	for n != 0 {
		digits := n % len(base)

		res = string(base[digits]) + res

		n /= len(base)
	}

	if negative {
		res = "-" + res
		return res
	}

	return res
}
