package shoppinglist

import (
	"log"
	"testing"
	"time"
)

func Test_initSession(t *testing.T) {

	initSession()

	item := CartItem{"testItem", "testCategory", time.Now(), "testUser"}
	addNewItem(item)

	res := getItemsForUser("testUser")

	log.Println(res)

}
