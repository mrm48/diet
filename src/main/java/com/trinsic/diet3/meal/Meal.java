package com.trinsic.diet3.meal;

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
    private Integer calories;
    private Long dieterid;
    private String dieter;
    private List<String> food;

    public Meal() {
    }

    public Meal(String name, LocalDate day, Integer calories, Long dieterid, String dieter) {
        this.name = name;
        this.day = day;
        this.calories = calories;
        this.dieterid = dieterid;
        this.dieter = dieter;
    }

    public Meal(Long id, String name, LocalDate day, Integer calories, Long dieterid, String dieter) {
        this.id = id;
        this.name = name;
        this.day = day;
        this.calories = calories;
        this.dieterid = dieterid;
        this.dieter = dieter;
    }

    public Meal(Long id, String name, LocalDate day, Integer calories, Long dieterid, String dieter, List<String> food) {
        this.id = id;
        this.name = name;
        this.day = day;
        this.calories = calories;
        this.dieterid = dieterid;
        this.dieter = dieter;
        this.food = food;
    }
    public Meal(String name, Integer calories, Long dieterid, String dieter) {
        this.name = name;
        this.day = LocalDate.now();
        this.calories = calories;
        this.dieterid = dieterid;
        this.dieter = dieter;
    }

    public Meal(String name, String dieter, String food){
        this.name = name;
        this.dieter = dieter; 
        this.food.add(food);
    }

    public Meal(String name, Integer calories, String dieter){
        this.name = name;
        this.day = LocalDate.now();
        this.calories = calories;
        this.dieter = dieter;
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

    public Long getDieterId(){
        return dieterid;
    }

    public void setDieterId(Long dieterid){
        this.dieterid = dieterid;
    }

    public String getDieter(){
        return dieter;
    }

    public void setDieter(String dieter){
        this.dieter = dieter;
    }

    public Long getId(){
        return id;
    }

    public void setFood(List<String> food){
        this.food = food;
    }

    public void addFood(List<String> food){
        if(this.food != null){
            this.food.addAll(food);
        }
        else {
            this.food = food;
        }
    }

    public List<String> getFood(){
        return food;
    }

}
