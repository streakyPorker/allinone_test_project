package com.example.learn.spring.aop.mixin;

import org.aspectj.lang.annotation.After;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.DeclareParents;
import org.springframework.core.Ordered;
import org.springframework.stereotype.Component;


@Aspect
@Component
public class MixinInterceptor implements Ordered {

    @DeclareParents(value = "com.example.learn.spring.aop.AopEntity",defaultImpl = TestServiceImpl.class)
    public TestService testService;

    @After(" @args(com.example.learn.spring.aop.mixin.Asd,com.example.learn.spring.aop.mixin.Asd,..)")
    public void pfm(){

    }

    @Override
    public int getOrder() {
        return 2;
    }
}
