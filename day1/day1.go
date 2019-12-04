package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func calcFuelFromMass(mass int) int{
    fuelMass := mass/3-2
    if (fuelMass < 0) {
        return 0;
    }
    return fuelMass + calcFuelFromMass(fuelMass)
}

func main() {
    dat, err := ioutil.ReadFile("input.txt")
    if (err != nil) {
        fmt.Println(err)
    }
    fuelSum := 0
    for _, v := range(strings.Split(string(dat), "\n")) {
        moduleMass, err := strconv.Atoi(v)
        if (err != nil){
            fmt.Println("Error parsing to int " + v, err)
            continue
        }
        fuel := calcFuelFromMass(moduleMass)
        fuelSum += fuel
    }
    fmt.Println("FUEL SUM: ")
    fmt.Println(fuelSum)
}
