package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/nlopes/slack"
)

var ticker *time.Ticker = nil

var businessCalendar = map[string][]string{
	"Jan 31":       []string{"Send 1099’s to Contractors", "Send W2’s to Employees", "Send 1095-B and 1095-C forms to Employees"},
	"March 01":     []string{"Delaware Annual Franchise Report filing due: Pay a min. of $400, +more if you have significant funding", "IRS ACA Compliance 1094-B, 1095-B, 1095-C and 1095-C filings are due (efiling)"},
	"March 31":     []string{"1099’s and W2’s must be e-filed with the IRS by this due date"},
	"April 15":     []string{"C Corp Form 1120 Tax Return due. Can extend to Sept 15.", "File the R&D; Tax Credit Form 6765 with your tax return. Can extend to Sept 15.", "NY and NYC estimated tax payment due."},
	"April 20":     []string{"Claim your R&D; Tax Credits on Form 941."},
	"May 31":       []string{"NYC Business Registration check for current status."},
	"June 15":      []string{"NY and NYC estimated tax payment due."},
	"July 31":      []string{"Claim your R&D; Tax Credits on Form 941."},
	"September 15": []string{"C Corp Form 1120 Tax Return final due date if extension was filed.", "NY and NYC estimated tax payment due."},
	"October 31":   []string{"Claim your R&D; Tax Credits on Form 941."},
	"December 15":  []string{"NY and NYC estimated tax payment due."},
}

func sendMessageToSlack() {
	for k, v := range businessCalendar {
		if time.Now().Format("January 02") == k {
			api := slack.New("xoxp-21288980865-21288980897-171681523890-b2f412bb4ed8aa57ee36568471f4d48f")
			params := slack.PostMessageParameters{}
			params.Username = "Lord Business"
			params.IconURL = "http://vignette2.wikia.nocookie.net/jadensadventures/images/6/65/PrezBiz_filmstill1.jpg/revision/latest?cb=20140222064357"
			channelID, timestamp, err := api.PostMessage("#business-reminders", k+" reminder: "+strings.Join(v, ", "), params)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
			fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
		}
	}
}

func main() {
	sendMessageToSlack()
	time.AfterFunc(duration(), sendMessageToSlack)
	wg.Add(1)
	// do normal task here
	wg.Wait()
}

func duration() time.Duration {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), 9, 0, 0, 0, t.Location())
	if t.After(n) {
		n = n.Add(24 * time.Hour)
	}
	d := n.Sub(t)
	return d
}

var wg sync.WaitGroup
