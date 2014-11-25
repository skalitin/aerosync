package main

import "fmt"

func main() {
	//resp, err := http.Get("http://google.com")
	//if err != nil {
	//	panic(err)
	//}

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Printf(string(body))

	for x := 0; x < 10; x++ {
		fmt.Printf("X=%d\n", x)
	}
}
