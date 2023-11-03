package com.trinsic.diet3.meal;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import java.util.List;
import com.trinsic.diet3.meal.Meal;
import com.trinsic.diet3.meal.MealService;

@RestController
@RequestMapping(path = "api/v1/meal")
public class MealController{

    private final MealService mealService;

    public MealController(MealService mealService){
        this.mealService = mealService;
    }

	@GetMapping("/listmeals")
	public List<Meal> listMeals(){
        return mealService.listMeals();
	}

    @PostMapping("/addmeal")
    public void addMeal(){
    
    }

    @GetMapping("/listtodaysmeals")
    public List<Meal> listTodaysMeals(){
        return List.of(new Meal());
    }

    @GetMapping("/getRemainingCalories")
    public Integer getremainingcalories(){
        return Integer.valueOf(0);
    }

}
