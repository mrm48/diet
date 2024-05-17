package com.trinsic.diet3.food;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

/**
 * Food service to communicate between the REST API endpoints and the database.
 * @author Matt Miller
 */
@Service
public class FoodService{

    /**
     * The food repository that runs queries on the database.
     */
    private final FoodRepository foodRepository;

    /**
     * Constructor taking the food repository for communicating with the database
     * @param foodRepository Food repository responsible for running queries.
     */
    public FoodService(FoodRepository foodRepository){
        this.foodRepository = foodRepository;
    }

    /**
     * Get the number of calories by the name of the food
     * @param requestFood Food object specifying the name of the food
     * @return Food object with the number of calories, null if the food could not be found
     */
	public Food listCalories(Food requestFood){
        Optional<Food> food = foodRepository.findFoodByName(requestFood.getName());
        if (food.isPresent()){
            return food.get();
        }
        return null;
	}

    /**
     * Update the number of calories by the name of the food
     * @param requestFood Food object specifying the name of the food and the new calories
     * @return The food object after the number of calories have been updated, null if not found
     */
	@Transactional
    public Food updateCaloriesByName(Food requestFood){
        Food responseFood = this.listCalories(requestFood);
        if(responseFood != null){
            responseFood.setCalories(requestFood.getCalories());
            foodRepository.addCaloriesByName(responseFood.getName(), responseFood.getCalories());    
            return responseFood;
        }
        return null;
    }


    /**
     * Set the number of calories by name
     * @param requestFood Food object specifying the name and number of calories
     * @return The food object after it has been modified, null if it is not found
     */
	@Transactional
    public Food addCaloriesByName(Food requestFood){
        Food responseFood = this.listCalories(requestFood);
        if(responseFood != null){
            responseFood.setCalories(requestFood.getCalories());
            foodRepository.addCaloriesByName(responseFood.getName(), responseFood.getCalories());    
            return responseFood;
        }
        else{
            foodRepository.addFood(requestFood.getName(), requestFood.getUnits(), requestFood.getCalories());
        }
        return null;
    }

    /**
     * Add a new food to the database
     * @param requestFood The new food, specifying the name, units and number of calories
     * @return The new food that has been added, null if unsuccessful
     */
    @Transactional
    public Food addFood(Food requestFood){
        Optional<Food> food = foodRepository.findFoodByName(requestFood.getName());
        if (food.isEmpty()) {
            foodRepository.addFood(requestFood.getName(), requestFood.getUnits(), requestFood.getCalories());  
            return listCalories(requestFood);
        }
        return null;
    }

    /**
     * Remove the food where the name matches the food provided
     * @param requestFood Food to remove (must have name field defined)
     * @return The food removed or null if it was not found.
     */
    @Transactional
    public Food removeFood(Food requestFood){
        Optional<Food> food = foodRepository.findFoodByName(requestFood.getName());
        if (food.isPresent()){
            foodRepository.removeFoodByName(requestFood.getName());
            return food.get();
        }
        return null;
    }

}
