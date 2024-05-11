package com.trinsic.diet3.food;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;

@Entity
@Table
/**
* Food is an Entity for storing calorie and quantity information for a food. 
* 
* See the {@link com.trinsic.diet3.food.FoodService} class for definitions of  
* interactions with the table, see {@link com.trinsic.diet3.food.FoodController} 
* for the REST API enpoints, see {@link com.trinsic.diet3.food.FoodRepository}
* for queries run on the Postgresql database.
* @author Matt Miller
* 
*/
public class Food {
    @Id
    @SequenceGenerator(
        name = "food_sequence",
        sequenceName = "food_sequence",
        allocationSize = 1
    )
    @GeneratedValue(
        strategy = GenerationType.SEQUENCE,
        generator = "food_sequence"
    )
    /**
     * The primary key.
     */
    private Long id;
    /**
     * The name of the food.
     */
    private String name;
    /**
     * The quantity of food for the number of calories.
     */
    private Integer units;
    /**
     * The number of calories consumed when eating this food.
     */
    private Integer calories;

    public Food() {
    }

    /**
     *  Food constructor expecting all fields. 
     *  @param id The primary key for the new food.
     *  @param name The name of the Food: Cheerios, Toast, Pizza etc. 
     *  @param units The quantity of the food to be consumed for the number of calories logged.
     *  @param calories The number of calories consumed when eating the item.
     *  @return A new Food object initialized with the parameters passed. 
     */
    public Food(Long id, String name, Integer units, Integer calories) {
        this.id = id;
        this.name = name;
        this.units = units;
        this.calories = calories;
    }

    /**
     *  Food constructor expecting all fields other than the primary key.
     *  @param name The name of the Food: Cheerios, Toast, Pizza etc. 
     *  @param units The quantity of the food to be consumed for the number of calories logged.
     *  @param calories The number of calories consumed when eating the item.
     *  @return A new Food object initialized with the parameters passed. 
     */
    public Food(String name, Integer units, Integer calories) {
        this.name = name;
        this.units = units;
        this.calories = calories;
    }

    /**
     *  Set the name of the food
     *  @param name The name of the Food: Cheerios, Toast, Pizza etc. 
     */
    public void setName(String name){
        this.name = name;
    }

    /**
     *  Get the name of the food
     *  @return The name of the food: Cheerios, Toast, Pizza etc.
     */
    public String getName(){
        return this.name;
    }

    /**
     *  Set the units consumed for the number of calories specified in this food.
     *  @param units The quantity of the food to be consumed for the number of calories logged.
     */
    public void setUnits(Integer units){
        this.units = units;
    }

    /**
     *  Get the units consumed for the number of calories specified in this food.
     *  Food constructor expecting all fields other than the primary key.
     *  @return The quantity of the food to be consumed for the number of calories logged.
     */
    public Integer getUnits(){
        return this.units;
    }

    /**
     *  Set the number of calories consumed when eating the item.
     *  @param calories The number of calories consumed when eating the item.
     */
    public void setCalories(Integer calories){
        this.calories = calories;
    }

    /**
     *  Get the number of calories consumed when eating the item.
     *  @return The number of calories consumed when eating the item.
     */
    public Integer getCalories(){
        return this.calories;
    }

    /**
     * Get the primary key for this food.
     *  @return the primary key for this food.
     */
    public Long getID(){
        return this.id;
    }
}

