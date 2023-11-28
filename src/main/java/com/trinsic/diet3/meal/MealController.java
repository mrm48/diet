package com.trinsic.diet3.meal;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.Optional;
import java.time.LocalDate;

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

    @PostMapping("/addfood")
    @ResponseBody
    public Integer addFood(String food, String dietername){
        // Get food from string
        Optional<Food> foundFood = foodRepository.findFoodByName(food);
        if(foundFood.isPresent()){
            return mealService.addCalories(foundFood.get(), dietername);
        }
        return Integer.valueOf(-1);
    }

    @PostMapping("/addmeal")
    @ResponseBody
    public Integer addMeal(@RequestBody Meal hMeal){
        hMeal.setDay(LocalDate.now());
        return mealService.addMeal(hMeal, hMeal.getDieter());
    }

    @GetMapping("/getremainingcalories")
    public Integer getremainingcalories(@RequestBody Dieter dieter){

        String name = dieter.getName();

        Integer usedCalories = mealService.getCaloriesByDay(name, LocalDate.now());
        Integer totalCalories = dieterService.getCaloriesByDay(name, LocalDate.now());

        return totalCalories - usedCalories;
    }

}
