package core

import (
	"testing"
	"time"

	"github.com/hexbotio/hex/models"
)

func setupPipelines() (config models.Config) {
	config.Pipelines = make([]models.Pipeline, 2)
	config.Pipelines[0] = models.Pipeline{Name: "Test1"}
	config.Pipelines[1] = models.Pipeline{Name: "Test2"}
	return config
}

func TestStateLength(t *testing.T) {
	config := setupPipelines()
	state = NewState(&config)
	e := 2
	a := len(state.States)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestStateRunningFalse(t *testing.T) {
	config := setupPipelines()
	state = NewState(&config)
	state.SetRunning("Test1", false)
	e := false
	a := state.States["Test1"].Running
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestStateRunningTrue(t *testing.T) {
	config := setupPipelines()
	state = NewState(&config)
	state.SetRunning("Test1", true)
	e := true
	a := state.States["Test1"].Running
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestStateSuccessFalse(t *testing.T) {
	config := setupPipelines()
	state = NewState(&config)
	state.SetState("Test1", false)
	e := false
	a := state.States["Test1"].Success
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestStateSuccessTrue(t *testing.T) {
	config := setupPipelines()
	state = NewState(&config)
	state.SetState("Test1", true)
	e := true
	a := state.States["Test1"].Success
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestStateChange(t *testing.T) {
	config := setupPipelines()
	state = NewState(&config)
	state.SetState("Test1", false)
	e := state.States["Test1"].LastChange
	time.Sleep(1 * time.Second)
	state.SetState("Test1", true)
	a := state.States["Test1"].LastChange
	if e >= a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}
