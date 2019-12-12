package main

import (
    _ "errors"
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    _ "math"
)

type Coord struct {
    x int
    y int
    dist int
}

func (i *Coord) Equals(j Coord) bool {
    return i.x == j.x && i.y == j.y
}

func getDelta(dir string) int {
    d, _ := strconv.Atoi(dir[1:])
    return d
}

func getDir(dir string) string {
    return dir[:1]
}

func getIntersect(wire1 []Coord, wire2 []Coord) []Coord {
    intersect := make([]Coord, 0, 0)
    for _, a := range(wire1) {
        for _, b := range(wire2) {
            if (a.Equals(b)) {
                intersect = append(intersect,Coord{a.x, a.y, a.dist+b.dist})
            }
        }
    }
    return intersect
}

func getCoordSet(wire []string) []Coord {
    x := 0
    y := 0
    dist := 0
    coords := make([]Coord, 0, 0)
    for _,v := range(wire) {
        x2 := x
        y2 := y
        dir := getDir(v)
        delta := getDelta(v)
        if(dir == "L") {
            x2 = x - delta
            for i := x-1; i > x2-1; i--{
                dist += 1
                coords = append(coords, Coord{i,y, dist})
            }
            x = x2
        } else if (dir == "R") {
            x2 = x + delta
            for i := x+1; i < x2+1; i++{
                dist += 1
                coords = append(coords, Coord{i,y, dist})
            }
            x = x2
        } else if (dir == "U") {
            y2 = y + delta
            for i := y+1; i < y2 + 1; i++{
                dist += 1
                coords = append(coords, Coord{x,i,dist})
            }
            y = y2
        } else if (dir == "D") {
            y2 = y - delta
            for i := y-1; i > y2-1; i--{
                dist += 1
                coords = append(coords, Coord{x,i,dist})
            }
            y = y2
        } 
    }
    return coords
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getDist(c Coord) int {
    return Abs(c.x) + Abs(c.y)
}

func getMinDist(coords []Coord) int {
    min := -1
    for _, v := range(coords) {
        dist := v.dist 
        if dist < min || min == -1 {
            min = dist
        }
    }
    return min
}

func main() {
    dat, err := ioutil.ReadFile("input.txt")
    if (err != nil) {
        fmt.Println(err)
    }
    wires := strings.Split(string(dat), "\n")
    wire1 := strings.Split(wires[0],",")
    wire2:= strings.Split(wires[1],",")
    coords1 := getCoordSet(wire1)
    coords2 := getCoordSet(wire2)
    intersect := getIntersect(coords1, coords2)
    fmt.Println(intersect)
    minDist := getMinDist(intersect)
    fmt.Println(minDist)
}
