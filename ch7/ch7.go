package main

import (
	"io"
	"os"
	"slices"
)

type Team struct {
	TeamName    string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) _DefineWinningTeam(firstTeam string, firstScore int, secondTeam string, secondScore int) string {
	if firstScore > secondScore {
		return firstTeam
	}
	return secondTeam
}

func (l *League) MatchResult(firstTeam string, firstScore int, secondTeam string, secondScore int) {

	winningTeam := l._DefineWinningTeam(firstTeam, firstScore, secondTeam, secondScore)

	l.Wins[winningTeam]++
}

func (l League) Ranking() []string {
	var teams []string

	for _, t := range l.Teams {
		teams = append(teams, t.TeamName)
	}

	slices.SortFunc(teams, func(a, b string) int {
		if l.Wins[a] < l.Wins[b] {
			return 1
		}
		if l.Wins[b] < l.Wins[a] {
			return -1
		}
		return 0
	})

	return teams
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for _, s := range results {
		io.WriteString(w, s)
		w.Write([]byte("\n"))
	}
}

func main() {
	league := League{
		Teams: []Team{
			{
				TeamName:    "Lions",
				PlayerNames: []string{"Ana", "Bruno", "Caio", "Davi", "Eli"},
			},
			{
				TeamName:    "Tigers",
				PlayerNames: []string{"Fernanda", "Gabi", "Heitor", "Igor", "Julia"},
			},
			{
				TeamName:    "Bulls",
				PlayerNames: []string{"Kai", "Lara", "Mateus", "Nina", "Otavio"},
			},
			{
				TeamName:    "Hawks",
				PlayerNames: []string{"Paula", "Quintino", "Rafa", "Sofia", "Tiago"},
			},
		},
		Wins: map[string]int{
			"Lions":  0,
			"Tigers": 0,
			"Bulls":  0,
			"Hawks":  0,
		},
	}

	league.MatchResult("Lions", 88, "Tigers", 79) // Lions win
	league.MatchResult("Bulls", 91, "Hawks", 85)  // Bulls win
	league.MatchResult("Lions", 95, "Bulls", 90)  // Lions win
	league.MatchResult("Hawks", 84, "Tigers", 80) // Hawks win
	league.MatchResult("Lions", 99, "Hawks", 89)  // Lions win

	// Expected wins:
	// Lions: 3
	// Bulls: 1
	// Hawks: 1
	// Tigers: 0

	println("Ranking:")
	RankPrinter(league, os.Stdout)
}
