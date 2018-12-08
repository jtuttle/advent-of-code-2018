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

type Job struct {
	node *Node
	worker int
	seconds int
}

func getSeconds(step string, base int) (int) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return base + strings.Index(alphabet, step) + 1
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

func findSecondsToComplete(roots []*Node, workerCount int, baseSeconds int) int {
	ready := roots
	completed := make(map[string]bool)
	order := make([]string, 0)

	workers := make([]int, 0)
	for i := 0; i < workerCount; i++ {
		workers = append(workers, i)
	}
	
	jobs := make([]*Job, 0)

	seconds := 0
	
	var nextStep *Node

	for len(ready) > 0 || len(jobs) > 0 {
		sort.Slice(ready, func(i, j int) bool {
			return ready[i].step < ready[j].step
		})

		// assign jobs to workers
		for i := len(ready) - 1; i >= 0; i-- {
			if len(workers) > 0 {
				var nextWorker int
				
				nextStep, ready = ready[0], ready[1:]
				nextWorker, workers = workers[0], workers[1:]

				nextJob := &Job{
					node: nextStep,
					worker: nextWorker,
					seconds: getSeconds(nextStep.step, baseSeconds),
				}

				jobs = append(jobs, nextJob)
			}
		}

		// decrement job seconds and complete jobs
		for i := len(jobs) - 1; i >= 0; i-- {
			job := jobs[i]
			
			job.seconds = job.seconds - 1

			if job.seconds == 0 {
				jobs = append(jobs[:i], jobs[i+1:]...)
				
				jobStep := job.node
				completed[jobStep.step] = true
				order = append(order, jobStep.step)

				workers = append(workers, job.worker)

				// check job children for ready jobs
				for _, child := range jobStep.children {
					nodeReady := true

					for _, parent := range child.parents {
						if completed[parent.step] != true {
							nodeReady = false
							break
						}
					}

					if nodeReady {
						ready = append(ready, child)
					}
				}
			}
		}

		seconds++
	}
	
	return seconds
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
	seconds := findSecondsToComplete(roots, 5, 60)

	fmt.Println("Seconds to complete:", seconds)
}
