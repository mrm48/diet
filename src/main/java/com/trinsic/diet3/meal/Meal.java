package com.trinsic.diet3.meal;

//import com.trinsic.diet3.food.Food;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;

import java.time.LocalDate;
import java.util.List;

@Entity
@Table
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
    private Long id;
    private String name;
    private LocalDate day;
    //private List<Food> items;
    private Integer calories;

    public Meal() {
    }

    public Meal(String name, LocalDate day, Integer calories) {
        this.name = name;
        this.day = day;
        //this.items = items;
        this.calories = calories;
    }

    public Meal(Long id, String name, LocalDate day, Integer calories) {
        this.id = id;
        this.name = name;
        this.day = day;
        this.calories = calories;
    }

    public Meal(String name, Integer calories) {
        this.name = name;
        this.day = LocalDate.now();
        this.calories = calories;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public LocalDate getDay() {
        return day;
    }

    public void setDay(LocalDate day) {
        this.day = day;
    }

    public Integer getCalories() {
        return calories;
    }

    public void setCalories(Integer calories) {
        this.calories = calories;
    }
}
