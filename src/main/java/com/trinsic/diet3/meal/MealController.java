package com.trinsic.diet3.meal;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@RestController
@RequestMapping(path = "api/v1/meal")
public class MealController{

    private final MealService mealService;

    public MealController(MealService mealService){
        this.mealService = mealService;
    }

    @PostMapping("/food")
    @ResponseBody
    public Integer addFood(@RequestBody String foodBlock){
        return mealService.addCalories(foodBlock);
    }

    @PostMapping("/")
    @ResponseBody
    public Integer addMeal(@RequestBody Meal meal){
        return mealService.addMeal(meal, meal.getDieter());
    }

}
