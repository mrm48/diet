package com.trinsic.diet3.entrycreaterequest;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.meal.Meal;

public class Entrycreaterequest {

    private Food food;
    private Meal meal;

    public Entrycreaterequest() {
    }

    public Entrycreaterequest(Food food, Meal meal) {
        this.food = food;
        this.meal = meal;
    }

    public Food getFood(){
        return this.food;
    }

    public Meal getMeal(){
        return this.meal;
    }
}

