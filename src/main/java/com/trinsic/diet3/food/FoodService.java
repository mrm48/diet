package com.trinsic.diet3.food;
import java.util.List;
import java.util.Optional;
import java.util.function.Consumer;

import org.springframework.stereotype.Service;

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
}
