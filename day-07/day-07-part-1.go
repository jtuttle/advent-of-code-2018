package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Node struct {
	step string
	parents []*Node
	children []*Node
}

func createNode(step string) Node {
	parents := make([]*Node, 0)
	children := make([]*Node, 0)

	node := Node{
		step: step,
		parents: parents,
		children: children,
	}

	return node
}

func parseStep(line string, nodeMap map[string]*Node) {
	parentStepRegex := regexp.MustCompile(`(?:Step\s)([A-Z])(?:\smust)`)
	parentStep := parentStepRegex.FindStringSubmatch(line)[1]

	childStepRegex := regexp.MustCompile(`(?:step\s)([A-Z])(?:\scan)`)
	childStep := childStepRegex.FindStringSubmatch(line)[1]

	if nodeMap[parentStep] == nil {
		newNode := createNode(parentStep)
		nodeMap[parentStep] = &newNode
	}

	if nodeMap[childStep] == nil {
		newNode := createNode(childStep)
		nodeMap[childStep] = &newNode
	}

	parent := nodeMap[parentStep]
	child := nodeMap[childStep]

	parent.children = append(parent.children, child)
	child.parents = append(child.parents, parent)
}

func findRoots(nodeMap map[string]*Node) []*Node {
	var roots []*Node
	
	for _, v := range nodeMap {
		parents := v.parents
		
		if len(parents) == 0 {
			roots = append(roots, v)
		}
	}

	return roots
}

func findStepOrder(roots []*Node) string {
	ready := roots
	completed := make(map[string]bool)
	order := make([]string, 0)
	
	var nextStep *Node

	for len(ready) > 0 {
		sort.Slice(ready, func(i, j int) bool {
			return ready[i].step < ready[j].step
		})
		
		nextStep, ready = ready[0], ready[1:]

		completed[(*nextStep).step] = true
		order = append(order, nextStep.step)

		for _, child := range (*nextStep).children {
			nodeReady := true

			for _, parent := range (*child).parents {
				if completed[(*parent).step] != true {
					nodeReady = false
					break
				}
			}

			if nodeReady {
				ready = append(ready, child)
			}
		}
	}

	return strings.Join(order, "")
}

func main() {
	file, _ := os.Open("./day-07-input.txt")
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nodeMap := make(map[string]*Node, 0)
	
	for scanner.Scan() {
		parseStep(scanner.Text(), nodeMap)
	}

	roots := findRoots(nodeMap)
	order := findStepOrder(roots)

	fmt.Println("Step order:", order)
}
