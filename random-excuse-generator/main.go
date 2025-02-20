package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var excuses = []string{
	"My cat needed a ride to the vet.",
	"I accidentally fell into a time portal.",
	"I got stuck in traffic... on a bicycle.",
	"My alarm clock ran out of battery.",
	"The dog ate my car keys.",
	"I was abducted by aliens... again.",
	"I lost track of time while saving the world.",
	"I had a flat tire... on my skateboard.",
	"The Wi-Fi was down... at my grandma's house.",
	"I got caught up in a game of hide and seek.",
	"My pet hamster had a medical emergency.",
	"I was practicing my ninja skills.",
	"I tripped over a cloud and fell into the future.",
	"My code works, but only if you run it on a full moon.",
	"I swear it worked on the first try... on a different machine.",
	"My code was perfect until I hit the 'Save' button.",
	"I had a bug, but then I fixed it by deleting all the code.",
	"The internet told me to use a 'for loop' but I misunderstood.",
	"Git told me there were no conflicts, but my code and my sanity disagree.",
	"My keyboard was typing in a different language than me.",
	"I tried to refactor, but I think I accidentally invented a new programming language.",
	"That bug wasn’t a bug, it was just a 'feature' I hadn’t documented yet.",
}


func getRandomExcuse(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	excuse := excuses[rand.Intn(len(excuses))]
	c.JSON(http.StatusOK, gin.H{
		"excuse": excuse,
	})
}

func main() {
	r := gin.Default()
	r.GET("/excuse", getRandomExcuse)
	r.Run(":8080")
}
