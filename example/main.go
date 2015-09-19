package main

import (
	"fmt"
	"github.com/pims/assist"
	"log"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	client := assist.NewDefaultClient()
	user, err := client.Users.Get("simplebits")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Name)

	members, err := client.Teams.Members("dribbble")
	checkErr(err)

	for _, member := range members {
		log.Printf("[M] %s\n", member.Username)
	}

	shots, err := client.Teams.Shots("dribbble")
	checkErr(err)

	for _, shot := range shots {
		log.Printf("[S] %s\n", shot.Title)
	}

	config := assist.NewConfig(os.Getenv("DRIBBBLE_TOKEN"), assist.DefaultApiEndpoint)
	client = assist.NewClient(config)

	user, err = client.Users.Get("simplebits")
	checkErr(err)

	fmt.Println(user.Name)

	shots, err = client.Shots.List(nil)
	checkErr(err)

	for _, shot := range shots {
		log.Printf("%s : %s\n", shot.Title, strings.Join(shot.Tags, ", #"))
	}

	likes, err := client.Users.Likes("simplebits")
	checkErr(err)

	for _, like := range likes {
		log.Printf("[L] %s : %s\n", like.Title, strings.Join(like.Tags, ", #"))
	}

	buckets, err := client.Users.Buckets("simplebits")
	checkErr(err)

	for _, bucket := range buckets {

		log.Printf("[B] %s : %s (%d)\n", bucket.Name, bucket.Description, bucket.ShotsCount)

	}

	followers, err := client.Users.Followers("simplebits")
	checkErr(err)

	for _, follower := range followers {
		log.Printf("[F <-] %s\n", follower.Name)
	}

	followings, err := client.Users.Following("simplebits")
	checkErr(err)

	for _, following := range followings {
		log.Printf("[F ->] %s\n", following.Name)
	}

	teams, err := client.Users.Teams("simplebits")
	checkErr(err)

	for _, team := range teams {
		log.Printf("[T] %s\n", team.Name)
	}

	fmt.Println(client.Status())
}
