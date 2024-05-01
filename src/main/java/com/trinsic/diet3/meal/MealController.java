package com.trinsic.diet3.meal;

import org.springframework.web.bind.annotation.RestController;

import com.trinsic.diet3.entrycreaterequest.Entrycreaterequest;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.DeleteMapping;
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
    public Meal addFood(@RequestBody Entrycreaterequest ecr){
        return mealService.addCalories(ecr.getMeal(), ecr.getFood());
    }

    @DeleteMapping("/food")
    @ResponseBody
    public Meal removeFood(@RequestBody Entrycreaterequest ecr){
        return mealService.removeCalories(ecr.getMeal(), ecr.getFood());
    }

    @PostMapping("/")
    @ResponseBody
    public Meal addMeal(@RequestBody Meal meal){
        return mealService.addMeal(meal);
    }

    @GetMapping("/")
    @ResponseBody
    public Meal getMeal(@RequestBody Meal meal){
        return mealService.getMeal(meal);
    }

}
