package com.trinsic.diet3.meal;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import org.json.*;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.food.FoodRepository;

import com.trinsic.diet3.dieter.*;

@Service
public class MealService{

    private final MealRepository mealRepository;
    private final DieterRepository dieterRepository;
    private final FoodRepository foodRepository;

    public MealService(MealRepository mealRepository, DieterRepository dieterRepository, FoodRepository foodRepository){
        this.mealRepository = mealRepository;
        this.dieterRepository = dieterRepository;
        this.foodRepository = foodRepository;
    }

    @Transactional
    public Meal addMeal(Meal newMeal, String dietername){
        Optional<Meal> searchMeal = mealRepository.findMealByName(newMeal.getName(), newMeal.getDay(), newMeal.getDieterId());
        if (searchMeal.isEmpty()) {
            newMeal.setDay(LocalDate.now());
            Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dietername);
            if (searchDieter.isPresent()) {
                newMeal.setDieterId(searchDieter.get().getId());
                newMeal.setDieter(dietername);
                mealRepository.addMeal(newMeal.getCalories(), newMeal.getName(), newMeal.getDay(), searchDieter.get().getId(), dietername);
                return getMeal(newMeal);
            }
        }
        return null;
    }

    @Transactional
    public Meal addCalories(String foodData){
       Long dieterid;
       String food;
       String dieter;
       String meal;
       JSONObject requestObject = new JSONObject(foodData);
       food = requestObject.get("name").toString();
       dieter = requestObject.get("dietername").toString();
       meal = requestObject.get("mealname").toString();
       Optional<Food> foundFood = foodRepository.findFoodByName(food);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter);
        if (searchDieter.isPresent() && foundFood.isPresent()){
            dieterid = searchDieter.get().getId();
            Optional<Meal> searchMeal = mealRepository.findMealByName(meal, LocalDate.now(), searchDieter.get().getId());
            if (searchMeal.isPresent()){
                mealRepository.addFood(foundFood.get().getCalories(),searchMeal.get().getId(),foundFood.get().getName(),LocalDate.now(),dieterid,dieter);
                searchMeal.get().setCalories(searchMeal.get().getCalories() + foundFood.get().getCalories());
                return searchMeal.get();
            }
            else{
                mealRepository.addMeal(foundFood.get().getCalories(),meal,LocalDate.now(),dieterid,dieter);
                Optional<Meal> newMeal = mealRepository.findMealByName(meal, LocalDate.now(), searchDieter.get().getId());
                return newMeal.get();
            }
        }
        return null;
    }

    public Integer getCalories(String name, LocalDate day, String dieterName){
        Long dieterid;
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieterName);
        if (searchDieter.isPresent()){
            dieterid = searchDieter.get().getId();
            Optional<Meal> meal = mealRepository.findMealByName(name, day, dieterid);
            if (meal.isPresent()){
                return meal.get().getCalories();
            }
        }
        return 0;
    }

    public Integer getCaloriesByDay(String dieterName, LocalDate day){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieterName);
        if (searchDieter.isPresent()){
            Integer meal = mealRepository.findCaloriesByDay(dieterName, day);
            return meal;
        }
        return 0;
    }

    public Meal getMeal(Meal mealData){
        String dieter = mealData.getDieter();
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter);
        if (searchDieter.isPresent()){
            Optional<Meal> meal = mealRepository.findMealByDay(LocalDate.now(), searchDieter.get().getId(), mealData.getName());
            if (meal.isPresent()){
                return meal.get();
            }
        }
        return null;
    }

}
