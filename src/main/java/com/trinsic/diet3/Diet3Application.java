package com.trinsic.diet3;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * Entry point for diet3
 * @author Matt Miller
 */
@SpringBootApplication
public class Diet3Application {

	/**
	 * No fields required for Diet3Application
	 */
	public Diet3Application(){

	}

	/**
	 * Start the REST API application using Spring Boot
	 * @param args Command line arguments passed to the application.
	 */
	public static void main(String[] args) {
		SpringApplication.run(Diet3Application.class, args);
	}

}
