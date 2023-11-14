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
    private Integer totalcalories;

    public Dieter() {
    }

    public Dieter (Long id, String name, Integer totalCalories){
        this.id = id;
        this.name = name;
        this.totalcalories = totalCalories;
    }

    public Dieter(String name, Integer totalCalories){
        this.name = name;
        this.totalcalories = totalCalories;
    }

    public String getName(){
        return this.name;
    }

    public void setName(String name){
        this.name = name;
    }

    public Integer getTotalCalories(){
        return this.totalcalories;
    }

    public void setTotalCalories(Integer totalCalories){
        this.totalcalories = totalCalories;
    }
}
