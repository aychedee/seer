package main

type ActivationFunction func(n float64) float64

type Synapse struct {
	Weight float64
	In     float64
	Out    float64
}

func (s *Synapse) Activate(value float64) {
	s.In = value
	s.Out = value * s.Weight
}

type Neuron struct {
	Out                float64
	Outgoing           []*Synapse
	Incoming           []*Synapse
	ActivationFunction ActivationFunction
}

func (n *Neuron) Connect(neuron *Neuron, weight float64) {
	synapse := &Synapse{Weight: weight}
	n.Outgoing = append(n.Outgoing, synapse)
	neuron.Incoming = append(neuron.Incoming, synapse)
}

func (n *Neuron) Calculate() {

	var sum float64
	for _, s := range n.Incoming {
		sum += s.Out
	}

	n.Out = n.ActivationFunction(sum)

	for _, s := range n.Outgoing {
		s.Activate(n.Out)
	}

}

func NewNeuron(af ActivationFunction) *Neuron {
	return &Neuron{ActivationFunction: af}
}

func main() {
}
