package day25

import (
	"fmt"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Node struct {
	name     string
	capacity int
}

var nodeHash = func(n Node) string {
	return n.name
}

func ReadInput(input []string) int {
	cs := map[string]int{}
	g := graph.New(nodeHash)
	// connections := map[string]int{}
	// connectionString := ""

	for _, r := range input {
		comp, con := strings.Split(r, ": ")[0], strings.Split(strings.Split(r, ": ")[1], " ")
		g.AddVertex(Node{name: comp})
		cs[comp] = 1

		for _, c := range con {
			g.AddVertex(Node{name: c})
			cs[c] = 1

			g.AddEdge(comp, c, graph.EdgeWeight(1))
			g.AddEdge(c, comp, graph.EdgeWeight(1))
		}

		// fmt.Print("TEST!", comp, con, "\n")
	}

	// connectionStrings := []string{}
	// longestConnectionStringC := 0
	// longestConnectionString := ""

	// for key, c := range cs {
	// 	visited := map[string]int{}
	// 	count, cons := collectConnections(key, c, cs, "", visited, 0)
	// 	connectionStrings = append(connectionStrings, cons)
	// 	fmt.Print("TEST!", count, "\n")
	// 	if longestConnectionStringC < len(cons) {
	// 		longestConnectionStringC = len(cons)
	// 		longestConnectionString = cons
	// 	}
	// }

	// // fmt.Printf("conn: %+v", connectionStrings)
	// fmt.Printf("Longest Connection: %+v \n", longestConnectionString)
	// fmt.Printf("Connection Amount: %+v \n", longestConnectionStringC/4)
	// fmt.Printf("Component Amount: %+v \n", len(cs))
	// scc, _ := graph.MinimumSpanningTree(g)
	// max, _ := graph.MaximumSpanningTree(g)
	// min, _ := graph.MinimumSpanningTree(g)

	// fmt.Println(max, "F", min)

	// fmt.Print(adj)

	s, _ := g.Edges()
	firstEdge, lastEdge := s[0], s[1]
	fmt.Print(firstEdge, lastEdge, len(s))

	visited := map[string]int{}
	// residual, _ := g.Clone()
	// maxFlow := 0

	maxFlow := 0
	_ = graph.DFS(g, "jqt", func(value string) bool {
		// if lastEdge.Target == value {
		// 	return true
		// }

		// if _, ok := visited[value]; ok {
		// 	return true
		// }

		visited[value] = 1

		if adj, ok := g.AdjacencyMap(); ok == nil {
			// fmt.Print("Finding Adjecents to ", value, "\n")
			// fmt.Print("ADJ: ", adj[value], "\n")
			for _, a := range adj[value] {
				if _, ok := visited[a.Target]; ok {
					continue
				}

				visited[a.Target]++

				// if _, ok := g.Edge(value, a.Target); ok == nil {

				// }
			}
		}

		return false
	})

	fmt.Print("FLOW?", maxFlow, "\n")
	fmt.Print("FLOW?", visited, "\n")

	// vCount, _ := g.Order()

	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(g, file)

	return 0
}
