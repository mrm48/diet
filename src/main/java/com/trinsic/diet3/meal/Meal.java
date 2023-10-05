package com.trinsic.diet3.meal;

import com.trinsic.diet3.food.Food;

import java.time.LocalDate;
import java.util.List;

public class Meal {
    private Long id;
    private String name;
    private LocalDate day;
    private List<Food> items;
    private Integer calories;

    public Meal() {
    }

    public Meal(String name, LocalDate day, List<Food> items, Integer calories) {
        this.name = name;
        this.day = day;
        this.items = items;
        this.calories = calories;
    }

    public Meal(Long id, String name, LocalDate day, List<Food> items, Integer calories) {
        this.id = id;
        this.name = name;
        this.day = day;
        this.items = items;
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

    public List<Food> getItems() {
        return items;
    }

    public void setItems(List<Food> items) {
        this.items = items;
    }

    public void addItem(Food item){
        this.items.add(item);
        updateCalories();
    }

    public void deleteItem(String item){
        for (Food food : items) {
           if(food.getName() == item){
               this.items.remove(items.indexOf(food));
           }
        }
        updateCalories();
    }

    public Integer getCalories() {
        return calories;
    }

    public void setCalories(Integer calories) {
        this.calories = calories;
    }

    public void updateCalories(){
        calories = 0; 
        for (Food food : items) {
            calories += food.getCalories();
        }
    }
}
