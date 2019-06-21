package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

/*
	Vertice ... accepts generic type [T]
	Piece intended for GENERIC {T} type coercion (As we do in C++)
*/
type Vertice struct {
	Piece interface{}
}

//ProtoGraph rootie
type ProtoGraph struct {
	Vertices []*Vertice             //Pointer to a vector of Vertice
	Edges    map[Vertice][]*Vertice //Bunch of mapped edges
	lock     sync.RWMutex           //Locking mechanism
}

//PushVertice ...
func (ptr *ProtoGraph) PushVertice(vertice *Vertice) {
	//Lock it up
	ptr.lock.Lock()
	//Add vertice
	ptr.Vertices = append(ptr.Vertices, vertice)
	//Done
	ptr.lock.Unlock()
}

//PushEdge ...
func (ptr *ProtoGraph) PushEdge(e1 *Vertice, e2 *Vertice) {
	//Lock it up
	ptr.lock.Lock()

	//Check if first edge
	if ptr.Edges == nil {
		//Create mapper
		ptr.Edges = make(map[Vertice][]*Vertice)
	}

	//Connect edges
	ptr.Edges[*e1] = append(ptr.Edges[*e1], e2)
	ptr.Edges[*e2] = append(ptr.Edges[*e2], e1)

	//Done
	ptr.lock.Unlock()
}

//ToString ...
func (vertice *Vertice) ToString() string {
	return fmt.Sprintf("%v", vertice.Piece)
}

//ToString ...
func (ptr *ProtoGraph) ToString() {
	ptr.lock.RLock()

	var format string
	index := 0

	verticeSize := len(ptr.Vertices)

	for index < verticeSize {
		format += ptr.Vertices[index].ToString() + "->"
		adjacency := ptr.Edges[*ptr.Vertices[index]]

		//Add adjacent pieces
		for innerIndex := 0; innerIndex < len(adjacency); innerIndex++ {
			format += adjacency[innerIndex].ToString() + " "
		}
		format += "\n"
		index++
	}
	fmt.Println(format)
	ptr.lock.RUnlock()
}

//JSONData ...
func JSONData() string {
	return `{
			"graph":[
			{
				"node":"A",
				"edges":[
					{
						"node":"B"
					},
					{
						"node":"F"
					}
				]
			},
			{
				"node":"L",
				"edges":[
					{
						"node":"Z"
					},
					{
						"node":"F"
					}
				]
			},
			{
				"node":"F",
				"edges":[
					{
						"node":"A"
					},
					{
						"node":"L"
					},
					{
						"node":"V"
					}
				]
			},
			{
				"node":"B",
				"edges":[
					{
						"node":"A"
					}
				]
			},
			{
				"node":"Z",
				"edges":[
					{
						"node":"L"
					}
				]
			},
			{
				"node":"V",
				"edges":[
					{
						"node":"F"
					},
					{
						"node":"G"
					}
				]
			},
			{
				"node":"G",
				"edges":[
					{
						"node":"V"
					}
				]
			}
		]
	}`
}

func main() {
	myjsondata := JSONData()

	tobyte := []byte(myjsondata)

	grp := ProtoGraph{}
	err := json.Unmarshal(tobyte, &grp)

	fmt.Println(myjsondata)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(grp.Edges)
}
