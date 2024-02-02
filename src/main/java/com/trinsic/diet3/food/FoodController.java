package com.trinsic.diet3.food;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;



@RestController
@RequestMapping(path = "api/v1/food")
public class FoodController {

    private final FoodService foodService;

    public FoodController(FoodService foodService){
        this.foodService = foodService;
    }

	@GetMapping("/calories")
    @ResponseBody
	public Food listCalories(@RequestBody Food food){
        return foodService.listCalories(food);
	}

    @PostMapping("/calories")
    @ResponseBody
    public Food addCalories(@RequestBody Food food){
        return foodService.addCaloriesByName(food);
    }

    @PutMapping("/calories")
    @ResponseBody
    public Food setCaloriesByName(@RequestBody Food food){
        return foodService.addCaloriesByName(food);
    }

    @PostMapping("/")
    @ResponseBody
    public Food addFood(@RequestBody Food food){
        return foodService.addFood(food);
    }

}
