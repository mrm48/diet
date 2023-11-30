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
        Integer calories = Integer.valueOf(-1);
        if (searchFood.isPresent()){
            calories = searchFood.get().getCalories();
        }
        return calories;
	}

	@Transactional
    public Integer addCaloriesByName(String name, Integer cals){
        Integer queryStatus = Integer.valueOf(-1);
        System.out.println(name);
        System.out.println(cals);
        if(this.listCalories(name) != -1){
            System.out.println("changing " + name);
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
