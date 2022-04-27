package datastruct

import "sync"

type NodeQueue struct {
	Items []Vertex
	lock  sync.RWMutex
}

// Enqueue adds an Node to the end of the queue
func (s *NodeQueue) Enqueue(t Vertex) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		// add vertex distance less than travers's vertex distance
		if t.Distance < v.Distance {
			s.Items = append(s.Items[:k+1], s.Items[k:]...)
			s.Items[k] = t
			insertFlag = true
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, t)
	}
}

// Dequeue removes an Node from the start of the queue
func (s *NodeQueue) Dequeue() *Vertex {
	s.lock.Lock()
	defer s.lock.Unlock()
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	return &item
}

//NewQ Creates New Queue
func (s *NodeQueue) NewQ() *NodeQueue {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Items = []Vertex{}
	return s
}

// IsEmpty returns true if the queue is empty
func (s *NodeQueue) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *NodeQueue) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.Items)
}
