package main

import (
	"github.com/andradesampaio/api-rest-go-with-gin/database"
	"github.com/andradesampaio/api-rest-go-with-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
