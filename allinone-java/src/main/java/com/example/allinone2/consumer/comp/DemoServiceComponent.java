package com.example.allinone2.consumer.comp;

import com.example.allinone2.DemoService;

import org.apache.dubbo.config.annotation.DubboReference;
import org.springframework.stereotype.Component;

@Component("demoServiceComponent")
public class DemoServiceComponent implements DemoService {

    @DubboReference
    private DemoService demoService;

    @Override
    public String hello(String name) {
        return demoService.hello(name);
    }

}
