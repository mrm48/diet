package com.trinsic.diet3.meal;
import java.util.List;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.food.Food;

@Service
public class MealService{

    MealRepository mealRepository;

    @Transactional
    public Integer addMeal(Meal f){
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Meal> searchFood = mealRepository.findMealByName(f.getName(), f.getDay());
        if (searchFood.isEmpty()) {
            queryStatus = mealRepository.addMeal(f.getName(), f.getCalories(), f.getDay());  
        }
        return queryStatus;
    }

}
