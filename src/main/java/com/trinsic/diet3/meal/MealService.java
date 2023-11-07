package com.trinsic.diet3.meal;
import java.time.LocalDate;
import java.util.List;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.food.Food;

@Service
public class MealService{

    MealRepository mealRepository;
    Integer calories;

    @Transactional
    public Integer addMeal(Meal f){
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Meal> searchFood = mealRepository.findMealByName(f.getName(), f.getDay());
        if (searchFood.isEmpty()) {
            queryStatus = mealRepository.addMeal(f.getName(), f.getCalories(), f.getDay());  
        }
        return queryStatus;
    }

    public Integer getCalories(String name, LocalDate day){
        this.calories = -1;
        Optional<Meal> qMeal = mealRepository.findMealByName(name, day);
        qMeal.ifPresent(this::setCalories);
        return this.calories;
    }

    public void setCalories(Meal m){
        this.calories = m.getCalories();
    }

}
