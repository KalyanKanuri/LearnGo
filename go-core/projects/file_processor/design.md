# File Processor Project

-------------------------

Build a concurrent file processor

## Project Requirements

Read Files in Directory, count words in each file and show output as which file contains how many words

## High Level Design

### Init stage

1. initialize required jobs channel and results channel of type struct
2. initialize goroutines of certain pool length called workers along with waitgroup

### Flow stage

1. get files from directory and push to jobs channel
2. workers takes files from jobs channel
3. workers reads the file if no error else skips
4. workers counts the number of words and build the result struct
5. workers push the result struct into results channel

### Result Processing stage

1. in main wait for workers to be completed in a goroutine
2. after all the workers are completed close the results channel
3. simultaneously when waiting main should process the results pushed into main channel
4. main prints output
