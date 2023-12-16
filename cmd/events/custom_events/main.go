package main

func main() {
	rt := servicemesh.New()

	rt.Add(&sender{})
	rt.Add(&receiver{})

	rt.Run()
}
