package com.trinsic.diet3.user;

import java.time.LocalDate;
import java.util.Optional;

import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface UserRepository 
        extends JpaRepository<User, Long>{
        
       @Query("SELECT u FROM User u WHERE u.name = ?1")
        Optional<User> findUserByName(String name);

       @Query("SELECT u.totalcalories FROM User u WHERE u.name = ?1")
        Integer findUserCaloriesByDay(String name);

        @Modifying
        @Query("INSERT INTO User (totalcalories, name) VALUES (?2, ?1)")
        Integer addUser(String name, Integer cals);

        @Modifying
        @Query("UPDATE u FROM User SET u.totalcalories=?2 WHERE u.name = ?1")
        Integer addTotalCalories(String name, Integer cals);

}

