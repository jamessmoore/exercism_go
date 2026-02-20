package chessboard

// Declare a type named File which stores if a square is occupied by a piece - 
// this will be a slice of bools
type File int
// Declare a type named Chessboard which contains a map of eight Files, accessed 
// with keys from "A" to "H"
type Chessboard map[string] File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	countInFile := 0
	for v := range cb[file] {
		if v != 0 {
			countInFile++
		}
	}
	return countInFile
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	countInRank := 0
	for _, file := range cb {
		if file[rank] {
			countInRank++
		}
	}
	return countInRank
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	countAll := 0
	return countAll
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	countOccupied := 0
	return countOccupied
}
