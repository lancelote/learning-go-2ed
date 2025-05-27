package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (lg *League) MatchResult(name1 string, score1 int, name2 string, score2 int) {
	if score1 > score2 {
		lg.Wins[name1]++
	} else if score2 < score1 {
		lg.Wins[name2]++
	}
}

func (lg League) Ranking() []string {
	names := make([]string, 0, len(lg.Teams))
	for name := range lg.Teams {
		names = append(names, name)
	}

	sort.Slice(names, func(i, j int) bool {
		return lg.Wins[names[i]] > lg.Wins[names[j]]
	})

	return names
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()

	for _, v := range results {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func main() {
	lg := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	lg.MatchResult("Italy", 50, "France", 70)
	lg.MatchResult("India", 85, "Nigeria", 80)
	lg.MatchResult("Italy", 60, "India", 55)
	lg.MatchResult("France", 100, "Nigeria", 110)
	lg.MatchResult("Italy", 65, "Nigeria", 70)
	lg.MatchResult("France", 95, "India", 80)
	RankPrinter(lg, os.Stdout)
}
