# Advent Of Code - 2024

## Build

    $ go build -o app aoc/day01 && ./app -input=path/to/your/input.txt

## Test suite

    $ go test -v  ./...

### Day 1: Historian Hysteria

[Story](https://adventofcode.com/2024/day/1)

- Parse text file and return two slices of ints
- Sort both slices
- Sum all number distances
- Part 2: calculate a similarity score

### Day 2: Red-Nosed Reports

[Story](https://adventofcode.com/2024/day/2)

- Parse reports and levels
- Track down all reports that seems wrong

### Day 3: Mull It Over

[Story](https://adventofcode.com/2024/day/3)

- Created my first scanner and parser
- Parser is weirdly implemented but it works \o/

### Day 4: Ceres Search

[Story](https://adventofcode.com/2024/day/4)

- Simple implementation, maybe slow at scale
- I think the code is readable

### Day 5: Print Queue

[Story](https://adventofcode.com/2024/day/5)

- Rules validation is ok IMHO
- Part 2 is sloooooooow af.

### Day 6: Guard Gallivant

[Story](https://adventofcode.com/2024/day/6)

- Fun little game
- I bruteforced the search for infinite loop scenarios
- Code takes 2 minutes to complete

### Day 7: Bridge Repair

[Story](https://adventofcode.com/2024/day/7)

- Recursivity ftw \o/
- nice, clean and performant code

### Day 8: Resonant Collinearity

[Story](https://adventofcode.com/2024/day/8)

- Playing with vectors is somewhat fun
- Complex problem with elegant mathematical solution
- Crazy fast

### Day 9: Disk Fragmenter

[Story](https://adventofcode.com/2024/day/9)

- Simple algorithm for a simple problem
- Quite fast

### Day 10: Hoof It

[Story](https://adventofcode.com/2024/day/10)

- Recursivity fun
- Graph traversal
- Fast and Clean

### Day 11: Plutonian Pebbles

[Story](https://adventofcode.com/2024/day/11)

- First part easily done the dumb way
- Rewrite the recursive way for part 2
- removing useless allocation and add cache


### Day 12: Garden Groups

[Story](https://adventofcode.com/2024/day/12)

- I found the part 2 really hard
- Code is ugly and not optimized
- It works and I'm happy to have solved this one
