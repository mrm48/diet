package com.trinsic.diet3.food;
import java.util.List;
import java.util.Optional;
import java.util.function.Consumer;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class FoodService{
    FoodRepository foodRepository;
    private Integer calories;
    public FoodService(FoodRepository foodRepository){
        this.foodRepository = foodRepository;
    }
	public Integer listCalories(String name){
        Optional<Food> searchFood = foodRepository.findFoodByName(name);
        this.calories = Integer.valueOf(-1);
        searchFood.ifPresent(this::setCalories);
        return this.calories;
	}

    public void setCalories(Food f){
        this.calories = f.getCalories();
    }

    public Integer addCaloriesByName(String name, Integer cals){
        Integer queryStatus = Integer.valueOf(-1);
        if(this.listCalories(name) != -1){
            queryStatus = foodRepository.addCaloriesByName(name, cals);    
        }
        return queryStatus;
    }

    @Transactional
    public Integer addFood(Food f){
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Food> searchFood = foodRepository.findFoodByName(f.getName());
        if (searchFood.isEmpty()) {
            queryStatus = foodRepository.addFood(f.getName(), f.getUnits(), f.getCalories());  
        }
        return queryStatus;
    }

}
