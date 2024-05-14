package com.trinsic.diet3.entrycreaterequest;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.meal.Meal;

/**
 * Utility class to specify the meal and food for interacting with the entry table.
 * @author Matt Miller
 */
public class Entrycreaterequest {

    /**
     * A food object with at minimum the name field populated
     */
    private Food food;
    /**
     * A meal object with at minimum the dieter name and meal name fields populated
     */
    private Meal meal;

    /**
     * Default constructor setting no fields
     */
    public Entrycreaterequest() {
    }

    /**
     * Create a new Entrycreaterequest object with the specified food and meal objects.
     * @param food The food related to the entry
     * @param meal The meal related to the entry
     */
    public Entrycreaterequest(Food food, Meal meal) {
        this.food = food;
        this.meal = meal;
    }

    /**
     * Get the food object related to the entry
     * @return Food object within this entry
     */
    public Food getFood(){
        return this.food;
    }

    /**
     * Get the meal object related to the entry
     * @return Meal object within this entry
     */
    public Meal getMeal(){
        return this.meal;
    }
}

