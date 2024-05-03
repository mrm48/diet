package com.trinsic.diet3.meal;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;

import java.time.LocalDate;

@Entity
@Table
/**
* Meal is an entity for storing date, dieter, description and number of calories
* for a given meal on a day.
* 
* See the {@link com.trinsic.diet3.meal.MealService} class for definitions of  
* interactions with the table, see {@link com.trinsic.diet3.meal.MealController} 
* for the REST API enpoints, see {@link com.trinsic.diet3.meal.MealRepository}
* for queries run on the Postgresql database.
* @author Matt Miller
* 
*/
public class Meal {
    @Id
    @SequenceGenerator(
        name = "meal_sequence",
        sequenceName = "meal_sequence",
        allocationSize = 1
    )
    @GeneratedValue(
        strategy = GenerationType.SEQUENCE,
        generator = "meal_sequence"
    )
    
    /**
     * The primary key.
     */
    private Long id;
    /**
     * The name of the meal: Breakfast, Lunch, Dinner etc.
     */
    private String name;
    /**
     * The day when the meal was consumed.
     */
    private LocalDate day;
    /**
     * The number of calories consumed during the meal.
     */
    private Integer calories;
    /**
     * Dieter table primary key representing the dieter who consumed the meal.
     */
    private Long dieterid;
    /**
     * The name of the dieter who consumed the meal.
     */
    private String dieter;

    /**
     *  Meal constructor with no parameters, no fields are set by default.
     */
    public Meal() {
    }

    /**
     *  Meal constructor expecting all fields other than the primary key.
     *  @param name The name of the meal: Breakfast, Lunch, Dinner etc. 
     *  @param day The day when the meal was consumed
     *  @param calories The number of calories consumed during the meal
     *  @param dieterid The primary key representing the dieter in the dieter table
     *  @param dieter The name of the dieter
     *  @return A Meal object initialized with the parameters passed. 
     */
    public Meal(String name, LocalDate day, Integer calories, Long dieterid, String dieter) {
        this.name = name;
        this.day = day;
        this.calories = calories;
        this.dieterid = dieterid;
        this.dieter = dieter;
    }

    /**
     *  Meal constructor accepting all fields, including the primary key  
     *  @param id The primary key of the new Meal object
     *  @param name The name of the meal: Breakfast, Lunch, Dinner etc. 
     *  @param day The day when the meal was consumed
     *  @param calories The number of calories consumed during the meal
     *  @param dieterid The primary key representing the dieter in the dieter table
     *  @param dieter The name of the dieter
     *  @return A Meal object initialized with the parameters passed. 
     */
    public Meal(Long id, String name, LocalDate day, Integer calories, Long dieterid, String dieter) {
        this.id = id;
        this.name = name;
        this.day = day;
        this.calories = calories;
        this.dieterid = dieterid;
        this.dieter = dieter;
    }

    /**
     *  Meal constructor expecting all fields except the primary key and day
     *  @param name The name of the meal: Breakfast, Lunch, Dinner etc. 
     *  @param calories The number of calories consumed during the meal
     *  @param dieterid The primary key representing the dieter in the dieter table
     *  @param dieter The name of the dieter
     *  @return A Meal object initialized with the parameters passed. 
     */
    public Meal(String name, Integer calories, Long dieterid, String dieter) {
        this.name = name;
        this.day = LocalDate.now();
        this.calories = calories;
        this.dieterid = dieterid;
        this.dieter = dieter;
    }

    /**
     *  Meal constructor expecting only the name, calories and dieter name
     *  @param name The name of the meal: Breakfast, Lunch, Dinner etc. 
     *  @param calories The number of calories consumed during the meal
     *  @param dieter The name of the dieter
     *  @return A Meal object initialized with the parameters passed. 
     */
    public Meal(String name, Integer calories, String dieter){
        this.name = name;
        this.day = LocalDate.now();
        this.calories = calories;
        this.dieter = dieter;
    }

    /**
     *  Get the name of the meal: Breakfast, Lunch, Dinner etc.
     *  @return The name of the meal as entered by the user.
     */
    public String getName() {
        return name;
    }

    /**
     *  Set the name of the meal: Breakfast, Lunch, Dinner etc.
     *  @param name The new name of the meal.
     */
    public void setName(String name) {
        this.name = name;
    }

    /**
     *  Get the day the meal was consumed. 
     *  @return The {@link LocalDate} object representing when the meal was consumed. 
     */
    public LocalDate getDay() {
        return day;
    }

    /**
     *  Set the day when the meal was consumed.
     *  @param day {@link LocalDate} object representing when the meal was consumed.
     */
    public void setDay(LocalDate day) {
        this.day = day;
    }

    
    /**
     *  Get the number of calories consumed during the meal. 
     *  @return The number of calories consumed, represented as an {@link Integer}.
     */
    public Integer getCalories() {
        return calories;
    }

    /**
     *  Set the day when the meal was consumed.
     *  @param calories The number of calories consumed represented as an {@link Integer}.
     */
    public void setCalories(Integer calories) {
        this.calories = calories;
    }

    /**
     *  Get the primary key of the dieter within the dieter table.
     *  @return The dieter key represented as a {@link Long}.
     */
    public Long getDieterId(){
        return dieterid;
    }

    /**
     *  Set the primary key of the dieter within the dieter table.
     *  @param dieterid The {@link Long} dieter id. 
     */
    public void setDieterId(Long dieterid){
        this.dieterid = dieterid;
    }

    /**
     *  Get the name of the dieter.
     *  @return The dieter represented as a {@link String}.
     */
    public String getDieter(){
        return dieter;
    }

    /**
     *  Set the name of the dieter.
     *  @param dieter The {@link String} dieter name. 
     */
    public void setDieter(String dieter){
        this.dieter = dieter;
    }

    /**
     *  Get the Meal primary key.
     *  @return The Meal primary key as a {@link Long}.
     */
    public Long getId(){
        return id;
    }

}
