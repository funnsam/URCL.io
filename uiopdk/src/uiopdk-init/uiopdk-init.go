package main

func main() {
	err := UnzipHTTP("https://funnsam.github.io/URCL.io/uiopdk-res/latest.zip", "")
	if err != nil {
		panic(err)
	}
}
