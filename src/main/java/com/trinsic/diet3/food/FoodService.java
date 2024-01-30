package com.trinsic.diet3.food;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class FoodService{

    private final FoodRepository foodRepository;

    public FoodService(FoodRepository foodRepository){
        this.foodRepository = foodRepository;
    }

	public Integer listCalories(String name){
        Optional<Food> searchFood = foodRepository.findFoodByName(name);
        if (searchFood.isPresent()){
            return searchFood.get().getCalories();
        }
        return -1;
	}

	@Transactional
    public Integer addCaloriesByName(String name, Integer cals){
        if(this.listCalories(name) != -1){
            System.out.println("changing " + name);
            return foodRepository.addCaloriesByName(name, cals);    
        }
        return -1;
    }

    @Transactional
    public Food addFood(Food f){
        Optional<Food> searchFood = foodRepository.findFoodByName(f.getName());
        if (searchFood.isEmpty()) {
            foodRepository.addFood(f.getName(), f.getUnits(), f.getCalories());  
            return f;
        }
        return null;
    }

}
