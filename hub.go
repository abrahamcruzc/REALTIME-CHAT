package main

type Message struct {
	ClientID string
	Text     string
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}
