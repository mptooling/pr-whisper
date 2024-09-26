package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type PrReviewer struct {
	url     string
	headers map[string]string
}

func NewPrReviewer(apiUrl, token, repo, pullRequestNumber string) *PrReviewer {
	url := fmt.Sprintf("%s/repos/%s/pulls/%s/reviews", apiUrl, repo, pullRequestNumber)
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"Authorization":        "Bearer " + token,
		"X-GitHub-Api-Version": "2022-11-28",
	}

	return &PrReviewer{
		url:     url,
		headers: headers,
	}
}

func (client PrReviewer) comment(comments []*Comment) error {
	if len(comments) == 0 {
		return client.commentInspiration()
	}

	return client.commentWhispers(comments)
}

func (client PrReviewer) commentInspiration() error {
	inspiration := PRInspiration{
		Body:  `> [!NOTE]` + "\n" + `> ` + client.randomInspiration() + " Good Job!\n",
		Event: "COMMENT",
	}

	jsonData, err := json.Marshal(inspiration)
	if err != nil {
		return err
	}

	return client.send(jsonData)
}

func (client PrReviewer) commentWhispers(comments []*Comment) error {
	var commentType string
	var cs []PrReviewComment
	commentTypeString := Note
	for _, c := range comments {
		if c.Type > commentTypeString {
			commentTypeString = c.Type
		}

		cs = append(cs, PrReviewComment{
			Path:     c.FilePath,
			Position: c.Position,
			Body:     c.Content,
		})
	}
	switch commentTypeString {
	case Important:
		commentType = "IMPORTANT"
	case Caution:
		commentType = "CAUTION"
	case Warning:
		commentType = "WARNING"
	case Tip:
		commentType = "TIP"
	default:
		commentType = "NOTE"
	}

	review := PRReview{
		Body:     `> [!` + commentType + `]` + "\n" + `> ` + client.randomIntroString(commentTypeString) + "\n",
		Event:    "COMMENT",
		Comments: cs,
	}

	jsonData, err := json.Marshal(review)
	if err != nil {
		return err
	}

	return client.send(jsonData)
}

func (client PrReviewer) send(jsonData []byte) error {
	req, err := http.NewRequest("POST", client.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	for key, value := range client.headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	if resp.Status != "200 OK" {
		return fmt.Errorf("error commenting on PR: %s", resp.Status)
	}

	return nil
}

func (client PrReviewer) randomIntroString(severityLevel int) string {
	switch severityLevel {
	case Caution:
		return client.randomCautionString()
	case Important:
		return client.randomImportantString()
	case Warning:
		return client.randomWarningString()
	default:
		return client.randomNoteString()
	}
}

func (client PrReviewer) randomCautionString() string {
	phrases := []string{
		"Before you push this code to run,\n> Beware, my friend, you’ve just begun.\n> Tread lightly here, the ground is thin,\n> One false move, and bugs creep in.",
		"Proceed with care, for danger’s near,\n> The code’s not bad, but not quite clear.\n> Refactor now, or soon you’ll find,\n> It’s chaos left for future kind.",
		"A word of caution, soft and fair,\n> Your code is close, but needs repair.\n> Not yet the beast, but seeds you sow—\n> Of bugs to come, you soon may know.",
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	return phrases[r.Intn(len(phrases))]
}

func (client PrReviewer) randomImportantString() string {
	phrases := []string{
		"An important note, I must insist,\n> This code you’ve built—don’t let it twist.\n> For in its heart, a flaw resides,\n> Fix it now, or chaos abides.",
		"A crucial point, I raise my voice,\n> Your code demands a wise, firm choice.\n> Neglect it now, and soon you’ll find,\n> A world of bugs, unkind and blind.",
		"This is important, heed my plea,\n> Your code’s a ship, adrift at sea.\n> With one swift change, it sails so true,\n> Ignore it now? You’ll rue the view.",
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	return phrases[r.Intn(len(phrases))]
}

func (client PrReviewer) randomWarningString() string {
	phrases := []string{
		"A warning rises, bold and true,\n> This code, I fear, won’t quite pull through.\n> Errors wait in lines of dread,\n> Tread here, and soon you’ll pull your head.",
		"This warning sounds, a siren’s call,\n> For errors hide behind the wall.\n> Fix it now, or risk the fall,\n> Of a PR that fails us all.",
		"A warning, coder, hear it well,\n> Your logic leads to coding hell.\n> Take heed and fix before too late,\n> Lest broken builds shall be your fate.",
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	return phrases[r.Intn(len(phrases))]
}

func (client PrReviewer) randomNoteString() string {
	phrases := []string{
		"A gentle note, if I may share,\n> A tiny tweak to show I care.\n> This line is fine, but here’s the catch,\n> It could be cleaner, it’s no match.",
		"A note, a whisper in the wind,\n> Your code is good, but there’s a pin—\n> Just one small thing, to help you out,\n> Consider this, or have some doubt.",
		"Just a note, no need for stress,\n> Your work’s not bad, but here’s my guess—\n> If you’d adjust this single line,\n> The code would shine, and all be fine.",
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	return phrases[r.Intn(len(phrases))]
}
func (client PrReviewer) randomInspiration() string {
	phrases := []string{
		"🌟 Your creativity in solving challenges inspires those around you to think outside the box.",
		"🚀 Every line of code you write contributes to the greater vision we’re building together.",
		"🙌 Your dedication to continuous learning sets a remarkable example for the entire team.",
		"✨ The passion you bring to your work fosters a collaborative spirit that elevates us all.",
		"🌻 Your ability to embrace challenges transforms obstacles into opportunities for growth.",
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	return phrases[r.Intn(len(phrases))]
}
