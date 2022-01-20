package com.example.learn.spring;


import lombok.Data;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.EnvironmentAware;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Component;

@Component
@Data
public class LearnEnv implements  EnvironmentAware {


    @Value("ad")
    private String asd;

    private Environment environment;

    public void instanceMethod(){
        System.out.println("I`m instance method");
    }

    @Override
    public void setEnvironment(Environment environment) {
        this.environment = environment;
    }
}


