package com.trinsic.diet3.dieter;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class DieterService{

    private final DieterRepository dieterRepository;

    public DieterService(DieterRepository dieterRepository){
        this.dieterRepository = dieterRepository;
    }

    @Transactional
    public Integer addDieter(Dieter newDieter){
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(newDieter.getName());
        if (searchDieter.isEmpty()) {
            queryStatus = dieterRepository.addDieter(newDieter.getName(), newDieter.getCalories());  
        }
        return queryStatus;
    }

    @Transactional
    public Integer setCalories(String dieter){
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter);
        if (searchDieter.isPresent()){
            queryStatus = dieterRepository.addTotalCalories(searchDieter.get().getName(),searchDieter.get().getCalories());
        }
        return queryStatus;
    }

    public Integer getCaloriesByDay(String dieter, LocalDate day){
        Integer queryStatus = Integer.valueOf(-1);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter);
        if(searchDieter.isPresent()){
            Optional<Integer> currentCalories = dieterRepository.findDieterCaloriesByDay(searchDieter.get().getName());
            if (currentCalories.isPresent()){
                queryStatus = currentCalories.get();
            }
        }
        return queryStatus;
    }

    public Long getID(Dieter dieter){
        Long queryStatus = Long.valueOf(-1);
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter.getName());
        if (searchDieter.isPresent()){
            queryStatus = searchDieter.get().getId();
        }
        return queryStatus;
    }

}
