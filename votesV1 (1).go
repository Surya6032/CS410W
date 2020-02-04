package main

import (
        "fmt"
        "datafile"
        "log"
)

func main() {
        lines, err := datafile.GetStrings("votes.txt")
        if err != nil {
                log.Fatal(err)
        }

        var names []string              // slice of names
        var counts []int                // slice of counts

        for _, line := range lines {
                matched := false
                for i, name := range names {
                        if name==line {
                                counts[i]++
                                matched = true
                        }
                }
                if matched == false {
                        names  = append(names, line)
                        counts = append(counts, 1)
                }

        }
        for i, name := range names {
                fmt.Printf("Votes for %s: %d\n", name, counts[i])
        }
} 
