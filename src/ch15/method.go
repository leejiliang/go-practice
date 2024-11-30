package main

type CompanyString string

func (c *CompanyString) StringSuffix() string {

	//str := (*string)(c)
	return *(*string)(c) + "::Lotus"
}

func main() {

}
