package com.trinsic.diet3.entry;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;

/**
 * Entry object stored in the entry table referring to the meal, food and dieter
 * @author Matt Miller
 */
@Entity
@Table
public class Entry {
    @Id
    @SequenceGenerator(
        name = "entry_sequence",
        sequenceName = "entry_sequence",
        allocationSize = 1
    )
    @GeneratedValue(
        strategy = GenerationType.SEQUENCE,
        generator = "entry_sequence"
    )
    private Long id;
    /**
     * The primary key in the food table related to the food in this entry
     */
    private Long food_id;
    /**
     * The primary key in the meal table related to the meal in this entry
     */
    private Long meal_id;
    /**
     * The number of calories consumed for this entry
     */
    private Integer calories;

    /**
     * Default constructor taking no arguments, does not set any fields
     */
    public Entry() {
    }

    /**
     * Constructor setting all fields
     * @param id The primary key within the entry table
     * @param food_id The primary key of the food consumed in the food table
     * @param meal_id The primary key of the meal related to this entry in the meal table
     * @param calories The number of calories consumed for this entry
     */
    public Entry(Long id, Long food_id, Long meal_id, Integer calories) {
        this.id = id;
        this.food_id = food_id;
        this.meal_id = meal_id;
        this.calories = calories;
    }

    /**
     * Constructor without requiring the primary key for the entry table
     * @param food_id The primary key of the food consumed in the food table
     * @param meal_id The primary key of the meal related to this entry in the meal table
     * @param calories The number of calories consumed for this entry
     */
    public Entry(Long food_id, Long meal_id, Integer calories) {
        this.food_id = food_id;
        this.meal_id = meal_id;
        this.calories = calories;
    }

    /**
     * Sets the food id from the food table to the parameter passed
     * @param food_id The primary key in the food table
     */
    public void setFood_Id(Long food_id){
        this.food_id = food_id;
    }

    /**
     * Gets the food id related to this entry
     * @return The primary key of the food consumed in the food table
     */
    public Long getFood_Id(){
        return this.food_id;
    }

    /**
     * Set the meal primary key from the meal table related to this entry
     * @param meal_id The primary key of the meal from the meal table
     */
    public void setMeal_Id(Long meal_id){
        this.meal_id = meal_id;
    }

    /**
     * Get the meal primary key from the meal table related to this entry
     * @return The meal primary key from the meal table related to this entry
     */
    public Long getMeal_Id(){
        return this.meal_id;
    }

    /**
     * Set the number of calories consumed with this entry
     * @param calories The number of calories consumed
     */
    public void setCalories(Integer calories){
        this.calories = calories;
    }

    /**
     * Get the number of calories consumed with this entry
     * @return The number of calories consumed
     */
    public Integer getCalories(){
        return this.calories;
    }

    /**
     * Get the primary key of this entry from the entry table
     * @return The primary key associated with this entry
     */
    public Long getID(){
        return this.id;
    }
}

