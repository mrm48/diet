package com.trinsic.diet3.user;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.food.Food;

@Service
public class UserService{

    UserRepository userRepository;

    public UserService(UserRepository userRepository){
        this.userRepository = userRepository;
    }

    @Transactional
    public Integer addUser(User newUser){
        // Only add user if there is not a meal with the same name for the day
        Integer queryStatus = Integer.valueOf(-1);
        Optional<User> searchUser = userRepository.findUserByName(newUser.getName());
        if (searchUser.isEmpty()) {
            queryStatus = userRepository.addUser(newUser.getName(), newUser.getTotalCalories());  
        }
        return queryStatus;
    }

    @Transactional
    public Integer addCalories(User user){
        Integer queryStatus = Integer.valueOf(-1);
        Optional<User> searchUser = userRepository.findUserByName(user.getName());
        if (searchUser.isPresent()){
            queryStatus = userRepository.addTotalCalories(user.getName(),user.getTotalCalories());
        }
        return queryStatus;
    }

    public Integer getCaloriesByDay(User searchUser, LocalDate day){
        Optional<Integer> currentCalories = userRepository.findUserCaloriesByDay(searchUser, day);
        if (currentCalories.isPresent()){
            return currentCalories.get();
        }
        return 0;
    }

}