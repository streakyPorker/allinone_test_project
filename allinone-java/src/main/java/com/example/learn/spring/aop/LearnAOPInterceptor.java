package com.example.learn.spring.aop;

import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.After;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.springframework.core.Ordered;
import org.springframework.stereotype.Component;

import java.io.IOException;

@Component
@Aspect
public class LearnAOPInterceptor implements Ordered {

    @After("this(com.example.learn.spring.aop.mixin.TestService)")
    public void pfmThisTestService() throws Exception{
        System.out.println("aop of this as mixin mode activated here");
    }

    @After("target(com.example.learn.spring.aop.mixin.TestService)")
    public void pfmTargetTestService() throws Exception{
        System.out.println("aop of target as mixin mode activated here");
    }

    @After("this(com.example.learn.spring.aop.thisntarget.IDummy)")
    public void pfmThisSelfMethod() throws Exception{
        System.out.println("aop of this activated here");
    }

    @After("target(com.example.learn.spring.aop.thisntarget.IDummy)")
    public void pfmTargetSelfMethod() throws Exception{
        System.out.println("aop of target activated here");
    }

    @After("execution(* com.example.learn.spring.aop.ExecService.*(..))")
    public void pfmExecMethod() throws Exception{
        System.out.println("aop of execution activated here");
    }

    @After("within(com.example.learn.spring.aop.ExecService)")
    public void pfmWithinMethod() throws Exception{
        System.out.println("aop of within activated here");
    }




    @Override
    public int getOrder() {
        return 1;
    }
}

