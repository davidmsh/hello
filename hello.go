package hello

import "rsc.io/quote/v3"

// Export function Hello()
func Hello() string {
	return quote.HelloV3()
}

// Exprot function Proverb()
func Proverb() string {
	return quote.Concurrency()
}
