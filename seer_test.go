package main

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"testing"
)

type Point struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Class int     `json:"class"`
}

func TestConnectTwoNeurons(t *testing.T) {
	neuron1 := NewNeuron(math.Tanh)
	neuron2 := NewNeuron(math.Tanh)

	neuron1.Connect(neuron2, 1.0)

	if neuron1.Outgoing[0] != neuron2.Incoming[0] {
		t.Error(
			"Neuron's not connected: %v, %v",
			neuron1.Outgoing,
			neuron2.Incoming,
		)
	}

	if neuron1.Outgoing[0].Weight != 1.0 {
		t.Error("Connection not assigned correct weight")
	}
}

func TestFireConnectedNeurons(t *testing.T) {
	neuron1 := NewNeuron(math.Tanh)
	neuron2 := NewNeuron(math.Tanh)
	neuron3 := NewNeuron(math.Tanh)
	neuron4 := NewNeuron(math.Tanh)

	neuron1.Connect(neuron2, 10.0)
	neuron2.Connect(neuron3, 5.0)
	neuron3.Connect(neuron4, 5.0)

	neuron1.Incoming = []*Synapse{&Synapse{Out: 2}}
	neuron1.Calculate()
	neuron2.Calculate()
	neuron3.Calculate()
	neuron4.Calculate()

	if neuron4.Out < 0.9 {
		t.Errorf("Out not correct, was %f", neuron4.Out)
	}

}

func TestTrainNetwork(t *testing.T) {

	testData, err := ioutil.ReadFile("training_data.json")
	if err != nil {
		t.Errorf("Error getting test data: %v", err)
	}
	points := make([]Point, 0)
	json.Unmarshal(testData, &points)

	input1 := NewNeuron(math.Tanh)
	input2 := NewNeuron(math.Tanh)

	hiddenLayer := []*Neuron{
		NewNeuron(math.Tanh),
		NewNeuron(math.Tanh),
		NewNeuron(math.Tanh),
		NewNeuron(math.Tanh),
	}

	output1 := NewNeuron(math.Tanh)
	output2 := NewNeuron(math.Tanh)

	// Connect our layers
	for _, hl := range hiddenLayer {
		input1.Connect(hl, 1)
		input2.Connect(hl, 1)
		hl.Connect(output1, 0.5)
		hl.Connect(output2, 0.5)
	}

}
