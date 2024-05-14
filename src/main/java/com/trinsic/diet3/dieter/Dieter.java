package com.trinsic.diet3.dieter;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;

/**
 * Dieter class representing the dieter and their target number of calories
 * @author Matt Miller
 */
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
    /**
     * The name of this dieter.
     */
    private String name;
    /**
     * The target number of calories for this dieter.
     */
    private Integer calories;

    public Dieter() {
    }

    /**
     * Creates a dieter with all specified parameters.
     * @param id primary key for this dieter.
     * @param name name for this dieter.
     * @param calories the target number of calories for this dieter.
     */
    public Dieter (Long id, String name, Integer calories){
        this.id = id;
        this.name = name;
        this.calories = calories;
    }

    /**
     * Creates a dieter without the primary key parameter.
     * @param name name for this dieter.
     * @param calories the target number of calories for this dieter.
     */
    public Dieter(String name, Integer calories){
        this.name = name;
        this.calories = calories;
    }

    /**
     * Get the name of this dieter.
     * @return the String name for this dieter.
     */
    public String getName(){
        return this.name;
    }

    /**
     * Set the name of this dieter.
     * @param name the String name for this dieter.
     */
    public void setName(String name){
        this.name = name;
    }

    /**
     * Get the target number of calories for this dieter.
     * @return the Integer number of calories for this dieter.
     */
    public Integer getCalories(){
        return this.calories;
    }

    /**
     * Set the target number of calories for this dieter.
     * @param calories the Integer number of calories for this dieter.
     */
    public void setCalories(Integer calories){
        this.calories = calories;
    }

    /**
     * Get the primary key (ID) of this dieter.
     * @return the Long ID.
     */
    public Long getId(){
        return this.id;
    }
}
