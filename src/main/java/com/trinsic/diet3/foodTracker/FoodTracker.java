package com.trinsic.diet3.foodTracker;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;

@Entity
@Table
public class FoodTracker {
    @Id
    @SequenceGenerator(
        name = "foodtracker_sequence",
        sequenceName = "foodtracker_sequence",
        allocationSize = 1
    )
    @GeneratedValue(
        strategy = GenerationType.SEQUENCE,
        generator = "foodtracker_sequence"
    )
    private Long id;
    private Long food_id;
    private Long meal_id;
    private Integer servings;
    private Integer calories;

    public FoodTracker() {
    }

    public FoodTracker(Long id, Long food_id, Long meal_id, Integer servings, Integer calories) {
        this.id = id;
        this.food_id = food_id;
        this.meal_id = meal_id;
        this.servings = servings;
        this.calories = calories;
    }

    public FoodTracker(Long food_id, Long meal_id, Integer servings, Integer calories) {
        this.food_id = food_id;
        this.meal_id = meal_id;
        this.servings = servings;
        this.calories = calories;
    }

    public void setFood_Id(Long food_id){
        this.food_id = food_id;
    }

    public Long getFood_Id(){
        return this.food_id;
    }

    public void setMeal_Id(Long meal_id){
        this.meal_id = meal_id;
    }

    public Long getMeal_Id(){
        return this.meal_id;
    }

    public void setServings(Integer servings){
        this.servings = servings;
    }

    public Integer getServings(){
        return this.servings;
    }

    public void setCalories(Integer calories){
        this.calories = calories;
    }

    public Integer getCalories(){
        return this.calories;
    }

    public Long getID(){
        return this.id;
    }
}

