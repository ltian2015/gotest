package test
import(
//  "rsc.io/quote"
//  quoteV3 "rsc.io/quote/v3"
"rsc.io/quote/v3"
)
//return "Hello, world."
func Hello() string {
  //  return "Hello, world."
//  return quote.Hello()
   return quote.HelloV3()
}
//return go Concurrency  Proverb
func Proverb() string {
    return quote.Concurrency()
}
