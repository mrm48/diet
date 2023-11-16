package com.trinsic.diet3.dieter;

import java.util.Optional;

import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface DieterRepository 
        extends JpaRepository<Dieter, Long>{
        
       @Query("SELECT d FROM Dieter d WHERE d.name = ?1")
        Optional<Dieter> findDieterByName(String name);

       @Query("SELECT d.totalcalories FROM Dieter d WHERE d.name = ?1")
        Optional<Integer> findDieterCaloriesByDay(String name);

        @Modifying
        @Query("INSERT INTO Dieter (name, totalcalories) VALUES (?1, ?2)")
        Integer addDieter(String name, Integer cals);

        @Modifying
        @Query("UPDATE Dieter SET totalcalories = ?2 WHERE name = ?1")
        Integer addTotalCalories(String name, Integer cals);

}

