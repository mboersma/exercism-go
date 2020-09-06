// Package tournament tallies up and summarizes the results of a
// football competition.
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Stats manages a team's results from matches played.
type Stats struct {
	Team   string // Team is the unique name identifying a group of players.
	Wins   int    // Wins are how many matches the team has won.
	Losses int    // Losses are how many matches the team has lost.
	Draws  int    // Draws are how many matches the team has tied.
}

// MatchesPlayed are the total number of matches played by this team.
func (s *Stats) MatchesPlayed() int { return s.Wins + s.Losses + s.Draws }

// Points are the team's score, calculated as 3 points per win and 1 per draw.
func (s *Stats) Points() int { return s.Wins*3 + s.Draws }

// Tally reads in win/loss/draw records and writes out a summary of competition.
func Tally(reader io.Reader, writer io.Writer) error {
	competition := map[string]Stats{}

	// read input data and build a map of Team name -> Stats
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return errors.New("invalid input line: " + line)
		}
		team1, team2, result := parts[0], parts[1], parts[2]
		team1Stats, team2Stats := competition[team1], competition[team2]
		team1Stats.Team, team2Stats.Team = team1, team2
		switch result {
		case "win":
			team1Stats.Wins++
			team2Stats.Losses++
		case "loss":
			team1Stats.Losses++
			team2Stats.Wins++
		case "draw":
			team1Stats.Draws++
			team2Stats.Draws++
		default:
			return errors.New("invalid input line: " + line)
		}
		competition[team1], competition[team2] = team1Stats, team2Stats
	}

	// put stats into a slice, sorted by Points, then Team name
	stats := []Stats{}
	for _, stat := range competition {
		stats = append(stats, stat)
	}
	sort.Slice(stats, func(i, j int) bool {
		if stats[i].Points() != stats[j].Points() {
			return stats[i].Points() > stats[j].Points()
		}
		return stats[i].Team < stats[j].Team
	})

	// write out a formatted summary of competition
	header := "Team                           | MP |  W |  D |  L |  P\n"
	if _, err := writer.Write([]byte(header)); err != nil {
		return err
	}
	for _, stat := range stats {
		line := fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", stat.Team,
			stat.MatchesPlayed(), stat.Wins, stat.Draws, stat.Losses, stat.Points())
		if _, err := writer.Write([]byte(line)); err != nil {
			return err
		}
	}
	return nil
}
