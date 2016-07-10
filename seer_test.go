package main

import "testing"

func TestConnectTwoNeurons(t *testing.T) {
	neuron1 := NewNeuron()
	neuron2 := NewNeuron()

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
