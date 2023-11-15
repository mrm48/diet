package com.trinsic.diet3.dieter;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class DieterService{

    DieterRepository dieterRepository;

    public DieterService(DieterRepository dieterRepository){
        this.dieterRepository = dieterRepository;
    }

    @Transactional
    public Integer addDieter(Dieter newDieter){
        // Only add dieter if there is not a meal with the same name for the day
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(newDieter.getName());
        if (searchDieter.isEmpty()) {
            queryStatus = dieterRepository.addDieter(newDieter.getName(), newDieter.getTotalCalories());  
        }
        return queryStatus;
    }

    @Transactional
    public Integer addCalories(Dieter dieter){
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter.getName());
        if (searchDieter.isPresent()){
            queryStatus = dieterRepository.addTotalCalories(dieter.getName(),dieter.getTotalCalories());
        }
        return queryStatus;
    }

    public Integer getCaloriesByDay(Dieter searchDieter, LocalDate day){
        Optional<Integer> currentCalories = dieterRepository.findDieterCaloriesByDay(searchDieter.getName());
        if (currentCalories.isPresent()){
            return currentCalories.get();
        }
        return 0;
    }

}
