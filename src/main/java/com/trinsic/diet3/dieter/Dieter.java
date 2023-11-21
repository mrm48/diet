package com.trinsic.diet3.dieter;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;

@Entity
@Table
public class Dieter {
    @Id
    @SequenceGenerator(
        name = "dieter_sequence",
        sequenceName = "dieter_sequence",
        allocationSize = 1
    )
    @GeneratedValue(
        strategy = GenerationType.SEQUENCE,
        generator = "dieter_sequence"
    ) 
    private Long id;
    private String name;
    private Integer calories;

    public Dieter() {
    }

    public Dieter (Long id, String name, Integer calories){
        this.id = id;
        this.name = name;
        this.calories = calories;
    }

    public Dieter(String name, Integer calories){
        this.name = name;
        this.calories = calories;
    }

    public String getName(){
        return this.name;
    }

    public void setName(String name){
        this.name = name;
    }

    public Integer getCalories(){
        return this.calories;
    }

    public void setCalories(Integer calories){
        this.calories = calories;
    }

    public Long getId(){
        return this.id;
    }
}
