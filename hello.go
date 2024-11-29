package main

const spanish = "Spanish"
const french = "French"
const engishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjor, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := engishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}
