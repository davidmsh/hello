package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

// Export function Hello()
func Hello() string {
	return quote.Hello()
}

// Exprot function Proverb()
func Proverb() string {
	return quoteV3.Concurrency()
}
