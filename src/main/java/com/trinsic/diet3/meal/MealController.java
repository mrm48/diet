package com.trinsic.diet3.meal;

import org.springframework.web.bind.annotation.RestController;

import com.trinsic.diet3.entrycreaterequest.Entrycreaterequest;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@RestController
@RequestMapping(path = "api/v1/meal")
/**
* MealController is a REST controller that interacts with {@link com.trinsic.diet3.meal.MealService}
* 
* See the {@link com.trinsic.diet3.meal.MealService} class for definitions of  
* interactions with the table, see {@link com.trinsic.diet3.meal.Meal} 
* for the Entity representing items in the table, see {@link com.trinsic.diet3.meal.MealRepository}
* for queries run on the Postgresql database.
* @author Matt Miller
* 
*/
public class MealController{

    /**
     * The service utilized to interact with the meal table.
     */
    private final MealService mealService;

    /**
     *  MealController constructor accepting the MealService to interact with the database.
     *
     *  @param mealService MealService object that will interact with {@link com.trinsic.diet3.meal.MealRepository}
     */
    public MealController(MealService mealService){
        this.mealService = mealService;
    }

    @PostMapping("/food")
    @ResponseBody
    /**
     *  Add food to the meal with an {@link com.trinsic.diet3.entrycreaterequest} object.
     *
     *  @param ecr Entrycreaterequest containing the meal and the food to add.
     *  @return Meal object from the database after the food has been added.
     */
    public Meal addFood(@RequestBody Entrycreaterequest ecr){
        return mealService.addCalories(ecr.getMeal(), ecr.getFood());
    }

    @DeleteMapping("/food")
    @ResponseBody
    /**
     *  Remove food from the meal with an {@link com.trinsic.diet3.entrycreaterequest} object.
     *
     *  @param ecr Entrycreaterequest containing the meal and the food to remove.
     *  @return Meal object from the database after the food has been removed.
     */
    public Meal removeFood(@RequestBody Entrycreaterequest ecr){
        return mealService.removeCalories(ecr.getMeal(), ecr.getFood());
    }

    @PostMapping("/")
    @ResponseBody
    /**
     *  Add meal to the database.
     *
     *  @param meal Meal object to add to the database.  
     *  @return Meal object from the database after the meal has been added.
     */
    public Meal addMeal(@RequestBody Meal meal){
        return mealService.addMeal(meal);
    }

    @GetMapping("/")
    @ResponseBody
    /**
     *  Retrieves the requested meal from the database. 
     *
     *  @param meal Meal to retrieve from the database, requires the day and name and dieter name.
     *  @return Meal object from the database or null if not found
     */
    public Meal getMeal(@RequestBody Meal meal){
        return mealService.getMeal(meal);
    }

}
