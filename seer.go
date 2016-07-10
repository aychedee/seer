package main

type Synapse struct {
	Weight float64
}

type Neuron struct {
	Outgoing []*Synapse
	Incoming []*Synapse
}

func (n *Neuron) Connect(neuron *Neuron, weight float64) {
	synapse := &Synapse{Weight: weight}
	n.Outgoing = append(n.Outgoing, synapse)
	neuron.Incoming = append(neuron.Incoming, synapse)
}

func NewNeuron() *Neuron {
	return &Neuron{}
}

func main() {
}
