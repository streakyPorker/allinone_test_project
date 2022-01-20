package com.example.learn.spring.aop.mixin;

import org.springframework.stereotype.Component;

@Component
public class TestServiceImpl implements TestService{
    @Override
    public void methodInTestService() {
        System.out.println("impl for bark");
    }
}
