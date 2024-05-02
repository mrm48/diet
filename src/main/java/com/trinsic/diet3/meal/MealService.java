
package com.trinsic.diet3.meal;
import java.time.LocalDate;
import java.util.Optional;
import java.util.List;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.food.FoodRepository;
import com.trinsic.diet3.entry.EntryRepository;
import com.trinsic.diet3.entry.Entry;
import com.trinsic.diet3.dieter.*;

@Service
/**
* MealService is intended to take requests from MealController and pass back
* Meal objects or Integers. 
* 
* See the {@link com.trinsic.diet3.meal.Meal} class for definitions of objects 
* passed
* back to {@link com.trinsic.diet3.meal.MealController}
* @author Matt Miller
* 
*/
public class MealService{

    /**
     * A MealRepository object to interact with the meal table
     */
    private final MealRepository mealRepository;

    /**
     * A DieterRepository object to interact with the dieter table
     */
    private final DieterRepository dieterRepository;
    
    /**
     * A FoodRepository object to interact with the food table
     */
    private final FoodRepository foodRepository;

    /**
     * An EntryRepository object to interact with the entry table
     */
    private final EntryRepository entryRepository;

    /**
     *  MealService constructor accepting fields for each repository type 
     *  @param mealRepository A MealRepository object to interact with the meal table
     *  @param dieterRepository A DieterRepository object to interact with the dieter table
     *  @param foodRepository A FoodRepository object to interact with the food table
     *  @param entryRepository An EntryRepository object to interact with the entry table
     */
    public MealService(MealRepository mealRepository, DieterRepository dieterRepository, FoodRepository foodRepository, EntryRepository entryRepository){
        this.mealRepository = mealRepository;
        this.dieterRepository = dieterRepository;
        this.foodRepository = foodRepository;
        this.entryRepository = entryRepository;
    }

    /**
     *  Add a new meal for the dieter, if it doesn't exist
     *  @param requestMeal A meal object to add to the database
     */
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

    /**
     *  Add a food item to a meal, if it exists
     *  @param requestMeal The meal object where food is being added 
     *  @param requestFood A food object to add to the meal
     *  @return The meal object from the database after the food is added
     */
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
                entryRepository.addFoodEntry(food.get().getID(), meal.get().getId(), food.get().getCalories());
                Integer newCalories = meal.get().getCalories() + food.get().getCalories();
                mealRepository.addFood(newCalories, meal.get().getId(), requestMealName, LocalDate.now(), dieter.get().getId(), dieter.get().getName());
                meal.get().setCalories(meal.get().getCalories() + food.get().getCalories());
                return meal.get();
            }
            else{
                Integer newMealStatus = mealRepository.addMeal(food.get().getCalories(),requestMealName,LocalDate.now(),dieter.get().getId(),requestDieter);
                if (newMealStatus != 0) {
                    Optional<Meal> newMeal = mealRepository.findMealByDay(LocalDate.now(), dieter.get().getId(), requestMealName);
                    if (newMeal.isPresent()){               
                        entryRepository.addFoodEntry(food.get().getID(), newMeal.get().getId(), food.get().getCalories());
                        return newMeal.get();
                    }
                }
            }
        }
        return null;
    }

    /**
     *  Remove a food item from a meal, if it exists
     *  @param requestMeal The meal object where food is being removed 
     *  @param requestFood A food object to remove from the meal
     *  @return The meal object from the database after the food is removed 
     */
    @Transactional
    public Meal removeCalories(Meal requestMeal, Food requestFood){
       String requestDieter;
       String requestMealName;
       requestDieter = requestMeal.getDieter();
       requestMealName = requestMeal.getName();
       Optional<Food> food = foodRepository.findFoodByName(requestFood.getName());
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
        if (dieter.isPresent() && food.isPresent()){
            Optional<Meal> meal = mealRepository.findMealByName(requestMealName, LocalDate.now(), dieter.get().getId());
            if (meal.isPresent()){
                List<Entry> entry = entryRepository.findEntryById(meal.get().getId(), food.get().getID(), food.get().getCalories());
                if (!entry.isEmpty()){
                    entryRepository.removeFoodEntry(entry.get(0).getID());
                    meal.get().setCalories(meal.get().getCalories() - food.get().getCalories());
                    return meal.get();
                }
            }
        }
        return null;
    }

    /**
     *  Check number of calories for a meal
     *  @param requestMeal The name of the meal for which calories are being checked, as a string
     *  @param requestDay A LocalDate object when the meal was consumed
     *  @param requestDieter The name of the dieter who ate the meal, as a string
     *  @return The number of calories for the meal 
     */
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

    /**
     *  Get the number of calories consumed by day for a dieter
     *  @param requestDieter The dieter's name, as a string
     *  @param requestDay A LocalDate object when the calories were consumed
     *  @return The number of calories consumed for the day by the dieter
     */
    public Integer getCaloriesByDay(String requestDieter, LocalDate requestDay){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
        if (dieter.isPresent()){
            Integer meal = mealRepository.findCaloriesByDay(requestDieter, requestDay);
            return meal;
        }
        return 0;
    }

    /**
     *  Get a meal object from the database, if it exists.
     *  @param requestMeal The meal object to retrieve from the database
     *  @return The meal object from the database or null if not found
     */
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
