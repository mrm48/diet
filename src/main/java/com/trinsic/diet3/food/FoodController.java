package com.trinsic.diet3.food;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;



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
	public Integer listCalories(@RequestBody Food hFood){
        json = hFood.getName();
        return foodService.listCalories(json);
	}

    @PostMapping("/addcalories")
    @ResponseBody
    public Integer addCalories(@RequestBody Food hFood){
        return foodService.addCaloriesByName(hFood.getName(), hFood.getCalories());
    }

    @PostMapping("/addfood")
    @ResponseBody
    public Integer addFood(@RequestBody Food hFood){
        return foodService.addFood(new Food(hFood.getName(), hFood.getUnits(), hFood.getCalories()));
    }

    @GetMapping("/getRemainingCalories")
    public Integer getremainingcalories(){
        return Integer.valueOf(0);
    }
}
