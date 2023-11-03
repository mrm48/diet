package com.trinsic.diet3.meal;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import java.util.List;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.meal.Meal;
import com.trinsic.diet3.meal.MealService;

@RestController
@RequestMapping(path = "api/v1/meal")
public class MealController{

    private final MealService mealService;

    public MealController(MealService mealService){
        this.mealService = mealService;
    }

/* 	@GetMapping("/listfood")
	public List<Food> listFood(){
        return mealService.listFood();
	} */

    @PostMapping("/addfood")
    public Integer addFood(Food f){
        return f.getCalories();
    }

    @GetMapping("/getRemainingCalories")
    public Integer getremainingcalories(){
        return Integer.valueOf(0);
    }

}
