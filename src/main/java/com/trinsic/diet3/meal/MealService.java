package com.trinsic.diet3.meal;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.food.FoodRepository;
import com.trinsic.diet3.entry.Entry;
import com.trinsic.diet3.entry.EntryRepository;
import com.trinsic.diet3.dieter.*;

@Service
public class MealService{

    private final MealRepository mealRepository;
    private final DieterRepository dieterRepository;
    private final FoodRepository foodRepository;
    private final EntryRepository entryRepository;

    public MealService(MealRepository mealRepository, DieterRepository dieterRepository, FoodRepository foodRepository, EntryRepository entryRepository){
        this.mealRepository = mealRepository;
        this.dieterRepository = dieterRepository;
        this.foodRepository = foodRepository;
        this.entryRepository = entryRepository;
    }

    @Transactional
    public Meal addMeal(Meal requestMeal){
        String requestDieter;
        requestDieter = requestMeal.getDieter();
        Optional<Meal> meal = mealRepository.findMealByName(requestMeal.getName(), requestMeal.getDay(), requestMeal.getDieterId());
        if (meal.isEmpty()) {
            requestMeal.setDay(LocalDate.now());
            Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
            if (dieter.isPresent()) {
                requestMeal.setDieterId(dieter.get().getId());
                requestMeal.setDieter(requestDieter);
                Integer newMealStatus = mealRepository.addMeal(0, requestMeal.getName(), requestMeal.getDay(), dieter.get().getId(), requestDieter);
                if (newMealStatus != 0) {
                    Optional<Meal> newMeal = mealRepository.findMealByDay(requestMeal.getDay(), dieter.get().getId(), requestMeal.getName());               
                    return newMeal.get();
                }
            }
        }
        return null;
    }

    @Transactional
    public Meal addCalories(Meal requestMeal, Food requestFood){
       String requestDieter;
       String requestMealName;
       requestDieter = requestMeal.getDieter();
       requestMealName = requestMeal.getName();
       Optional<Food> food = foodRepository.findFoodByName(requestFood.getName());
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
        if (dieter.isPresent() && food.isPresent()){
            Optional<Meal> meal = mealRepository.findMealByName(requestMealName, LocalDate.now(), dieter.get().getId());
            if (meal.isPresent()){
                entryRepository.addFoodEntry(meal.get().getId(), food.get().getID(), food.get().getCalories());
                Integer newCalories = meal.get().getCalories() + food.get().getCalories();
                mealRepository.addFood(newCalories, meal.get().getId(), requestMealName, LocalDate.now(), dieter.get().getId(), dieter.get().getName());
                meal.get().setCalories(meal.get().getCalories() + food.get().getCalories());
                return meal.get();
            }
            else{
                Integer newMealStatus = mealRepository.addMeal(food.get().getCalories(),requestMealName,LocalDate.now(),dieter.get().getId(),requestDieter);
                if (newMealStatus != 0) {
                    Optional<Meal> newMeal = mealRepository.findMealByDay(requestMeal.getDay(), dieter.get().getId(), requestMeal.getName());               
                    entryRepository.addFoodEntry(food.get().getID(), newMeal.get().getId(), food.get().getCalories());
                    return newMeal.get();
                }
            }
        }
        return null;
    }

    public Integer getCalories(String requestMeal, LocalDate requestDay, String requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
        if (dieter.isPresent()){
            Optional<Meal> meal = mealRepository.findMealByName(requestMeal, requestDay, dieter.get().getId());
            if (meal.isPresent()){
                return entryRepository.findCaloriesByMeal(meal.get().getId());
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
