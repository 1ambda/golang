package main

func main() {

	// cases break automatically, no fallthrough by default
	switch os := getRuntimeOS(); os {
	case "darwin":
		pass(true)
	case "window":
		pass(false)
	case "linux":
		pass(false)
	default:
		pass(false)
	}
}

func pass(expr bool) {
	if !expr {
		panic(0)
	}
}

func getRuntimeOS() string {
	return "darwin"
}
