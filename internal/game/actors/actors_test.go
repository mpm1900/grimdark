package actors

import (
	"grimdark/internal/game"
	"grimdark/internal/game/weapons"
	"testing"

	"github.com/google/uuid"
)

func TestHydrateActorConfigIgnoresEmptyWeaponSlots(t *testing.T) {
	if weapons.TomeOfSacrifice().ID == uuid.Nil {
		t.Fatal("Tome of Sacrifice must have a stable non-nil ID")
	}

	classID := Bloodknight.ID
	greatswordSlot := uuid.New()
	emptySlot := uuid.New()

	actor, ok := HydrateActorConfig(game.ActorConfig{
		Class: &classID,
		Weapons: map[uuid.UUID]uuid.UUID{
			greatswordSlot: weapons.Greatsword().ID,
			emptySlot:      uuid.Nil,
		},
	})
	if !ok {
		t.Fatal("expected actor config to hydrate")
	}

	if len(actor.Weapons) != 1 {
		t.Fatalf("expected one hydrated weapon, got %d", len(actor.Weapons))
	}

	weapon, ok := actor.Weapons[greatswordSlot]
	if !ok {
		t.Fatal("expected Greatsword slot to be hydrated")
	}
	if weapon.ID != weapons.Greatsword().ID {
		t.Fatalf("expected Greatsword, got %s", weapon.ID)
	}
	if _, ok := actor.Weapons[emptySlot]; ok {
		t.Fatal("expected empty weapon slot to be ignored")
	}
}

func TestHydrateActorConfigFiltersWeaponsByStrength(t *testing.T) {
	classID := Bloodknight.ID
	tomeSlot := uuid.New()
	greatswordSlot := uuid.New()

	actor, ok := HydrateActorConfig(game.ActorConfig{
		Class: &classID,
		Weapons: map[uuid.UUID]uuid.UUID{
			tomeSlot:       weapons.TomeOfSacrifice().ID,
			greatswordSlot: weapons.Greatsword().ID,
		},
	})
	if !ok {
		t.Fatal("expected actor config to hydrate")
	}

	if len(actor.Weapons) != 1 {
		t.Fatalf("expected one hydrated weapon, got %d", len(actor.Weapons))
	}

	weapon, ok := actor.Weapons[greatswordSlot]
	if !ok {
		t.Fatal("expected Greatsword slot to be kept")
	}
	if weapon.ID != weapons.Greatsword().ID {
		t.Fatalf("expected Greatsword, got %s", weapon.ID)
	}
	if _, ok := actor.Weapons[tomeSlot]; ok {
		t.Fatal("expected overweight Tome slot to be ignored")
	}
}
