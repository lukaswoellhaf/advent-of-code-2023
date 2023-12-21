package day3

//--- Day 3: Gear Ratios ---
// The engine schematic (your puzzle input) consists of a visual representation of the engine.
// There are lots of numbers and symbols you don't really understand, but apparently any number adjacent
// to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not
// count as a symbol.)

//  What is the sum of all of the part numbers in the engine schematic?

// Parse input into 2D array
// For each x y in 2D array:
// 2.1 Find digit
// 2.2 If digit found: look for neighbour digits
// 2.3 Save coordinates of found PartNumber
// 2.4 Create hit box for found part number of specific length
// For each digit in hitbox:
// 2.4.1 Hitbox must search for matches: (x0 y+1; x0 y-1; x+1 y0; x-1 y0)
// 2.4.2 Save Part Number as soon as a symbol is found in a hitbox of a digit
// 2.5 Mark found Part Numbers as '.' to be not considered anymore

//2.6 Sum stored PartNumbers
