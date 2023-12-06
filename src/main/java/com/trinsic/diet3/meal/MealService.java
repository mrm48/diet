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
    public Integer addMeal(Meal newMeal, String dietername){
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Meal> searchMeal = mealRepository.findMealByName(newMeal.getName(), newMeal.getDay(), newMeal.getDieterId());
        if (searchMeal.isEmpty()) {
            newMeal.setDay(LocalDate.now());
            Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dietername);
            if (searchDieter.isPresent()) {
                queryStatus = mealRepository.addMeal(newMeal.getCalories(), newMeal.getName(), newMeal.getDay(), searchDieter.get().getId(), dietername);
            }
        }
        return queryStatus;
    }

    @Transactional
    public Integer addCalories(String foodBlock){
       Integer queryStatus = Integer.valueOf(-1);
       Long dieterid;
       String food;
       String dieter;
       JSONObject requestObject = new JSONObject(foodBlock);
       food = requestObject.get("name").toString();
       dieter = requestObject.get("dietername").toString();
       Optional<Food> foundFood = foodRepository.findFoodByName(food);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter);
        if (searchDieter.isPresent() && foundFood.isPresent()){
            dieterid = searchDieter.get().getId();
            queryStatus = mealRepository.addFood(foundFood.get().getCalories(),foundFood.get().getName(),LocalDate.now(),dieterid,dieter);
        }
        return queryStatus;
    }

    public Integer getCalories(String name, LocalDate day, String dieterName){
        Long dieterid;
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieterName);
        if (searchDieter.isPresent()){
            dieterid = searchDieter.get().getId();
            Optional<Meal> meal = mealRepository.findMealByName(name, day, dieterid);
            if (meal.isPresent()){
                queryStatus = meal.get().getCalories();
            }
        }
        return queryStatus;
    }

    public Integer getCaloriesByDay(String dieterName, LocalDate day){
        Long dieterid;
        Integer queryStatus = Integer.valueOf(0);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieterName);
        if (searchDieter.isPresent()){
            dieterid = searchDieter.get().getId();
            Optional<Meal> meal = mealRepository.findMealByDay(day, dieterid);
            if (meal.isPresent()){
                queryStatus = meal.get().getCalories();
            }
        }
        return queryStatus;
    }

    public Meal getMeal(String mealBlock){
        JSONObject jsonObject = new JSONObject(mealBlock);
        String dieter = jsonObject.get("dieterName").toString();
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter);
        if (searchDieter.isPresent()){
            Optional<Meal> meal = mealRepository.findMealByDay(LocalDate.now(), searchDieter.get().getId());
            if (meal.isPresent()){
                return meal.get();
            }
        }
        return null;
    }

}
