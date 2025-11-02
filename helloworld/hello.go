package helloworld

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}
