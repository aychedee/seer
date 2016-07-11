package main

import (
	"math"
	"testing"
)

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
