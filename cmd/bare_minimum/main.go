package main

func main() {
	rt := servicemesh.New("my runtime")

	rt.Add(&example{name: "foo"})

	rt.Run()
}
