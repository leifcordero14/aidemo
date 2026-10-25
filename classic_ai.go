package main

type ClassicAI struct{}

func (c ClassicAI) RecommendActivity(weather string, temperature float64, isWeekend bool) string {
	if weather == "sunny" {
		if temperature > 30 {
			if isWeekend {
				return "Go to the swimming pool"
			}

			return "Have a refreshing drink"
		}

		if temperature > 20 {
			if isWeekend {
				return "Go for a drive"
			}

			return "Go for a walk nearby"
		}

		return "Dress warmly but go out anyway"
	}

	if weather == "rainy" {
		if isWeekend {
			return "Watch movies at home"
		}

		return "Stay home and work"
	}

	if weather == "cloudy" {
		if temperature > 25 {
			return "Go out anyway, it's not that hot"
		}

		return "Stay home or go out for a short while"
	}

	return "Unable to determine the activity"
}
