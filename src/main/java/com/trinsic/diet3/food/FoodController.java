package com.trinsic.diet3.food;

import org.springframework.web.bind.annotation.*;


/**
 * REST API endpoints for working with the food table
 * @author Matt Miller
 */
@RestController
@RequestMapping(path = "api/v1/food")
public class FoodController {

    /**
     * FoodService object that specifies communication to the database
     */
    private final FoodService foodService;

    /**
     * Default constructor setting no fields
     */
    public FoodController(){
        foodService = null;
    }

    /**
     * Create a new food controller passing in the food service to communicate to the database
     * @param foodService The service to communicate with the database
     */
    public FoodController(FoodService foodService){
        this.foodService = foodService;
    }

    /**
     * Get the number of calories associated with the food specified by name
     * @param food The food object from the body of the request
     * @return The food object with the calories field populated from the food table
     */
	@GetMapping("/calories")
    @ResponseBody
	public Food listCalories(@RequestBody Food food){
        return foodService.listCalories(food);
	}

    /**
     * Add calories to a food object by name
     * @param food The food object from the body of the request
     * @return The food object from the database after the calories have been updated
     */
    @PostMapping("/calories")
    @ResponseBody
    public Food addCalories(@RequestBody Food food){
        return foodService.addCaloriesByName(food);
    }

    /**
     * Set the number of calories by name
     * @param food Food to update
     * @return The food after the change to the number of calories has been made
     */
    @PutMapping("/calories")
    @ResponseBody
    public Food setCaloriesByName(@RequestBody Food food){
        return foodService.updateCaloriesByName(food);
    }

    /**
     * Add a new food
     * @param food The food object to add to the database
     * @return The food object after it has been added to the database, null if it could not be added
     */
    @PostMapping("/")
    @ResponseBody
    public Food addFood(@RequestBody Food food){
        return foodService.addFood(food);
    }

    /**
     * Removes food by name of food provided in the body of the request
     * @param food Food object specified in the body of the request
     * @return The food object removed from the database or null if the food was not found.
     */
    @DeleteMapping("/")
    @ResponseBody
    public Food removeFood(@RequestBody Food food){
        return foodService.removeFood(food);
    }

}
