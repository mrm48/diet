
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

/**
 * MealService is intended to take requests from MealController and pass back
 * Meal objects or Integers.
 * See the {@link com.trinsic.diet3.meal.Meal} class for definitions of objects
 * passed
 * back to {@link com.trinsic.diet3.meal.MealController}
 * @author Matt Miller
 *
 */
@Service
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
     *  @return The meal object added to the database.
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
                    Optional<Meal> newMeal = mealRepository.findMealByName(requestMeal.getName(), requestMeal.getDay(), dieter.get().getId());
                    if (newMeal.isPresent()) {
                        return newMeal.get();
                    }
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
                return addEntry(meal.get(), food.get(), requestMealName, dieter.get());
            }
            else{
                Integer newMealStatus = mealRepository.addMeal(food.get().getCalories(),requestMealName,LocalDate.now(),dieter.get().getId(),requestDieter);
                if (newMealStatus != 0) {
                    Optional<Meal> newMeal = mealRepository.findMealByName(requestMealName, LocalDate.now(), dieter.get().getId());
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
     *  Add an entry to a meal that has been found in addCalories
     *  @param meal The meal object where food is being added 
     *  @param food A food object to add to the meal
     *  @param requestMealName Name of the meal as a String
     *  @oaram dieter Dieter object who is adding the meal
     *  @return The meal object from the database after the food is added
     */
    @Transactional
    private Meal addEntry(Meal meal, Food food, String requestMealName, Dieter dieter){
        entryRepository.addFoodEntry(food.getID(), meal.getId(), food.getCalories());
        Integer newCalories = meal.getCalories() + food.getCalories();
        mealRepository.addFood(newCalories, meal.getId(), requestMealName, LocalDate.now(), dieter.getId(), dieter.getName());
        meal.setCalories(meal.getCalories() + food.getCalories());
        return meal;
    }

    /**
     *  Remove a food item from a meal, if it exists
     *  @param requestMeal The meal object where food is being removed 
     *  @param requestFood A food object to remove from the meal
     *  @return The meal object from the database after the food is removed 
     */
    @Transactional
    public Meal removeCalories(Meal requestMeal, Food requestFood){
       String requestMealName;
       String requestDieter = requestMeal.getDieter();
       requestMealName = requestMeal.getName();
       Optional<Food> food = foodRepository.findFoodByName(requestFood.getName());
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter);
        if (dieter.isPresent() && food.isPresent()){
            Optional<Meal> meal = mealRepository.findMealByName(requestMealName, LocalDate.now(), dieter.get().getId());
            if (meal.isPresent()){
                List<Entry> entry = entryRepository.findEntryById(meal.get().getId(), food.get().getID(), food.get().getCalories());
                if (!entry.isEmpty()){
                    entryRepository.removeFoodEntry(entry.getFirst().getID());
                    meal.get().setCalories(meal.get().getCalories() - food.get().getCalories());
                    return meal.get();
                }
            }
        }
        return null;
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
            Optional<Meal> meal = mealRepository.findMealByName(requestMeal.getName(), LocalDate.now(), dieter.get().getId());
            if (meal.isPresent()){
                return meal.get();
            }
        }
        return null;
    }

    /**
     * Remove a meal from the database (requires meal name, day and dietername set correctly).
     * @param requestMeal The meal to be deleted
     * @return The meal that was deleted, null if the meal was not found.
     */
    public Meal removeMeal(Meal requestMeal){
       Optional<Meal> meal = mealRepository.findMealByDieter(requestMeal.getName(), requestMeal.getDay(), requestMeal.getDieter());
       if (meal.isPresent()){
           entryRepository.removeFoodFromMeal(meal.get().getId());
           mealRepository.deleteMealById(meal.get().getId());
           return meal.get();
       }
       return null;
    }

}
