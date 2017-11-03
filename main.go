package main

import "gitlab/nefco/mail-service/src/api"

func main() {
	api.NewApi(
		":1536",
	)
}