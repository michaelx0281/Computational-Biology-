package main

func main() {

}

/*
Let's add some general comments and notes about the Game of Life here! Let's try to finish all of these algorithms today. Also, need to compile a document for what I want to / need to do for the rest of the summer or build on the previous one. Additionally, definitely try to research different people who may or may not be able to take me as a computational bio research student.

The rules of John Conway's Game of Life is deterministic, or applicable across the entire board.

Below are the 5 rules:

1) If a cell is aive and has either two or three live neighbors, then it remains alive. (neighbors are supposed to describe the square of 8 cells around the observed cell)

2) If a cell is alive and has zero or one live neighbors, then it dies out

3) If a cell is alive and has four or more live neighbors, then it dies out

4) If a cell is dead and has more than or fewer than three live neighbors, then it remains dead

5) If a cell is dead and has exactly three live neighbors, then it becomes alive

1 => A Propagation, this allows life to continue if there are enough suitable mates close by
2 => B Lack of mates
3 => C Overpopulation
4 => D rest in peace (most dead cells don't spontaneously regenerate)
5 => E Zombie

There are a few stable forms to the Game of Life where it does not change from each generation to the next

There are also some forms dubbed 'oscillators' which have different periods of length each

In 2023, scientists proved that every positive integer is the period of some Game of Life oscillator (which is pretty cool! --> at least in terms of the sheer number of different things that were able to come out of the simple rules of this model)

**The curious case of the R pentomino**
Most of the simple patterns for the Game of Life either 'die out' or eventually produce an oscillator. There are a few obsessions that distinguish it and make it more than just a random novelty, however

A) Glider
They are small units that seem to move linearly and diagonally by 'gliding' across the board (very interesting!)
B) R pentomino
The interior settles into several stable froms as well as oscillators of p=2

However, 5 gliders are formed early on and one additional sometime later (for a total of 6 gliders!)

C) Gosper's gun

This answers the question of 'Can an automaton's population grow *infinitely large?*

This question remained unsolvec until 1970, when Bill Gosper found teh following pattern named the Gosper glider gun

There are about 2 components or parts which seem to coordinate to create endless amounts of glider's

However, this still does not quite reach the exponential self replication of real human cells, despite the increased number of interesting patterns which were discovered.

Computional Biology Problem:

Input: An initial configuration of the Game of Life board initialBoard, and an integer numGens

Output: All numGens + 1 configurations of this board over numGens generations of the Game of Life, starting with initalBoard

Always try to plan ahead before starting on the code!
-------------

We will explore the paradignm for solving computational problems that is called top-down programming. We will first write a function which solves the problem assuming any subroutines that we need are implemented correctly. We will then write each of these subroutines assuming any other functions they may call as subroutines in return.

With this, we may start out initial planning by creating a function hierarchy or a network where functions are connected by what they invoke through the usage of arrows.

Initial thoughts and ideas--
CreateBoard:

Update inputs:

CreateNewBoard:

Here is the pseudocode provided by the website!

PlayGameOfLife(initialBoard, numGens)
    boards ← array of numGens + 1 game boards
    boards[0] ← initialBoard
    for every integer i from 1 to numGens
        boards[i] ← UpdateBoard(boards[i–1])
    return boards

This is essentially a more crystalized version of what my first thought for implementation of this was.const

We can create a 2D array of bools which will signal which cells are alive or dead!


*/
