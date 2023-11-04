package com.trinsic.diet3.food;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface FoodRepository 
    extends JpaRepository<Food, Long>{

        @Query("Select f FROM Food f WHERE f.name = ?1")
        Optional<Food> findFoodByName(String name);

        @Modifying
        @Query("UPDATE Food f SET f.calories = ?2 WHERE name = ?1")
        Integer addCaloriesByName(String name, Integer cals);
    
}
