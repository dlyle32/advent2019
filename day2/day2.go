package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func executeOpCode(program []int, idx int) (bool, error) {
    opCode := program[idx]
    if(opCode == 99){
    //    fmt.Println("Terminating")
        return true, nil
    } else if(opCode == 1){
     //   fmt.Println(fmt.Sprintf("program[%d] = %d + %d", program[idx+3], program[program[idx+1]], program[program[idx+2]]))
        program[program[idx+3]] = program[program[idx+1]] + program[program[idx+2]]
        return false, nil
    } else if(opCode == 2){
      //  fmt.Println(fmt.Sprintf("program[%d] = %d * %d", program[idx+3], program[program[idx+1]], program[program[idx+2]]))
        program[program[idx+3]] = program[program[idx+1]] * program[program[idx+2]]
        return false, nil
    }
    return false, errors.New(fmt.Sprintf("Unknown Op Code: %d", opCode)) 
}

func executeProgram(program []int) {
    for i := 0; i < len(program); i+=4 {
        term, err := executeOpCode(program, i)
        if term {
            return
        }
        if err != nil {
            fmt.Println(err)
            break;
        }
    }
}

func main() {
    dat, err := ioutil.ReadFile("input.txt")
    if (err != nil) {
        fmt.Println(err)
    }
    strProgram := strings.Split(string(dat), ",")
    program := []int{}
    for _, v := range(strProgram) {
        i, err := strconv.Atoi(strings.TrimSuffix(v,"\n"))
        if (err != nil) {
            fmt.Println(err)
            break;
        }
        program = append(program, i)
    }
    runProgram := make([]int, len(program), cap(program))
    for n := 0; n <= 99; n++ {
        for v := 0; v <= 99; v++ {
            copy(runProgram, program)
            runProgram[1] = n
            runProgram[2] = v
            executeProgram(runProgram)
            if (runProgram[0] == 19690720) {
                fmt.Println(fmt.Sprintf("Noun: %d, Verb: %d", n, v))
                return;
            }
        }
    }
    fmt.Println(program)
}
