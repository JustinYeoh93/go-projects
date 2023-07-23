package main

import (
	"sync"

	"github.com/olahol/melody"
)

type Manager struct {
	clients ClientList

	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

func (m *Manager) addClient(c *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[c] = true
}

func (m *Manager) removeClient(c *Client) {
	m.Lock()
	defer m.Unlock()

	delete(m.clients, c)
}

type ClientList map[*Client]bool

type Client struct {
	connection *melody.Session

	manager *Manager
}

func NewClient(s *melody.Session, m *Manager) *Client {
	return &Client{
		connection: s,
		manager:    m,
	}
}
