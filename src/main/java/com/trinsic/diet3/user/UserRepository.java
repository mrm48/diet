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
        
       @Query("SELECT m FROM User m WHERE m.name = ?1 AND m.day = ?2")
        Optional<User> findUserByName(String name, LocalDate day);

        @Modifying
        @Query("INSERT INTO User (calories, name, day) VALUES (?2, ?1, ?3)")
        Integer addUser(String name, Integer cals, LocalDate day);
}

