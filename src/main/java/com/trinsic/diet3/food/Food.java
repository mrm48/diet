package com.trinsic.diet3.food;

import java.time.LocalDate;

public class Food {
    private Long id;
    private String name;
    private Integer units;
    private Integer calories;

    public Food() {
    }

    public Food(Long id, String name, Integer units, Integer calories) {
        this.id = id;
        this.name = name;
        this.units = units;
        this.calories = calories;
    }

    public Food(String name, Integer units, Integer calories) {
        this.name = name;
        this.units = units;
        this.calories = calories;
    }

    public void setName(String name){
        this.name = name;
    }

    public String getName(){
        return this.name;
    }

    public void setUnits(Integer units){
        this.units = units;
    }

    public Integer getUnits(){
        return this.units;
    }

    public void setCalories(Integer calories){
        this.calories = calories;
    }

    public Integer getCalories(){
        return this.calories;
    }
}

