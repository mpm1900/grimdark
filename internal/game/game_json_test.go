package game

import (
	"testing"

	"github.com/google/uuid"
)

func TestDiffGameJSONBuildsSparsePatch(t *testing.T) {
	actorID := uuid.New()
	firstLog := Bindable[Log]{
		ID:      uuid.New(),
		Context: NewContext(),
		Payload: NewLog("first", map[string]string{}),
	}
	secondLog := Bindable[Log]{
		ID:      uuid.New(),
		Context: NewContext(),
		Payload: NewLog("second", map[string]string{}),
	}

	previous := GameJSON{
		Actors:    []actorJSON{{ID: actorID, Name: "old"}},
		Logs:      []Bindable[Log]{firstLog},
		Positions: []Position{{ID: uuid.New()}},
		Prompts: []Bindable[actionJSON]{
			{ID: uuid.New(), Context: NewContext(), Payload: actionJSON{ID: uuid.New()}},
		},
	}
	next := previous
	next.Actors = []actorJSON{{ID: actorID, Name: "new"}}
	next.Logs = []Bindable[Log]{firstLog, secondLog}
	next.Prompts = []Bindable[actionJSON]{}

	patch, ok := DiffGameJSON(previous, next)
	if !ok {
		t.Fatal("expected diff to be safe")
	}
	if len(patch.Actors) != 1 || patch.Actors[0].Name != "new" {
		t.Fatalf("expected only changed actor in patch, got %#v", patch.Actors)
	}
	if len(patch.Logs) != 1 || patch.Logs[0].ID != secondLog.ID {
		t.Fatalf("expected only appended log in patch, got %#v", patch.Logs)
	}
	if patch.Positions != nil {
		t.Fatalf("expected unchanged positions to be nil, got %#v", patch.Positions)
	}
	if patch.Prompts == nil || len(patch.Prompts) != 0 {
		t.Fatalf("expected changed empty prompts to be non-nil empty slice, got %#v", patch.Prompts)
	}
}

func TestDiffGameJSONRejectsNonAppendOnlyLogs(t *testing.T) {
	previous := GameJSON{
		Logs: []Bindable[Log]{
			{ID: uuid.New(), Context: NewContext(), Payload: NewLog("old", map[string]string{})},
		},
	}
	next := GameJSON{
		Logs: []Bindable[Log]{
			{ID: uuid.New(), Context: NewContext(), Payload: NewLog("new", map[string]string{})},
		},
	}

	if _, ok := DiffGameJSON(previous, next); ok {
		t.Fatal("expected non-append-only logs to be rejected")
	}
}

func TestDiffGameJSONAllowsUnchangedState(t *testing.T) {
	actorID := uuid.New()
	actionID := uuid.New()
	state := GameJSON{
		Actors: []actorJSON{{ID: actorID, Name: "same"}},
		Commands: []Bindable[actionJSON]{
			{ID: uuid.New(), Context: NewContext(), Payload: actionJSON{ID: actionID}},
		},
		Logs: []Bindable[Log]{
			{ID: uuid.New(), Context: NewContext(), Payload: NewLog("same", map[string]string{})},
		},
	}
	next := state
	next.Actors = []actorJSON{{ID: actorID, Name: "same", Seen: true}}
	next.Commands = []Bindable[actionJSON]{
		{ID: uuid.New(), Context: NewContext(), Payload: actionJSON{ID: actionID}},
	}

	patch, ok := DiffGameJSON(state, next)
	if !ok {
		t.Fatal("expected unchanged state to be safe to patch")
	}
	if patch.Actors != nil {
		t.Fatalf("expected unchanged actors to be nil, got %#v", patch.Actors)
	}
	if patch.Logs != nil {
		t.Fatalf("expected unchanged logs to be nil, got %#v", patch.Logs)
	}
	if patch.Commands != nil {
		t.Fatalf("expected commands with only regenerated bind IDs to be nil, got %#v", patch.Commands)
	}
}
