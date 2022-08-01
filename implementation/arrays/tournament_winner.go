package arrays

const HomeTeamWon = 1

func TournamentWinner(competitions [][]string, results []int) string {
	lb := make(map[string]int)

	var currWinner string
	currWinnerScore := 0

	for idx, comp := range competitions{
		winner := comp[1]
		if results[idx] == HomeTeamWon {
			winner = comp[0];
		}

		lb[winner] = lb[winner]+3

		if lb[winner] >= currWinnerScore{
			currWinnerScore = lb[winner]
			currWinner = winner
		}
	}

	return currWinner
}