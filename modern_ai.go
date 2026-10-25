package main

type ModernAI struct {
	mean float64
}

func (m ModernAI) ClassifyTemperature(newTemperature float64) string {
	if newTemperature >= m.mean {
		return "High"
	}

	return "Low"
}

func (m *ModernAI) Train(temperatures []TemperatureTrainingData) {
	length := len(temperatures)

	if length == 0 {
		m.mean = 0
		return
	}

	sum := 0.0

	for _, temperature := range temperatures {
		sum += temperature.getValue()
	}

	m.mean = sum / float64(length)
}

type TemperatureTrainingData struct {
	value float64
	label string
}

func (t TemperatureTrainingData) getValue() float64 {
	return t.value
}
