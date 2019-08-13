package hello

import "rsc.io/quote"

// Export function Hello()
func Hello() string {
	return quote.Hello()
}
