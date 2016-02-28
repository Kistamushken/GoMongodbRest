package main

import ("users")
/*
Create a new MongoDB session, using a database
named "signatures". Create a new server using
that session, then begin listening for HTTP requests.
*/
func main() {
	session := users.NewSession("users")
	server := users.NewServer(session)
	server.Run()
}