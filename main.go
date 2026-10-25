package main

import "fmt"

func main() {
	// ===== CLASSIC AI =====
	// Classic AI makes decisions using explicitly defined rules and conditions.
	// If it receives a combination of data that is not covered by those rules, it will not be able
	// to provide a specific recommendation and will return a default message.

	weather := "sunny"
	temp := 25.0
	isWeekend := true

	fmt.Println("The plan I recommend is:", ClassicAI{}.RecommendActivity(weather, temp, isWeekend))

	// ===== MODERN AI =====
	// Modern AI uses training data to calculate a temperature threshold.
	// In this example, the threshold is obtained by calculating the average of the
	// provided temperatures. Therefore, even though the value 30.0 has the label "High",
	// the model may classify it as "Low" if its value is below the calculated average.
	// This demonstrates that the result depends on the training data and the method used.

	data := []TemperatureTrainingData{
		{value: 10.0, label: "Low"},
		{value: 20.0, label: "Low"},
		{value: 30.0, label: "High"},
		{value: 80.0, label: "High"},
		{value: 90.0, label: "High"},
	}

	newTemperature := 30.0

	modernAI := ModernAI{}

	modernAI.Train(data)
	fmt.Println("The result is:", modernAI.ClassifyTemperature(newTemperature))
}
