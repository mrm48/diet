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
    private String[] food;

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

    public Meal(Long id, String name, LocalDate day, Integer calories, Long dieterid, String dieter, String[] food) {
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

    public Meal(String name, String dieter, String[] food){
        this.name = name;
        this.dieter = dieter; 
        this.food = food;
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

    public void setFood(String[] food){
        this.food = food;
    }

    public void addFood(String[] food){
        String[] temp = this.food;
        int lastIdx = 0;
        System.out.println(this.food.length);
        this.food = new String[this.food.length + food.length];
        System.out.println(this.food.length);
        System.out.println(food.length);
        for (int i = 0; i < temp.length; i++) {
           this.food[i] = temp[i];
           lastIdx = i;
        }
        for (int j = lastIdx; j - lastIdx < food.length; j++) {
            System.out.println("j: " + j);
            this.food[j] = food[j - lastIdx];
        }
    }

    public String[] getFood(){
        return food;
    }

}
