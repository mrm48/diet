package com.trinsic.diet3.dieter;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.meal.MealRepository;

@Service
public class DieterService{

    private final DieterRepository dieterRepository;
    private final MealRepository mealRepository;

    public DieterService(DieterRepository dieterRepository, MealRepository mealRepository){
        this.dieterRepository = dieterRepository;
        this.mealRepository = mealRepository;
    }

    @Transactional
    public Dieter addDieter(Dieter requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        if (dieter.isEmpty()) {
            dieterRepository.addDieter(requestDieter.getName(), requestDieter.getCalories());  
            return getID(requestDieter);
        }
        return null;
    }

    @Transactional
    public Dieter setCalories(Dieter requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        if (dieter.isPresent()){
            dieterRepository.addTotalCalories(requestDieter.getName(),requestDieter.getCalories());
            dieter.get().setCalories(requestDieter.getCalories());
            return dieter.get();
        }
        return null;
    }

    public Dieter getCaloriesByDay(Dieter requestDieter, LocalDate requestDay){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        Dieter foundDieter = null; 
        if(dieter.isPresent()){
            foundDieter = dieter.get();
            Integer currentCalories = mealRepository.findCaloriesByDay(foundDieter.getName(), requestDay);
            if (currentCalories != null) {
                foundDieter.setCalories(currentCalories);
            }
            else {
                foundDieter.setCalories(0);
            }
            return foundDieter;
        }
        return foundDieter;
    }

    public Dieter getRemainingCalories(Dieter requestDieter){
        LocalDate day = LocalDate.now();
        String dieterName = requestDieter.getName();
        Optional<Dieter> dieter = dieterRepository.findDieterByName(dieterName);
        Integer totalCalories = dieter.get().getCalories();
        Dieter responseDieter = new Dieter();
        responseDieter = getCaloriesByDay(requestDieter, day);
        if(dieter.isPresent()){
            responseDieter.setName(dieterName);
            if(responseDieter.getCalories() != null){
                responseDieter.setCalories(totalCalories - responseDieter.getCalories());
            }
            else{
                responseDieter.setCalories(totalCalories);
            }
            return responseDieter;
        }
        return responseDieter;
    }

    public Dieter getID(Dieter requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        if (dieter.isPresent()){
            return dieter.get();
        }
        return null;
    }

    public Dieter getDieterByName(Dieter requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        if (dieter.isPresent()){
            return dieter.get();
        }
        return null;
    }
    

}
