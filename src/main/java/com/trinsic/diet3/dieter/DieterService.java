package com.trinsic.diet3.dieter;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.meal.MealService;

@Service
public class DieterService{

    private final DieterRepository dieterRepository;
    private final MealService mealService;

    public DieterService(DieterRepository dieterRepository, MealService mealService){
        this.dieterRepository = dieterRepository;
        this.mealService = mealService;
    }

    @Transactional
    public Dieter addDieter(Dieter newDieter){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(newDieter.getName());
        if (searchDieter.isEmpty()) {
            dieterRepository.addDieter(newDieter.getName(), newDieter.getCalories());  
            return newDieter;
        }
        return null;
    }

    @Transactional
    public Integer setCalories(String dieter){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter);
        if (searchDieter.isPresent()){
            return dieterRepository.addTotalCalories(searchDieter.get().getName(),searchDieter.get().getCalories());
        }
        return -1;
    }

    public Integer getCaloriesByDay(String dieter, LocalDate day){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter);
        if(searchDieter.isPresent()){
            Optional<Integer> currentCalories = dieterRepository.findDieterCaloriesByDay(searchDieter.get().getName());
            if (currentCalories.isPresent()){
                return currentCalories.get();
            }
        }
        return -1;
    }

    public Integer getRemainingCalories(Dieter dieter){
        LocalDate day = LocalDate.now();
        String dieterName = dieter.getName();
        Integer totalCalories = getCaloriesByDay(dieterName, day);
        Integer usedCalories = mealService.getCaloriesByDay(dieterName, day);
        return totalCalories - usedCalories;
    }

    public Long getID(Dieter dieter){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter.getName());
        if (searchDieter.isPresent()){
            return searchDieter.get().getId();
        }
        return Long.valueOf(-1);
    }

}
