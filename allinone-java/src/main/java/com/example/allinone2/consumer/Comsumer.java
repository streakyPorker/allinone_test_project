package com.example.allinone2.consumer;

import com.alibaba.dubbo.config.spring.context.annotation.EnableDubbo;

import com.example.allinone2.DemoService;
import com.example.allinone2.consumer.comp.DemoServiceComponent;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.PropertySource;

public class Comsumer {
    public static void main(String[] args) {
        AnnotationConfigApplicationContext context = new AnnotationConfigApplicationContext(ConsumerConfiguration.class);
        context.start();
        DemoService service = context.getBean("demoServiceComponent", DemoServiceComponent.class);
        String hello = service.hello("world");
        System.out.println("result :" + hello);
    }

    @Configuration
    @EnableDubbo(scanBasePackages = "com.example.consumer.comp")
    @PropertySource("classpath:static/dubbo-consumer.properties")
    @ComponentScan(value = {"com.example.consumer.comp"})
    static class ConsumerConfiguration {

    }
}
