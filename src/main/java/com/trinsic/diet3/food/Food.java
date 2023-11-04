package com.trinsic.diet3.food;

import java.time.LocalDate;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;

@Entity
@Table
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

    public Long getID(){
        return this.id;
    }
}

