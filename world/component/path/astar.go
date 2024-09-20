package path

import (
	"errors"
	"reflect"
	"tanks/data"
	"tanks/world/board"
	components "tanks/world/component"
)

// Node represents a given point on a map
// g is the total distance of the Node from the start
// h is the estimated distance of the Node from the ending
// f is the total value of the Node (g + h)
type Node struct {
	Parent   *Node
	Position *components.Position
	g        int
	h        int
	f        int
}

func newNode(parent *Node, position *components.Position) *Node {
	n := Node{}
	n.Parent = parent
	n.Position = position
	n.g = 0
	n.h = 0
	n.f = 0

	return &n
}

func (n *Node) IsEqual(other *Node) bool {
	return n.Position.IsEqual(other.Position)
}

type AStar struct{}

// GetPath takes a level, the starting position and an ending position (the goal) and returns
// a list of Positions which is the path between the points.
func (as AStar) GetPath(level board.Level, start *components.Position, end *components.Position) []components.Position {
	openList := make([]*Node, 0)
	closedList := make([]*Node, 0)

	//Create our starting point
	startNode := newNode(nil, start)
	startNode.g = 0
	startNode.h = 0
	startNode.f = 0

	//Create this Node just for ease of dropping into our isEqual function to see if we are at the end
	//May be worth a refactor of changing the isEqual to test on Position.
	endNodePlaceholder := newNode(nil, end)

	openList = append(openList, startNode)
	for {
		if len(openList) == 0 {
			break
		}

		//Get the current Node
		currentNode := openList[0]
		currentIndex := 0

		//Get the Node with the smallest f value
		for index, item := range openList {
			if item.f < currentNode.f {
				currentNode = item
				currentIndex = index
			}
		}

		//Move from open to closed list
		openList = append(openList[:currentIndex], openList[currentIndex+1:]...)
		closedList = append(closedList, currentNode)

		//Check to see if we reached our end
		//If so, we are done here
		if currentNode.IsEqual(endNodePlaceholder) {
			path := make([]components.Position, 0)
			current := currentNode
			for {
				if current == nil {
					break
				}
				path = append(path, *current.Position)
				current = current.Parent
			}

			//Reverse the Path and Return it
			ReverseSlice(path)
			return path
		}

		edges := make([]*Node, 0)

		if currentNode.Position.Y > 0 {
			tile := level.Tiles[level.GetIndexFromXY(currentNode.Position.X, currentNode.Position.Y-1)]
			if tile.TileType != data.WALL {
				//The location is in the map bounds and is walkable
				upNodePosition := components.Position{
					X: currentNode.Position.X,
					Y: currentNode.Position.Y - 1,
				}
				newNode := newNode(currentNode, &upNodePosition)
				edges = append(edges, newNode)

			}
		}

		if currentNode.Position.Y < data.ScreenHeight {
			tile := level.Tiles[level.GetIndexFromXY(currentNode.Position.X, currentNode.Position.Y+1)]
			if tile.TileType != data.WALL {
				//The location is in the map bounds and is walkable
				downNodePosition := components.Position{
					X: currentNode.Position.X,
					Y: currentNode.Position.Y + 1,
				}
				newNode := newNode(currentNode, &downNodePosition)
				edges = append(edges, newNode)
			}
		}

		if currentNode.Position.X > 0 {
			tile := level.Tiles[level.GetIndexFromXY(currentNode.Position.X-1, currentNode.Position.Y)]
			if tile.TileType != data.WALL {
				//The location is in the map bounds and is walkable
				leftNodePosition := components.Position{
					X: currentNode.Position.X - 1,
					Y: currentNode.Position.Y,
				}
				newNode := newNode(currentNode, &leftNodePosition)
				edges = append(edges, newNode)
			}
		}

		if currentNode.Position.X < data.ScreenWidth {
			tile := level.Tiles[level.GetIndexFromXY(currentNode.Position.X+1, currentNode.Position.Y)]
			if tile.TileType != data.WALL {
				//The location is in the map bounds and is walkable
				rightNodePosition := components.Position{
					X: currentNode.Position.X + 1,
					Y: currentNode.Position.Y,
				}
				newNode := newNode(currentNode, &rightNodePosition)
				edges = append(edges, newNode)
			}
		}

		for _, edge := range edges {
			if IsInSlice(closedList, edge) {
				continue
			}
			edge.g = currentNode.g + 1
			edge.h = edge.Position.GetManhattanDistance(endNodePlaceholder.Position)
			edge.f = edge.g + edge.h

			if IsInSlice(openList, edge) {
				//Loop through and check g values
				isFurther := false
				for _, n := range openList {
					if edge.g > n.g {
						isFurther = true
						break
					}
				}
				if isFurther {
					continue
				}
			}
			openList = append(openList, edge)
		}
	}
	return nil
}

func ReverseSlice(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		panic(errors.New("data must be a slice type"))
	}
	valueLen := value.Len()
	for i := 0; i <= int((valueLen-1)/2); i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
}

func IsInSlice(s []*Node, target *Node) bool {
	for _, n := range s {
		if n.IsEqual(target) {
			return true
		}
	}
	return false
}
