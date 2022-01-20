package com.example.learn.spring;

import com.example.learn.spring.aop.AopEntity;
import com.example.learn.spring.aop.ExecService;
import com.example.learn.spring.aop.mixin.TestService;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;

@Component
public class Runner implements ApplicationRunner {

    @Resource
    private AopEntity aopEntity;

    @Resource
    ExecService execService;


    @Override
    public void run(ApplicationArguments args) throws Exception {
        ((TestService)aopEntity).methodInTestService();
        aopEntity.implMethod();
    }
}
