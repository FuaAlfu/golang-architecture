package hello

import "rsc.io/quote"

func Greeting() string{
	//return "hello world"
	return quote.HelloV3()
}

func Proverb()string{
	return quote.Concurrency()
}