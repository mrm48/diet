package com.trinsic.diet3.food;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import java.util.List;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.meal.Meal;
import com.trinsic.diet3.meal.MealService;

@RestController
@RequestMapping(path = "api/v1/food")
public class FoodController {
    private final FoodService foodService;

    public FoodController(FoodService foodService){
        this.foodService = foodService;
    }

	@GetMapping("/listcalories")
	public Integer listCalories(){
        return foodService.listCalories();
	}

    @PostMapping("/addcalories")
    public void addCalories(){
        
    }

    @GetMapping("/getRemainingCalories")
    public Integer getremainingcalories(){
        return Integer.valueOf(0);
    }
}
