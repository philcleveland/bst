package main

import "fmt"

//PostType is the type of posting
type PostType int

const (
	//ForSale is a PostType
	ForSale PostType = iota + 1
)

//PostCategory is the category of the posting
type PostCategory int

//Post is the data structure with the core post data for all posts
type Post struct {
	Type        PostType     `json:"type"`
	Category    PostCategory `json:"category"`
	Title       string       `json:"title"`
	Price       float32      `json:"price"`
	City        string       `json:"city"`
	PostalCode  string       `json:"postalCode"`
	Description string       `json:"description"`
}

func (p *Post) details() string {
	return fmt.Sprintf("Type: %d\n", p.Type) +
		fmt.Sprintf("Category: %d\n", p.Category) +
		fmt.Sprintf("Title: %s\n", p.Title) +
		fmt.Sprintf("Price: %f\n", p.Price) +
		fmt.Sprintf("City: %s\n", p.City) +
		fmt.Sprintf("PostalCode: %s\n", p.PostalCode) +
		fmt.Sprintf("Description: %s\n", p.Description)
}
