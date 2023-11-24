package com.trinsic.diet3.meal;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.food.Food;

import com.trinsic.diet3.dieter.*;

@Service
public class MealService{

    MealRepository mealRepository;
    DieterRepository dieterRepository;

    public MealService(MealRepository mealRepository, DieterRepository dieterRepository){
        this.mealRepository = mealRepository;
        this.dieterRepository = dieterRepository;
    }

    @Transactional
    public Integer addMeal(Meal newMeal, String dietername){
        // Only add meal if there is not a meal with the same name for the day
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Meal> searchFood = mealRepository.findMealByName(newMeal.getName(), newMeal.getDay(), newMeal.getDieterId());
        if (searchFood.isEmpty()) {
            Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dietername);
            if (searchDieter.isPresent()) {
                queryStatus = mealRepository.addMeal(newMeal.getCalories(), newMeal.getName(), newMeal.getDay(), searchDieter.get().getId(), dietername);
            }
        }
        return queryStatus;
    }

    @Transactional
    public Integer addCalories(Food newFood, String dieterName){
        Integer queryStatus = Integer.valueOf(-1);
        Long dieterid;
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieterName);
        if (searchDieter.isPresent()){
            dieterid = searchDieter.get().getId();
            queryStatus = mealRepository.addFood(newFood.getCalories(),newFood.getName(),LocalDate.now(),dieterid,dieterName);
        }
        return queryStatus;
    }

    public Integer getCalories(String name, LocalDate day, String dieterName){
        // Respond with number of calories or -1 for meal not found
        Long dieterid;
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieterName);
        if (searchDieter.isPresent()){
            dieterid = searchDieter.get().getId();
            Optional<Meal> meal = mealRepository.findMealByName(name, day, dieterid);
            if (meal.isPresent()){
                return meal.get().getCalories();
            }
        }
        return -1;
    }

    public Integer getCaloriesByDay(String dieterName, LocalDate day){
        Long dieterid;
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieterName);
        if (searchDieter.isPresent()){
            dieterid = searchDieter.get().getId();
            Optional<Meal> meal = mealRepository.findMealByDay(day, dieterid);
            if (meal.isPresent()){
                return meal.get().getCalories();
            }
        }
        return 0;
    }

}
