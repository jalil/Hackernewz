package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client struct {
	BaseUrl string
	Version string
	Suffix  string
}

type Item struct {
	Id          int
	Type        string
	By          string
	Time        int
	Text        string
	Parent      int
	Kids        []int
	Score       int
	Url         string
	Title       string
	Descendants int
}

type Story struct {
	By    string
	Id    int
	Kids  []int
	Score int
	Time  int
	Title string
	Type  string
	Url   string
}

type Part struct {
	By     string
	Id     int
	Parent int
	Score  int
	Text   string
	Time   int
	Type   string
}
type Comment struct {
	By     string
	Id     int
	Kids   []int
	Parent int
	Time   int
	Text   string
	Title  string
	Type   string
}

type Ask struct {
	By          string
	Descendants int
	Id          int
	Kids        []int
	Score       int
	Text        string
	Time        int
	Title       string
	Type        string
	Url         string
}

type Job struct {
	By    string
	Id    int
	Score int
	Text  string
	Time  int
	Title string
	Type  string
	Url   string
}

type Poll struct {
	By          string
	Descendants int
	Id          int
	Kids        []int
	Score       int
	Text        string
	Time        int
	Title       string
	Type        string
}

type User struct {
	About     string
	Created   int
	Delat     int
	Id        string
	Karma     int
	Submitted []int
}

func NewClient() *Client {
	var c Client
	c.BaseUrl = "https://hacker-news.firebaseio.com/"
	c.Version = "v0"
	c.Suffix = ".json"
	return &c
}

func (c *Client) GetItem(id int) Item {
	url := c.BaseUrl + c.Version + "/item/" + strconv.Itoa(id) + c.Suffix
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	var i Item
	err = json.Unmarshal(body, &i)
	if err != nil {
		panic(err)
	}
	return i

}

func (i Item) ConvertToStory() Story {
	var story Story
	story.By = i.By
	story.Id = i.Id
	story.Kids = i.Kids
	story.Score = i.Score
	story.Time = i.Time
	story.Title = i.Title
	story.Type = i.Type
	story.Url = i.Url
	return story

}

func (c Client) GetStory(id int) (Story, error) {
	item := c.GetItem(id)

	if item.Type != "story" {
		return Story{}, fmt.Errorf("Called GetStory on ID #%v which is not a _story_. "+
			"Item is of type %v.", id, item.Type)
	} else {
		return item.ConvertToStory(), nil
	}
}

func (i Item) ConvertToComment() Comment {
	var comment Comment
	comment.By = i.By
	comment.Id = i.Id
	comment.Kids = i.Kids
	comment.Parent = i.Parent
	comment.Text = i.Text
	comment.Time = i.Time
	comment.Type = i.Type
	return comment

}

func (c Client) GetComment(id int) (Comment, error) {
	item := c.GetItem(id)

	if item.Type != "comment" {
		return Comment{}, fmt.Errorf("Called GetStory on ID #%v which is not a _story_. "+
			"Item is of type %v.", id, item.Type)
	} else {
		return item.ConvertToComment(), nil
	}
}

func (i Item) ConvertToJob() Job {
	var job Job
	job.By = i.By
	job.Id = i.Id
	job.Score = i.Score
	job.Text = i.Text
	job.Time = i.Time
	job.Title = i.Title
	job.Type = i.Url
	return job

}
func (c Client) GetJob(id int) (Job, error) {
	item := c.GetItem(id)

	if item.Type != "job" {
		return Job{}, fmt.Errorf("Called GetStory on ID #%v which is not a _story_. "+
			"Item is of type %v.", id, item.Type)
	} else {
		return item.ConvertToJob(), nil
	}
}

func (i Item) ConvertToPoll() Poll {
	var poll Poll
	poll.By = i.By
	poll.Descendants = i.Descendants
	poll.Id = i.Id
	poll.Kids = i.Kids
	poll.Score = i.Score
	poll.Text = i.Text
	poll.Time = i.Time
	poll.Title = i.Title
	poll.Type = i.Url
	return poll

}
func (c Client) GetPoll(id int) (Poll, error) {
	item := c.GetItem(id)

	if item.Type != "poll" {
		return Poll{}, fmt.Errorf("Called GetStory on ID #%v which is not a _story_. "+
			"Item is of type %v.", id, item.Type)
	} else {
		return item.ConvertToPoll(), nil
	}
}

func (i Item) ConvertToPart() Part {
	var part Part
	part.By = i.By
	part.Id = i.Id
	part.Parent = i.Parent
	part.Score = i.Score
	part.Text = i.Text
	part.Time = i.Time
	part.Type = i.Type
	return part

}
func (c Client) GetPart(id int) (Part, error) {
	item := c.GetItem(id)

	if item.Type != "pollopt" {
		return Part{}, fmt.Errorf("Called GetStory on ID #%v which is not a _story_. "+
			"Item is of type %v.", id, item.Type)
	} else {
		return item.ConvertToPart(), nil
	}
}
func main() {
	cl := NewClient()
	fmt.Println(cl.GetPoll(2921983))
}
