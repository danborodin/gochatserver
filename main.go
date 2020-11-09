package main

import (
	"log"
	"net/http"

	"github.com/danborodin/go-chat-server/router"
)

func main() {
	// setupRoutes()
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// msg1 := models.Message{
	// 	Sender: "nil",
	// 	Text:   "Sergan pidor",
	// 	Date:   "nov 9 20:10",
	// }
	// msg2 := models.Message{
	// 	Sender: "nil",
	// 	Text:   "Sergan pidor2",
	// 	Date:   "nov 9 20:11",
	// }
	// msg3 := models.Message{
	// 	Sender: "nil",
	// 	Text:   "Sergan pidor3",
	// 	Date:   "nov 9 20:12",
	// }

	// database.AddMessage("5fa7fee7bff4b94efc5b0288", msg1)
	// database.AddMessage("5fa7fee7bff4b94efc5b0288", msg2)
	// database.AddMessage("5fa7fee7bff4b94efc5b0288", msg3)

	router.SetupRoutes()
	log.Printf("Server running at address %s", router.Host)
	log.Fatal(http.ListenAndServe(router.Host, router.Router))
}
