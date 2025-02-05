package collection

import (
	"errors"

	"github.com/uor-framework/client/model"
)

// AddNode adds a new node to the graph.
func (c *Collection) AddNode(node model.Node) error {
	if _, exists := c.nodes[node.ID()]; exists {
		return errors.New("node ID collision")
	}
	c.nodes[node.ID()] = node
	return nil
}

// AddEdge adds an edge between two nodes in the graph
func (c *Collection) AddEdge(edge model.Edge) error {
	from := edge.From().ID()
	to := edge.To().ID()

	if from == to {
		return errors.New("adding self edge")
	}

	n1 := c.nodes[from]
	n2 := c.nodes[to]

	// return an error if one of the nodes doesn't exist
	if n1 == nil || n2 == nil {
		return errors.New("not all nodes exist")
	}

	if c.HasEdgeFromTo(from, to) {
		return nil
	}

	c.setEdgeFrom(edge)
	c.setEdgeTo(edge)

	return nil
}

func (c *Collection) setEdgeFrom(edge model.Edge) {
	from, ok := c.from[edge.From().ID()]
	if !ok {
		c.from[edge.From().ID()] = map[string]model.Edge{edge.To().ID(): edge}
		return
	}
	from[edge.To().ID()] = edge
	c.from[edge.From().ID()] = from
}

func (c *Collection) setEdgeTo(edge model.Edge) {
	to, ok := c.to[edge.To().ID()]
	if !ok {
		c.to[edge.To().ID()] = map[string]model.Edge{edge.From().ID(): edge}
		return
	}
	to[edge.From().ID()] = edge
	c.to[edge.To().ID()] = to
}
