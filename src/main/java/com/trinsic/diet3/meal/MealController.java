package com.trinsic.diet3.meal;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.Optional;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.food.FoodRepository;

@RestController
@RequestMapping(path = "api/v1/meal")
public class MealController{

    private final MealService mealService;
    private final FoodRepository foodRepository;

    public MealController(MealService mealService, FoodRepository foodRepository){
        this.mealService = mealService;
        this.foodRepository = foodRepository;
    }

    @PostMapping("/addfood")
    public Integer addFood(String food){
        // Get food from string
        Optional<Food> f = foodRepository.findFoodByName(food);
        if(f.isPresent()){
            return mealService.addCalories(f.get());
        }
        return Integer.valueOf(-1);
    }

    @PostMapping("/addmeal")
    @ResponseBody
    public Integer addMeal(@RequestBody Meal hMeal){
        // call the Meal constructor that uses the current date as the value
        return mealService.addMeal(new Meal(hMeal.getName(), hMeal.getCalories()));
    }

    @GetMapping("/getRemainingCalories")
    public Integer getremainingcalories(){
        return Integer.valueOf(0);
    }

}
