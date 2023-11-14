package com.trinsic.diet3.meal;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.food.Food;

@Service
public class MealService{

    MealRepository mealRepository;

    public MealService(MealRepository mealRepository){
        this.mealRepository = mealRepository;
    }

    @Transactional
    public Integer addMeal(Meal newMeal){
        // Only add meal if there is not a meal with the same name for the day
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Meal> searchFood = mealRepository.findMealByName(newMeal.getName(), newMeal.getDay());
        if (searchFood.isEmpty()) {
            queryStatus = mealRepository.addMeal(newMeal.getName(), newMeal.getCalories(), newMeal.getDay());  
        }
        return queryStatus;
    }

    @Transactional
    public Integer addCalories(Food newFood){
        Integer queryStatus = Integer.valueOf(-1);
        queryStatus = mealRepository.addFood(newFood.getName(),newFood.getCalories(),LocalDate.now());
        return queryStatus;
    }

    public Integer getCalories(String name, LocalDate day){
        // Respond with number of calories or -1 for meal not found
        Optional<Meal> meal = mealRepository.findMealByName(name, day);
        if (meal.isPresent()){
            return meal.get().getCalories();
        }
        return -1;
    }

    public Integer getCaloriesByDay(LocalDate day){
        Optional<Meal> meal = mealRepository.findMealByDay(day);
        if (meal.isPresent()){
            return meal.get().getCalories();
        }
        return 0;
    }

    public Integer getTotalCalories(){
        // call to user table and grab total calories set
        return 0;
    }

}
