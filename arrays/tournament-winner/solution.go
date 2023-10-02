package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	scores := make(map[string]int)
	winner := ""
	for i := 0; i < len(competitions); i++ {
		team := competitions[i][HOME_TEAM_WON-results[i]]
		scores[team] += 3
		if scores[team] > scores[winner] {
			winner = team
		}
	}
	return winner
}