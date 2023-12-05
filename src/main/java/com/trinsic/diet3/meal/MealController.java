package com.trinsic.diet3.meal;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.Optional;
import java.time.LocalDate;

import org.json.*;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.food.FoodRepository;
import com.trinsic.diet3.dieter.DieterService;
import com.trinsic.diet3.dieter.Dieter;

@RestController
@RequestMapping(path = "api/v1/meal")
public class MealController{

    private final MealService mealService;
    private final FoodRepository foodRepository;
    private final DieterService dieterService;

    public MealController(MealService mealService, FoodRepository foodRepository, DieterService dieterService){
        this.mealService = mealService;
        this.foodRepository = foodRepository;
        this.dieterService = dieterService;
    }

    @PostMapping("/food")
    @ResponseBody
    public Integer addFood(@RequestBody String foodBlock){
       String food;
       String dieter;
       JSONObject requestObject = new JSONObject(foodBlock);
       food = requestObject.get("name").toString();
       dieter = requestObject.get("dietername").toString();
       Optional<Food> foundFood = foodRepository.findFoodByName(food);
        if(foundFood.isPresent()){
            return mealService.addCalories(foundFood.get(), dieter);
        }
        return Integer.valueOf(-1);
    }

    @PostMapping("/")
    @ResponseBody
    public Integer addMeal(@RequestBody Meal meal){
        meal.setDay(LocalDate.now());
        return mealService.addMeal(meal, meal.getDieter());
    }

}
