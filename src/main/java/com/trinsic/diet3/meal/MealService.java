package com.trinsic.diet3.meal;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.food.FoodRepository;
import com.trinsic.diet3.foodEntry.FoodEntryRepository;
import com.trinsic.diet3.dieter.*;

@Service
public class MealService{

    private final MealRepository mealRepository;
    private final DieterRepository dieterRepository;
    private final FoodRepository foodRepository;
    private final FoodEntryRepository foodEntryRepository;

    public MealService(MealRepository mealRepository, DieterRepository dieterRepository, FoodRepository foodRepository, FoodEntryRepository foodEntryRepository){
        this.mealRepository = mealRepository;
        this.dieterRepository = dieterRepository;
        this.foodRepository = foodRepository;
        this.foodEntryRepository = foodEntryRepository;
    }

    @Transactional
    public Meal addMeal(Meal requestMeal){
        String requestDieter;
        requestDieter = requestMeal.getDieter();
        Optional<Meal> meal = mealRepository.findMealByName(requestMeal.getName(), requestMeal.getDay(), requestMeal.getDieterId());
        if (meal.isEmpty()) {
            requestMeal.setDay(LocalDate.now());
            Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
            Optional<Food> food = foodRepository.findFoodByName(requestMeal.getFood()[0]);
            if (dieter.isPresent() && food.isPresent()) {
                requestMeal.setDieterId(dieter.get().getId());
                requestMeal.setDieter(requestDieter);
                Meal newMeal = mealRepository.addMeal(food.get().getCalories(), requestMeal.getName(), requestMeal.getDay(), dieter.get().getId(), requestDieter, requestMeal.getFood());
                foodEntryRepository.addFoodEntry(food.get().getID(), newMeal.getId(), food.get().getCalories());
                return newMeal;
            }
        }
        return null;
    }

    @Transactional
    public Meal addCalories(Meal requestMeal){
       String requestFood;
       String requestDieter;
       String requestMealName;
       requestFood = requestMeal.getFood()[0];
       requestDieter = requestMeal.getDieter();
       requestMealName = requestMeal.getName();
       Optional<Food> food = foodRepository.findFoodByName(requestFood);
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
        if (dieter.isPresent() && food.isPresent()){
            Optional<Meal> meal = mealRepository.findMealByName(requestMealName, LocalDate.now(), dieter.get().getId());
            if (meal.isPresent()){
                meal.get().addFood(requestMeal.getFood());
                mealRepository.addFood(food.get().getCalories(),meal.get().getId(),food.get().getName(),LocalDate.now(),dieter.get().getId(),requestDieter,meal.get().getFood());
                meal.get().setCalories(meal.get().getCalories() + food.get().getCalories());
                meal.get().addFood(requestMeal.getFood());
                return meal.get();
            }
            else{
                mealRepository.addMeal(food.get().getCalories(),requestMealName,LocalDate.now(),dieter.get().getId(),requestDieter,requestMeal.getFood());
                Optional<Meal> newMeal = mealRepository.findMealByName(requestMealName, LocalDate.now(), dieter.get().getId());
                return newMeal.get();
            }
        }
        return null;
    }

    public Integer getCalories(String requestMeal, LocalDate requestDay, String requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
        if (dieter.isPresent()){
            Optional<Meal> meal = mealRepository.findMealByName(requestMeal, requestDay, dieter.get().getId());
            if (meal.isPresent()){
                return meal.get().getCalories();
            }
        }
        return 0;
    }

    public Integer getCaloriesByDay(String requestDieter, LocalDate requestDay){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
        if (dieter.isPresent()){
            Integer meal = mealRepository.findCaloriesByDay(requestDieter, requestDay);
            return meal;
        }
        return 0;
    }

    public Meal getMeal(Meal requestMeal){
        String requestDieter = requestMeal.getDieter();
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
        if (dieter.isPresent()){
            Optional<Meal> meal = mealRepository.findMealByDay(LocalDate.now(), dieter.get().getId(), requestMeal.getName());
            if (meal.isPresent()){
                return meal.get();
            }
        }
        return null;
    }

}
