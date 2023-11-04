package com.trinsic.diet3.food;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.http.HttpEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.List;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.meal.Meal;
import com.trinsic.diet3.meal.MealService;

@RestController
@RequestMapping(path = "api/v1/food")
public class FoodController {
    private final FoodService foodService;
    private String json; 
    public FoodController(FoodService foodService){
        this.foodService = foodService;
    }

	@RequestMapping(value = "/listcalories")
    @ResponseBody
	public Integer listCalories(@RequestBody Food httpEntity){
        json = httpEntity.getName();
        return foodService.listCalories(json);
	}

    @PostMapping("/addcalories")
    public void addCalories(){
        
    }

    @GetMapping("/getRemainingCalories")
    public Integer getremainingcalories(){
        return Integer.valueOf(0);
    }
}
