package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type issueStruct struct {
	ID          int           `json:"id"`
	Iid         int           `json:"iid"`
	ProjectID   int           `json:"project_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	State       string        `json:"state"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	ClosedAt    interface{}   `json:"closed_at"`
	ClosedBy    interface{}   `json:"closed_by"`
	Labels      []interface{} `json:"labels"`
	Milestone   struct {
		ID          int         `json:"id"`
		Iid         int         `json:"iid"`
		ProjectID   int         `json:"project_id"`
		Title       string      `json:"title"`
		Description interface{} `json:"description"`
		State       string      `json:"state"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
		DueDate     interface{} `json:"due_date"`
		StartDate   interface{} `json:"start_date"`
	} `json:"milestone"`
	Assignees []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"assignees"`
	Author struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	Assignee struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"assignee"`
	UserNotesCount   int         `json:"user_notes_count"`
	Upvotes          int         `json:"upvotes"`
	Downvotes        int         `json:"downvotes"`
	DueDate          interface{} `json:"due_date"`
	Confidential     bool        `json:"confidential"`
	DiscussionLocked interface{} `json:"discussion_locked"`
	WebURL           string      `json:"web_url"`
	TimeStats        struct {
		TimeEstimate        int         `json:"time_estimate"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
	} `json:"time_stats"`
	Links struct {
		Self       string `json:"self"`
		Notes      string `json:"notes"`
		AwardEmoji string `json:"award_emoji"`
		Project    string `json:"project"`
	} `json:"_links"`
	Subscribed bool `json:"subscribed"`
}

func main() {
	var issue_struct issueStruct
	issue_json, err := ioutil.ReadFile("issue.src")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(issue_json, &issue_struct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", issue_struct)

}
