package com.example.learn.spring.aop;
import com.example.learn.spring.aop.thisntarget.IDummy;
import org.springframework.stereotype.Component;


@Component

public class AopEntity implements ExecService,IDummy{
    // 实现接口会使得spring使用JDK动态代理实现AOP
    public void selfFunc(){
        System.out.println("func of mine");
    }

    @Override
    public void implMethod() {
        System.out.println("this is impl method");
    }
}
