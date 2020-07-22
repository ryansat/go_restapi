package main

func main() {
	a := App{}
	a.Initialize("root", "root", "db_lokasi")

	a.Run(":8080")
}
