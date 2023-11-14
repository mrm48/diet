package com.trinsic.diet3.user;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.Optional;
import java.time.LocalDate;

import com.trinsic.diet3.food.Food;
import com.trinsic.diet3.food.FoodRepository;

@RestController
@RequestMapping(path = "api/v1/user")
public class UserController{

    private final UserService userService;

    public UserController(UserService userService, FoodRepository foodRepository){
        this.userService = userService;
    }

    @PostMapping("/adduser")
    @ResponseBody
    public Integer addUser(@RequestBody User hUser){
        // call the User constructor that uses the current date as the value
        return userService.addUser(new User(hUser.getName(), hUser.getTotalCalories()));
    }
}
